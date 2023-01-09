// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: oracle/network_constant.proto

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

type NetworkConstantCallData struct {
	Symbols    []string `protobuf:"bytes,1,rep,name=symbols,proto3" json:"symbols,omitempty"`
	Multiplier uint64   `protobuf:"varint,2,opt,name=multiplier,proto3" json:"multiplier,omitempty"`
}

func (m *NetworkConstantCallData) Reset()         { *m = NetworkConstantCallData{} }
func (m *NetworkConstantCallData) String() string { return proto.CompactTextString(m) }
func (*NetworkConstantCallData) ProtoMessage()    {}
func (*NetworkConstantCallData) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab72fcfbb1bb4158, []int{0}
}
func (m *NetworkConstantCallData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NetworkConstantCallData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NetworkConstantCallData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NetworkConstantCallData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkConstantCallData.Merge(m, src)
}
func (m *NetworkConstantCallData) XXX_Size() int {
	return m.Size()
}
func (m *NetworkConstantCallData) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkConstantCallData.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkConstantCallData proto.InternalMessageInfo

func (m *NetworkConstantCallData) GetSymbols() []string {
	if m != nil {
		return m.Symbols
	}
	return nil
}

func (m *NetworkConstantCallData) GetMultiplier() uint64 {
	if m != nil {
		return m.Multiplier
	}
	return 0
}

type NetworkConstantResult struct {
	Rates []uint64 `protobuf:"varint,1,rep,packed,name=rates,proto3" json:"rates,omitempty"`
}

func (m *NetworkConstantResult) Reset()         { *m = NetworkConstantResult{} }
func (m *NetworkConstantResult) String() string { return proto.CompactTextString(m) }
func (*NetworkConstantResult) ProtoMessage()    {}
func (*NetworkConstantResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab72fcfbb1bb4158, []int{1}
}
func (m *NetworkConstantResult) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NetworkConstantResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NetworkConstantResult.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NetworkConstantResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkConstantResult.Merge(m, src)
}
func (m *NetworkConstantResult) XXX_Size() int {
	return m.Size()
}
func (m *NetworkConstantResult) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkConstantResult.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkConstantResult proto.InternalMessageInfo

func (m *NetworkConstantResult) GetRates() []uint64 {
	if m != nil {
		return m.Rates
	}
	return nil
}

func init() {
	proto.RegisterType((*NetworkConstantCallData)(nil), "mandechain.oracle.NetworkConstantCallData")
	proto.RegisterType((*NetworkConstantResult)(nil), "mandechain.oracle.NetworkConstantResult")
}

func init() { proto.RegisterFile("oracle/network_constant.proto", fileDescriptor_ab72fcfbb1bb4158) }

var fileDescriptor_ab72fcfbb1bb4158 = []byte{
	// 207 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcd, 0x2f, 0x4a, 0x4c,
	0xce, 0x49, 0xd5, 0xcf, 0x4b, 0x2d, 0x29, 0xcf, 0x2f, 0xca, 0x8e, 0x4f, 0xce, 0xcf, 0x2b, 0x2e,
	0x49, 0xcc, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0xcc, 0x4d, 0xcc, 0x4b, 0x49,
	0x4d, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x83, 0xa8, 0x54, 0x0a, 0xe6, 0x12, 0xf7, 0x83, 0x28, 0x76,
	0x86, 0xaa, 0x75, 0x4e, 0xcc, 0xc9, 0x71, 0x49, 0x2c, 0x49, 0x14, 0x92, 0xe0, 0x62, 0x2f, 0xae,
	0xcc, 0x4d, 0xca, 0xcf, 0x29, 0x96, 0x60, 0x54, 0x60, 0xd6, 0xe0, 0x0c, 0x82, 0x71, 0x85, 0xe4,
	0xb8, 0xb8, 0x72, 0x4b, 0x73, 0x4a, 0x32, 0x0b, 0x72, 0x32, 0x53, 0x8b, 0x24, 0x98, 0x14, 0x18,
	0x35, 0x58, 0x82, 0x90, 0x44, 0x94, 0x74, 0xb9, 0x44, 0xd1, 0x0c, 0x0d, 0x4a, 0x2d, 0x2e, 0xcd,
	0x29, 0x11, 0x12, 0xe1, 0x62, 0x2d, 0x4a, 0x2c, 0x49, 0x85, 0x18, 0xc8, 0x12, 0x04, 0xe1, 0x38,
	0x99, 0x9c, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e,
	0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x14, 0xd8, 0xc1, 0xba,
	0x60, 0x17, 0xeb, 0x57, 0xe8, 0x43, 0x7d, 0x57, 0x52, 0x59, 0x90, 0x5a, 0x9c, 0xc4, 0x06, 0xf6,
	0x93, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xc0, 0x7a, 0x44, 0xc3, 0xf4, 0x00, 0x00, 0x00,
}

func (m *NetworkConstantCallData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NetworkConstantCallData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NetworkConstantCallData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Multiplier != 0 {
		i = encodeVarintNetworkConstant(dAtA, i, uint64(m.Multiplier))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Symbols) > 0 {
		for iNdEx := len(m.Symbols) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Symbols[iNdEx])
			copy(dAtA[i:], m.Symbols[iNdEx])
			i = encodeVarintNetworkConstant(dAtA, i, uint64(len(m.Symbols[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *NetworkConstantResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NetworkConstantResult) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NetworkConstantResult) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Rates) > 0 {
		dAtA2 := make([]byte, len(m.Rates)*10)
		var j1 int
		for _, num := range m.Rates {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintNetworkConstant(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintNetworkConstant(dAtA []byte, offset int, v uint64) int {
	offset -= sovNetworkConstant(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *NetworkConstantCallData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Symbols) > 0 {
		for _, s := range m.Symbols {
			l = len(s)
			n += 1 + l + sovNetworkConstant(uint64(l))
		}
	}
	if m.Multiplier != 0 {
		n += 1 + sovNetworkConstant(uint64(m.Multiplier))
	}
	return n
}

func (m *NetworkConstantResult) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Rates) > 0 {
		l = 0
		for _, e := range m.Rates {
			l += sovNetworkConstant(uint64(e))
		}
		n += 1 + sovNetworkConstant(uint64(l)) + l
	}
	return n
}

func sovNetworkConstant(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNetworkConstant(x uint64) (n int) {
	return sovNetworkConstant(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *NetworkConstantCallData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNetworkConstant
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
			return fmt.Errorf("proto: NetworkConstantCallData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NetworkConstantCallData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbols", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkConstant
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
				return ErrInvalidLengthNetworkConstant
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNetworkConstant
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbols = append(m.Symbols, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Multiplier", wireType)
			}
			m.Multiplier = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkConstant
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Multiplier |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipNetworkConstant(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNetworkConstant
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
func (m *NetworkConstantResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNetworkConstant
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
			return fmt.Errorf("proto: NetworkConstantResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NetworkConstantResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowNetworkConstant
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Rates = append(m.Rates, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowNetworkConstant
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthNetworkConstant
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthNetworkConstant
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Rates) == 0 {
					m.Rates = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowNetworkConstant
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Rates = append(m.Rates, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Rates", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipNetworkConstant(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNetworkConstant
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
func skipNetworkConstant(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNetworkConstant
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
					return 0, ErrIntOverflowNetworkConstant
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
					return 0, ErrIntOverflowNetworkConstant
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
				return 0, ErrInvalidLengthNetworkConstant
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNetworkConstant
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNetworkConstant
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNetworkConstant        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNetworkConstant          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNetworkConstant = fmt.Errorf("proto: unexpected end of group")
)
