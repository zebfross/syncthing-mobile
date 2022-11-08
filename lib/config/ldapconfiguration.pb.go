// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lib/config/ldapconfiguration.proto

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

type LDAPConfiguration struct {
	Address            string        `protobuf:"bytes,1,opt,name=address,proto3" json:"address" xml:"address,omitempty"`
	BindDN             string        `protobuf:"bytes,2,opt,name=bind_dn,json=bindDn,proto3" json:"bindDN" xml:"bindDN,omitempty"`
	Transport          LDAPTransport `protobuf:"varint,3,opt,name=transport,proto3,enum=config.LDAPTransport" json:"transport" xml:"transport,omitempty"`
	InsecureSkipVerify bool          `protobuf:"varint,4,opt,name=insecure_skip_verify,json=insecureSkipVerify,proto3" json:"insecureSkipVerify" xml:"insecureSkipVerify,omitempty" default:"false"`
	SearchBaseDN       string        `protobuf:"bytes,5,opt,name=search_base_dn,json=searchBaseDn,proto3" json:"searchBaseDN" xml:"searchBaseDN,omitempty"`
	SearchFilter       string        `protobuf:"bytes,6,opt,name=search_filter,json=searchFilter,proto3" json:"searchFilter" xml:"searchFilter,omitempty"`
}

func (m *LDAPConfiguration) Reset()         { *m = LDAPConfiguration{} }
func (m *LDAPConfiguration) String() string { return proto.CompactTextString(m) }
func (*LDAPConfiguration) ProtoMessage()    {}
func (*LDAPConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_9681ad7e41c73956, []int{0}
}
func (m *LDAPConfiguration) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LDAPConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LDAPConfiguration.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LDAPConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LDAPConfiguration.Merge(m, src)
}
func (m *LDAPConfiguration) XXX_Size() int {
	return m.ProtoSize()
}
func (m *LDAPConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_LDAPConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_LDAPConfiguration proto.InternalMessageInfo

func init() {
	proto.RegisterType((*LDAPConfiguration)(nil), "config.LDAPConfiguration")
}

func init() {
	proto.RegisterFile("lib/config/ldapconfiguration.proto", fileDescriptor_9681ad7e41c73956)
}

var fileDescriptor_9681ad7e41c73956 = []byte{
	// 510 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0x31, 0x6f, 0xd3, 0x40,
	0x14, 0x8e, 0x81, 0xba, 0xc4, 0x2a, 0x15, 0x35, 0x50, 0x42, 0x55, 0xf9, 0xa2, 0xc8, 0x43, 0x06,
	0x88, 0xa5, 0xb2, 0x95, 0xa9, 0xa6, 0x62, 0x40, 0x08, 0x81, 0x0b, 0x1d, 0x58, 0x22, 0x3b, 0x3e,
	0x27, 0xa7, 0xda, 0x77, 0xd6, 0xdd, 0xb9, 0x6a, 0xfa, 0x2b, 0xa0, 0xbf, 0xa0, 0x1b, 0x7f, 0xa5,
	0x5b, 0x3c, 0x32, 0x9d, 0xd4, 0x64, 0xf3, 0xe8, 0x91, 0x09, 0xe5, 0xec, 0x34, 0x76, 0x1a, 0xb1,
	0xbd, 0xf7, 0x7d, 0xef, 0x7d, 0xef, 0xbb, 0x7b, 0x7a, 0x5a, 0x27, 0x44, 0x9e, 0x35, 0x20, 0x38,
	0x40, 0x43, 0x2b, 0xf4, 0xdd, 0xb8, 0x08, 0x13, 0xea, 0x72, 0x44, 0x70, 0x2f, 0xa6, 0x84, 0x13,
	0x5d, 0x2d, 0xc0, 0x3d, 0x63, 0xa5, 0x96, 0x53, 0x17, 0xb3, 0x98, 0x50, 0x5e, 0xd4, 0xed, 0x35,
	0xe1, 0x45, 0x19, 0x76, 0x7e, 0xa9, 0xda, 0xce, 0xa7, 0xe3, 0xa3, 0x2f, 0xef, 0xab, 0x72, 0xfa,
	0x77, 0x6d, 0xd3, 0xf5, 0x7d, 0x0a, 0x19, 0x6b, 0x29, 0x6d, 0xa5, 0xdb, 0xb4, 0xdf, 0x65, 0x02,
	0x2c, 0xa0, 0x5c, 0x80, 0x97, 0x17, 0x51, 0x78, 0xd8, 0x29, 0xf3, 0xd7, 0x24, 0x42, 0x1c, 0x46,
	0x31, 0x1f, 0x77, 0xb2, 0x89, 0xb9, 0x73, 0x0f, 0x75, 0x16, 0x8d, 0x3a, 0xd1, 0x36, 0x3d, 0x84,
	0xfd, 0xbe, 0x8f, 0x5b, 0x0f, 0xa4, 0xec, 0xe9, 0x54, 0x00, 0xd5, 0x46, 0xd8, 0x3f, 0xfe, 0x9c,
	0x09, 0xa0, 0x7a, 0x32, 0xca, 0x05, 0xd8, 0x95, 0xfa, 0x45, 0x5a, 0x97, 0x7f, 0xba, 0x0a, 0xe6,
	0x13, 0xb3, 0xec, 0xbb, 0x4a, 0xcd, 0x52, 0xcb, 0x29, 0x10, 0xac, 0x9f, 0x6b, 0xcd, 0xbb, 0xb7,
	0xb7, 0x1e, 0xb6, 0x95, 0xee, 0xf6, 0xc1, 0x8b, 0x5e, 0xf1, 0x31, 0xbd, 0xf9, 0xab, 0xbf, 0x2d,
	0x48, 0xfb, 0x28, 0x13, 0x60, 0x59, 0x9b, 0x0b, 0xf0, 0x4a, 0x5a, 0xb8, 0x43, 0xea, 0x2e, 0x9e,
	0xad, 0xc1, 0x9d, 0x65, 0xbb, 0xfe, 0x5b, 0xd1, 0x9e, 0x23, 0xcc, 0xe0, 0x20, 0xa1, 0xb0, 0xcf,
	0xce, 0x50, 0xdc, 0x3f, 0x87, 0x14, 0x05, 0xe3, 0xd6, 0xa3, 0xb6, 0xd2, 0x7d, 0x6c, 0x27, 0x99,
	0x00, 0xfa, 0x82, 0x3f, 0x39, 0x43, 0xf1, 0xa9, 0x64, 0x73, 0x01, 0x0e, 0xe4, 0xd4, 0xfb, 0x54,
	0x65, 0x7c, 0xdb, 0x87, 0x81, 0x9b, 0x84, 0xfc, 0xb0, 0x13, 0xb8, 0x21, 0x83, 0x73, 0x3b, 0xfb,
	0xff, 0x6b, 0xf8, 0x3b, 0x31, 0x37, 0x64, 0xa5, 0xb3, 0x66, 0xa4, 0x7e, 0xad, 0x68, 0xdb, 0x0c,
	0xba, 0x74, 0x30, 0xea, 0x7b, 0x2e, 0x83, 0xf3, 0xd5, 0x6c, 0xc8, 0xd5, 0x5c, 0x4e, 0x05, 0xd8,
	0x3a, 0x91, 0x8c, 0xed, 0x32, 0x28, 0x17, 0xb4, 0xc5, 0x2a, 0x79, 0x2e, 0xc0, 0xbe, 0x74, 0x5b,
	0x05, 0xeb, 0xdf, 0xb4, 0xbb, 0x9e, 0xca, 0x27, 0x66, 0x4d, 0xe9, 0x2a, 0x35, 0x6b, 0x93, 0x9c,
	0x2a, 0x8b, 0x75, 0xa2, 0x3d, 0x29, 0x1d, 0x06, 0x28, 0xe4, 0x90, 0xb6, 0x54, 0x69, 0xf0, 0xe3,
	0xd2, 0xd0, 0x07, 0x89, 0xaf, 0x18, 0x2a, 0xc0, 0xb5, 0x86, 0x56, 0x29, 0xa7, 0xa6, 0x63, 0x7f,
	0xbd, 0xb9, 0x35, 0x1a, 0xe9, 0xad, 0xd1, 0xb8, 0x99, 0x1a, 0x4a, 0x3a, 0x35, 0x94, 0x9f, 0x33,
	0xa3, 0x71, 0x3d, 0x33, 0x94, 0x74, 0x66, 0x34, 0xfe, 0xcc, 0x8c, 0xc6, 0x0f, 0x6b, 0x88, 0xf8,
	0x28, 0xf1, 0x7a, 0x03, 0x12, 0x59, 0x97, 0xd0, 0x0b, 0x28, 0x61, 0xcc, 0x62, 0x63, 0x3c, 0xe0,
	0x23, 0x84, 0x87, 0x6f, 0x22, 0xe2, 0xa1, 0x10, 0x5a, 0xcb, 0x2b, 0xf4, 0x54, 0x79, 0x6d, 0x6f,
	0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0xb0, 0x2a, 0x9b, 0x7b, 0xc6, 0x03, 0x00, 0x00,
}

func (m *LDAPConfiguration) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LDAPConfiguration) MarshalTo(dAtA []byte) (int, error) {
	size := m.ProtoSize()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LDAPConfiguration) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SearchFilter) > 0 {
		i -= len(m.SearchFilter)
		copy(dAtA[i:], m.SearchFilter)
		i = encodeVarintLdapconfiguration(dAtA, i, uint64(len(m.SearchFilter)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.SearchBaseDN) > 0 {
		i -= len(m.SearchBaseDN)
		copy(dAtA[i:], m.SearchBaseDN)
		i = encodeVarintLdapconfiguration(dAtA, i, uint64(len(m.SearchBaseDN)))
		i--
		dAtA[i] = 0x2a
	}
	if m.InsecureSkipVerify {
		i--
		if m.InsecureSkipVerify {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if m.Transport != 0 {
		i = encodeVarintLdapconfiguration(dAtA, i, uint64(m.Transport))
		i--
		dAtA[i] = 0x18
	}
	if len(m.BindDN) > 0 {
		i -= len(m.BindDN)
		copy(dAtA[i:], m.BindDN)
		i = encodeVarintLdapconfiguration(dAtA, i, uint64(len(m.BindDN)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintLdapconfiguration(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintLdapconfiguration(dAtA []byte, offset int, v uint64) int {
	offset -= sovLdapconfiguration(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LDAPConfiguration) ProtoSize() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovLdapconfiguration(uint64(l))
	}
	l = len(m.BindDN)
	if l > 0 {
		n += 1 + l + sovLdapconfiguration(uint64(l))
	}
	if m.Transport != 0 {
		n += 1 + sovLdapconfiguration(uint64(m.Transport))
	}
	if m.InsecureSkipVerify {
		n += 2
	}
	l = len(m.SearchBaseDN)
	if l > 0 {
		n += 1 + l + sovLdapconfiguration(uint64(l))
	}
	l = len(m.SearchFilter)
	if l > 0 {
		n += 1 + l + sovLdapconfiguration(uint64(l))
	}
	return n
}

func sovLdapconfiguration(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLdapconfiguration(x uint64) (n int) {
	return sovLdapconfiguration(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LDAPConfiguration) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLdapconfiguration
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
			return fmt.Errorf("proto: LDAPConfiguration: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LDAPConfiguration: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLdapconfiguration
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
				return ErrInvalidLengthLdapconfiguration
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLdapconfiguration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BindDN", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLdapconfiguration
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
				return ErrInvalidLengthLdapconfiguration
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLdapconfiguration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BindDN = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Transport", wireType)
			}
			m.Transport = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLdapconfiguration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Transport |= LDAPTransport(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InsecureSkipVerify", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLdapconfiguration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.InsecureSkipVerify = bool(v != 0)
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SearchBaseDN", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLdapconfiguration
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
				return ErrInvalidLengthLdapconfiguration
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLdapconfiguration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SearchBaseDN = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SearchFilter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLdapconfiguration
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
				return ErrInvalidLengthLdapconfiguration
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLdapconfiguration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SearchFilter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLdapconfiguration(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLdapconfiguration
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
func skipLdapconfiguration(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLdapconfiguration
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
					return 0, ErrIntOverflowLdapconfiguration
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
					return 0, ErrIntOverflowLdapconfiguration
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
				return 0, ErrInvalidLengthLdapconfiguration
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLdapconfiguration
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLdapconfiguration
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLdapconfiguration        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLdapconfiguration          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLdapconfiguration = fmt.Errorf("proto: unexpected end of group")
)