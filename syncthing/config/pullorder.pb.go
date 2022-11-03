// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: config/pullorder.proto

package config

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

type PullOrder int32

const (
	PullOrder_PULL_ORDER_RANDOM         PullOrder = 0
	PullOrder_PULL_ORDER_ALPHABETIC     PullOrder = 1
	PullOrder_PULL_ORDER_SMALLEST_FIRST PullOrder = 2
	PullOrder_PULL_ORDER_LARGEST_FIRST  PullOrder = 3
	PullOrder_PULL_ORDER_OLDEST_FIRST   PullOrder = 4
	PullOrder_PULL_ORDER_NEWEST_FIRST   PullOrder = 5
)

// Enum value maps for PullOrder.
var (
	PullOrder_name = map[int32]string{
		0: "PULL_ORDER_RANDOM",
		1: "PULL_ORDER_ALPHABETIC",
		2: "PULL_ORDER_SMALLEST_FIRST",
		3: "PULL_ORDER_LARGEST_FIRST",
		4: "PULL_ORDER_OLDEST_FIRST",
		5: "PULL_ORDER_NEWEST_FIRST",
	}
	PullOrder_value = map[string]int32{
		"PULL_ORDER_RANDOM":         0,
		"PULL_ORDER_ALPHABETIC":     1,
		"PULL_ORDER_SMALLEST_FIRST": 2,
		"PULL_ORDER_LARGEST_FIRST":  3,
		"PULL_ORDER_OLDEST_FIRST":   4,
		"PULL_ORDER_NEWEST_FIRST":   5,
	}
)

func (x PullOrder) Enum() *PullOrder {
	p := new(PullOrder)
	*p = x
	return p
}

func (x PullOrder) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PullOrder) Descriptor() protoreflect.EnumDescriptor {
	return file_config_pullorder_proto_enumTypes[0].Descriptor()
}

func (PullOrder) Type() protoreflect.EnumType {
	return &file_config_pullorder_proto_enumTypes[0]
}

func (x PullOrder) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PullOrder.Descriptor instead.
func (PullOrder) EnumDescriptor() ([]byte, []int) {
	return file_config_pullorder_proto_rawDescGZIP(), []int{0}
}

var File_config_pullorder_proto protoreflect.FileDescriptor

var file_config_pullorder_proto_rawDesc = []byte{
	0x0a, 0x16, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x70, 0x75, 0x6c, 0x6c, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2a, 0xb4, 0x01, 0x0a, 0x09, 0x50, 0x75, 0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x15,
	0x0a, 0x11, 0x50, 0x55, 0x4c, 0x4c, 0x5f, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x52, 0x41, 0x4e,
	0x44, 0x4f, 0x4d, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x50, 0x55, 0x4c, 0x4c, 0x5f, 0x4f, 0x52,
	0x44, 0x45, 0x52, 0x5f, 0x41, 0x4c, 0x50, 0x48, 0x41, 0x42, 0x45, 0x54, 0x49, 0x43, 0x10, 0x01,
	0x12, 0x1d, 0x0a, 0x19, 0x50, 0x55, 0x4c, 0x4c, 0x5f, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x53,
	0x4d, 0x41, 0x4c, 0x4c, 0x45, 0x53, 0x54, 0x5f, 0x46, 0x49, 0x52, 0x53, 0x54, 0x10, 0x02, 0x12,
	0x1c, 0x0a, 0x18, 0x50, 0x55, 0x4c, 0x4c, 0x5f, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x4c, 0x41,
	0x52, 0x47, 0x45, 0x53, 0x54, 0x5f, 0x46, 0x49, 0x52, 0x53, 0x54, 0x10, 0x03, 0x12, 0x1b, 0x0a,
	0x17, 0x50, 0x55, 0x4c, 0x4c, 0x5f, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x4f, 0x4c, 0x44, 0x45,
	0x53, 0x54, 0x5f, 0x46, 0x49, 0x52, 0x53, 0x54, 0x10, 0x04, 0x12, 0x1b, 0x0a, 0x17, 0x50, 0x55,
	0x4c, 0x4c, 0x5f, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x4e, 0x45, 0x57, 0x45, 0x53, 0x54, 0x5f,
	0x46, 0x49, 0x52, 0x53, 0x54, 0x10, 0x05, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x65, 0x62, 0x66, 0x72, 0x6f, 0x73, 0x73, 0x2f, 0x73,
	0x79, 0x6e, 0x63, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x2d, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2f,
	0x73, 0x79, 0x6e, 0x63, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_pullorder_proto_rawDescOnce sync.Once
	file_config_pullorder_proto_rawDescData = file_config_pullorder_proto_rawDesc
)

func file_config_pullorder_proto_rawDescGZIP() []byte {
	file_config_pullorder_proto_rawDescOnce.Do(func() {
		file_config_pullorder_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_pullorder_proto_rawDescData)
	})
	return file_config_pullorder_proto_rawDescData
}

var file_config_pullorder_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_config_pullorder_proto_goTypes = []interface{}{
	(PullOrder)(0), // 0: config.PullOrder
}
var file_config_pullorder_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_config_pullorder_proto_init() }
func file_config_pullorder_proto_init() {
	if File_config_pullorder_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_config_pullorder_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_pullorder_proto_goTypes,
		DependencyIndexes: file_config_pullorder_proto_depIdxs,
		EnumInfos:         file_config_pullorder_proto_enumTypes,
	}.Build()
	File_config_pullorder_proto = out.File
	file_config_pullorder_proto_rawDesc = nil
	file_config_pullorder_proto_goTypes = nil
	file_config_pullorder_proto_depIdxs = nil
}
