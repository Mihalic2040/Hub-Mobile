// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: protocols/bridge.proto

package protocols

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

type PeersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Peers string `protobuf:"bytes,1,opt,name=peers,proto3" json:"peers,omitempty"`
}

func (x *PeersResponse) Reset() {
	*x = PeersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocols_bridge_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeersResponse) ProtoMessage() {}

func (x *PeersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protocols_bridge_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeersResponse.ProtoReflect.Descriptor instead.
func (*PeersResponse) Descriptor() ([]byte, []int) {
	return file_protocols_bridge_proto_rawDescGZIP(), []int{0}
}

func (x *PeersResponse) GetPeers() string {
	if x != nil {
		return x.Peers
	}
	return ""
}

type RequestData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Peer    string `protobuf:"bytes,1,opt,name=peer,proto3" json:"peer,omitempty"`
	Handler string `protobuf:"bytes,2,opt,name=handler,proto3" json:"handler,omitempty"`
	User    string `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	Payload string `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *RequestData) Reset() {
	*x = RequestData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocols_bridge_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestData) ProtoMessage() {}

func (x *RequestData) ProtoReflect() protoreflect.Message {
	mi := &file_protocols_bridge_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestData.ProtoReflect.Descriptor instead.
func (*RequestData) Descriptor() ([]byte, []int) {
	return file_protocols_bridge_proto_rawDescGZIP(), []int{1}
}

func (x *RequestData) GetPeer() string {
	if x != nil {
		return x.Peer
	}
	return ""
}

func (x *RequestData) GetHandler() string {
	if x != nil {
		return x.Handler
	}
	return ""
}

func (x *RequestData) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *RequestData) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

type ResponseData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload string `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	Status  int64  `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ResponseData) Reset() {
	*x = ResponseData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocols_bridge_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseData) ProtoMessage() {}

func (x *ResponseData) ProtoReflect() protoreflect.Message {
	mi := &file_protocols_bridge_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseData.ProtoReflect.Descriptor instead.
func (*ResponseData) Descriptor() ([]byte, []int) {
	return file_protocols_bridge_proto_rawDescGZIP(), []int{2}
}

func (x *ResponseData) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

func (x *ResponseData) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

var File_protocols_bridge_proto protoreflect.FileDescriptor

var file_protocols_bridge_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x2f, 0x62, 0x72, 0x69, 0x64,
	0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x25, 0x0a, 0x0d, 0x50, 0x65, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x65, 0x65,
	0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x22,
	0x69, 0x0a, 0x0b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x65, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x65,
	0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x40, 0x0a, 0x0c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0x30, 0x0a, 0x05,
	0x50, 0x65, 0x65, 0x72, 0x73, 0x12, 0x27, 0x0a, 0x05, 0x50, 0x65, 0x65, 0x72, 0x73, 0x12, 0x0e,
	0x2e, 0x50, 0x65, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x0e,
	0x2e, 0x50, 0x65, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x38,
	0x0a, 0x0e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x26, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0c, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x0d, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x42, 0x4d, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e,
	0x42, 0x72, 0x69, 0x64, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73,
	0x42, 0x06, 0x42, 0x72, 0x69, 0x64, 0x67, 0x65, 0x50, 0x01, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x69, 0x68, 0x61, 0x6c, 0x69, 0x63, 0x32, 0x30,
	0x34, 0x30, 0x2f, 0x48, 0x75, 0x62, 0x2d, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protocols_bridge_proto_rawDescOnce sync.Once
	file_protocols_bridge_proto_rawDescData = file_protocols_bridge_proto_rawDesc
)

func file_protocols_bridge_proto_rawDescGZIP() []byte {
	file_protocols_bridge_proto_rawDescOnce.Do(func() {
		file_protocols_bridge_proto_rawDescData = protoimpl.X.CompressGZIP(file_protocols_bridge_proto_rawDescData)
	})
	return file_protocols_bridge_proto_rawDescData
}

var file_protocols_bridge_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_protocols_bridge_proto_goTypes = []interface{}{
	(*PeersResponse)(nil), // 0: PeersResponse
	(*RequestData)(nil),   // 1: RequestData
	(*ResponseData)(nil),  // 2: ResponseData
}
var file_protocols_bridge_proto_depIdxs = []int32{
	0, // 0: Peers.Peers:input_type -> PeersResponse
	1, // 1: RequestService.Request:input_type -> RequestData
	0, // 2: Peers.Peers:output_type -> PeersResponse
	2, // 3: RequestService.Request:output_type -> ResponseData
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protocols_bridge_proto_init() }
func file_protocols_bridge_proto_init() {
	if File_protocols_bridge_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protocols_bridge_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeersResponse); i {
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
		file_protocols_bridge_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestData); i {
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
		file_protocols_bridge_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseData); i {
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
			RawDescriptor: file_protocols_bridge_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_protocols_bridge_proto_goTypes,
		DependencyIndexes: file_protocols_bridge_proto_depIdxs,
		MessageInfos:      file_protocols_bridge_proto_msgTypes,
	}.Build()
	File_protocols_bridge_proto = out.File
	file_protocols_bridge_proto_rawDesc = nil
	file_protocols_bridge_proto_goTypes = nil
	file_protocols_bridge_proto_depIdxs = nil
}
