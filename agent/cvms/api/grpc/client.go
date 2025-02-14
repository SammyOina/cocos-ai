// Copyright (c) Ultraviolet
// SPDX-License-Identifier: Apache-2.0
package grpc

import (
	"context"
	"log/slog"
	"sync"
	"time"

	"github.com/absmach/magistrala/pkg/errors"
	"github.com/ultravioletrs/cocos/agent"
	"github.com/ultravioletrs/cocos/agent/cvms"
	"github.com/ultravioletrs/cocos/agent/cvms/server"
	"golang.org/x/sync/errgroup"
	"google.golang.org/protobuf/proto"
)

var (
	errCorruptedManifest  = errors.New("received manifest may be corrupted")
	errUnknonwMessageType = errors.New("unknown message type")
	sendTimeout           = 5 * time.Second
)

type CVMSClient struct {
	mu            sync.Mutex
	stream        cvms.Service_ProcessClient
	svc           agent.Service
	messageQueue  chan *cvms.ClientStreamMessage
	logger        *slog.Logger
	runReqManager *runRequestManager
	sp            server.AgentServer
}

// NewClient returns new gRPC client instance.
func NewClient(stream cvms.Service_ProcessClient, svc agent.Service, messageQueue chan *cvms.ClientStreamMessage, logger *slog.Logger, sp server.AgentServer) CVMSClient {
	return CVMSClient{
		stream:        stream,
		svc:           svc,
		messageQueue:  messageQueue,
		logger:        logger,
		runReqManager: newRunRequestManager(),
		sp:            sp,
	}
}

func (client *CVMSClient) Process(ctx context.Context, cancel context.CancelFunc) error {
	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		return client.handleIncomingMessages(ctx)
	})

	eg.Go(func() error {
		return client.handleOutgoingMessages(ctx)
	})

	return eg.Wait()
}

func (client *CVMSClient) handleIncomingMessages(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			req, err := client.stream.Recv()
			if err != nil {
				return err
			}
			if err := client.processIncomingMessage(ctx, req); err != nil {
				return err
			}
		}
	}
}

func (client *CVMSClient) processIncomingMessage(ctx context.Context, req *cvms.ServerStreamMessage) error {
	switch mes := req.Message.(type) {
	case *cvms.ServerStreamMessage_RunReqChunks:
		return client.handleRunReqChunks(ctx, mes)
	case *cvms.ServerStreamMessage_StopComputation:
		go client.handleStopComputation(ctx, mes)
	default:
		return errUnknonwMessageType
	}
	return nil
}

func (client *CVMSClient) handleRunReqChunks(ctx context.Context, msg *cvms.ServerStreamMessage_RunReqChunks) error {
	buffer, complete := client.runReqManager.addChunk(msg.RunReqChunks.Id, msg.RunReqChunks.Data, msg.RunReqChunks.IsLast)

	if complete {
		var runReq cvms.ComputationRunReq
		if err := proto.Unmarshal(buffer, &runReq); err != nil {
			return errors.Wrap(err, errCorruptedManifest)
		}

		go client.executeRun(ctx, &runReq)
	}

	return nil
}

func (client *CVMSClient) executeRun(ctx context.Context, runReq *cvms.ComputationRunReq) {
	ac := agent.Computation{
		ID:          runReq.Id,
		Name:        runReq.Name,
		Description: runReq.Description,
	}

	if runReq.Algorithm != nil {
		ac.Algorithm = agent.Algorithm{
			Hash:    [32]byte(runReq.Algorithm.Hash),
			UserKey: runReq.Algorithm.UserKey,
		}
	}

	for _, ds := range runReq.Datasets {
		ac.Datasets = append(ac.Datasets, agent.Dataset{
			Hash:    [32]byte(ds.Hash),
			UserKey: ds.UserKey,
		})
	}

	for _, rc := range runReq.ResultConsumers {
		ac.ResultConsumers = append(ac.ResultConsumers, agent.ResultConsumer{
			UserKey: rc.UserKey,
		})
	}

	if err := client.svc.InitComputation(ctx, ac); err != nil {
		client.logger.Warn(err.Error())
		return
	}

	client.mu.Lock()
	defer client.mu.Unlock()

	if runReq.AgentConfig == nil {
		runReq.AgentConfig = &cvms.AgentConfig{}
	}

	runRes := &cvms.ClientStreamMessage_RunRes{
		RunRes: &cvms.RunResponse{
			ComputationId: runReq.Id,
		},
	}

	err := client.sp.Start(ctx, agent.AgentConfig{
		Port:         runReq.AgentConfig.Port,
		Host:         runReq.AgentConfig.Host,
		CertFile:     runReq.AgentConfig.CertFile,
		KeyFile:      runReq.AgentConfig.KeyFile,
		ServerCAFile: runReq.AgentConfig.ServerCaFile,
		ClientCAFile: runReq.AgentConfig.ClientCaFile,
		AttestedTls:  runReq.AgentConfig.AttestedTls,
	}, ac)
	if err != nil {
		client.logger.Warn(err.Error())
		runRes.RunRes.Error = err.Error()
	}

	client.sendMessage(&cvms.ClientStreamMessage{Message: runRes})
}

func (client *CVMSClient) handleStopComputation(ctx context.Context, mes *cvms.ServerStreamMessage_StopComputation) {
	msg := &cvms.ClientStreamMessage_StopComputationRes{
		StopComputationRes: &cvms.StopComputationResponse{
			ComputationId: mes.StopComputation.ComputationId,
		},
	}
	if err := client.svc.StopComputation(ctx); err != nil {
		msg.StopComputationRes.Message = err.Error()
	}

	client.mu.Lock()
	defer client.mu.Unlock()

	if err := client.sp.Stop(); err != nil {
		msg.StopComputationRes.Message = err.Error()
	}

	client.sendMessage(&cvms.ClientStreamMessage{Message: msg})
}

func (client *CVMSClient) handleOutgoingMessages(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case mes := <-client.messageQueue:
			if err := client.stream.Send(mes); err != nil {
				return err
			}
		}
	}
}

func (client *CVMSClient) sendMessage(mes *cvms.ClientStreamMessage) {
	ctx, cancel := context.WithTimeout(context.Background(), sendTimeout)
	defer cancel()

	select {
	case client.messageQueue <- mes:
	case <-ctx.Done():
		client.logger.Warn("Failed to send message: timeout exceeded")
	}
}

type runRequestManager struct {
	requests map[string]*runRequest
	mu       sync.Mutex
}

type runRequest struct {
	buffer    []byte
	lastChunk time.Time
	timer     *time.Timer
}

func newRunRequestManager() *runRequestManager {
	return &runRequestManager{
		requests: make(map[string]*runRequest),
	}
}

func (m *runRequestManager) addChunk(id string, chunk []byte, isLast bool) ([]byte, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	req, exists := m.requests[id]
	if !exists {
		req = &runRequest{
			buffer:    make([]byte, 0),
			lastChunk: time.Now(),
			timer:     time.AfterFunc(runReqTimeout, func() { m.timeoutRequest(id) }),
		}
		m.requests[id] = req
	}

	req.buffer = append(req.buffer, chunk...)
	req.lastChunk = time.Now()
	req.timer.Reset(runReqTimeout)

	if isLast {
		delete(m.requests, id)
		req.timer.Stop()
		return req.buffer, true
	}

	return nil, false
}

func (m *runRequestManager) timeoutRequest(id string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.requests, id)
	// Log timeout or handle it as needed
}
