// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lib/discover/local.proto

package discover

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_zebfross_syncthing_mobile_lib_protocol "github.com/zebfross/syncthing-mobile/lib/protocol"
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

type Announce struct {
	ID         github_com_zebfross_syncthing_mobile_lib_protocol.DeviceID `protobuf:"bytes,1,opt,name=id,proto3,customtype=github.com/zebfross/syncthing-mobile/lib/protocol.DeviceID" json:"id" xml:"id"`
	Addresses  []string                                                   `protobuf:"bytes,2,rep,name=addresses,proto3" json:"addresses" xml:"address"`
	InstanceID int64                                                      `protobuf:"varint,3,opt,name=instance_id,json=instanceId,proto3" json:"instanceId" xml:"instanceId"`
}

func (m *Announce) Reset()         { *m = Announce{} }
func (m *Announce) String() string { return proto.CompactTextString(m) }
func (*Announce) ProtoMessage()    {}
func (*Announce) Descriptor() ([]byte, []int) {
	return fileDescriptor_18afca46562fdaf4, []int{0}
}
func (m *Announce) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Announce) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Announce.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Announce) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Announce.Merge(m, src)
}
func (m *Announce) XXX_Size() int {
	return m.ProtoSize()
}
func (m *Announce) XXX_DiscardUnknown() {
	xxx_messageInfo_Announce.DiscardUnknown(m)
}

var xxx_messageInfo_Announce proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Announce)(nil), "discover.Announce")
}

func init() { proto.RegisterFile("lib/discover/local.proto", fileDescriptor_18afca46562fdaf4) }

var fileDescriptor_18afca46562fdaf4 = []byte{
	// 342 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x31, 0x6e, 0xf2, 0x30,
	0x18, 0x86, 0x93, 0x20, 0xfd, 0x02, 0xff, 0xad, 0x54, 0x65, 0x8a, 0x18, 0x6c, 0x94, 0x32, 0xb0,
	0x14, 0xab, 0xea, 0x86, 0xaa, 0x4a, 0x8d, 0xb2, 0x64, 0xa5, 0x5b, 0x17, 0x44, 0x6c, 0x13, 0x2c,
	0x05, 0x1b, 0xc5, 0x01, 0xd1, 0x9e, 0xa0, 0x63, 0xc5, 0x05, 0xda, 0xe3, 0x30, 0x66, 0xac, 0x3a,
	0x58, 0x22, 0xd9, 0x18, 0x39, 0x41, 0x85, 0x81, 0xc2, 0xd8, 0xed, 0xf3, 0xe3, 0xd7, 0xd6, 0xa3,
	0xf7, 0x03, 0x5e, 0xca, 0x63, 0x4c, 0xb9, 0x22, 0x72, 0xce, 0x32, 0x9c, 0x4a, 0x32, 0x4c, 0xbb,
	0xd3, 0x4c, 0xe6, 0xd2, 0xad, 0x1f, 0x69, 0xf3, 0x3a, 0x63, 0x53, 0xa9, 0xb0, 0xc1, 0xf1, 0x6c,
	0x84, 0x13, 0x99, 0x48, 0x73, 0x30, 0xd3, 0x3e, 0xde, 0x6c, 0xb0, 0x45, 0xbe, 0x1f, 0xfd, 0x0f,
	0x07, 0xd4, 0x1f, 0x85, 0x90, 0x33, 0x41, 0x98, 0x9b, 0x03, 0x87, 0x53, 0xcf, 0x6e, 0xd9, 0x9d,
	0x8b, 0x80, 0xae, 0x34, 0xb2, 0xbe, 0x35, 0xea, 0x25, 0x3c, 0x1f, 0xcf, 0xe2, 0x2e, 0x91, 0x13,
	0xfc, 0xca, 0xe2, 0x51, 0x26, 0x95, 0xc2, 0xea, 0x45, 0x90, 0x7c, 0xcc, 0x45, 0x72, 0x33, 0x91,
	0x31, 0x4f, 0x19, 0xde, 0x99, 0x99, 0x0f, 0x89, 0x4c, 0xbb, 0x21, 0x9b, 0x73, 0xc2, 0xa2, 0xb0,
	0xd4, 0xc8, 0x89, 0xc2, 0x8d, 0x46, 0x0e, 0xa7, 0x5b, 0x8d, 0xea, 0x8b, 0x49, 0xda, 0xf3, 0x39,
	0xf5, 0xdf, 0x8a, 0xb6, 0xbd, 0x2c, 0xda, 0x4e, 0x14, 0xf6, 0x1d, 0x4e, 0xdd, 0x7b, 0xd0, 0x18,
	0x52, 0x9a, 0x31, 0xa5, 0x98, 0xf2, 0x9c, 0x56, 0xad, 0xd3, 0x08, 0xe0, 0x46, 0xa3, 0x13, 0xdc,
	0x6a, 0x74, 0x69, 0xde, 0x1e, 0x88, 0xdf, 0x3f, 0xdd, 0xb9, 0x03, 0xf0, 0x9f, 0x0b, 0x95, 0x0f,
	0x05, 0x61, 0x03, 0x4e, 0xbd, 0x5a, 0xcb, 0xee, 0xd4, 0x82, 0x87, 0x52, 0x23, 0x10, 0x1d, 0xb0,
	0x51, 0x00, 0xc7, 0x50, 0xb4, 0x53, 0xb9, 0xda, 0xab, 0xfc, 0x22, 0x7f, 0x59, 0xb4, 0xcf, 0xf2,
	0xfd, 0xb3, 0x74, 0xf0, 0xb4, 0x5a, 0x43, 0xab, 0x58, 0x43, 0x6b, 0x55, 0x42, 0xbb, 0x28, 0xa1,
	0xfd, 0x5e, 0x41, 0xeb, 0xb3, 0x82, 0x76, 0x51, 0x41, 0xeb, 0xab, 0x82, 0xd6, 0xf3, 0xed, 0x9f,
	0x2b, 0x3a, 0xae, 0x29, 0xfe, 0x67, 0xca, 0xba, 0xfb, 0x09, 0x00, 0x00, 0xff, 0xff, 0x47, 0x40,
	0x39, 0x48, 0xd3, 0x01, 0x00, 0x00,
}

func (m *Announce) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Announce) MarshalTo(dAtA []byte) (int, error) {
	size := m.ProtoSize()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Announce) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.InstanceID != 0 {
		i = encodeVarintLocal(dAtA, i, uint64(m.InstanceID))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Addresses) > 0 {
		for iNdEx := len(m.Addresses) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Addresses[iNdEx])
			copy(dAtA[i:], m.Addresses[iNdEx])
			i = encodeVarintLocal(dAtA, i, uint64(len(m.Addresses[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size := m.ID.ProtoSize()
		i -= size
		if _, err := m.ID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLocal(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintLocal(dAtA []byte, offset int, v uint64) int {
	offset -= sovLocal(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Announce) ProtoSize() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ID.ProtoSize()
	n += 1 + l + sovLocal(uint64(l))
	if len(m.Addresses) > 0 {
		for _, s := range m.Addresses {
			l = len(s)
			n += 1 + l + sovLocal(uint64(l))
		}
	}
	if m.InstanceID != 0 {
		n += 1 + sovLocal(uint64(m.InstanceID))
	}
	return n
}

func sovLocal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLocal(x uint64) (n int) {
	return sovLocal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Announce) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLocal
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
			return fmt.Errorf("proto: Announce: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Announce: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLocal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthLocal
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthLocal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Addresses", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLocal
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
				return ErrInvalidLengthLocal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLocal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Addresses = append(m.Addresses, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InstanceID", wireType)
			}
			m.InstanceID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLocal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.InstanceID |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipLocal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLocal
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
func skipLocal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLocal
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
					return 0, ErrIntOverflowLocal
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
					return 0, ErrIntOverflowLocal
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
				return 0, ErrInvalidLengthLocal
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLocal
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLocal
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLocal        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLocal          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLocal = fmt.Errorf("proto: unexpected end of group")
)
