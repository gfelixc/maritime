// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: v1/port_service.proto

package portv1

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

type CreateOrUpdateFromPortsDataFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename string `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
}

func (x *CreateOrUpdateFromPortsDataFileRequest) Reset() {
	*x = CreateOrUpdateFromPortsDataFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_port_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrUpdateFromPortsDataFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrUpdateFromPortsDataFileRequest) ProtoMessage() {}

func (x *CreateOrUpdateFromPortsDataFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_port_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrUpdateFromPortsDataFileRequest.ProtoReflect.Descriptor instead.
func (*CreateOrUpdateFromPortsDataFileRequest) Descriptor() ([]byte, []int) {
	return file_v1_port_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateOrUpdateFromPortsDataFileRequest) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

type CreateOrUpdateFromPortsDataFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PortsProcessed uint32   `protobuf:"varint,1,opt,name=ports_processed,json=portsProcessed,proto3" json:"ports_processed,omitempty"`
	PortsSucceeded uint32   `protobuf:"varint,2,opt,name=ports_succeeded,json=portsSucceeded,proto3" json:"ports_succeeded,omitempty"`
	PortsFailed    uint32   `protobuf:"varint,3,opt,name=ports_failed,json=portsFailed,proto3" json:"ports_failed,omitempty"`
	Errors         []string `protobuf:"bytes,4,rep,name=errors,proto3" json:"errors,omitempty"`
}

func (x *CreateOrUpdateFromPortsDataFileResponse) Reset() {
	*x = CreateOrUpdateFromPortsDataFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_port_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrUpdateFromPortsDataFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrUpdateFromPortsDataFileResponse) ProtoMessage() {}

func (x *CreateOrUpdateFromPortsDataFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_port_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrUpdateFromPortsDataFileResponse.ProtoReflect.Descriptor instead.
func (*CreateOrUpdateFromPortsDataFileResponse) Descriptor() ([]byte, []int) {
	return file_v1_port_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateOrUpdateFromPortsDataFileResponse) GetPortsProcessed() uint32 {
	if x != nil {
		return x.PortsProcessed
	}
	return 0
}

func (x *CreateOrUpdateFromPortsDataFileResponse) GetPortsSucceeded() uint32 {
	if x != nil {
		return x.PortsSucceeded
	}
	return 0
}

func (x *CreateOrUpdateFromPortsDataFileResponse) GetPortsFailed() uint32 {
	if x != nil {
		return x.PortsFailed
	}
	return 0
}

func (x *CreateOrUpdateFromPortsDataFileResponse) GetErrors() []string {
	if x != nil {
		return x.Errors
	}
	return nil
}

var File_v1_port_service_proto protoreflect.FileDescriptor

var file_v1_port_service_proto_rawDesc = []byte{
	0x0a, 0x15, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x31,
	0x22, 0x44, 0x0a, 0x26, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x44, 0x61, 0x74, 0x61, 0x46,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69,
	0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69,
	0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xb6, 0x01, 0x0a, 0x27, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x50, 0x6f, 0x72,
	0x74, 0x73, 0x44, 0x61, 0x74, 0x61, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x70, 0x6f, 0x72,
	0x74, 0x73, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x70,
	0x6f, 0x72, 0x74, 0x73, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x65, 0x64, 0x65, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x53, 0x75, 0x63, 0x63, 0x65,
	0x65, 0x64, 0x65, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x5f, 0x66, 0x61,
	0x69, 0x6c, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x70, 0x6f, 0x72, 0x74,
	0x73, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x32,
	0x9a, 0x01, 0x0a, 0x11, 0x50, 0x6f, 0x72, 0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x84, 0x01, 0x0a, 0x1f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x50, 0x6f, 0x72, 0x74,
	0x73, 0x44, 0x61, 0x74, 0x61, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x2f, 0x2e, 0x70, 0x6f, 0x72, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x44, 0x61, 0x74, 0x61, 0x46,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x70, 0x6f, 0x72,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x44, 0x61, 0x74, 0x61,
	0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x83, 0x01, 0x0a,
	0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x10, 0x50, 0x6f,
	0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x66, 0x65,
	0x6c, 0x69, 0x78, 0x63, 0x2f, 0x6d, 0x61, 0x72, 0x69, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31,
	0x3b, 0x70, 0x6f, 0x72, 0x74, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x50, 0x58, 0x58, 0xaa, 0x02, 0x07,
	0x50, 0x6f, 0x72, 0x74, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x07, 0x50, 0x6f, 0x72, 0x74, 0x5c, 0x56,
	0x31, 0xe2, 0x02, 0x13, 0x50, 0x6f, 0x72, 0x74, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x08, 0x50, 0x6f, 0x72, 0x74, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_port_service_proto_rawDescOnce sync.Once
	file_v1_port_service_proto_rawDescData = file_v1_port_service_proto_rawDesc
)

func file_v1_port_service_proto_rawDescGZIP() []byte {
	file_v1_port_service_proto_rawDescOnce.Do(func() {
		file_v1_port_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_port_service_proto_rawDescData)
	})
	return file_v1_port_service_proto_rawDescData
}

var file_v1_port_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_v1_port_service_proto_goTypes = []interface{}{
	(*CreateOrUpdateFromPortsDataFileRequest)(nil),  // 0: port.v1.CreateOrUpdateFromPortsDataFileRequest
	(*CreateOrUpdateFromPortsDataFileResponse)(nil), // 1: port.v1.CreateOrUpdateFromPortsDataFileResponse
}
var file_v1_port_service_proto_depIdxs = []int32{
	0, // 0: port.v1.PortDomainService.CreateOrUpdateFromPortsDataFile:input_type -> port.v1.CreateOrUpdateFromPortsDataFileRequest
	1, // 1: port.v1.PortDomainService.CreateOrUpdateFromPortsDataFile:output_type -> port.v1.CreateOrUpdateFromPortsDataFileResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v1_port_service_proto_init() }
func file_v1_port_service_proto_init() {
	if File_v1_port_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_port_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrUpdateFromPortsDataFileRequest); i {
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
		file_v1_port_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrUpdateFromPortsDataFileResponse); i {
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
			RawDescriptor: file_v1_port_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_port_service_proto_goTypes,
		DependencyIndexes: file_v1_port_service_proto_depIdxs,
		MessageInfos:      file_v1_port_service_proto_msgTypes,
	}.Build()
	File_v1_port_service_proto = out.File
	file_v1_port_service_proto_rawDesc = nil
	file_v1_port_service_proto_goTypes = nil
	file_v1_port_service_proto_depIdxs = nil
}
