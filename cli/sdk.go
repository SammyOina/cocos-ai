// Copyright (c) Ultraviolet
// SPDX-License-Identifier: Apache-2.0
package cli

import (
	"context"
	"time"

	"github.com/spf13/cobra"
	"github.com/ultravioletrs/cocos/manager"
	"github.com/ultravioletrs/cocos/pkg/attestation/cmdconfig"
	"github.com/ultravioletrs/cocos/pkg/clients"
	"github.com/ultravioletrs/cocos/pkg/clients/grpc"
	"github.com/ultravioletrs/cocos/pkg/clients/grpc/agent"
	managergrpc "github.com/ultravioletrs/cocos/pkg/clients/grpc/manager"
	"github.com/ultravioletrs/cocos/pkg/sdk"
)

type AgentVMConfig struct {
	CVMServerURL string
	CVMServerCA  string
	CVMClientKey string
	CVMClientCrt string
	CVMCaURL     string
	LogLevel     string
	Ttl          time.Duration
}

type AWSConfig struct {
	AccessKeyID     string
	SecretAccessKey string
	EndpointURL     string
	Region          string
}

type AttestationConfig struct {
	KbsParams string
}

type CLI struct {
	agentSDK      sdk.SDK
	agentConfig   clients.AttestedClientConfig
	managerConfig clients.StandardClientConfig
	client        grpc.Client
	managerClient manager.ManagerServiceClient
	connectErr    error
	measurement   cmdconfig.MeasurementProvider
	Verbose       bool
	IsManifest    bool
	ToBase64      bool
	KeyType       string
	AgentVM       AgentVMConfig
	AWS           AWSConfig
	Attestation   AttestationConfig
}

func New(agentConfig clients.AttestedClientConfig, managerConfig clients.StandardClientConfig, measurement cmdconfig.MeasurementProvider) *CLI {
	return &CLI{
		agentConfig:   agentConfig,
		managerConfig: managerConfig,
		measurement:   measurement,
	}
}

func (c *CLI) InitializeAgentSDK(cmd *cobra.Command) error {
	agentGRPCClient, agentClient, err := agent.NewAgentClient(context.Background(), c.agentConfig)
	if err != nil {
		c.connectErr = err
		return err
	}
	cmd.Println("🔗 Connected to agent ", agentGRPCClient.Secure())
	c.client = agentGRPCClient

	c.agentSDK = sdk.NewAgentSDK(agentClient)
	return nil
}

func (c *CLI) InitializeManagerClient(cmd *cobra.Command) error {
	managerGRPCClient, managerClient, err := managergrpc.NewManagerClient(c.managerConfig)
	if err != nil {
		c.connectErr = err
		return err
	}

	cmd.Println("🔗 Connected to manager using ", managerGRPCClient.Secure())
	c.client = managerGRPCClient

	c.managerClient = managerClient
	return nil
}

func (c *CLI) Close() {
	if c.client != nil {
		c.client.Close()
	}
}
