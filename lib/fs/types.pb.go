// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lib/fs/types.proto

package fs

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type FilesystemType int32

const (
	FilesystemTypeBasic FilesystemType = 0
	FilesystemTypeFake  FilesystemType = 1
)

var FilesystemType_name = map[int32]string{
	0: "FILESYSTEM_TYPE_BASIC",
	1: "FILESYSTEM_TYPE_FAKE",
}

var FilesystemType_value = map[string]int32{
	"FILESYSTEM_TYPE_BASIC": 0,
	"FILESYSTEM_TYPE_FAKE":  1,
}

func (FilesystemType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b556f45c4309ad5d, []int{0}
}

func init() {
	proto.RegisterEnum("fs.FilesystemType", FilesystemType_name, FilesystemType_value)
}

func init() { proto.RegisterFile("lib/fs/types.proto", fileDescriptor_b556f45c4309ad5d) }

var fileDescriptor_b556f45c4309ad5d = []byte{
	// 243 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xca, 0xc9, 0x4c, 0xd2,
	0x4f, 0x2b, 0xd6, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62,
	0x4a, 0x2b, 0x96, 0x52, 0x2e, 0x4a, 0x2d, 0xc8, 0x2f, 0xd6, 0x07, 0x0b, 0x24, 0x95, 0xa6, 0xe9,
	0xa7, 0xe7, 0xa7, 0xe7, 0x83, 0x39, 0x60, 0x16, 0x44, 0xa1, 0x56, 0x0d, 0x17, 0x9f, 0x5b, 0x66,
	0x4e, 0x6a, 0x71, 0x65, 0x71, 0x49, 0x6a, 0x6e, 0x48, 0x65, 0x41, 0xaa, 0x90, 0x11, 0x97, 0xa8,
	0x9b, 0xa7, 0x8f, 0x6b, 0x70, 0x64, 0x70, 0x88, 0xab, 0x6f, 0x7c, 0x48, 0x64, 0x80, 0x6b, 0xbc,
	0x93, 0x63, 0xb0, 0xa7, 0xb3, 0x00, 0x83, 0x94, 0x78, 0xd7, 0x5c, 0x05, 0x61, 0x54, 0xe5, 0x4e,
	0x89, 0xc5, 0x99, 0xc9, 0x42, 0x06, 0x5c, 0x22, 0xe8, 0x7a, 0xdc, 0x1c, 0xbd, 0x5d, 0x05, 0x18,
	0xa5, 0xc4, 0xba, 0xe6, 0x2a, 0x08, 0xa1, 0x6a, 0x71, 0x4b, 0xcc, 0x4e, 0x95, 0x62, 0x59, 0xb1,
	0x44, 0x8e, 0xc1, 0xc9, 0xf7, 0xc4, 0x43, 0x39, 0x86, 0x0b, 0x0f, 0xe5, 0x18, 0x4e, 0x3c, 0x92,
	0x63, 0xbc, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x05, 0x8f, 0xe5, 0x18, 0x2f, 0x3c,
	0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0x4a, 0x3b, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f,
	0x39, 0x3f, 0x57, 0xbf, 0x2a, 0x35, 0x29, 0xad, 0x28, 0xbf, 0xb8, 0x58, 0xbf, 0xb8, 0x32, 0x2f,
	0xb9, 0x24, 0x23, 0x33, 0x2f, 0x5d, 0x37, 0x37, 0x3f, 0x29, 0x33, 0x27, 0x55, 0x1f, 0x12, 0x02,
	0x49, 0x6c, 0x60, 0x3f, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x1a, 0xaa, 0xc5, 0x5a, 0x12,
	0x01, 0x00, 0x00,
}