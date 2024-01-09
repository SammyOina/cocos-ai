// Copyright (c) Ultraviolet
// SPDX-License-Identifier: Apache-2.0
package grpc

type algoRes struct {
	AlgorithmID string `json:"algorithmId,omitempty"`
}

type dataRes struct {
	DatasetID string `json:"datasetId,omitempty"`
}

type resultRes struct {
	File []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

type attestationRes struct {
	File []byte
}
