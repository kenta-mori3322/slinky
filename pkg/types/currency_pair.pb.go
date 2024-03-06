// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: slinky/types/v1/currency_pair.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// CurrencyPair is the standard representation of a pair of assets, where one
// (Base) is priced in terms of the other (Quote)
type CurrencyPair struct {
	Base      string `protobuf:"bytes,1,opt,name=Base,proto3" json:"Base,omitempty"`
	Quote     string `protobuf:"bytes,2,opt,name=Quote,proto3" json:"Quote,omitempty"`
	Delimiter string `protobuf:"bytes,3,opt,name=Delimiter,proto3" json:"Delimiter,omitempty"`
}

func (m *CurrencyPair) Reset()      { *m = CurrencyPair{} }
func (*CurrencyPair) ProtoMessage() {}
func (*CurrencyPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfeb23a2dc8e58da, []int{0}
}
func (m *CurrencyPair) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CurrencyPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CurrencyPair.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CurrencyPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CurrencyPair.Merge(m, src)
}
func (m *CurrencyPair) XXX_Size() int {
	return m.Size()
}
func (m *CurrencyPair) XXX_DiscardUnknown() {
	xxx_messageInfo_CurrencyPair.DiscardUnknown(m)
}

var xxx_messageInfo_CurrencyPair proto.InternalMessageInfo

func (m *CurrencyPair) GetBase() string {
	if m != nil {
		return m.Base
	}
	return ""
}

func (m *CurrencyPair) GetQuote() string {
	if m != nil {
		return m.Quote
	}
	return ""
}

func (m *CurrencyPair) GetDelimiter() string {
	if m != nil {
		return m.Delimiter
	}
	return ""
}

func init() {
	proto.RegisterType((*CurrencyPair)(nil), "slinky.types.v1.CurrencyPair")
}

func init() {
	proto.RegisterFile("slinky/types/v1/currency_pair.proto", fileDescriptor_bfeb23a2dc8e58da)
}

var fileDescriptor_bfeb23a2dc8e58da = []byte{
	// 217 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2e, 0xce, 0xc9, 0xcc,
	0xcb, 0xae, 0xd4, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x2f, 0x33, 0xd4, 0x4f, 0x2e, 0x2d, 0x2a,
	0x4a, 0xcd, 0x4b, 0xae, 0x8c, 0x2f, 0x48, 0xcc, 0x2c, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0xe2, 0x87, 0x28, 0xd2, 0x03, 0x2b, 0xd2, 0x2b, 0x33, 0x94, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07,
	0xcb, 0xe9, 0x83, 0x58, 0x10, 0x65, 0x4a, 0x09, 0x5c, 0x3c, 0xce, 0x50, 0xdd, 0x01, 0x89, 0x99,
	0x45, 0x42, 0x42, 0x5c, 0x2c, 0x4e, 0x89, 0xc5, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41,
	0x60, 0xb6, 0x90, 0x08, 0x17, 0x6b, 0x60, 0x69, 0x7e, 0x49, 0xaa, 0x04, 0x13, 0x58, 0x10, 0xc2,
	0x11, 0x92, 0xe1, 0xe2, 0x74, 0x49, 0xcd, 0xc9, 0xcc, 0xcd, 0x2c, 0x49, 0x2d, 0x92, 0x60, 0x06,
	0xcb, 0x20, 0x04, 0xac, 0x38, 0x66, 0x2c, 0x90, 0x67, 0x68, 0xb8, 0xa3, 0xc0, 0xe0, 0x64, 0x77,
	0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7,
	0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51, 0x2a, 0xe9, 0x99, 0x25, 0x19, 0xa5,
	0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0xc5, 0xd9, 0x99, 0x05, 0xba, 0xb9, 0xa9, 0x65, 0xfa, 0x50,
	0xbf, 0x15, 0x64, 0xa7, 0x43, 0xfc, 0x97, 0xc4, 0x06, 0x76, 0xa8, 0x31, 0x20, 0x00, 0x00, 0xff,
	0xff, 0xae, 0xcc, 0xa7, 0x6b, 0xf6, 0x00, 0x00, 0x00,
}

func (m *CurrencyPair) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CurrencyPair) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CurrencyPair) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Delimiter) > 0 {
		i -= len(m.Delimiter)
		copy(dAtA[i:], m.Delimiter)
		i = encodeVarintCurrencyPair(dAtA, i, uint64(len(m.Delimiter)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Quote) > 0 {
		i -= len(m.Quote)
		copy(dAtA[i:], m.Quote)
		i = encodeVarintCurrencyPair(dAtA, i, uint64(len(m.Quote)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Base) > 0 {
		i -= len(m.Base)
		copy(dAtA[i:], m.Base)
		i = encodeVarintCurrencyPair(dAtA, i, uint64(len(m.Base)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCurrencyPair(dAtA []byte, offset int, v uint64) int {
	offset -= sovCurrencyPair(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CurrencyPair) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Base)
	if l > 0 {
		n += 1 + l + sovCurrencyPair(uint64(l))
	}
	l = len(m.Quote)
	if l > 0 {
		n += 1 + l + sovCurrencyPair(uint64(l))
	}
	l = len(m.Delimiter)
	if l > 0 {
		n += 1 + l + sovCurrencyPair(uint64(l))
	}
	return n
}

func sovCurrencyPair(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCurrencyPair(x uint64) (n int) {
	return sovCurrencyPair(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CurrencyPair) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCurrencyPair
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
			return fmt.Errorf("proto: CurrencyPair: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CurrencyPair: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Base", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCurrencyPair
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
				return ErrInvalidLengthCurrencyPair
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCurrencyPair
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Base = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Quote", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCurrencyPair
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
				return ErrInvalidLengthCurrencyPair
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCurrencyPair
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Quote = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Delimiter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCurrencyPair
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
				return ErrInvalidLengthCurrencyPair
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCurrencyPair
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Delimiter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCurrencyPair(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCurrencyPair
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
func skipCurrencyPair(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCurrencyPair
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
					return 0, ErrIntOverflowCurrencyPair
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
					return 0, ErrIntOverflowCurrencyPair
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
				return 0, ErrInvalidLengthCurrencyPair
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCurrencyPair
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCurrencyPair
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCurrencyPair        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCurrencyPair          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCurrencyPair = fmt.Errorf("proto: unexpected end of group")
)
