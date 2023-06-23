// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: slinky/oracle/v1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// CurrencyPair is the standard representation of a pair of assets, where one
// (Base) is priced in terms of the other (Quote)
type CurrencyPair struct {
	Base  string `protobuf:"bytes,1,opt,name=Base,proto3" json:"Base,omitempty"`
	Quote string `protobuf:"bytes,2,opt,name=Quote,proto3" json:"Quote,omitempty"`
}

func (m *CurrencyPair) Reset()         { *m = CurrencyPair{} }
func (m *CurrencyPair) String() string { return proto.CompactTextString(m) }
func (*CurrencyPair) ProtoMessage()    {}
func (*CurrencyPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_de36a97821ccc13b, []int{0}
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

// QuotePrice is the representation of the aggregated prices for a CurrencyPair,
// where price represents the price of Base in terms of Quote
type QuotePrice struct {
	Price github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,1,opt,name=price,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"price"`
	// BlockTimestamp tracks the block height associated with this price update.
	// We include block timestamp alongside the price to ensure that smart
	// contracts and applications are not utilizing stale oracle prices
	BlockTimestamp time.Time `protobuf:"bytes,2,opt,name=block_timestamp,json=blockTimestamp,proto3,stdtime" json:"block_timestamp"`
	// BlockHeight is height of block mentioned above
	BlockHeight uint64 `protobuf:"varint,3,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
}

func (m *QuotePrice) Reset()         { *m = QuotePrice{} }
func (m *QuotePrice) String() string { return proto.CompactTextString(m) }
func (*QuotePrice) ProtoMessage()    {}
func (*QuotePrice) Descriptor() ([]byte, []int) {
	return fileDescriptor_de36a97821ccc13b, []int{1}
}
func (m *QuotePrice) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuotePrice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuotePrice.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuotePrice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuotePrice.Merge(m, src)
}
func (m *QuotePrice) XXX_Size() int {
	return m.Size()
}
func (m *QuotePrice) XXX_DiscardUnknown() {
	xxx_messageInfo_QuotePrice.DiscardUnknown(m)
}

var xxx_messageInfo_QuotePrice proto.InternalMessageInfo

func (m *QuotePrice) GetBlockTimestamp() time.Time {
	if m != nil {
		return m.BlockTimestamp
	}
	return time.Time{}
}

func (m *QuotePrice) GetBlockHeight() uint64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

// CurrencyPairGenesis is the information necessary for initialization of a
// CurrencyPair.
type CurrencyPairGenesis struct {
	// The CurrencyPair to be added to module state
	CurrencyPair CurrencyPair `protobuf:"bytes,1,opt,name=currency_pair,json=currencyPair,proto3" json:"currency_pair"`
	// A genesis price if one exists (note this will be empty, unless it results
	// from forking the state of this module)
	CurrencyPairPrice *QuotePrice `protobuf:"bytes,2,opt,name=currency_pair_price,json=currencyPairPrice,proto3" json:"currency_pair_price,omitempty"`
	// nonce is the nonce (number of updates) for the CP (same case as above,
	// likely 0 unless it results from fork of module)
	Nonce uint64 `protobuf:"varint,3,opt,name=nonce,proto3" json:"nonce,omitempty"`
}

func (m *CurrencyPairGenesis) Reset()         { *m = CurrencyPairGenesis{} }
func (m *CurrencyPairGenesis) String() string { return proto.CompactTextString(m) }
func (*CurrencyPairGenesis) ProtoMessage()    {}
func (*CurrencyPairGenesis) Descriptor() ([]byte, []int) {
	return fileDescriptor_de36a97821ccc13b, []int{2}
}
func (m *CurrencyPairGenesis) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CurrencyPairGenesis) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CurrencyPairGenesis.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CurrencyPairGenesis) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CurrencyPairGenesis.Merge(m, src)
}
func (m *CurrencyPairGenesis) XXX_Size() int {
	return m.Size()
}
func (m *CurrencyPairGenesis) XXX_DiscardUnknown() {
	xxx_messageInfo_CurrencyPairGenesis.DiscardUnknown(m)
}

var xxx_messageInfo_CurrencyPairGenesis proto.InternalMessageInfo

func (m *CurrencyPairGenesis) GetCurrencyPair() CurrencyPair {
	if m != nil {
		return m.CurrencyPair
	}
	return CurrencyPair{}
}

func (m *CurrencyPairGenesis) GetCurrencyPairPrice() *QuotePrice {
	if m != nil {
		return m.CurrencyPairPrice
	}
	return nil
}

func (m *CurrencyPairGenesis) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

// GenesisState is the genesis-state for the x/oracle module, it takes a set of
// predefined CurrencyPairGeneses
type GenesisState struct {
	CurrencyPairGenesis []CurrencyPairGenesis `protobuf:"bytes,1,rep,name=currency_pair_genesis,json=currencyPairGenesis,proto3" json:"currency_pair_genesis"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_de36a97821ccc13b, []int{3}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetCurrencyPairGenesis() []CurrencyPairGenesis {
	if m != nil {
		return m.CurrencyPairGenesis
	}
	return nil
}

func init() {
	proto.RegisterType((*CurrencyPair)(nil), "slinky.oracle.v1.CurrencyPair")
	proto.RegisterType((*QuotePrice)(nil), "slinky.oracle.v1.QuotePrice")
	proto.RegisterType((*CurrencyPairGenesis)(nil), "slinky.oracle.v1.CurrencyPairGenesis")
	proto.RegisterType((*GenesisState)(nil), "slinky.oracle.v1.GenesisState")
}

func init() { proto.RegisterFile("slinky/oracle/v1/genesis.proto", fileDescriptor_de36a97821ccc13b) }

var fileDescriptor_de36a97821ccc13b = []byte{
	// 462 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0xcd, 0x8e, 0x12, 0x41,
	0x10, 0xa6, 0x5d, 0x30, 0x5a, 0xe0, 0x5f, 0x83, 0x09, 0x12, 0x33, 0x20, 0x89, 0x06, 0x0f, 0xf4,
	0x64, 0xf1, 0xe2, 0xc1, 0x13, 0x7b, 0x50, 0x0e, 0x26, 0x6b, 0xeb, 0xc9, 0xcb, 0x64, 0x68, 0xdb,
	0xa1, 0x03, 0x33, 0x3d, 0x99, 0x6e, 0x88, 0xbc, 0xc5, 0x3e, 0x8c, 0x0f, 0xb1, 0xf1, 0xb4, 0xd9,
	0x93, 0xf1, 0xb0, 0x1a, 0x78, 0x11, 0x43, 0x57, 0xb3, 0xce, 0xaa, 0xd9, 0xd3, 0x54, 0xd5, 0xd7,
	0xf5, 0x55, 0x7d, 0x55, 0x35, 0x10, 0x98, 0x85, 0xca, 0xe6, 0xeb, 0x50, 0x17, 0xb1, 0x58, 0xc8,
	0x70, 0x75, 0x18, 0x26, 0x32, 0x93, 0x46, 0x19, 0x96, 0x17, 0xda, 0x6a, 0x7a, 0x1f, 0x71, 0x86,
	0x38, 0x5b, 0x1d, 0x76, 0x5a, 0x89, 0x4e, 0xb4, 0x03, 0xc3, 0x9d, 0x85, 0xef, 0x3a, 0xdd, 0x44,
	0xeb, 0x64, 0x21, 0x43, 0xe7, 0x4d, 0x97, 0x9f, 0x43, 0xab, 0x52, 0x69, 0x6c, 0x9c, 0xe6, 0xfe,
	0xc1, 0x23, 0xa1, 0x4d, 0xaa, 0x4d, 0x84, 0x99, 0xe8, 0x20, 0xd4, 0x7f, 0x09, 0x8d, 0xa3, 0x65,
	0x51, 0xc8, 0x4c, 0xac, 0x8f, 0x63, 0x55, 0x50, 0x0a, 0xd5, 0x71, 0x6c, 0x64, 0x9b, 0xf4, 0xc8,
	0xe0, 0x36, 0x77, 0x36, 0x6d, 0x41, 0xed, 0xdd, 0x52, 0x5b, 0xd9, 0xbe, 0xe1, 0x82, 0xe8, 0xf4,
	0xcf, 0x09, 0x80, 0xb3, 0x8e, 0x0b, 0x25, 0x24, 0xe5, 0x50, 0xcb, 0x77, 0x06, 0x66, 0x8e, 0x5f,
	0x9d, 0x5e, 0x74, 0x2b, 0x3f, 0x2e, 0xba, 0xcf, 0x12, 0x65, 0x67, 0xcb, 0x29, 0x13, 0x3a, 0xf5,
	0x85, 0xfd, 0x67, 0x68, 0x3e, 0xcd, 0x43, 0xbb, 0xce, 0xa5, 0x61, 0x93, 0xcc, 0x9e, 0x7f, 0x1d,
	0x82, 0xef, 0x6b, 0x92, 0x59, 0x8e, 0x54, 0xf4, 0x2d, 0xdc, 0x9b, 0x2e, 0xb4, 0x98, 0x47, 0x97,
	0x82, 0x5c, 0x0b, 0xf5, 0x51, 0x87, 0xa1, 0x64, 0xb6, 0x97, 0xcc, 0x3e, 0xec, 0x5f, 0x8c, 0x6f,
	0xed, 0x2a, 0x9f, 0xfc, 0xec, 0x12, 0x7e, 0xd7, 0x25, 0x5f, 0x22, 0xf4, 0x09, 0x34, 0x90, 0x6e,
	0x26, 0x55, 0x32, 0xb3, 0xed, 0x83, 0x1e, 0x19, 0x54, 0x79, 0xdd, 0xc5, 0xde, 0xb8, 0x50, 0xff,
	0x1b, 0x81, 0x66, 0x79, 0x1e, 0xaf, 0x71, 0x21, 0x74, 0x02, 0x77, 0x84, 0x0f, 0x47, 0x79, 0xac,
	0x0a, 0xa7, 0xb2, 0x3e, 0x0a, 0xd8, 0xdf, 0x2b, 0x62, 0xe5, 0xec, 0x71, 0x75, 0xd7, 0x0b, 0x6f,
	0x88, 0xf2, 0x84, 0x39, 0x34, 0xaf, 0x50, 0x45, 0x38, 0x36, 0x14, 0xf6, 0xf8, 0x5f, 0xc2, 0x3f,
	0x33, 0x76, 0x74, 0x84, 0x3f, 0x28, 0xd3, 0xe1, 0xf0, 0x5b, 0x50, 0xcb, 0x74, 0x26, 0xa4, 0x97,
	0x84, 0x4e, 0x5f, 0x43, 0xc3, 0xf7, 0xff, 0xde, 0xc6, 0x56, 0xd2, 0x08, 0x1e, 0x5e, 0xad, 0xec,
	0xcf, 0xad, 0x4d, 0x7a, 0x07, 0x83, 0xfa, 0xe8, 0xe9, 0xf5, 0x62, 0x3c, 0x95, 0xd7, 0xd4, 0x14,
	0xff, 0x81, 0x8e, 0x4e, 0x37, 0x01, 0x39, 0xdb, 0x04, 0xe4, 0xd7, 0x26, 0x20, 0x27, 0xdb, 0xa0,
	0x72, 0xb6, 0x0d, 0x2a, 0xdf, 0xb7, 0x41, 0xe5, 0xe3, 0xf3, 0xd2, 0x19, 0x98, 0xb9, 0xca, 0x87,
	0xa9, 0x5c, 0x85, 0xfe, 0xfc, 0xbf, 0xec, 0x7f, 0x00, 0x77, 0x0d, 0xd3, 0x9b, 0x6e, 0xa7, 0x2f,
	0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0x93, 0xe7, 0x95, 0x93, 0x1e, 0x03, 0x00, 0x00,
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
	if len(m.Quote) > 0 {
		i -= len(m.Quote)
		copy(dAtA[i:], m.Quote)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Quote)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Base) > 0 {
		i -= len(m.Base)
		copy(dAtA[i:], m.Base)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Base)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QuotePrice) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuotePrice) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuotePrice) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BlockHeight != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x18
	}
	n1, err1 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.BlockTimestamp, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.BlockTimestamp):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintGenesis(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x12
	{
		size := m.Price.Size()
		i -= size
		if _, err := m.Price.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *CurrencyPairGenesis) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CurrencyPairGenesis) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CurrencyPairGenesis) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Nonce != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.Nonce))
		i--
		dAtA[i] = 0x18
	}
	if m.CurrencyPairPrice != nil {
		{
			size, err := m.CurrencyPairPrice.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.CurrencyPair.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
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
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.Quote)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func (m *QuotePrice) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Price.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.BlockTimestamp)
	n += 1 + l + sovGenesis(uint64(l))
	if m.BlockHeight != 0 {
		n += 1 + sovGenesis(uint64(m.BlockHeight))
	}
	return n
}

func (m *CurrencyPairGenesis) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.CurrencyPair.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if m.CurrencyPairPrice != nil {
		l = m.CurrencyPairPrice.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.Nonce != 0 {
		n += 1 + sovGenesis(uint64(m.Nonce))
	}
	return n
}

func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.CurrencyPairGenesis) > 0 {
		for _, e := range m.CurrencyPairGenesis {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CurrencyPair) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
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
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Quote = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *QuotePrice) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: QuotePrice: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuotePrice: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Price.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockTimestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.BlockTimestamp, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *CurrencyPairGenesis) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: CurrencyPairGenesis: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CurrencyPairGenesis: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrencyPair", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CurrencyPair.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrencyPairPrice", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CurrencyPairPrice == nil {
				m.CurrencyPairPrice = &QuotePrice{}
			}
			if err := m.CurrencyPairPrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			m.Nonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Nonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrencyPairGenesis", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CurrencyPairGenesis = append(m.CurrencyPairGenesis, CurrencyPairGenesis{})
			if err := m.CurrencyPairGenesis[len(m.CurrencyPairGenesis)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
