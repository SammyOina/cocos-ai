// Copyright (c) Ultraviolet
// SPDX-License-Identifier: Apache-2.0
package agent

import (
	"context"
	"encoding/json"
	"fmt"

	"google.golang.org/grpc/metadata"
)

var _ fmt.Stringer = (*Datasets)(nil)

type AgentConfig struct {
	LogLevel     string `json:"log_level,omitempty"`
	Host         string `json:"host,omitempty"`
	Port         string `json:"port,omitempty"`
	CertFile     string `json:"cert_file,omitempty"`
	KeyFile      string `json:"server_key,omitempty"`
	ServerCAFile string `json:"server_ca_file,omitempty"`
	ClientCAFile string `json:"client_ca_file,omitempty"`
	AttestedTls  bool   `json:"attested_tls,omitempty"`
}

type Computation struct {
	ID              string           `json:"id,omitempty"`
	Name            string           `json:"name,omitempty"`
	Description     string           `json:"description,omitempty"`
	Datasets        Datasets         `json:"datasets,omitempty"`
	Algorithm       Algorithm        `json:"algorithm,omitempty"`
	ResultConsumers []ResultConsumer `json:"result_consumers,omitempty"`
	AgentConfig     AgentConfig      `json:"agent_config,omitempty"`
}

type ResultConsumer struct {
	UserKey []byte `json:"user_key,omitempty"`
}

func (d *Datasets) String() string {
	dat, err := json.Marshal(d)
	if err != nil {
		return ""
	}
	return string(dat)
}

type Dataset struct {
	Dataset  []byte   `json:"-"`
	Hash     [32]byte `json:"hash,omitempty"`
	UserKey  []byte   `json:"user_key,omitempty"`
	Filename string   `json:"filename,omitempty"`
}

type Datasets []Dataset

type Algorithm struct {
	Algorithm    []byte   `json:"-"`
	Hash         [32]byte `json:"hash,omitempty"`
	UserKey      []byte   `json:"user_key,omitempty"`
	Requirements []byte   `json:"-"`
}

type ManifestIndexKey struct{}

func IndexToContext(ctx context.Context, index int) context.Context {
	return context.WithValue(ctx, ManifestIndexKey{}, index)
}

func IndexFromContext(ctx context.Context) (int, bool) {
	index, ok := ctx.Value(ManifestIndexKey{}).(int)
	return index, ok
}

const DecompressKey = "decompress"

func DecompressFromContext(ctx context.Context) bool {
	vals := metadata.ValueFromIncomingContext(ctx, DecompressKey)
	if len(vals) == 0 {
		return false
	}

	return vals[0] == "true"
}

func DecompressToContext(ctx context.Context, decompress bool) context.Context {
	return metadata.AppendToOutgoingContext(ctx, DecompressKey, fmt.Sprintf("%t", decompress))
}

//go:generate mockery --name=AgentService_AlgoClient --output=mocks --filename agent_grpc_algo.go --quiet --note "Copyright (c) Ultraviolet \n // SPDX-License-Identifier: Apache-2.0"
//go:generate mockery --name=AgentService_DataClient --output=mocks --filename agent_grpc_data.go --quiet --note "Copyright (c) Ultraviolet \n // SPDX-License-Identifier: Apache-2.0"
