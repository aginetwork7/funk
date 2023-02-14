// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: errorpb/status.proto

package errorpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Code int32

const (
	Code_OK                 Code = 0
	Code_Canceled           Code = 1
	Code_Unknown            Code = 2
	Code_InvalidArgument    Code = 3
	Code_DeadlineExceeded   Code = 4
	Code_NotFound           Code = 5
	Code_AlreadyExists      Code = 6
	Code_PermissionDenied   Code = 7
	Code_ResourceExhausted  Code = 8
	Code_FailedPrecondition Code = 9
	Code_Aborted            Code = 10
	Code_OutOfRange         Code = 11
	Code_Unimplemented      Code = 12
	Code_Internal           Code = 13
	Code_Unavailable        Code = 14
	Code_DataLoss           Code = 15
	Code_Unauthenticated    Code = 16
	Code_TooManyRequests    Code = 17
)

// Enum value maps for Code.
var (
	Code_name = map[int32]string{
		0:  "OK",
		1:  "Canceled",
		2:  "Unknown",
		3:  "InvalidArgument",
		4:  "DeadlineExceeded",
		5:  "NotFound",
		6:  "AlreadyExists",
		7:  "PermissionDenied",
		8:  "ResourceExhausted",
		9:  "FailedPrecondition",
		10: "Aborted",
		11: "OutOfRange",
		12: "Unimplemented",
		13: "Internal",
		14: "Unavailable",
		15: "DataLoss",
		16: "Unauthenticated",
		17: "TooManyRequests",
	}
	Code_value = map[string]int32{
		"OK":                 0,
		"Canceled":           1,
		"Unknown":            2,
		"InvalidArgument":    3,
		"DeadlineExceeded":   4,
		"NotFound":           5,
		"AlreadyExists":      6,
		"PermissionDenied":   7,
		"ResourceExhausted":  8,
		"FailedPrecondition": 9,
		"Aborted":            10,
		"OutOfRange":         11,
		"Unimplemented":      12,
		"Internal":           13,
		"Unavailable":        14,
		"DataLoss":           15,
		"Unauthenticated":    16,
		"TooManyRequests":    17,
	}
)

func (x Code) Enum() *Code {
	p := new(Code)
	*p = x
	return p
}

func (x Code) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Code) Descriptor() protoreflect.EnumDescriptor {
	return file_errorpb_status_proto_enumTypes[0].Descriptor()
}

func (Code) Type() protoreflect.EnumType {
	return &file_errorpb_status_proto_enumTypes[0]
}

func (x Code) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Code.Descriptor instead.
func (Code) EnumDescriptor() ([]byte, []int) {
	return file_errorpb_status_proto_rawDescGZIP(), []int{0}
}

type GenStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GenEnable bool `protobuf:"varint,1,opt,name=gen_enable,json=genEnable,proto3" json:"gen_enable,omitempty"`
	Code      Code `protobuf:"varint,2,opt,name=code,proto3,enum=status.Code" json:"code,omitempty"`
}

func (x *GenStatus) Reset() {
	*x = GenStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_errorpb_status_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenStatus) ProtoMessage() {}

func (x *GenStatus) ProtoReflect() protoreflect.Message {
	mi := &file_errorpb_status_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenStatus.ProtoReflect.Descriptor instead.
func (*GenStatus) Descriptor() ([]byte, []int) {
	return file_errorpb_status_proto_rawDescGZIP(), []int{0}
}

func (x *GenStatus) GetGenEnable() bool {
	if x != nil {
		return x.GenEnable
	}
	return false
}

func (x *GenStatus) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_OK
}

var file_errorpb_status_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.EnumValueOptions)(nil),
		ExtensionType: (*GenStatus)(nil),
		Field:         100002,
		Name:          "status.field",
		Tag:           "bytes,100002,opt,name=field",
		Filename:      "errorpb/status.proto",
	},
	{
		ExtendedType:  (*descriptorpb.EnumOptions)(nil),
		ExtensionType: (*GenStatus)(nil),
		Field:         100001,
		Name:          "status.opts",
		Tag:           "bytes,100001,opt,name=opts",
		Filename:      "errorpb/status.proto",
	},
}

// Extension fields to descriptorpb.EnumValueOptions.
var (
	// optional status.GenStatus field = 100002;
	E_Field = &file_errorpb_status_proto_extTypes[0]
)

// Extension fields to descriptorpb.EnumOptions.
var (
	// optional status.GenStatus opts = 100001;
	E_Opts = &file_errorpb_status_proto_extTypes[1]
)

var File_errorpb_status_proto protoreflect.FileDescriptor

var file_errorpb_status_proto_rawDesc = []byte{
	0x0a, 0x14, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x70, 0x62, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x1a, 0x20,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x4c, 0x0a, 0x09, 0x47, 0x65, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a,
	0x0a, 0x67, 0x65, 0x6e, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x09, 0x67, 0x65, 0x6e, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x2a, 0xc1,
	0x02, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12,
	0x0c, 0x0a, 0x08, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x65, 0x64, 0x10, 0x01, 0x12, 0x0b, 0x0a,
	0x07, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x49, 0x6e,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x41, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x10, 0x03, 0x12,
	0x14, 0x0a, 0x10, 0x44, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x45, 0x78, 0x63, 0x65, 0x65,
	0x64, 0x65, 0x64, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x4e, 0x6f, 0x74, 0x46, 0x6f, 0x75, 0x6e,
	0x64, 0x10, 0x05, 0x12, 0x11, 0x0a, 0x0d, 0x41, 0x6c, 0x72, 0x65, 0x61, 0x64, 0x79, 0x45, 0x78,
	0x69, 0x73, 0x74, 0x73, 0x10, 0x06, 0x12, 0x14, 0x0a, 0x10, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x6e, 0x69, 0x65, 0x64, 0x10, 0x07, 0x12, 0x15, 0x0a, 0x11,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x45, 0x78, 0x68, 0x61, 0x75, 0x73, 0x74, 0x65,
	0x64, 0x10, 0x08, 0x12, 0x16, 0x0a, 0x12, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x50, 0x72, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x09, 0x12, 0x0b, 0x0a, 0x07, 0x41,
	0x62, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x10, 0x0a, 0x12, 0x0e, 0x0a, 0x0a, 0x4f, 0x75, 0x74, 0x4f,
	0x66, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x10, 0x0b, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x6e, 0x69, 0x6d,
	0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x65, 0x64, 0x10, 0x0c, 0x12, 0x0c, 0x0a, 0x08, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x10, 0x0d, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x6e, 0x61,
	0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x10, 0x0e, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x61,
	0x74, 0x61, 0x4c, 0x6f, 0x73, 0x73, 0x10, 0x0f, 0x12, 0x13, 0x0a, 0x0f, 0x55, 0x6e, 0x61, 0x75,
	0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x10, 0x10, 0x12, 0x13, 0x0a,
	0x0f, 0x54, 0x6f, 0x6f, 0x4d, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73,
	0x10, 0x11, 0x3a, 0x4c, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x21, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6e,
	0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xa2,
	0x8d, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e,
	0x47, 0x65, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x3a, 0x45, 0x0a, 0x04, 0x6f, 0x70, 0x74, 0x73, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xa1, 0x8d, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x47, 0x65, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x04, 0x6f, 0x70, 0x74, 0x73, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x75, 0x62, 0x67, 0x6f, 0x2f, 0x66, 0x75, 0x6e, 0x6b,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x70, 0x62, 0x3b, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_errorpb_status_proto_rawDescOnce sync.Once
	file_errorpb_status_proto_rawDescData = file_errorpb_status_proto_rawDesc
)

func file_errorpb_status_proto_rawDescGZIP() []byte {
	file_errorpb_status_proto_rawDescOnce.Do(func() {
		file_errorpb_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_errorpb_status_proto_rawDescData)
	})
	return file_errorpb_status_proto_rawDescData
}

var file_errorpb_status_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_errorpb_status_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_errorpb_status_proto_goTypes = []interface{}{
	(Code)(0),                             // 0: status.Code
	(*GenStatus)(nil),                     // 1: status.GenStatus
	(*descriptorpb.EnumValueOptions)(nil), // 2: google.protobuf.EnumValueOptions
	(*descriptorpb.EnumOptions)(nil),      // 3: google.protobuf.EnumOptions
}
var file_errorpb_status_proto_depIdxs = []int32{
	0, // 0: status.GenStatus.code:type_name -> status.Code
	2, // 1: status.field:extendee -> google.protobuf.EnumValueOptions
	3, // 2: status.opts:extendee -> google.protobuf.EnumOptions
	1, // 3: status.field:type_name -> status.GenStatus
	1, // 4: status.opts:type_name -> status.GenStatus
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	3, // [3:5] is the sub-list for extension type_name
	1, // [1:3] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_errorpb_status_proto_init() }
func file_errorpb_status_proto_init() {
	if File_errorpb_status_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_errorpb_status_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenStatus); i {
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
			RawDescriptor: file_errorpb_status_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 2,
			NumServices:   0,
		},
		GoTypes:           file_errorpb_status_proto_goTypes,
		DependencyIndexes: file_errorpb_status_proto_depIdxs,
		EnumInfos:         file_errorpb_status_proto_enumTypes,
		MessageInfos:      file_errorpb_status_proto_msgTypes,
		ExtensionInfos:    file_errorpb_status_proto_extTypes,
	}.Build()
	File_errorpb_status_proto = out.File
	file_errorpb_status_proto_rawDesc = nil
	file_errorpb_status_proto_goTypes = nil
	file_errorpb_status_proto_depIdxs = nil
}
