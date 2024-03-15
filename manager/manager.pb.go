// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
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

type WriteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      uint64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Content string  `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Client  *Client `protobuf:"bytes,3,opt,name=client,proto3" json:"client,omitempty"`
}

func (x *WriteRequest) Reset() {
	*x = WriteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manager_manager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteRequest) ProtoMessage() {}

func (x *WriteRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use WriteRequest.ProtoReflect.Descriptor instead.
func (*WriteRequest) Descriptor() ([]byte, []int) {
	return file_manager_manager_proto_rawDescGZIP(), []int{0}
}

func (x *WriteRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *WriteRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *WriteRequest) GetClient() *Client {
	if x != nil {
		return x.Client
	}
	return nil
}

type ListenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      uint64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Type    string  `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Content string  `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Client  *Client `protobuf:"bytes,4,opt,name=client,proto3" json:"client,omitempty"`
}

func (x *ListenResponse) Reset() {
	*x = ListenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manager_manager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListenResponse) ProtoMessage() {}

func (x *ListenResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ListenResponse.ProtoReflect.Descriptor instead.
func (*ListenResponse) Descriptor() ([]byte, []int) {
	return file_manager_manager_proto_rawDescGZIP(), []int{1}
}

func (x *ListenResponse) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ListenResponse) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ListenResponse) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *ListenResponse) GetClient() *Client {
	if x != nil {
		return x.Client
	}
	return nil
}

type StartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path   uint64  `protobuf:"varint,1,opt,name=path,proto3" json:"path,omitempty"`
	Client *Client `protobuf:"bytes,2,opt,name=client,proto3" json:"client,omitempty"`
}

func (x *StartRequest) Reset() {
	*x = StartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manager_manager_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartRequest) ProtoMessage() {}

func (x *StartRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use StartRequest.ProtoReflect.Descriptor instead.
func (*StartRequest) Descriptor() ([]byte, []int) {
	return file_manager_manager_proto_rawDescGZIP(), []int{2}
}

func (x *StartRequest) GetPath() uint64 {
	if x != nil {
		return x.Path
	}
	return 0
}

func (x *StartRequest) GetClient() *Client {
	if x != nil {
		return x.Client
	}
	return nil
}

type Client struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Client) Reset() {
	*x = Client{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manager_manager_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Client) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Client) ProtoMessage() {}

func (x *Client) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Client.ProtoReflect.Descriptor instead.
func (*Client) Descriptor() ([]byte, []int) {
	return file_manager_manager_proto_rawDescGZIP(), []int{3}
}

func (x *Client) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type StatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Usedmemory uint64 `protobuf:"varint,1,opt,name=usedmemory,proto3" json:"usedmemory,omitempty"`
}

func (x *StatusResponse) Reset() {
	*x = StatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manager_manager_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusResponse) ProtoMessage() {}

func (x *StatusResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use StatusResponse.ProtoReflect.Descriptor instead.
func (*StatusResponse) Descriptor() ([]byte, []int) {
	return file_manager_manager_proto_rawDescGZIP(), []int{4}
}

func (x *StatusResponse) GetUsedmemory() uint64 {
	if x != nil {
		return x.Usedmemory
	}
	return 0
}

var File_manager_manager_proto protoreflect.FileDescriptor

var file_manager_manager_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x59, 0x0a, 0x0c, 0x57, 0x72, 0x69, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x12, 0x1f, 0x0a, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x07, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x22, 0x6f, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x07, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x22, 0x43, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x1f, 0x0a, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x52, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x22, 0x18, 0x0a, 0x06, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x30, 0x0a, 0x0e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x64, 0x6d, 0x65, 0x6d, 0x6f,
	0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x64, 0x6d, 0x65,
	0x6d, 0x6f, 0x72, 0x79, 0x32, 0xfa, 0x01, 0x0a, 0x07, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x12, 0x1a, 0x0a, 0x04, 0x4c, 0x6f, 0x63, 0x6b, 0x12, 0x07, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x1a, 0x07, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x1c, 0x0a, 0x06,
	0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x07, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x1a,
	0x07, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x27, 0x0a, 0x05, 0x57, 0x72,
	0x69, 0x74, 0x65, 0x12, 0x0d, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x00, 0x12, 0x27, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x07,
	0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x1a, 0x0f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x21, 0x0a, 0x05,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x0d, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x07, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x12,
	0x1a, 0x0a, 0x04, 0x53, 0x74, 0x6f, 0x70, 0x12, 0x07, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x1a, 0x07, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x24, 0x0a, 0x06, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x07, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x1a, 0x0f,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x34, 0x5a, 0x32, 0x63, 0x67, 0x69, 0x74, 0x2e, 0x62, 0x62, 0x61, 0x61, 0x2e, 0x66,
	0x75, 0x6e, 0x2f, 0x62, 0x62, 0x61, 0x61, 0x2f, 0x6d, 0x69, 0x6e, 0x65, 0x63, 0x72, 0x61, 0x66,
	0x74, 0x2d, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_manager_manager_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_manager_manager_proto_goTypes = []interface{}{
	(*WriteRequest)(nil),   // 0: WriteRequest
	(*ListenResponse)(nil), // 1: ListenResponse
	(*StartRequest)(nil),   // 2: StartRequest
	(*Client)(nil),         // 3: Client
	(*StatusResponse)(nil), // 4: StatusResponse
}
var file_manager_manager_proto_depIdxs = []int32{
	3,  // 0: WriteRequest.client:type_name -> Client
	3,  // 1: ListenResponse.client:type_name -> Client
	3,  // 2: StartRequest.client:type_name -> Client
	3,  // 3: Manager.Lock:input_type -> Client
	3,  // 4: Manager.Unlock:input_type -> Client
	0,  // 5: Manager.Write:input_type -> WriteRequest
	3,  // 6: Manager.Message:input_type -> Client
	2,  // 7: Manager.Start:input_type -> StartRequest
	3,  // 8: Manager.Stop:input_type -> Client
	3,  // 9: Manager.Status:input_type -> Client
	3,  // 10: Manager.Lock:output_type -> Client
	3,  // 11: Manager.Unlock:output_type -> Client
	0,  // 12: Manager.Write:output_type -> WriteRequest
	1,  // 13: Manager.Message:output_type -> ListenResponse
	3,  // 14: Manager.Start:output_type -> Client
	3,  // 15: Manager.Stop:output_type -> Client
	4,  // 16: Manager.Status:output_type -> StatusResponse
	10, // [10:17] is the sub-list for method output_type
	3,  // [3:10] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_manager_manager_proto_init() }
func file_manager_manager_proto_init() {
	if File_manager_manager_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_manager_manager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteRequest); i {
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
			switch v := v.(*ListenResponse); i {
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
			switch v := v.(*StartRequest); i {
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
			switch v := v.(*Client); i {
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
			switch v := v.(*StatusResponse); i {
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
			NumMessages:   5,
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
