// Copyright (c) Ultraviolet
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: agent/agent.proto

package agent

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AlgoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Algorithm []byte `protobuf:"bytes,1,opt,name=algorithm,proto3" json:"algorithm,omitempty"`
	Provider  string `protobuf:"bytes,2,opt,name=provider,proto3" json:"provider,omitempty"`
	Id        string `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AlgoRequest) Reset() {
	*x = AlgoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_agent_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlgoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlgoRequest) ProtoMessage() {}

func (x *AlgoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_agent_agent_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlgoRequest.ProtoReflect.Descriptor instead.
func (*AlgoRequest) Descriptor() ([]byte, []int) {
	return file_agent_agent_proto_rawDescGZIP(), []int{0}
}

func (x *AlgoRequest) GetAlgorithm() []byte {
	if x != nil {
		return x.Algorithm
	}
	return nil
}

func (x *AlgoRequest) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *AlgoRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type AlgoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AlgorithmID string `protobuf:"bytes,1,opt,name=algorithmID,proto3" json:"algorithmID,omitempty"`
}

func (x *AlgoResponse) Reset() {
	*x = AlgoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_agent_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlgoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlgoResponse) ProtoMessage() {}

func (x *AlgoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_agent_agent_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlgoResponse.ProtoReflect.Descriptor instead.
func (*AlgoResponse) Descriptor() ([]byte, []int) {
	return file_agent_agent_proto_rawDescGZIP(), []int{1}
}

func (x *AlgoResponse) GetAlgorithmID() string {
	if x != nil {
		return x.AlgorithmID
	}
	return ""
}

type DataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dataset  []byte `protobuf:"bytes,1,opt,name=dataset,proto3" json:"dataset,omitempty"`
	Provider string `protobuf:"bytes,2,opt,name=provider,proto3" json:"provider,omitempty"`
	Id       string `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DataRequest) Reset() {
	*x = DataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_agent_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataRequest) ProtoMessage() {}

func (x *DataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_agent_agent_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataRequest.ProtoReflect.Descriptor instead.
func (*DataRequest) Descriptor() ([]byte, []int) {
	return file_agent_agent_proto_rawDescGZIP(), []int{2}
}

func (x *DataRequest) GetDataset() []byte {
	if x != nil {
		return x.Dataset
	}
	return nil
}

func (x *DataRequest) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *DataRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DatasetID string `protobuf:"bytes,1,opt,name=datasetID,proto3" json:"datasetID,omitempty"`
}

func (x *DataResponse) Reset() {
	*x = DataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_agent_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataResponse) ProtoMessage() {}

func (x *DataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_agent_agent_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataResponse.ProtoReflect.Descriptor instead.
func (*DataResponse) Descriptor() ([]byte, []int) {
	return file_agent_agent_proto_rawDescGZIP(), []int{3}
}

func (x *DataResponse) GetDatasetID() string {
	if x != nil {
		return x.DatasetID
	}
	return ""
}

type ResultRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Consumer string `protobuf:"bytes,1,opt,name=consumer,proto3" json:"consumer,omitempty"`
}

func (x *ResultRequest) Reset() {
	*x = ResultRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_agent_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultRequest) ProtoMessage() {}

func (x *ResultRequest) ProtoReflect() protoreflect.Message {
	mi := &file_agent_agent_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultRequest.ProtoReflect.Descriptor instead.
func (*ResultRequest) Descriptor() ([]byte, []int) {
	return file_agent_agent_proto_rawDescGZIP(), []int{4}
}

func (x *ResultRequest) GetConsumer() string {
	if x != nil {
		return x.Consumer
	}
	return ""
}

type ResultResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	File []byte `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
}

func (x *ResultResponse) Reset() {
	*x = ResultResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_agent_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultResponse) ProtoMessage() {}

func (x *ResultResponse) ProtoReflect() protoreflect.Message {
	mi := &file_agent_agent_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultResponse.ProtoReflect.Descriptor instead.
func (*ResultResponse) Descriptor() ([]byte, []int) {
	return file_agent_agent_proto_rawDescGZIP(), []int{5}
}

func (x *ResultResponse) GetFile() []byte {
	if x != nil {
		return x.File
	}
	return nil
}

type AttestationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReportData []byte `protobuf:"bytes,1,opt,name=report_data,json=reportData,proto3" json:"report_data,omitempty"`
}

func (x *AttestationRequest) Reset() {
	*x = AttestationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_agent_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttestationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttestationRequest) ProtoMessage() {}

func (x *AttestationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_agent_agent_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttestationRequest.ProtoReflect.Descriptor instead.
func (*AttestationRequest) Descriptor() ([]byte, []int) {
	return file_agent_agent_proto_rawDescGZIP(), []int{6}
}

func (x *AttestationRequest) GetReportData() []byte {
	if x != nil {
		return x.ReportData
	}
	return nil
}

type AttestationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	File []byte `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
}

func (x *AttestationResponse) Reset() {
	*x = AttestationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_agent_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttestationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttestationResponse) ProtoMessage() {}

func (x *AttestationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_agent_agent_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttestationResponse.ProtoReflect.Descriptor instead.
func (*AttestationResponse) Descriptor() ([]byte, []int) {
	return file_agent_agent_proto_rawDescGZIP(), []int{7}
}

func (x *AttestationResponse) GetFile() []byte {
	if x != nil {
		return x.File
	}
	return nil
}

var File_agent_agent_proto protoreflect.FileDescriptor

var file_agent_agent_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x22, 0x57, 0x0a, 0x0b, 0x41, 0x6c,
	0x67, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x6c, 0x67,
	0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x61, 0x6c,
	0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x30, 0x0a, 0x0c, 0x41, 0x6c, 0x67, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69,
	0x74, 0x68, 0x6d, 0x49, 0x44, 0x22, 0x53, 0x0a, 0x0b, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2c, 0x0a, 0x0c, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x61,
	0x74, 0x61, 0x73, 0x65, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64,
	0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x49, 0x44, 0x22, 0x2b, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6e,
	0x73, 0x75, 0x6d, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x6e,
	0x73, 0x75, 0x6d, 0x65, 0x72, 0x22, 0x24, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x35, 0x0a, 0x12, 0x41,
	0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x22, 0x29, 0x0a, 0x13, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x69, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x32, 0xf5, 0x01,
	0x0a, 0x0c, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x31,
	0x0a, 0x04, 0x41, 0x6c, 0x67, 0x6f, 0x12, 0x12, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x41,
	0x6c, 0x67, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x2e, 0x41, 0x6c, 0x67, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x31, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x2e, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x14,
	0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a,
	0x0b, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x2e, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e,
	0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_agent_agent_proto_rawDescOnce sync.Once
	file_agent_agent_proto_rawDescData = file_agent_agent_proto_rawDesc
)

func file_agent_agent_proto_rawDescGZIP() []byte {
	file_agent_agent_proto_rawDescOnce.Do(func() {
		file_agent_agent_proto_rawDescData = protoimpl.X.CompressGZIP(file_agent_agent_proto_rawDescData)
	})
	return file_agent_agent_proto_rawDescData
}

var file_agent_agent_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_agent_agent_proto_goTypes = []interface{}{
	(*AlgoRequest)(nil),         // 0: agent.AlgoRequest
	(*AlgoResponse)(nil),        // 1: agent.AlgoResponse
	(*DataRequest)(nil),         // 2: agent.DataRequest
	(*DataResponse)(nil),        // 3: agent.DataResponse
	(*ResultRequest)(nil),       // 4: agent.ResultRequest
	(*ResultResponse)(nil),      // 5: agent.ResultResponse
	(*AttestationRequest)(nil),  // 6: agent.AttestationRequest
	(*AttestationResponse)(nil), // 7: agent.AttestationResponse
}
var file_agent_agent_proto_depIdxs = []int32{
	0, // 0: agent.AgentService.Algo:input_type -> agent.AlgoRequest
	2, // 1: agent.AgentService.Data:input_type -> agent.DataRequest
	4, // 2: agent.AgentService.Result:input_type -> agent.ResultRequest
	6, // 3: agent.AgentService.Attestation:input_type -> agent.AttestationRequest
	1, // 4: agent.AgentService.Algo:output_type -> agent.AlgoResponse
	3, // 5: agent.AgentService.Data:output_type -> agent.DataResponse
	5, // 6: agent.AgentService.Result:output_type -> agent.ResultResponse
	7, // 7: agent.AgentService.Attestation:output_type -> agent.AttestationResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_agent_agent_proto_init() }
func file_agent_agent_proto_init() {
	if File_agent_agent_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_agent_agent_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlgoRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_agent_agent_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlgoResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_agent_agent_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_agent_agent_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_agent_agent_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResultRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_agent_agent_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResultResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_agent_agent_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttestationRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_agent_agent_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttestationResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_agent_agent_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_agent_agent_proto_goTypes,
		DependencyIndexes: file_agent_agent_proto_depIdxs,
		MessageInfos:      file_agent_agent_proto_msgTypes,
	}.Build()
	File_agent_agent_proto = out.File
	file_agent_agent_proto_rawDesc = nil
	file_agent_agent_proto_goTypes = nil
	file_agent_agent_proto_depIdxs = nil
}
