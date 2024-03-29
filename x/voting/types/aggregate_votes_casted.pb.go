// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: voting/aggregate_votes_casted.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

type AggregateVotesCasted struct {
	Index    string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Positive uint64 `protobuf:"varint,2,opt,name=positive,proto3" json:"positive,omitempty"`
	Negative uint64 `protobuf:"varint,3,opt,name=negative,proto3" json:"negative,omitempty"`
}

func (m *AggregateVotesCasted) Reset()         { *m = AggregateVotesCasted{} }
func (m *AggregateVotesCasted) String() string { return proto.CompactTextString(m) }
func (*AggregateVotesCasted) ProtoMessage()    {}
func (*AggregateVotesCasted) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7a75b62d847bf3b, []int{0}
}
func (m *AggregateVotesCasted) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AggregateVotesCasted) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AggregateVotesCasted.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AggregateVotesCasted) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AggregateVotesCasted.Merge(m, src)
}
func (m *AggregateVotesCasted) XXX_Size() int {
	return m.Size()
}
func (m *AggregateVotesCasted) XXX_DiscardUnknown() {
	xxx_messageInfo_AggregateVotesCasted.DiscardUnknown(m)
}

var xxx_messageInfo_AggregateVotesCasted proto.InternalMessageInfo

func (m *AggregateVotesCasted) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *AggregateVotesCasted) GetPositive() uint64 {
	if m != nil {
		return m.Positive
	}
	return 0
}

func (m *AggregateVotesCasted) GetNegative() uint64 {
	if m != nil {
		return m.Negative
	}
	return 0
}

func init() {
	proto.RegisterType((*AggregateVotesCasted)(nil), "mandechain.voting.AggregateVotesCasted")
}

func init() {
	proto.RegisterFile("voting/aggregate_votes_casted.proto", fileDescriptor_a7a75b62d847bf3b)
}

var fileDescriptor_a7a75b62d847bf3b = []byte{
	// 190 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2e, 0xcb, 0x2f, 0xc9,
	0xcc, 0x4b, 0xd7, 0x4f, 0x4c, 0x4f, 0x2f, 0x4a, 0x4d, 0x4f, 0x2c, 0x49, 0x8d, 0x2f, 0xcb, 0x2f,
	0x49, 0x2d, 0x8e, 0x4f, 0x4e, 0x2c, 0x2e, 0x49, 0x4d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x12, 0xcc, 0x4d, 0xcc, 0x4b, 0x49, 0x4d, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x83, 0xa8, 0x57, 0x4a,
	0xe1, 0x12, 0x71, 0x84, 0x69, 0x09, 0x03, 0xe9, 0x70, 0x06, 0x6b, 0x10, 0x12, 0xe1, 0x62, 0xcd,
	0xcc, 0x4b, 0x49, 0xad, 0x90, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x70, 0x84, 0xa4, 0xb8,
	0x38, 0x0a, 0xf2, 0x8b, 0x33, 0x4b, 0x32, 0xcb, 0x52, 0x25, 0x98, 0x14, 0x18, 0x35, 0x58, 0x82,
	0xe0, 0x7c, 0x90, 0x5c, 0x1e, 0xc8, 0x18, 0x90, 0x1c, 0x33, 0x44, 0x0e, 0xc6, 0x77, 0x32, 0x39,
	0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63,
	0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0x29, 0xb0, 0x93, 0x74, 0xc1, 0x6e,
	0xd2, 0xaf, 0xd0, 0x87, 0xfa, 0xa2, 0xa4, 0xb2, 0x20, 0xb5, 0x38, 0x89, 0x0d, 0xec, 0x6a, 0x63,
	0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfe, 0xc4, 0xaf, 0xa4, 0xdc, 0x00, 0x00, 0x00,
}

func (m *AggregateVotesCasted) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AggregateVotesCasted) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AggregateVotesCasted) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Negative != 0 {
		i = encodeVarintAggregateVotesCasted(dAtA, i, uint64(m.Negative))
		i--
		dAtA[i] = 0x18
	}
	if m.Positive != 0 {
		i = encodeVarintAggregateVotesCasted(dAtA, i, uint64(m.Positive))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintAggregateVotesCasted(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAggregateVotesCasted(dAtA []byte, offset int, v uint64) int {
	offset -= sovAggregateVotesCasted(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AggregateVotesCasted) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovAggregateVotesCasted(uint64(l))
	}
	if m.Positive != 0 {
		n += 1 + sovAggregateVotesCasted(uint64(m.Positive))
	}
	if m.Negative != 0 {
		n += 1 + sovAggregateVotesCasted(uint64(m.Negative))
	}
	return n
}

func sovAggregateVotesCasted(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAggregateVotesCasted(x uint64) (n int) {
	return sovAggregateVotesCasted(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AggregateVotesCasted) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAggregateVotesCasted
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
			return fmt.Errorf("proto: AggregateVotesCasted: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AggregateVotesCasted: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAggregateVotesCasted
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
				return ErrInvalidLengthAggregateVotesCasted
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAggregateVotesCasted
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Positive", wireType)
			}
			m.Positive = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAggregateVotesCasted
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Positive |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Negative", wireType)
			}
			m.Negative = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAggregateVotesCasted
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Negative |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAggregateVotesCasted(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAggregateVotesCasted
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
func skipAggregateVotesCasted(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAggregateVotesCasted
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
					return 0, ErrIntOverflowAggregateVotesCasted
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
					return 0, ErrIntOverflowAggregateVotesCasted
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
				return 0, ErrInvalidLengthAggregateVotesCasted
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAggregateVotesCasted
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAggregateVotesCasted
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAggregateVotesCasted        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAggregateVotesCasted          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAggregateVotesCasted = fmt.Errorf("proto: unexpected end of group")
)
