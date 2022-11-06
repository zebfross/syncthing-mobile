// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lib/config/ldaptransport.proto

package config

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/zebfross/syncthing-mobile/proto/ext"
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

type LDAPTransport int32

const (
	LDAPTransportPlain    LDAPTransport = 0
	LDAPTransportTLS      LDAPTransport = 2
	LDAPTransportStartTLS LDAPTransport = 3
)

var LDAPTransport_name = map[int32]string{
	0: "LDAP_TRANSPORT_PLAIN",
	2: "LDAP_TRANSPORT_TLS",
	3: "LDAP_TRANSPORT_START_TLS",
}

var LDAPTransport_value = map[string]int32{
	"LDAP_TRANSPORT_PLAIN":     0,
	"LDAP_TRANSPORT_TLS":       2,
	"LDAP_TRANSPORT_START_TLS": 3,
}

func (LDAPTransport) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_79795fc8505b82bf, []int{0}
}

func init() {
	proto.RegisterEnum("config.LDAPTransport", LDAPTransport_name, LDAPTransport_value)
}

func init() { proto.RegisterFile("lib/config/ldaptransport.proto", fileDescriptor_79795fc8505b82bf) }

var fileDescriptor_79795fc8505b82bf = []byte{
	// 290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcb, 0xc9, 0x4c, 0xd2,
	0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0xcf, 0x49, 0x49, 0x2c, 0x28, 0x29, 0x4a, 0xcc, 0x2b,
	0x2e, 0xc8, 0x2f, 0x2a, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xc8, 0x49, 0x29,
	0x17, 0xa5, 0x16, 0xe4, 0x17, 0xeb, 0x83, 0x05, 0x93, 0x4a, 0xd3, 0xf4, 0xd3, 0xf3, 0xd3, 0xf3,
	0xc1, 0x1c, 0x30, 0x0b, 0xa2, 0x58, 0x8a, 0x33, 0xb5, 0x02, 0xaa, 0x4f, 0xeb, 0x23, 0x23, 0x17,
	0xaf, 0x8f, 0x8b, 0x63, 0x40, 0x08, 0xcc, 0x3c, 0x21, 0x37, 0x2e, 0x11, 0x90, 0x40, 0x7c, 0x48,
	0x90, 0xa3, 0x5f, 0x70, 0x80, 0x7f, 0x50, 0x48, 0x7c, 0x80, 0x8f, 0xa3, 0xa7, 0x9f, 0x00, 0x83,
	0x94, 0x4e, 0xd7, 0x5c, 0x05, 0x21, 0x14, 0xc5, 0x01, 0x39, 0x89, 0x99, 0x79, 0x97, 0xfa, 0x54,
	0xb1, 0x88, 0x0a, 0x39, 0x70, 0x09, 0xa1, 0x99, 0x13, 0xe2, 0x13, 0x2c, 0xc0, 0x24, 0xa5, 0xd1,
	0x35, 0x57, 0x41, 0x00, 0x45, 0x7d, 0x88, 0x4f, 0xf0, 0xa5, 0x3e, 0x55, 0x0c, 0x31, 0xa1, 0x00,
	0x2e, 0x09, 0x34, 0x13, 0x82, 0x43, 0x1c, 0xa1, 0xe6, 0x30, 0x4b, 0x19, 0x75, 0xcd, 0x55, 0x10,
	0x45, 0xd1, 0x13, 0x5c, 0x92, 0x08, 0x33, 0x0c, 0xbb, 0x84, 0x14, 0xcb, 0x8a, 0x25, 0x72, 0x0c,
	0x4e, 0xe1, 0x27, 0x1e, 0xca, 0x31, 0x5c, 0x78, 0x28, 0xc7, 0x70, 0xe2, 0x91, 0x1c, 0xe3, 0x85,
	0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x2c, 0x78, 0x2c, 0xc7, 0x78, 0xe1, 0xb1, 0x1c, 0xc3,
	0x8d, 0xc7, 0x72, 0x0c, 0x51, 0xa6, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9,
	0xfa, 0x55, 0xa9, 0x49, 0x69, 0x45, 0xf9, 0xc5, 0xc5, 0xfa, 0xc5, 0x95, 0x79, 0xc9, 0x25, 0x19,
	0x99, 0x79, 0xe9, 0xba, 0xb9, 0xf9, 0x49, 0x99, 0x39, 0xa9, 0x10, 0x52, 0x1f, 0x11, 0x31, 0x49,
	0x6c, 0xe0, 0x30, 0x35, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x30, 0x02, 0x3d, 0xb4, 0xad, 0x01,
	0x00, 0x00,
}
