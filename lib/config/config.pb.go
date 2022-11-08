// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lib/config/config.proto

package config

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/zebfross/syncthing-mobile/proto/ext"
	io "io"
	math "math"
	math_bits "math/bits"
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

type Configuration struct {
	Version                  int                   `protobuf:"varint,1,opt,name=version,proto3,casttype=int" json:"version" xml:"version,attr"`
	Folders                  []FolderConfiguration `protobuf:"bytes,2,rep,name=folders,proto3" json:"folders" xml:"folder"`
	Devices                  []DeviceConfiguration `protobuf:"bytes,3,rep,name=devices,proto3" json:"devices" xml:"device"`
	GUI                      GUIConfiguration      `protobuf:"bytes,4,opt,name=gui,proto3" json:"gui" xml:"gui"`
	LDAP                     LDAPConfiguration     `protobuf:"bytes,5,opt,name=ldap,proto3" json:"ldap" xml:"ldap"`
	Options                  OptionsConfiguration  `protobuf:"bytes,6,opt,name=options,proto3" json:"options" xml:"options"`
	IgnoredDevices           []ObservedDevice      `protobuf:"bytes,7,rep,name=ignored_devices,json=ignoredDevices,proto3" json:"remoteIgnoredDevices" xml:"remoteIgnoredDevice"`
	DeprecatedPendingDevices []ObservedDevice      `protobuf:"bytes,8,rep,name=pending_devices,json=pendingDevices,proto3" json:"-" xml:"pendingDevice,omitempty"` // Deprecated: Do not use.
	Defaults                 Defaults              `protobuf:"bytes,9,opt,name=defaults,proto3" json:"defaults" xml:"defaults"`
}

func (m *Configuration) Reset()         { *m = Configuration{} }
func (m *Configuration) String() string { return proto.CompactTextString(m) }
func (*Configuration) ProtoMessage()    {}
func (*Configuration) Descriptor() ([]byte, []int) {
	return fileDescriptor_baadf209193dc627, []int{0}
}
func (m *Configuration) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Configuration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Configuration.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Configuration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Configuration.Merge(m, src)
}
func (m *Configuration) XXX_Size() int {
	return m.ProtoSize()
}
func (m *Configuration) XXX_DiscardUnknown() {
	xxx_messageInfo_Configuration.DiscardUnknown(m)
}

var xxx_messageInfo_Configuration proto.InternalMessageInfo

type Defaults struct {
	Folder  FolderConfiguration `protobuf:"bytes,1,opt,name=folder,proto3" json:"folder" xml:"folder"`
	Device  DeviceConfiguration `protobuf:"bytes,2,opt,name=device,proto3" json:"device" xml:"device"`
	Ignores Ignores             `protobuf:"bytes,3,opt,name=ignores,proto3" json:"ignores" xml:"ignores"`
}

func (m *Defaults) Reset()         { *m = Defaults{} }
func (m *Defaults) String() string { return proto.CompactTextString(m) }
func (*Defaults) ProtoMessage()    {}
func (*Defaults) Descriptor() ([]byte, []int) {
	return fileDescriptor_baadf209193dc627, []int{1}
}
func (m *Defaults) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Defaults) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Defaults.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Defaults) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Defaults.Merge(m, src)
}
func (m *Defaults) XXX_Size() int {
	return m.ProtoSize()
}
func (m *Defaults) XXX_DiscardUnknown() {
	xxx_messageInfo_Defaults.DiscardUnknown(m)
}

var xxx_messageInfo_Defaults proto.InternalMessageInfo

type Ignores struct {
	Lines []string `protobuf:"bytes,1,rep,name=lines,proto3" json:"lines" xml:"line"`
}

func (m *Ignores) Reset()         { *m = Ignores{} }
func (m *Ignores) String() string { return proto.CompactTextString(m) }
func (*Ignores) ProtoMessage()    {}
func (*Ignores) Descriptor() ([]byte, []int) {
	return fileDescriptor_baadf209193dc627, []int{2}
}
func (m *Ignores) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Ignores) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Ignores.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Ignores) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ignores.Merge(m, src)
}
func (m *Ignores) XXX_Size() int {
	return m.ProtoSize()
}
func (m *Ignores) XXX_DiscardUnknown() {
	xxx_messageInfo_Ignores.DiscardUnknown(m)
}

var xxx_messageInfo_Ignores proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Configuration)(nil), "config.Configuration")
	proto.RegisterType((*Defaults)(nil), "config.Defaults")
	proto.RegisterType((*Ignores)(nil), "config.Ignores")
}

func init() { proto.RegisterFile("lib/config/config.proto", fileDescriptor_baadf209193dc627) }

var fileDescriptor_baadf209193dc627 = []byte{
	// 721 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x94, 0xbf, 0x6f, 0xd3, 0x40,
	0x14, 0xc7, 0xe3, 0xa6, 0x8d, 0x9b, 0xeb, 0x2f, 0x64, 0x10, 0x75, 0xf9, 0xe1, 0x0b, 0xa7, 0x80,
	0x22, 0xd4, 0x36, 0x52, 0x59, 0x2a, 0x36, 0x42, 0x44, 0xa9, 0x40, 0xa2, 0x18, 0x15, 0x01, 0x0b,
	0x4a, 0xe2, 0x8b, 0x7b, 0x92, 0x63, 0x5b, 0xb6, 0x53, 0xb5, 0x6c, 0x8c, 0x6c, 0x88, 0xbf, 0x80,
	0x09, 0x89, 0xff, 0xa4, 0x5b, 0x33, 0x32, 0x9d, 0xd4, 0x66, 0xf3, 0xe8, 0x91, 0x09, 0xdd, 0x0f,
	0x3b, 0xb6, 0x6a, 0x60, 0x4a, 0xde, 0xfb, 0x7e, 0xdf, 0xe7, 0x4e, 0xef, 0xde, 0x33, 0x58, 0x77,
	0x48, 0xbf, 0x3d, 0xf0, 0xdc, 0x21, 0xb1, 0xe5, 0xcf, 0xb6, 0x1f, 0x78, 0x91, 0xa7, 0xd5, 0x44,
	0x74, 0xab, 0x99, 0x33, 0x0c, 0x3d, 0xc7, 0xc2, 0x81, 0x08, 0xc6, 0x41, 0x2f, 0x22, 0x9e, 0x2b,
	0xdc, 0x05, 0x97, 0x85, 0x8f, 0xc9, 0x00, 0x97, 0xb9, 0xee, 0xe5, 0x5c, 0xf6, 0x98, 0x94, 0x59,
	0x50, 0xce, 0xe2, 0x58, 0x3d, 0xbf, 0xcc, 0x73, 0x3f, 0xe7, 0xf1, 0x7c, 0x26, 0x84, 0x65, 0xb6,
	0x8d, 0xbc, 0xad, 0x1f, 0xe2, 0xe0, 0x18, 0x5b, 0x52, 0xaa, 0xe3, 0x93, 0x48, 0xfc, 0x45, 0x3f,
	0x54, 0xb0, 0xf2, 0x34, 0x5f, 0xad, 0x99, 0x40, 0x3d, 0xc6, 0x41, 0x48, 0x3c, 0x57, 0x57, 0x1a,
	0x4a, 0x6b, 0xa1, 0xb3, 0x1b, 0x53, 0x98, 0xa6, 0x12, 0x0a, 0xb5, 0x93, 0x91, 0xf3, 0x18, 0xc9,
	0x78, 0xb3, 0x17, 0x45, 0x01, 0xfa, 0x4d, 0x61, 0x95, 0xb8, 0x51, 0x7c, 0xde, 0x5c, 0xce, 0xe7,
	0xcd, 0xb4, 0x4a, 0x7b, 0x0b, 0x54, 0xd1, 0xbc, 0x50, 0x9f, 0x6b, 0x54, 0x5b, 0x4b, 0x3b, 0xb7,
	0xb7, 0x65, 0xb7, 0x9f, 0xf1, 0x74, 0xe1, 0x06, 0x1d, 0x78, 0x46, 0x61, 0x85, 0x1d, 0x2a, 0x6b,
	0x12, 0x0a, 0x97, 0xf9, 0xa1, 0x22, 0x46, 0x66, 0x2a, 0x30, 0xae, 0x68, 0x77, 0xa8, 0x57, 0x8b,
	0xdc, 0x2e, 0x4f, 0xff, 0x85, 0x2b, 0x6b, 0x32, 0xae, 0x88, 0x91, 0x99, 0x0a, 0x9a, 0x09, 0xaa,
	0xf6, 0x98, 0xe8, 0xf3, 0x0d, 0xa5, 0xb5, 0xb4, 0xa3, 0xa7, 0xcc, 0xbd, 0xc3, 0xfd, 0x22, 0xf0,
	0x01, 0x03, 0x5e, 0x52, 0x58, 0xdd, 0x3b, 0xdc, 0x8f, 0x29, 0x64, 0x35, 0x09, 0x85, 0x75, 0xce,
	0xb4, 0xc7, 0x04, 0x7d, 0x9b, 0x34, 0x99, 0x64, 0x32, 0x41, 0x7b, 0x0f, 0xe6, 0xd9, 0x8b, 0xea,
	0x0b, 0x1c, 0xba, 0x91, 0x42, 0x5f, 0x76, 0x9f, 0x1c, 0x14, 0xa9, 0x0f, 0x25, 0x75, 0x9e, 0x49,
	0x31, 0x85, 0xbc, 0x2c, 0xa1, 0x10, 0x70, 0x2e, 0x0b, 0x18, 0x98, 0xab, 0x26, 0xd7, 0xb4, 0x77,
	0x40, 0x95, 0x83, 0xa0, 0xd7, 0x38, 0xfd, 0x4e, 0x4a, 0x7f, 0x25, 0xd2, 0xc5, 0x03, 0x1a, 0x69,
	0x1f, 0x64, 0x51, 0x42, 0xe1, 0x0a, 0x67, 0xcb, 0x18, 0x99, 0xa9, 0xa2, 0xfd, 0x54, 0xc0, 0x1a,
	0xb1, 0x5d, 0x2f, 0xc0, 0xd6, 0xc7, 0xb4, 0xd3, 0x2a, 0xef, 0xf4, 0xcd, 0xec, 0x08, 0x39, 0x5b,
	0xa2, 0xe3, 0x9d, 0x23, 0x09, 0xbf, 0x11, 0xe0, 0x91, 0x17, 0xe1, 0x7d, 0x51, 0xdc, 0xcd, 0x3a,
	0xbe, 0xc1, 0x4f, 0x2a, 0x11, 0x51, 0x7c, 0xde, 0xbc, 0x5e, 0x92, 0x4f, 0xce, 0x9b, 0xa5, 0x2c,
	0x73, 0x95, 0x14, 0x62, 0xed, 0x8b, 0x02, 0xd6, 0x7c, 0xec, 0x5a, 0xc4, 0xb5, 0xb3, 0xbb, 0x2e,
	0xfe, 0xf3, 0xae, 0xcf, 0x65, 0xa7, 0xf5, 0x2e, 0xf6, 0x03, 0x3c, 0xe8, 0x45, 0xd8, 0x3a, 0x10,
	0x00, 0xc9, 0x8c, 0x29, 0x54, 0xb6, 0x12, 0x0a, 0xef, 0xf2, 0x4b, 0xfb, 0x79, 0x6d, 0xd3, 0x1b,
	0x91, 0x08, 0x8f, 0xfc, 0xe8, 0x14, 0xe9, 0x8a, 0xb9, 0x5a, 0xd0, 0x42, 0xed, 0x00, 0x2c, 0x5a,
	0x78, 0xd8, 0x1b, 0x3b, 0x51, 0xa8, 0xd7, 0xf9, 0x93, 0x5c, 0x9b, 0x4d, 0xa6, 0xc8, 0x77, 0x90,
	0xec, 0x54, 0xe6, 0x4c, 0x28, 0x5c, 0x95, 0xf3, 0x28, 0x12, 0xc8, 0xcc, 0x34, 0xf4, 0x79, 0x0e,
	0x2c, 0xa6, 0xa5, 0xda, 0x1b, 0x50, 0x13, 0x2b, 0xc0, 0x57, 0xf4, 0x3f, 0xeb, 0x64, 0xc8, 0x73,
	0x64, 0xc9, 0x95, 0x6d, 0x92, 0x79, 0x06, 0x15, 0x6d, 0xd3, 0xe7, 0x8a, 0xd0, 0xb2, 0x5d, 0xca,
	0xa0, 0xa2, 0xe4, 0xca, 0x2a, 0xc9, 0xbc, 0xf6, 0x02, 0xa8, 0xe2, 0x99, 0xd8, 0x86, 0x32, 0xea,
	0x5a, 0x4a, 0x15, 0xaf, 0x19, 0xce, 0xa6, 0x51, 0xfa, 0xb2, 0x69, 0x94, 0x31, 0x32, 0x53, 0x05,
	0xed, 0x02, 0x55, 0x56, 0x69, 0x5b, 0x60, 0xc1, 0x21, 0x2e, 0x0e, 0x75, 0xa5, 0x51, 0x6d, 0xd5,
	0x3b, 0xeb, 0x31, 0x85, 0x22, 0x31, 0x5b, 0x14, 0xe2, 0x62, 0x64, 0x8a, 0x64, 0xe7, 0xf5, 0xd9,
	0x85, 0x51, 0x99, 0x5c, 0x18, 0x95, 0xb3, 0x4b, 0x43, 0x99, 0x5c, 0x1a, 0xca, 0xd7, 0xa9, 0x51,
	0xf9, 0x3e, 0x35, 0x94, 0xc9, 0xd4, 0xa8, 0xfc, 0x9a, 0x1a, 0x95, 0x0f, 0x6d, 0x9b, 0x44, 0x47,
	0xe3, 0xfe, 0xf6, 0xc0, 0x1b, 0xb5, 0x3f, 0xe1, 0xfe, 0x30, 0xf0, 0xc2, 0xb0, 0x1d, 0x9e, 0xba,
	0x83, 0xe8, 0x88, 0xb8, 0xf6, 0xd6, 0xc8, 0xeb, 0x13, 0x07, 0xb7, 0x67, 0xdf, 0xd4, 0x7e, 0x8d,
	0x7f, 0x40, 0x1f, 0xfd, 0x09, 0x00, 0x00, 0xff, 0xff, 0x7c, 0xf4, 0x98, 0x35, 0x43, 0x06, 0x00,
	0x00,
}

func (m *Configuration) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Configuration) MarshalTo(dAtA []byte) (int, error) {
	size := m.ProtoSize()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Configuration) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Defaults.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintConfig(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	if len(m.DeprecatedPendingDevices) > 0 {
		for iNdEx := len(m.DeprecatedPendingDevices) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DeprecatedPendingDevices[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintConfig(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.IgnoredDevices) > 0 {
		for iNdEx := len(m.IgnoredDevices) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.IgnoredDevices[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintConfig(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	{
		size, err := m.Options.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintConfig(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size, err := m.LDAP.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintConfig(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size, err := m.GUI.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintConfig(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.Devices) > 0 {
		for iNdEx := len(m.Devices) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Devices[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintConfig(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Folders) > 0 {
		for iNdEx := len(m.Folders) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Folders[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintConfig(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Version != 0 {
		i = encodeVarintConfig(dAtA, i, uint64(m.Version))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Defaults) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Defaults) MarshalTo(dAtA []byte) (int, error) {
	size := m.ProtoSize()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Defaults) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Ignores.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintConfig(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.Device.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintConfig(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.Folder.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintConfig(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Ignores) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Ignores) MarshalTo(dAtA []byte) (int, error) {
	size := m.ProtoSize()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Ignores) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Lines) > 0 {
		for iNdEx := len(m.Lines) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Lines[iNdEx])
			copy(dAtA[i:], m.Lines[iNdEx])
			i = encodeVarintConfig(dAtA, i, uint64(len(m.Lines[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintConfig(dAtA []byte, offset int, v uint64) int {
	offset -= sovConfig(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Configuration) ProtoSize() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Version != 0 {
		n += 1 + sovConfig(uint64(m.Version))
	}
	if len(m.Folders) > 0 {
		for _, e := range m.Folders {
			l = e.ProtoSize()
			n += 1 + l + sovConfig(uint64(l))
		}
	}
	if len(m.Devices) > 0 {
		for _, e := range m.Devices {
			l = e.ProtoSize()
			n += 1 + l + sovConfig(uint64(l))
		}
	}
	l = m.GUI.ProtoSize()
	n += 1 + l + sovConfig(uint64(l))
	l = m.LDAP.ProtoSize()
	n += 1 + l + sovConfig(uint64(l))
	l = m.Options.ProtoSize()
	n += 1 + l + sovConfig(uint64(l))
	if len(m.IgnoredDevices) > 0 {
		for _, e := range m.IgnoredDevices {
			l = e.ProtoSize()
			n += 1 + l + sovConfig(uint64(l))
		}
	}
	if len(m.DeprecatedPendingDevices) > 0 {
		for _, e := range m.DeprecatedPendingDevices {
			l = e.ProtoSize()
			n += 1 + l + sovConfig(uint64(l))
		}
	}
	l = m.Defaults.ProtoSize()
	n += 1 + l + sovConfig(uint64(l))
	return n
}

func (m *Defaults) ProtoSize() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Folder.ProtoSize()
	n += 1 + l + sovConfig(uint64(l))
	l = m.Device.ProtoSize()
	n += 1 + l + sovConfig(uint64(l))
	l = m.Ignores.ProtoSize()
	n += 1 + l + sovConfig(uint64(l))
	return n
}

func (m *Ignores) ProtoSize() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Lines) > 0 {
		for _, s := range m.Lines {
			l = len(s)
			n += 1 + l + sovConfig(uint64(l))
		}
	}
	return n
}

func sovConfig(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozConfig(x uint64) (n int) {
	return sovConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Configuration) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Configuration: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Configuration: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Folders", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Folders = append(m.Folders, FolderConfiguration{})
			if err := m.Folders[len(m.Folders)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Devices", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Devices = append(m.Devices, DeviceConfiguration{})
			if err := m.Devices[len(m.Devices)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GUI", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.GUI.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LDAP", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LDAP.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Options", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Options.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IgnoredDevices", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IgnoredDevices = append(m.IgnoredDevices, ObservedDevice{})
			if err := m.IgnoredDevices[len(m.IgnoredDevices)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeprecatedPendingDevices", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DeprecatedPendingDevices = append(m.DeprecatedPendingDevices, ObservedDevice{})
			if err := m.DeprecatedPendingDevices[len(m.DeprecatedPendingDevices)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Defaults", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Defaults.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Defaults) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Defaults: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Defaults: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Folder", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Folder.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Device", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Device.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ignores", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Ignores.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Ignores) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Ignores: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Ignores: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lines", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Lines = append(m.Lines, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConfig
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthConfig
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupConfig
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthConfig
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthConfig        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConfig          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupConfig = fmt.Errorf("proto: unexpected end of group")
)
