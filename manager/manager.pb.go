// Copyright (c) Ultraviolet
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: manager/manager.proto

package manager

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

type RunRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Computation *Computation `protobuf:"bytes,1,opt,name=computation,proto3" json:"computation,omitempty"`
}

func (x *RunRequest) Reset() {
	*x = RunRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manager_manager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunRequest) ProtoMessage() {}

func (x *RunRequest) ProtoReflect() protoreflect.Message {
	mi := &file_manager_manager_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunRequest.ProtoReflect.Descriptor instead.
func (*RunRequest) Descriptor() ([]byte, []int) {
	return file_manager_manager_proto_rawDescGZIP(), []int{0}
}

func (x *RunRequest) GetComputation() *Computation {
	if x != nil {
		return x.Computation
	}
	return nil
}

type Computation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name            string       `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description     string       `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Datasets        []*Dataset   `protobuf:"bytes,4,rep,name=datasets,proto3" json:"datasets,omitempty"`
	Algorithms      []*Algorithm `protobuf:"bytes,5,rep,name=algorithms,proto3" json:"algorithms,omitempty"`
	ResultConsumers []string     `protobuf:"bytes,6,rep,name=result_consumers,json=resultConsumers,proto3" json:"result_consumers,omitempty"`
	AgentConfig     *AgentConfig `protobuf:"bytes,7,opt,name=agent_config,json=agentConfig,proto3" json:"agent_config,omitempty"`
}

func (x *Computation) Reset() {
	*x = Computation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manager_manager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Computation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Computation) ProtoMessage() {}

func (x *Computation) ProtoReflect() protoreflect.Message {
	mi := &file_manager_manager_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Computation.ProtoReflect.Descriptor instead.
func (*Computation) Descriptor() ([]byte, []int) {
	return file_manager_manager_proto_rawDescGZIP(), []int{1}
}

func (x *Computation) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Computation) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Computation) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Computation) GetDatasets() []*Dataset {
	if x != nil {
		return x.Datasets
	}
	return nil
}

func (x *Computation) GetAlgorithms() []*Algorithm {
	if x != nil {
		return x.Algorithms
	}
	return nil
}

func (x *Computation) GetResultConsumers() []string {
	if x != nil {
		return x.ResultConsumers
	}
	return nil
}

func (x *Computation) GetAgentConfig() *AgentConfig {
	if x != nil {
		return x.AgentConfig
	}
	return nil
}

type Dataset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Provider string `protobuf:"bytes,1,opt,name=provider,proto3" json:"provider,omitempty"`
	Id       string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Dataset) Reset() {
	*x = Dataset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manager_manager_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dataset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dataset) ProtoMessage() {}

func (x *Dataset) ProtoReflect() protoreflect.Message {
	mi := &file_manager_manager_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dataset.ProtoReflect.Descriptor instead.
func (*Dataset) Descriptor() ([]byte, []int) {
	return file_manager_manager_proto_rawDescGZIP(), []int{2}
}

func (x *Dataset) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *Dataset) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Algorithm struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Provider string `protobuf:"bytes,1,opt,name=provider,proto3" json:"provider,omitempty"`
	Id       string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Algorithm) Reset() {
	*x = Algorithm{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manager_manager_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Algorithm) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Algorithm) ProtoMessage() {}

func (x *Algorithm) ProtoReflect() protoreflect.Message {
	mi := &file_manager_manager_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Algorithm.ProtoReflect.Descriptor instead.
func (*Algorithm) Descriptor() ([]byte, []int) {
	return file_manager_manager_proto_rawDescGZIP(), []int{3}
}

func (x *Algorithm) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *Algorithm) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type RunResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AgentAddress string `protobuf:"bytes,1,opt,name=agent_address,json=agentAddress,proto3" json:"agent_address,omitempty"`
}

func (x *RunResponse) Reset() {
	*x = RunResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manager_manager_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunResponse) ProtoMessage() {}

func (x *RunResponse) ProtoReflect() protoreflect.Message {
	mi := &file_manager_manager_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunResponse.ProtoReflect.Descriptor instead.
func (*RunResponse) Descriptor() ([]byte, []int) {
	return file_manager_manager_proto_rawDescGZIP(), []int{4}
}

func (x *RunResponse) GetAgentAddress() string {
	if x != nil {
		return x.AgentAddress
	}
	return ""
}

type AgentConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Port     string `protobuf:"bytes,1,opt,name=port,proto3" json:"port,omitempty"`
	Host     string `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	CertFile string `protobuf:"bytes,3,opt,name=cert_file,json=certFile,proto3" json:"cert_file,omitempty"`
	KeyFile  string `protobuf:"bytes,4,opt,name=key_file,json=keyFile,proto3" json:"key_file,omitempty"`
	LogLevel string `protobuf:"bytes,5,opt,name=log_level,json=logLevel,proto3" json:"log_level,omitempty"`
}

func (x *AgentConfig) Reset() {
	*x = AgentConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manager_manager_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentConfig) ProtoMessage() {}

func (x *AgentConfig) ProtoReflect() protoreflect.Message {
	mi := &file_manager_manager_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentConfig.ProtoReflect.Descriptor instead.
func (*AgentConfig) Descriptor() ([]byte, []int) {
	return file_manager_manager_proto_rawDescGZIP(), []int{5}
}

func (x *AgentConfig) GetPort() string {
	if x != nil {
		return x.Port
	}
	return ""
}

func (x *AgentConfig) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *AgentConfig) GetCertFile() string {
	if x != nil {
		return x.CertFile
	}
	return ""
}

func (x *AgentConfig) GetKeyFile() string {
	if x != nil {
		return x.KeyFile
	}
	return ""
}

func (x *AgentConfig) GetLogLevel() string {
	if x != nil {
		return x.LogLevel
	}
	return ""
}

var File_manager_manager_proto protoreflect.FileDescriptor

var file_manager_manager_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x22, 0x44, 0x0a, 0x0a, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36,
	0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x43, 0x6f,
	0x6d, 0x70, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x75,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x99, 0x02, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x70, 0x75,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x08,
	0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74,
	0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x12, 0x32, 0x0a, 0x0a, 0x61, 0x6c,
	0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74,
	0x68, 0x6d, 0x52, 0x0a, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x73, 0x12, 0x29,
	0x0a, 0x10, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65,
	0x72, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x73, 0x12, 0x37, 0x0a, 0x0c, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0b, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x22, 0x35, 0x0a, 0x07, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x37, 0x0a, 0x09, 0x41, 0x6c, 0x67,
	0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x32, 0x0a, 0x0b, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x8a, 0x01, 0x0a, 0x0b, 0x41, 0x67, 0x65, 0x6e, 0x74,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f,
	0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x63, 0x65, 0x72, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x65, 0x72, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6b,
	0x65, 0x79, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b,
	0x65, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x6f, 0x67, 0x5f, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x67, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x32, 0x44, 0x0a, 0x0e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x03, 0x52, 0x75, 0x6e, 0x12, 0x13, 0x2e, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x14, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x52, 0x75, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_manager_manager_proto_rawDescOnce sync.Once
	file_manager_manager_proto_rawDescData = file_manager_manager_proto_rawDesc
)

func file_manager_manager_proto_rawDescGZIP() []byte {
	file_manager_manager_proto_rawDescOnce.Do(func() {
		file_manager_manager_proto_rawDescData = protoimpl.X.CompressGZIP(file_manager_manager_proto_rawDescData)
	})
	return file_manager_manager_proto_rawDescData
}

var file_manager_manager_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_manager_manager_proto_goTypes = []interface{}{
	(*RunRequest)(nil),  // 0: manager.RunRequest
	(*Computation)(nil), // 1: manager.Computation
	(*Dataset)(nil),     // 2: manager.Dataset
	(*Algorithm)(nil),   // 3: manager.Algorithm
	(*RunResponse)(nil), // 4: manager.RunResponse
	(*AgentConfig)(nil), // 5: manager.AgentConfig
}
var file_manager_manager_proto_depIdxs = []int32{
	1, // 0: manager.RunRequest.computation:type_name -> manager.Computation
	2, // 1: manager.Computation.datasets:type_name -> manager.Dataset
	3, // 2: manager.Computation.algorithms:type_name -> manager.Algorithm
	5, // 3: manager.Computation.agent_config:type_name -> manager.AgentConfig
	0, // 4: manager.ManagerService.Run:input_type -> manager.RunRequest
	4, // 5: manager.ManagerService.Run:output_type -> manager.RunResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_manager_manager_proto_init() }
func file_manager_manager_proto_init() {
	if File_manager_manager_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_manager_manager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunRequest); i {
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
		file_manager_manager_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Computation); i {
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
		file_manager_manager_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dataset); i {
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
		file_manager_manager_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Algorithm); i {
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
		file_manager_manager_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunResponse); i {
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
		file_manager_manager_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentConfig); i {
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
			RawDescriptor: file_manager_manager_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_manager_manager_proto_goTypes,
		DependencyIndexes: file_manager_manager_proto_depIdxs,
		MessageInfos:      file_manager_manager_proto_msgTypes,
	}.Build()
	File_manager_manager_proto = out.File
	file_manager_manager_proto_rawDesc = nil
	file_manager_manager_proto_goTypes = nil
	file_manager_manager_proto_depIdxs = nil
}
