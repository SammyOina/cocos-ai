// Copyright (c) Ultraviolet
// SPDX-License-Identifier: Apache-2.0
package cli

import (
	"context"

	"github.com/ultravioletrs/cocos/pkg/clients/grpc"
	"github.com/ultravioletrs/cocos/pkg/clients/grpc/agent"
	"github.com/ultravioletrs/cocos/pkg/sdk"
)

var Verbose bool

type CLI struct {
	agentSDK   sdk.SDK
	config     grpc.Config
	client     grpc.Client
	connectErr error
}

func New(config grpc.Config) *CLI {
	return &CLI{
		config: config,
	}
}

func (c *CLI) InitializeSDK() error {
	agentGRPCClient, agentClient, err := agent.NewAgentClient(context.Background(), c.config)
	if err != nil {
		c.connectErr = err
		return err
	}
	c.client = agentGRPCClient

	c.agentSDK = sdk.NewAgentSDK(agentClient)
	return nil
}

func (c *CLI) Close() {
	c.client.Close()
}
