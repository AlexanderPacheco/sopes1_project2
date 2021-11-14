// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: proto/demo.proto

package proto

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

type OperacionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operacion string `protobuf:"bytes,1,opt,name=operacion,proto3" json:"operacion,omitempty"`
	Valor1    string `protobuf:"bytes,2,opt,name=valor1,proto3" json:"valor1,omitempty"`
	Valor2    string `protobuf:"bytes,3,opt,name=valor2,proto3" json:"valor2,omitempty"`
}

func (x *OperacionRequest) Reset() {
	*x = OperacionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_demo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperacionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperacionRequest) ProtoMessage() {}

func (x *OperacionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_demo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperacionRequest.ProtoReflect.Descriptor instead.
func (*OperacionRequest) Descriptor() ([]byte, []int) {
	return file_proto_demo_proto_rawDescGZIP(), []int{0}
}

func (x *OperacionRequest) GetOperacion() string {
	if x != nil {
		return x.Operacion
	}
	return ""
}

func (x *OperacionRequest) GetValor1() string {
	if x != nil {
		return x.Valor1
	}
	return ""
}

func (x *OperacionRequest) GetValor2() string {
	if x != nil {
		return x.Valor2
	}
	return ""
}

type OperacionReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resultado string `protobuf:"bytes,1,opt,name=resultado,proto3" json:"resultado,omitempty"`
}

func (x *OperacionReply) Reset() {
	*x = OperacionReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_demo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperacionReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperacionReply) ProtoMessage() {}

func (x *OperacionReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_demo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperacionReply.ProtoReflect.Descriptor instead.
func (*OperacionReply) Descriptor() ([]byte, []int) {
	return file_proto_demo_proto_rawDescGZIP(), []int{1}
}

func (x *OperacionReply) GetResultado() string {
	if x != nil {
		return x.Resultado
	}
	return ""
}

var File_proto_demo_proto protoreflect.FileDescriptor

var file_proto_demo_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x60, 0x0a, 0x10, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x63, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x63, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x63, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x76,
	0x61, 0x6c, 0x6f, 0x72, 0x31, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x76, 0x61, 0x6c,
	0x6f, 0x72, 0x31, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x6f, 0x72, 0x32, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x6f, 0x72, 0x32, 0x22, 0x2e, 0x0a, 0x0e, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x63, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1c, 0x0a,
	0x09, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x61, 0x64, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x61, 0x64, 0x6f, 0x32, 0x58, 0x0a, 0x13, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x63, 0x69, 0x6f, 0x6e, 0x41, 0x72, 0x69, 0x74, 0x6d, 0x65, 0x74, 0x69,
	0x63, 0x61, 0x12, 0x41, 0x0a, 0x0d, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x72, 0x56, 0x61, 0x6c, 0x6f,
	0x72, 0x65, 0x73, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x63, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x63, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x63, 0x61, 0x72, 0x6c, 0x6f, 0x73, 0x64, 0x61, 0x76, 0x69,
	0x64, 0x2f, 0x64, 0x65, 0x6d, 0x6f, 0x2d, 0x67, 0x52, 0x50, 0x43, 0x2d, 0x6b, 0x75, 0x62, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x67, 0x52, 0x50, 0x43, 0x2d, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_demo_proto_rawDescOnce sync.Once
	file_proto_demo_proto_rawDescData = file_proto_demo_proto_rawDesc
)

func file_proto_demo_proto_rawDescGZIP() []byte {
	file_proto_demo_proto_rawDescOnce.Do(func() {
		file_proto_demo_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_demo_proto_rawDescData)
	})
	return file_proto_demo_proto_rawDescData
}

var file_proto_demo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_demo_proto_goTypes = []interface{}{
	(*OperacionRequest)(nil), // 0: proto.OperacionRequest
	(*OperacionReply)(nil),   // 1: proto.OperacionReply
}
var file_proto_demo_proto_depIdxs = []int32{
	0, // 0: proto.OperacionAritmetica.OperarValores:input_type -> proto.OperacionRequest
	1, // 1: proto.OperacionAritmetica.OperarValores:output_type -> proto.OperacionReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_demo_proto_init() }
func file_proto_demo_proto_init() {
	if File_proto_demo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_demo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperacionRequest); i {
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
		file_proto_demo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperacionReply); i {
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
			RawDescriptor: file_proto_demo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_demo_proto_goTypes,
		DependencyIndexes: file_proto_demo_proto_depIdxs,
		MessageInfos:      file_proto_demo_proto_msgTypes,
	}.Build()
	File_proto_demo_proto = out.File
	file_proto_demo_proto_rawDesc = nil
	file_proto_demo_proto_goTypes = nil
	file_proto_demo_proto_depIdxs = nil
}
