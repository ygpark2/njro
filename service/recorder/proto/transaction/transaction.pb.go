// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: service/recorder/proto/transaction/transaction.proto

package transaction

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type ReadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *ReadRequest) Reset() {
	*x = ReadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_recorder_proto_transaction_transaction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadRequest) ProtoMessage() {}

func (x *ReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_recorder_proto_transaction_transaction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadRequest.ProtoReflect.Descriptor instead.
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return file_service_recorder_proto_transaction_transaction_proto_rawDescGZIP(), []int{0}
}

func (x *ReadRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type WriteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string            `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Event *TransactionEvent `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *WriteRequest) Reset() {
	*x = WriteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_recorder_proto_transaction_transaction_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteRequest) ProtoMessage() {}

func (x *WriteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_recorder_proto_transaction_transaction_proto_msgTypes[1]
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
	return file_service_recorder_proto_transaction_transaction_proto_rawDescGZIP(), []int{1}
}

func (x *WriteRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *WriteRequest) GetEvent() *TransactionEvent {
	if x != nil {
		return x.Event
	}
	return nil
}

// Transaction Event
type TransactionEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// request
	Req []byte `protobuf:"bytes,1,opt,name=req,proto3" json:"req,omitempty"`
	// responce
	Rsp []byte `protobuf:"bytes,2,opt,name=rsp,proto3" json:"rsp,omitempty"`
}

func (x *TransactionEvent) Reset() {
	*x = TransactionEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_recorder_proto_transaction_transaction_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionEvent) ProtoMessage() {}

func (x *TransactionEvent) ProtoReflect() protoreflect.Message {
	mi := &file_service_recorder_proto_transaction_transaction_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionEvent.ProtoReflect.Descriptor instead.
func (*TransactionEvent) Descriptor() ([]byte, []int) {
	return file_service_recorder_proto_transaction_transaction_proto_rawDescGZIP(), []int{2}
}

func (x *TransactionEvent) GetReq() []byte {
	if x != nil {
		return x.Req
	}
	return nil
}

func (x *TransactionEvent) GetRsp() []byte {
	if x != nil {
		return x.Rsp
	}
	return nil
}

var File_service_recorder_proto_transaction_transaction_proto protoreflect.FileDescriptor

var file_service_recorder_proto_transaction_transaction_proto_rawDesc = []byte{
	0x0a, 0x34, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x6d, 0x6b, 0x69, 0x74, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x24, 0x0a, 0x0b, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x3a, 0x03, 0xf8, 0x42, 0x01, 0x22, 0x67, 0x0a, 0x0c,
	0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x40,
	0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e,
	0x6d, 0x6b, 0x69, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x72, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x3a, 0x03, 0xf8, 0x42, 0x01, 0x22, 0x36, 0x0a, 0x10, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x71,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x72, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x72,
	0x73, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x72, 0x73, 0x70, 0x32, 0xb8, 0x01,
	0x0a, 0x12, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x59, 0x0a, 0x04, 0x52, 0x65, 0x61, 0x64, 0x12, 0x25, 0x2e, 0x6d,
	0x6b, 0x69, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x6d, 0x6b, 0x69, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12,
	0x47, 0x0a, 0x05, 0x57, 0x72, 0x69, 0x74, 0x65, 0x12, 0x26, 0x2e, 0x6d, 0x6b, 0x69, 0x74, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x5a, 0x0a, 0x18, 0x6d, 0x6b, 0x69, 0x74,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x79, 0x67, 0x70, 0x61, 0x72, 0x6b, 0x32, 0x2f, 0x6d, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_recorder_proto_transaction_transaction_proto_rawDescOnce sync.Once
	file_service_recorder_proto_transaction_transaction_proto_rawDescData = file_service_recorder_proto_transaction_transaction_proto_rawDesc
)

func file_service_recorder_proto_transaction_transaction_proto_rawDescGZIP() []byte {
	file_service_recorder_proto_transaction_transaction_proto_rawDescOnce.Do(func() {
		file_service_recorder_proto_transaction_transaction_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_recorder_proto_transaction_transaction_proto_rawDescData)
	})
	return file_service_recorder_proto_transaction_transaction_proto_rawDescData
}

var file_service_recorder_proto_transaction_transaction_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_service_recorder_proto_transaction_transaction_proto_goTypes = []interface{}{
	(*ReadRequest)(nil),      // 0: mkit.service.recorder.v1.ReadRequest
	(*WriteRequest)(nil),     // 1: mkit.service.recorder.v1.WriteRequest
	(*TransactionEvent)(nil), // 2: mkit.service.recorder.v1.TransactionEvent
	(*empty.Empty)(nil),      // 3: google.protobuf.Empty
}
var file_service_recorder_proto_transaction_transaction_proto_depIdxs = []int32{
	2, // 0: mkit.service.recorder.v1.WriteRequest.event:type_name -> mkit.service.recorder.v1.TransactionEvent
	0, // 1: mkit.service.recorder.v1.TransactionService.Read:input_type -> mkit.service.recorder.v1.ReadRequest
	1, // 2: mkit.service.recorder.v1.TransactionService.Write:input_type -> mkit.service.recorder.v1.WriteRequest
	2, // 3: mkit.service.recorder.v1.TransactionService.Read:output_type -> mkit.service.recorder.v1.TransactionEvent
	3, // 4: mkit.service.recorder.v1.TransactionService.Write:output_type -> google.protobuf.Empty
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_service_recorder_proto_transaction_transaction_proto_init() }
func file_service_recorder_proto_transaction_transaction_proto_init() {
	if File_service_recorder_proto_transaction_transaction_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_recorder_proto_transaction_transaction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadRequest); i {
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
		file_service_recorder_proto_transaction_transaction_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_service_recorder_proto_transaction_transaction_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionEvent); i {
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
			RawDescriptor: file_service_recorder_proto_transaction_transaction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_recorder_proto_transaction_transaction_proto_goTypes,
		DependencyIndexes: file_service_recorder_proto_transaction_transaction_proto_depIdxs,
		MessageInfos:      file_service_recorder_proto_transaction_transaction_proto_msgTypes,
	}.Build()
	File_service_recorder_proto_transaction_transaction_proto = out.File
	file_service_recorder_proto_transaction_transaction_proto_rawDesc = nil
	file_service_recorder_proto_transaction_transaction_proto_goTypes = nil
	file_service_recorder_proto_transaction_transaction_proto_depIdxs = nil
}