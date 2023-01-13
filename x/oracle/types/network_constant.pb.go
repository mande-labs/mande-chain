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
	Repeat uint64 `protobuf:"varint,1,opt,name=repeat,proto3" json:"repeat,omitempty"`
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

func (m *NetworkConstantCallData) GetRepeat() uint64 {
	if m != nil {
		return m.Repeat
	}
	return 0
}

type NetworkConstantResult struct {
	Response string `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
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

func (m *NetworkConstantResult) GetResponse() string {
	if m != nil {
		return m.Response
	}
	return ""
}

func init() {
	proto.RegisterType((*NetworkConstantCallData)(nil), "mandechain.oracle.NetworkConstantCallData")
	proto.RegisterType((*NetworkConstantResult)(nil), "mandechain.oracle.NetworkConstantResult")
}

func init() { proto.RegisterFile("oracle/network_constant.proto", fileDescriptor_ab72fcfbb1bb4158) }

var fileDescriptor_ab72fcfbb1bb4158 = []byte{
	// 187 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcd, 0x2f, 0x4a, 0x4c,
	0xce, 0x49, 0xd5, 0xcf, 0x4b, 0x2d, 0x29, 0xcf, 0x2f, 0xca, 0x8e, 0x4f, 0xce, 0xcf, 0x2b, 0x2e,
	0x49, 0xcc, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0xcc, 0x4d, 0xcc, 0x4b, 0x49,
	0x4d, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x83, 0xa8, 0x54, 0x32, 0xe4, 0x12, 0xf7, 0x83, 0x28, 0x76,
	0x86, 0xaa, 0x75, 0x4e, 0xcc, 0xc9, 0x71, 0x49, 0x2c, 0x49, 0x14, 0x12, 0xe3, 0x62, 0x2b, 0x4a,
	0x2d, 0x48, 0x4d, 0x2c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x09, 0x82, 0xf2, 0x94, 0x8c, 0xb9,
	0x44, 0xd1, 0xb4, 0x04, 0xa5, 0x16, 0x97, 0xe6, 0x94, 0x08, 0x49, 0x71, 0x71, 0x14, 0xa5, 0x16,
	0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x82, 0xb5, 0x70, 0x06, 0xc1, 0xf9, 0x4e, 0x26, 0x27, 0x1e, 0xc9,
	0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e,
	0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0x25, 0x05, 0x76, 0x94, 0x2e, 0xd8, 0x55, 0xfa, 0x15,
	0xfa, 0x50, 0x1f, 0x94, 0x54, 0x16, 0xa4, 0x16, 0x27, 0xb1, 0x81, 0xdd, 0x6d, 0x0c, 0x08, 0x00,
	0x00, 0xff, 0xff, 0x13, 0x2d, 0xb8, 0x5a, 0xd8, 0x00, 0x00, 0x00,
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
	if m.Repeat != 0 {
		i = encodeVarintNetworkConstant(dAtA, i, uint64(m.Repeat))
		i--
		dAtA[i] = 0x8
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
	if len(m.Response) > 0 {
		i -= len(m.Response)
		copy(dAtA[i:], m.Response)
		i = encodeVarintNetworkConstant(dAtA, i, uint64(len(m.Response)))
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
	if m.Repeat != 0 {
		n += 1 + sovNetworkConstant(uint64(m.Repeat))
	}
	return n
}

func (m *NetworkConstantResult) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Response)
	if l > 0 {
		n += 1 + l + sovNetworkConstant(uint64(l))
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Repeat", wireType)
			}
			m.Repeat = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkConstant
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Repeat |= uint64(b&0x7F) << shift
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Response", wireType)
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
			m.Response = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
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