//
//Copyright The containerd Authors.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: github.com/containerd/ttrpc/integration/streaming/test.proto

package streaming

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EchoPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Seq uint32 `protobuf:"varint,1,opt,name=seq,proto3" json:"seq,omitempty"`
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *EchoPayload) Reset() {
	*x = EchoPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_containerd_ttrpc_integration_streaming_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EchoPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoPayload) ProtoMessage() {}

func (x *EchoPayload) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_containerd_ttrpc_integration_streaming_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoPayload.ProtoReflect.Descriptor instead.
func (*EchoPayload) Descriptor() ([]byte, []int) {
	return file_github_com_containerd_ttrpc_integration_streaming_test_proto_rawDescGZIP(), []int{0}
}

func (x *EchoPayload) GetSeq() uint32 {
	if x != nil {
		return x.Seq
	}
	return 0
}

func (x *EchoPayload) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type Part struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Add int32 `protobuf:"varint,1,opt,name=add,proto3" json:"add,omitempty"`
}

func (x *Part) Reset() {
	*x = Part{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_containerd_ttrpc_integration_streaming_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Part) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Part) ProtoMessage() {}

func (x *Part) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_containerd_ttrpc_integration_streaming_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Part.ProtoReflect.Descriptor instead.
func (*Part) Descriptor() ([]byte, []int) {
	return file_github_com_containerd_ttrpc_integration_streaming_test_proto_rawDescGZIP(), []int{1}
}

func (x *Part) GetAdd() int32 {
	if x != nil {
		return x.Add
	}
	return 0
}

type Sum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sum int32 `protobuf:"varint,1,opt,name=sum,proto3" json:"sum,omitempty"`
	Num int32 `protobuf:"varint,2,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *Sum) Reset() {
	*x = Sum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_containerd_ttrpc_integration_streaming_test_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sum) ProtoMessage() {}

func (x *Sum) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_containerd_ttrpc_integration_streaming_test_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sum.ProtoReflect.Descriptor instead.
func (*Sum) Descriptor() ([]byte, []int) {
	return file_github_com_containerd_ttrpc_integration_streaming_test_proto_rawDescGZIP(), []int{2}
}

func (x *Sum) GetSum() int32 {
	if x != nil {
		return x.Sum
	}
	return 0
}

func (x *Sum) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

var File_github_com_containerd_ttrpc_integration_streaming_test_proto protoreflect.FileDescriptor

var file_github_com_containerd_ttrpc_integration_streaming_test_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x64, 0x2f, 0x74, 0x74, 0x72, 0x70, 0x63, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x69, 0x6e, 0x67, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b,
	0x74, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x31, 0x0a, 0x0b, 0x45, 0x63, 0x68, 0x6f,
	0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x71, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x73, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x18, 0x0a, 0x04, 0x50,
	0x61, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x64, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x03, 0x61, 0x64, 0x64, 0x22, 0x29, 0x0a, 0x03, 0x53, 0x75, 0x6d, 0x12, 0x10, 0x0a, 0x03,
	0x73, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x73, 0x75, 0x6d, 0x12, 0x10,
	0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6e, 0x75, 0x6d,
	0x32, 0xfa, 0x04, 0x0a, 0x09, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x12, 0x5a,
	0x0a, 0x04, 0x45, 0x63, 0x68, 0x6f, 0x12, 0x28, 0x2e, 0x74, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x69,
	0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x1a, 0x28, 0x2e, 0x74, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x45,
	0x63, 0x68, 0x6f, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x64, 0x0a, 0x0a, 0x45, 0x63,
	0x68, 0x6f, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x28, 0x2e, 0x74, 0x74, 0x72, 0x70, 0x63,
	0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x50, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x1a, 0x28, 0x2e, 0x74, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67,
	0x2e, 0x45, 0x63, 0x68, 0x6f, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x28, 0x01, 0x30, 0x01,
	0x12, 0x52, 0x0a, 0x09, 0x53, 0x75, 0x6d, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x21, 0x2e,
	0x74, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x61, 0x72, 0x74,
	0x1a, 0x20, 0x2e, 0x74, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x53,
	0x75, 0x6d, 0x28, 0x01, 0x12, 0x55, 0x0a, 0x0c, 0x44, 0x69, 0x76, 0x69, 0x64, 0x65, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x12, 0x20, 0x2e, 0x74, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x69, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69,
	0x6e, 0x67, 0x2e, 0x53, 0x75, 0x6d, 0x1a, 0x21, 0x2e, 0x74, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x69,
	0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x30, 0x01, 0x12, 0x4e, 0x0a, 0x08, 0x45,
	0x63, 0x68, 0x6f, 0x4e, 0x75, 0x6c, 0x6c, 0x12, 0x28, 0x2e, 0x74, 0x74, 0x72, 0x70, 0x63, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x28, 0x01, 0x12, 0x56, 0x0a, 0x0e, 0x45,
	0x63, 0x68, 0x6f, 0x4e, 0x75, 0x6c, 0x6c, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x28, 0x2e,
	0x74, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x45, 0x63, 0x68, 0x6f,
	0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x28,
	0x01, 0x30, 0x01, 0x12, 0x58, 0x0a, 0x12, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x50, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x28, 0x2e, 0x74, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2e,
	0x45, 0x63, 0x68, 0x6f, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x30, 0x01, 0x42, 0x3d, 0x5a,
	0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x64, 0x2f, 0x74, 0x74, 0x72, 0x70, 0x63, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69,
	0x6e, 0x67, 0x3b, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_containerd_ttrpc_integration_streaming_test_proto_rawDescOnce sync.Once
	file_github_com_containerd_ttrpc_integration_streaming_test_proto_rawDescData = file_github_com_containerd_ttrpc_integration_streaming_test_proto_rawDesc
)

func file_github_com_containerd_ttrpc_integration_streaming_test_proto_rawDescGZIP() []byte {
	file_github_com_containerd_ttrpc_integration_streaming_test_proto_rawDescOnce.Do(func() {
		file_github_com_containerd_ttrpc_integration_streaming_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_containerd_ttrpc_integration_streaming_test_proto_rawDescData)
	})
	return file_github_com_containerd_ttrpc_integration_streaming_test_proto_rawDescData
}

var file_github_com_containerd_ttrpc_integration_streaming_test_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_github_com_containerd_ttrpc_integration_streaming_test_proto_goTypes = []interface{}{
	(*EchoPayload)(nil),   // 0: ttrpc.integration.streaming.EchoPayload
	(*Part)(nil),          // 1: ttrpc.integration.streaming.Part
	(*Sum)(nil),           // 2: ttrpc.integration.streaming.Sum
	(*emptypb.Empty)(nil), // 3: google.protobuf.Empty
}
var file_github_com_containerd_ttrpc_integration_streaming_test_proto_depIdxs = []int32{
	0, // 0: ttrpc.integration.streaming.Streaming.Echo:input_type -> ttrpc.integration.streaming.EchoPayload
	0, // 1: ttrpc.integration.streaming.Streaming.EchoStream:input_type -> ttrpc.integration.streaming.EchoPayload
	1, // 2: ttrpc.integration.streaming.Streaming.SumStream:input_type -> ttrpc.integration.streaming.Part
	2, // 3: ttrpc.integration.streaming.Streaming.DivideStream:input_type -> ttrpc.integration.streaming.Sum
	0, // 4: ttrpc.integration.streaming.Streaming.EchoNull:input_type -> ttrpc.integration.streaming.EchoPayload
	0, // 5: ttrpc.integration.streaming.Streaming.EchoNullStream:input_type -> ttrpc.integration.streaming.EchoPayload
	3, // 6: ttrpc.integration.streaming.Streaming.EmptyPayloadStream:input_type -> google.protobuf.Empty
	0, // 7: ttrpc.integration.streaming.Streaming.Echo:output_type -> ttrpc.integration.streaming.EchoPayload
	0, // 8: ttrpc.integration.streaming.Streaming.EchoStream:output_type -> ttrpc.integration.streaming.EchoPayload
	2, // 9: ttrpc.integration.streaming.Streaming.SumStream:output_type -> ttrpc.integration.streaming.Sum
	1, // 10: ttrpc.integration.streaming.Streaming.DivideStream:output_type -> ttrpc.integration.streaming.Part
	3, // 11: ttrpc.integration.streaming.Streaming.EchoNull:output_type -> google.protobuf.Empty
	3, // 12: ttrpc.integration.streaming.Streaming.EchoNullStream:output_type -> google.protobuf.Empty
	0, // 13: ttrpc.integration.streaming.Streaming.EmptyPayloadStream:output_type -> ttrpc.integration.streaming.EchoPayload
	7, // [7:14] is the sub-list for method output_type
	0, // [0:7] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_github_com_containerd_ttrpc_integration_streaming_test_proto_init() }
func file_github_com_containerd_ttrpc_integration_streaming_test_proto_init() {
	if File_github_com_containerd_ttrpc_integration_streaming_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_containerd_ttrpc_integration_streaming_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EchoPayload); i {
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
		file_github_com_containerd_ttrpc_integration_streaming_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Part); i {
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
		file_github_com_containerd_ttrpc_integration_streaming_test_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sum); i {
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
			RawDescriptor: file_github_com_containerd_ttrpc_integration_streaming_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_github_com_containerd_ttrpc_integration_streaming_test_proto_goTypes,
		DependencyIndexes: file_github_com_containerd_ttrpc_integration_streaming_test_proto_depIdxs,
		MessageInfos:      file_github_com_containerd_ttrpc_integration_streaming_test_proto_msgTypes,
	}.Build()
	File_github_com_containerd_ttrpc_integration_streaming_test_proto = out.File
	file_github_com_containerd_ttrpc_integration_streaming_test_proto_rawDesc = nil
	file_github_com_containerd_ttrpc_integration_streaming_test_proto_goTypes = nil
	file_github_com_containerd_ttrpc_integration_streaming_test_proto_depIdxs = nil
}