// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: slinky/module/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Given an authority + a set of CurrencyPairGeneses, the x/oracle module will
// check to see that the authority has permissions to update the set of
// CurrencyPairs tracked in the oracle, and add the given CurrencyPairs to be
// tracked in each VoteExtension
type MsgSetCurrencyPairs struct {
	// authority is the address of the account that is authorized to update the
	// x/oracle's CurrencyPairs
	Authority string `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
	// set of CurrencyPairs to be added to the module (+ prices if they are to be
	// set)
	CurrencyPairGenesis []*CurrencyPairGenesis `protobuf:"bytes,2,rep,name=currency_pair_genesis,json=currencyPairGenesis,proto3" json:"currency_pair_genesis,omitempty"`
}

func (m *MsgSetCurrencyPairs) Reset()         { *m = MsgSetCurrencyPairs{} }
func (m *MsgSetCurrencyPairs) String() string { return proto.CompactTextString(m) }
func (*MsgSetCurrencyPairs) ProtoMessage()    {}
func (*MsgSetCurrencyPairs) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6759b8cd4a1f8c9, []int{0}
}
func (m *MsgSetCurrencyPairs) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSetCurrencyPairs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSetCurrencyPairs.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSetCurrencyPairs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetCurrencyPairs.Merge(m, src)
}
func (m *MsgSetCurrencyPairs) XXX_Size() int {
	return m.Size()
}
func (m *MsgSetCurrencyPairs) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetCurrencyPairs.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetCurrencyPairs proto.InternalMessageInfo

func (m *MsgSetCurrencyPairs) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

func (m *MsgSetCurrencyPairs) GetCurrencyPairGenesis() []*CurrencyPairGenesis {
	if m != nil {
		return m.CurrencyPairGenesis
	}
	return nil
}

type MsgSetCurrencyPairsResponse struct {
}

func (m *MsgSetCurrencyPairsResponse) Reset()         { *m = MsgSetCurrencyPairsResponse{} }
func (m *MsgSetCurrencyPairsResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSetCurrencyPairsResponse) ProtoMessage()    {}
func (*MsgSetCurrencyPairsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6759b8cd4a1f8c9, []int{1}
}
func (m *MsgSetCurrencyPairsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSetCurrencyPairsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSetCurrencyPairsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSetCurrencyPairsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetCurrencyPairsResponse.Merge(m, src)
}
func (m *MsgSetCurrencyPairsResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSetCurrencyPairsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetCurrencyPairsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetCurrencyPairsResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgSetCurrencyPairs)(nil), "slinky.module.v1.MsgSetCurrencyPairs")
	proto.RegisterType((*MsgSetCurrencyPairsResponse)(nil), "slinky.module.v1.MsgSetCurrencyPairsResponse")
}

func init() { proto.RegisterFile("slinky/module/v1/tx.proto", fileDescriptor_c6759b8cd4a1f8c9) }

var fileDescriptor_c6759b8cd4a1f8c9 = []byte{
	// 383 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2c, 0xce, 0xc9, 0xcc,
	0xcb, 0xae, 0xd4, 0xcf, 0xcd, 0x4f, 0x29, 0xcd, 0x49, 0xd5, 0x2f, 0x33, 0xd4, 0x2f, 0xa9, 0xd0,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x80, 0x48, 0xe9, 0x41, 0xa4, 0xf4, 0xca, 0x0c, 0xa5,
	0xe4, 0x30, 0x14, 0xa7, 0xa7, 0xe6, 0xa5, 0x16, 0x67, 0x16, 0x43, 0x74, 0x48, 0x49, 0x26, 0xe7,
	0x17, 0xe7, 0xe6, 0x17, 0xc7, 0x83, 0x79, 0xfa, 0x10, 0x0e, 0x54, 0x4a, 0x1c, 0xc2, 0xd3, 0xcf,
	0x2d, 0x4e, 0x07, 0xe9, 0xcb, 0x2d, 0x4e, 0x87, 0x4a, 0x08, 0x26, 0xe6, 0x66, 0xe6, 0xe5, 0xeb,
	0x83, 0x49, 0xa8, 0x90, 0x48, 0x7a, 0x7e, 0x7a, 0x3e, 0xc4, 0x0c, 0x10, 0x0b, 0x22, 0xaa, 0x74,
	0x86, 0x91, 0x4b, 0xd8, 0xb7, 0x38, 0x3d, 0x38, 0xb5, 0xc4, 0xb9, 0xb4, 0xa8, 0x28, 0x35, 0x2f,
	0xb9, 0x32, 0x20, 0x31, 0xb3, 0xa8, 0x58, 0xc8, 0x8c, 0x8b, 0x33, 0xb1, 0xb4, 0x24, 0x23, 0xbf,
	0x28, 0xb3, 0xa4, 0x52, 0x82, 0x51, 0x81, 0x51, 0x83, 0xd3, 0x49, 0xe2, 0xd2, 0x16, 0x5d, 0x11,
	0xa8, 0xf5, 0x8e, 0x29, 0x29, 0x45, 0xa9, 0xc5, 0xc5, 0xc1, 0x25, 0x45, 0x99, 0x79, 0xe9, 0x41,
	0x08, 0xa5, 0x42, 0x1e, 0x5c, 0xa2, 0xc9, 0x50, 0x83, 0xe2, 0x0b, 0x12, 0x33, 0x8b, 0xe2, 0xa1,
	0x7e, 0x91, 0x60, 0x52, 0x60, 0xd6, 0xe0, 0x36, 0x12, 0xd1, 0x43, 0xb6, 0xc6, 0x1d, 0x22, 0x17,
	0x24, 0x9c, 0x8c, 0x29, 0x68, 0x65, 0xfc, 0x62, 0x81, 0x3c, 0x43, 0xd3, 0xf3, 0x0d, 0x5a, 0x08,
	0xd3, 0xbb, 0x9e, 0x6f, 0xd0, 0x92, 0x81, 0x86, 0x56, 0x85, 0x7e, 0x7e, 0x51, 0x62, 0x72, 0x4e,
	0xaa, 0xbe, 0x6f, 0x71, 0xba, 0x63, 0x4a, 0x4a, 0x48, 0x66, 0x72, 0x76, 0x6a, 0x91, 0x92, 0x2c,
	0x97, 0x34, 0x16, 0xdf, 0x04, 0xa5, 0x16, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x1a, 0x95, 0x72, 0x31,
	0xfb, 0x16, 0xa7, 0x0b, 0xa5, 0x73, 0xf1, 0x3b, 0xa6, 0xa4, 0x20, 0x2b, 0x11, 0x52, 0xd5, 0x43,
	0x8f, 0x17, 0x3d, 0x2c, 0x06, 0x49, 0xe9, 0x12, 0xa5, 0x0c, 0x66, 0x9f, 0x14, 0x6b, 0xc3, 0xf3,
	0x0d, 0x5a, 0x8c, 0x4e, 0xce, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91,
	0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0xa5,
	0x99, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x5f, 0x9c, 0x9d, 0x59, 0xa0,
	0x9b, 0x9b, 0x5a, 0xa6, 0x8f, 0xee, 0xc3, 0x92, 0xca, 0x82, 0xd4, 0xe2, 0x24, 0x36, 0x70, 0x84,
	0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x80, 0x1c, 0xcf, 0xfc, 0x5c, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// This method will be used explicitly by governance to update the set of
	// available CurrencyPairs. Given a set of CurrencyPairGenesis objects, update
	// the available currecy pairs in the module + their genesis prices (if they
	// exist)
	AddCurrencyPair(ctx context.Context, in *MsgSetCurrencyPairs, opts ...grpc.CallOption) (*MsgSetCurrencyPairsResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) AddCurrencyPair(ctx context.Context, in *MsgSetCurrencyPairs, opts ...grpc.CallOption) (*MsgSetCurrencyPairsResponse, error) {
	out := new(MsgSetCurrencyPairsResponse)
	err := c.cc.Invoke(ctx, "/slinky.module.v1.Msg/AddCurrencyPair", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// This method will be used explicitly by governance to update the set of
	// available CurrencyPairs. Given a set of CurrencyPairGenesis objects, update
	// the available currecy pairs in the module + their genesis prices (if they
	// exist)
	AddCurrencyPair(context.Context, *MsgSetCurrencyPairs) (*MsgSetCurrencyPairsResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) AddCurrencyPair(ctx context.Context, req *MsgSetCurrencyPairs) (*MsgSetCurrencyPairsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCurrencyPair not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_AddCurrencyPair_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSetCurrencyPairs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).AddCurrencyPair(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/slinky.module.v1.Msg/AddCurrencyPair",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AddCurrencyPair(ctx, req.(*MsgSetCurrencyPairs))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "slinky.module.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddCurrencyPair",
			Handler:    _Msg_AddCurrencyPair_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "slinky/module/v1/tx.proto",
}

func (m *MsgSetCurrencyPairs) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSetCurrencyPairs) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSetCurrencyPairs) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.CurrencyPairGenesis) > 0 {
		for iNdEx := len(m.CurrencyPairGenesis) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CurrencyPairGenesis[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSetCurrencyPairsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSetCurrencyPairsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSetCurrencyPairsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgSetCurrencyPairs) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.CurrencyPairGenesis) > 0 {
		for _, e := range m.CurrencyPairGenesis {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgSetCurrencyPairsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgSetCurrencyPairs) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSetCurrencyPairs: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSetCurrencyPairs: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authority", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Authority = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrencyPairGenesis", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CurrencyPairGenesis = append(m.CurrencyPairGenesis, &CurrencyPairGenesis{})
			if err := m.CurrencyPairGenesis[len(m.CurrencyPairGenesis)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgSetCurrencyPairsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSetCurrencyPairsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSetCurrencyPairsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
