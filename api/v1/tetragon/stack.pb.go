// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Hubble

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.3
// source: tetragon/stack.proto

package tetragon

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

type StackAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address uint64 `protobuf:"varint,1,opt,name=address,proto3" json:"address,omitempty"`
	Symbol  string `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol,omitempty"`
}

func (x *StackAddress) Reset() {
	*x = StackAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tetragon_stack_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StackAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StackAddress) ProtoMessage() {}

func (x *StackAddress) ProtoReflect() protoreflect.Message {
	mi := &file_tetragon_stack_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StackAddress.ProtoReflect.Descriptor instead.
func (*StackAddress) Descriptor() ([]byte, []int) {
	return file_tetragon_stack_proto_rawDescGZIP(), []int{0}
}

func (x *StackAddress) GetAddress() uint64 {
	if x != nil {
		return x.Address
	}
	return 0
}

func (x *StackAddress) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

type StackTrace struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addresses []*StackAddress `protobuf:"bytes,1,rep,name=addresses,proto3" json:"addresses,omitempty"`
}

func (x *StackTrace) Reset() {
	*x = StackTrace{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tetragon_stack_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StackTrace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StackTrace) ProtoMessage() {}

func (x *StackTrace) ProtoReflect() protoreflect.Message {
	mi := &file_tetragon_stack_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StackTrace.ProtoReflect.Descriptor instead.
func (*StackTrace) Descriptor() ([]byte, []int) {
	return file_tetragon_stack_proto_rawDescGZIP(), []int{1}
}

func (x *StackTrace) GetAddresses() []*StackAddress {
	if x != nil {
		return x.Addresses
	}
	return nil
}

type StackTraceLabel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Count uint64 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *StackTraceLabel) Reset() {
	*x = StackTraceLabel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tetragon_stack_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StackTraceLabel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StackTraceLabel) ProtoMessage() {}

func (x *StackTraceLabel) ProtoReflect() protoreflect.Message {
	mi := &file_tetragon_stack_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StackTraceLabel.ProtoReflect.Descriptor instead.
func (*StackTraceLabel) Descriptor() ([]byte, []int) {
	return file_tetragon_stack_proto_rawDescGZIP(), []int{2}
}

func (x *StackTraceLabel) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *StackTraceLabel) GetCount() uint64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type StackTraceNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address  *StackAddress      `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Count    uint64             `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Labels   []*StackTraceLabel `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty"`
	Children []*StackTraceNode  `protobuf:"bytes,4,rep,name=children,proto3" json:"children,omitempty"`
}

func (x *StackTraceNode) Reset() {
	*x = StackTraceNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tetragon_stack_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StackTraceNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StackTraceNode) ProtoMessage() {}

func (x *StackTraceNode) ProtoReflect() protoreflect.Message {
	mi := &file_tetragon_stack_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StackTraceNode.ProtoReflect.Descriptor instead.
func (*StackTraceNode) Descriptor() ([]byte, []int) {
	return file_tetragon_stack_proto_rawDescGZIP(), []int{3}
}

func (x *StackTraceNode) GetAddress() *StackAddress {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *StackTraceNode) GetCount() uint64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *StackTraceNode) GetLabels() []*StackTraceLabel {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *StackTraceNode) GetChildren() []*StackTraceNode {
	if x != nil {
		return x.Children
	}
	return nil
}

var File_tetragon_stack_proto protoreflect.FileDescriptor

var file_tetragon_stack_proto_rawDesc = []byte{
	0x0a, 0x14, 0x74, 0x65, 0x74, 0x72, 0x61, 0x67, 0x6f, 0x6e, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x74, 0x65, 0x74, 0x72, 0x61, 0x67, 0x6f, 0x6e,
	0x22, 0x40, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79,
	0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62,
	0x6f, 0x6c, 0x22, 0x42, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x54, 0x72, 0x61, 0x63, 0x65,
	0x12, 0x34, 0x0a, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x65, 0x74, 0x72, 0x61, 0x67, 0x6f, 0x6e, 0x2e, 0x53,
	0x74, 0x61, 0x63, 0x6b, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x09, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x22, 0x39, 0x0a, 0x0f, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x54,
	0x72, 0x61, 0x63, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x22, 0xc1, 0x01, 0x0a, 0x0e, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x54, 0x72, 0x61, 0x63, 0x65,
	0x4e, 0x6f, 0x64, 0x65, 0x12, 0x30, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x65, 0x74, 0x72, 0x61, 0x67, 0x6f, 0x6e,
	0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x31, 0x0a, 0x06,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x74,
	0x65, 0x74, 0x72, 0x61, 0x67, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x54, 0x72, 0x61,
	0x63, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12,
	0x34, 0x0a, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x74, 0x65, 0x74, 0x72, 0x61, 0x67, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x61,
	0x63, 0x6b, 0x54, 0x72, 0x61, 0x63, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x08, 0x63, 0x68, 0x69,
	0x6c, 0x64, 0x72, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tetragon_stack_proto_rawDescOnce sync.Once
	file_tetragon_stack_proto_rawDescData = file_tetragon_stack_proto_rawDesc
)

func file_tetragon_stack_proto_rawDescGZIP() []byte {
	file_tetragon_stack_proto_rawDescOnce.Do(func() {
		file_tetragon_stack_proto_rawDescData = protoimpl.X.CompressGZIP(file_tetragon_stack_proto_rawDescData)
	})
	return file_tetragon_stack_proto_rawDescData
}

var file_tetragon_stack_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_tetragon_stack_proto_goTypes = []interface{}{
	(*StackAddress)(nil),    // 0: tetragon.StackAddress
	(*StackTrace)(nil),      // 1: tetragon.StackTrace
	(*StackTraceLabel)(nil), // 2: tetragon.StackTraceLabel
	(*StackTraceNode)(nil),  // 3: tetragon.StackTraceNode
}
var file_tetragon_stack_proto_depIdxs = []int32{
	0, // 0: tetragon.StackTrace.addresses:type_name -> tetragon.StackAddress
	0, // 1: tetragon.StackTraceNode.address:type_name -> tetragon.StackAddress
	2, // 2: tetragon.StackTraceNode.labels:type_name -> tetragon.StackTraceLabel
	3, // 3: tetragon.StackTraceNode.children:type_name -> tetragon.StackTraceNode
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_tetragon_stack_proto_init() }
func file_tetragon_stack_proto_init() {
	if File_tetragon_stack_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tetragon_stack_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StackAddress); i {
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
		file_tetragon_stack_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StackTrace); i {
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
		file_tetragon_stack_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StackTraceLabel); i {
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
		file_tetragon_stack_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StackTraceNode); i {
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
			RawDescriptor: file_tetragon_stack_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tetragon_stack_proto_goTypes,
		DependencyIndexes: file_tetragon_stack_proto_depIdxs,
		MessageInfos:      file_tetragon_stack_proto_msgTypes,
	}.Build()
	File_tetragon_stack_proto = out.File
	file_tetragon_stack_proto_rawDesc = nil
	file_tetragon_stack_proto_goTypes = nil
	file_tetragon_stack_proto_depIdxs = nil
}
