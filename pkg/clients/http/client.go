// Copyright (c) Ultraviolet
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"net/http"
	"time"

	"github.com/ultravioletrs/cocos/pkg/clients"
)

type Client interface {
	Transport() *http.Transport
	Secure() string
	Timeout() time.Duration
}

type client struct {
	transport *http.Transport
	cfg       clients.ClientConfiguration
	security  clients.Security
}

var _ Client = (*client)(nil)

func NewClient(cfg clients.ClientConfiguration) (Client, error) {
	transport, security, err := createTransport(cfg)
	if err != nil {
		return nil, err
	}

	return &client{
		transport: transport,
		cfg:       cfg,
		security:  security,
	}, nil
}

func (c *client) Transport() *http.Transport {
	return c.transport
}

func (c *client) Secure() string {
	return c.security.String()
}

func (c *client) Timeout() time.Duration {
	return c.cfg.Config().Timeout
}

func createTransport(cfg clients.ClientConfiguration) (*http.Transport, clients.Security, error) {
	transport := &http.Transport{
		MaxIdleConns:        100,
		IdleConnTimeout:     90 * time.Second,
		TLSHandshakeTimeout: 10 * time.Second,
	}

	security := clients.WithoutTLS

	if agcfg, ok := cfg.(*clients.AttestedClientConfig); ok && agcfg.AttestedTLS {
		result, err := clients.LoadATLSConfig(*agcfg)
		if err != nil {
			return nil, security, err
		}

		transport.TLSClientConfig = result.Config
		security = result.Security
	} else {
		conf := cfg.Config()

		result, err := clients.LoadBasicTLSConfig(conf.ServerCAFile, conf.ClientCert, conf.ClientKey)
		if err != nil {
			return nil, security, err
		}

		if result.Security != clients.WithoutTLS {
			transport.TLSClientConfig = result.Config
		}

		security = result.Security
	}

	return transport, security, nil
}
