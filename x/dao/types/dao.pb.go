// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: onomyprotocol/dao/v1/dao.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
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

// FundTreasuryProposal details a dao fund treasury proposal.
type FundTreasuryProposal struct {
	Sender      string                                   `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Title       string                                   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string                                   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Amount      github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,4,rep,name=amount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"amount" yaml:"amount"`
}

func (m *FundTreasuryProposal) Reset()      { *m = FundTreasuryProposal{} }
func (*FundTreasuryProposal) ProtoMessage() {}
func (*FundTreasuryProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_763c2cafbce2acfd, []int{0}
}
func (m *FundTreasuryProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FundTreasuryProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FundTreasuryProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FundTreasuryProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FundTreasuryProposal.Merge(m, src)
}
func (m *FundTreasuryProposal) XXX_Size() int {
	return m.Size()
}
func (m *FundTreasuryProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_FundTreasuryProposal.DiscardUnknown(m)
}

var xxx_messageInfo_FundTreasuryProposal proto.InternalMessageInfo

// ExchangeWithTreasuryProposal details a dao exchange with treasury proposal.
type ExchangeWithTreasuryProposal struct {
	Sender      string              `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Title       string              `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string              `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	CoinsPairs  []CoinsExchangePair `protobuf:"bytes,4,rep,name=coins_pairs,json=coinsPairs,proto3" json:"coins_pairs" yaml:"coins_pairs"`
}

func (m *ExchangeWithTreasuryProposal) Reset()      { *m = ExchangeWithTreasuryProposal{} }
func (*ExchangeWithTreasuryProposal) ProtoMessage() {}
func (*ExchangeWithTreasuryProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_763c2cafbce2acfd, []int{1}
}
func (m *ExchangeWithTreasuryProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExchangeWithTreasuryProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExchangeWithTreasuryProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExchangeWithTreasuryProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExchangeWithTreasuryProposal.Merge(m, src)
}
func (m *ExchangeWithTreasuryProposal) XXX_Size() int {
	return m.Size()
}
func (m *ExchangeWithTreasuryProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_ExchangeWithTreasuryProposal.DiscardUnknown(m)
}

var xxx_messageInfo_ExchangeWithTreasuryProposal proto.InternalMessageInfo

// CoinsExchangePair is an ask/bid coins pair to exchange.
type CoinsExchangePair struct {
	CoinAsk types.Coin `protobuf:"bytes,1,opt,name=coin_ask,json=coinAsk,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coin" json:"coin_ask" yaml:"coin_ask"`
	CoinBid types.Coin `protobuf:"bytes,2,opt,name=coin_bid,json=coinBid,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coin" json:"coin_bid" yaml:"coin_bid"`
}

func (m *CoinsExchangePair) Reset()         { *m = CoinsExchangePair{} }
func (m *CoinsExchangePair) String() string { return proto.CompactTextString(m) }
func (*CoinsExchangePair) ProtoMessage()    {}
func (*CoinsExchangePair) Descriptor() ([]byte, []int) {
	return fileDescriptor_763c2cafbce2acfd, []int{2}
}
func (m *CoinsExchangePair) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CoinsExchangePair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CoinsExchangePair.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CoinsExchangePair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CoinsExchangePair.Merge(m, src)
}
func (m *CoinsExchangePair) XXX_Size() int {
	return m.Size()
}
func (m *CoinsExchangePair) XXX_DiscardUnknown() {
	xxx_messageInfo_CoinsExchangePair.DiscardUnknown(m)
}

var xxx_messageInfo_CoinsExchangePair proto.InternalMessageInfo

func (m *CoinsExchangePair) GetCoinAsk() types.Coin {
	if m != nil {
		return m.CoinAsk
	}
	return types.Coin{}
}

func (m *CoinsExchangePair) GetCoinBid() types.Coin {
	if m != nil {
		return m.CoinBid
	}
	return types.Coin{}
}

// FundAccountProposal details a dao fund account proposal.
type FundAccountProposal struct {
	Recipient   string                                   `protobuf:"bytes,1,opt,name=recipient,proto3" json:"recipient,omitempty"`
	Title       string                                   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string                                   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Amount      github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,4,rep,name=amount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"amount" yaml:"amount"`
}

func (m *FundAccountProposal) Reset()      { *m = FundAccountProposal{} }
func (*FundAccountProposal) ProtoMessage() {}
func (*FundAccountProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_763c2cafbce2acfd, []int{3}
}
func (m *FundAccountProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FundAccountProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FundAccountProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FundAccountProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FundAccountProposal.Merge(m, src)
}
func (m *FundAccountProposal) XXX_Size() int {
	return m.Size()
}
func (m *FundAccountProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_FundAccountProposal.DiscardUnknown(m)
}

var xxx_messageInfo_FundAccountProposal proto.InternalMessageInfo

func init() {
	proto.RegisterType((*FundTreasuryProposal)(nil), "onomyprotocol.dao.v1.FundTreasuryProposal")
	proto.RegisterType((*ExchangeWithTreasuryProposal)(nil), "onomyprotocol.dao.v1.ExchangeWithTreasuryProposal")
	proto.RegisterType((*CoinsExchangePair)(nil), "onomyprotocol.dao.v1.CoinsExchangePair")
	proto.RegisterType((*FundAccountProposal)(nil), "onomyprotocol.dao.v1.FundAccountProposal")
}

func init() { proto.RegisterFile("onomyprotocol/dao/v1/dao.proto", fileDescriptor_763c2cafbce2acfd) }

var fileDescriptor_763c2cafbce2acfd = []byte{
	// 485 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x54, 0x41, 0x8b, 0x13, 0x31,
	0x14, 0x9e, 0xec, 0x6a, 0x75, 0x53, 0x45, 0x1c, 0x8b, 0xd4, 0xb2, 0x64, 0xca, 0x5c, 0xb6, 0x08,
	0x4e, 0xa8, 0xde, 0xf6, 0xd6, 0x11, 0xc5, 0xe3, 0x52, 0x04, 0xc1, 0xcb, 0x92, 0x49, 0x42, 0x1b,
	0xda, 0x49, 0x86, 0x24, 0x2d, 0xed, 0x2f, 0xd0, 0xa3, 0x47, 0x8f, 0x3d, 0x8a, 0xbf, 0x64, 0x8f,
	0x7b, 0xf4, 0xb4, 0xca, 0xf4, 0x22, 0x5e, 0x04, 0x7f, 0x81, 0x24, 0x33, 0xee, 0x76, 0x55, 0xc4,
	0x1e, 0x3c, 0x78, 0x4a, 0xde, 0x7b, 0x79, 0xef, 0x7b, 0xdf, 0xfb, 0x1e, 0x81, 0x48, 0x49, 0x95,
	0x2f, 0x0b, 0xad, 0xac, 0xa2, 0x6a, 0x8a, 0x19, 0x51, 0x78, 0xde, 0x77, 0x47, 0xe2, 0x7d, 0x61,
	0xeb, 0x52, 0x3c, 0x71, 0x81, 0x79, 0xbf, 0xd3, 0x1a, 0xa9, 0x91, 0xf2, 0x4e, 0xec, 0x6e, 0xd5,
	0xdb, 0x0e, 0xa2, 0xca, 0xe4, 0xca, 0xe0, 0x8c, 0x18, 0x8e, 0xe7, 0xfd, 0x8c, 0x5b, 0xd2, 0xc7,
	0x54, 0x09, 0x59, 0xc5, 0xe3, 0x2f, 0x00, 0xb6, 0x9e, 0xce, 0x24, 0x7b, 0xae, 0x39, 0x31, 0x33,
	0xbd, 0x3c, 0xd2, 0xaa, 0x50, 0x86, 0x4c, 0xc3, 0xbb, 0xb0, 0x61, 0xb8, 0x64, 0x5c, 0xb7, 0x41,
	0x17, 0xf4, 0xf6, 0x86, 0xb5, 0x15, 0xb6, 0xe0, 0x55, 0x2b, 0xec, 0x94, 0xb7, 0x77, 0xbc, 0xbb,
	0x32, 0xc2, 0x2e, 0x6c, 0x32, 0x6e, 0xa8, 0x16, 0x85, 0x15, 0x4a, 0xb6, 0x77, 0x7d, 0x6c, 0xd3,
	0x15, 0x5a, 0xd8, 0x20, 0xb9, 0x9a, 0x49, 0xdb, 0xbe, 0xd2, 0xdd, 0xed, 0x35, 0x1f, 0xde, 0x4b,
	0xaa, 0xce, 0x12, 0xd7, 0x59, 0x52, 0x77, 0x96, 0x3c, 0x56, 0x42, 0xa6, 0x83, 0x93, 0xb3, 0x28,
	0xf8, 0x76, 0x16, 0xdd, 0x5c, 0x92, 0x7c, 0x7a, 0x18, 0x57, 0x69, 0xf1, 0xfb, 0x8f, 0x51, 0x6f,
	0x24, 0xec, 0x78, 0x96, 0x25, 0x54, 0xe5, 0xb8, 0xe6, 0x55, 0x1d, 0x0f, 0x0c, 0x9b, 0x60, 0xbb,
	0x2c, 0xb8, 0xf1, 0x15, 0xcc, 0xb0, 0xc6, 0x3a, 0xbc, 0xf1, 0x7a, 0x15, 0x05, 0x6f, 0x57, 0x51,
	0xf0, 0x79, 0x15, 0x05, 0x71, 0x09, 0xe0, 0xfe, 0x93, 0x05, 0x1d, 0x13, 0x39, 0xe2, 0x2f, 0x84,
	0x1d, 0xff, 0x73, 0xd2, 0x0c, 0x36, 0xdd, 0xac, 0xcd, 0x71, 0x41, 0x84, 0x36, 0x35, 0xf3, 0x83,
	0xe4, 0x77, 0xfa, 0x55, 0x8d, 0xff, 0xe8, 0xee, 0x88, 0x08, 0x9d, 0x76, 0xea, 0x39, 0x84, 0xd5,
	0x1c, 0x36, 0x2a, 0xc5, 0x43, 0xe8, 0x2d, 0xf7, 0xcc, 0xfc, 0x44, 0xf2, 0xd5, 0x0e, 0xbc, 0xfd,
	0x4b, 0xad, 0x70, 0x01, 0xaf, 0xbb, 0x8c, 0x63, 0x62, 0x26, 0x9e, 0xdb, 0x1f, 0x05, 0x48, 0x6b,
	0xe0, 0x5b, 0x17, 0xc0, 0x2e, 0xd1, 0x49, 0x70, 0xf0, 0x97, 0x12, 0x0c, 0xaf, 0xb9, 0xac, 0x81,
	0x99, 0x9c, 0x23, 0x67, 0x82, 0xf9, 0xf1, 0x6d, 0x8d, 0x9c, 0x09, 0xb6, 0x3d, 0x72, 0x2a, 0x58,
	0xfc, 0x15, 0xc0, 0x3b, 0x6e, 0xb7, 0x07, 0x94, 0xba, 0x65, 0x38, 0x57, 0x79, 0x1f, 0xee, 0x69,
	0x4e, 0x45, 0x21, 0xb8, 0xb4, 0xb5, 0xd0, 0x17, 0x8e, 0xff, 0x79, 0xc1, 0xd3, 0x67, 0xef, 0x4a,
	0x04, 0x4e, 0x4a, 0x04, 0x4e, 0x4b, 0x04, 0x3e, 0x95, 0x08, 0xbc, 0x59, 0xa3, 0xe0, 0x74, 0x8d,
	0x82, 0x0f, 0x6b, 0x14, 0xbc, 0xbc, 0xbf, 0x51, 0xfd, 0xf2, 0x17, 0xe3, 0x2d, 0xbc, 0xf0, 0x5f,
	0x8d, 0x47, 0xc9, 0x1a, 0x3e, 0xf6, 0xe8, 0x7b, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9b, 0x99, 0xd3,
	0x78, 0x8c, 0x04, 0x00, 0x00,
}

func (this *CoinsExchangePair) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CoinsExchangePair)
	if !ok {
		that2, ok := that.(CoinsExchangePair)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.CoinAsk.Equal(&that1.CoinAsk) {
		return false
	}
	if !this.CoinBid.Equal(&that1.CoinBid) {
		return false
	}
	return true
}
func (m *FundTreasuryProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FundTreasuryProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FundTreasuryProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Amount) > 0 {
		for iNdEx := len(m.Amount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Amount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDao(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintDao(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintDao(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintDao(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ExchangeWithTreasuryProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExchangeWithTreasuryProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ExchangeWithTreasuryProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.CoinsPairs) > 0 {
		for iNdEx := len(m.CoinsPairs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CoinsPairs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDao(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintDao(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintDao(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintDao(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CoinsExchangePair) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CoinsExchangePair) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CoinsExchangePair) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.CoinBid.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintDao(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.CoinAsk.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintDao(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *FundAccountProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FundAccountProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FundAccountProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Amount) > 0 {
		for iNdEx := len(m.Amount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Amount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDao(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintDao(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintDao(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Recipient) > 0 {
		i -= len(m.Recipient)
		copy(dAtA[i:], m.Recipient)
		i = encodeVarintDao(dAtA, i, uint64(len(m.Recipient)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDao(dAtA []byte, offset int, v uint64) int {
	offset -= sovDao(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *FundTreasuryProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovDao(uint64(l))
	}
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovDao(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovDao(uint64(l))
	}
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovDao(uint64(l))
		}
	}
	return n
}

func (m *ExchangeWithTreasuryProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovDao(uint64(l))
	}
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovDao(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovDao(uint64(l))
	}
	if len(m.CoinsPairs) > 0 {
		for _, e := range m.CoinsPairs {
			l = e.Size()
			n += 1 + l + sovDao(uint64(l))
		}
	}
	return n
}

func (m *CoinsExchangePair) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.CoinAsk.Size()
	n += 1 + l + sovDao(uint64(l))
	l = m.CoinBid.Size()
	n += 1 + l + sovDao(uint64(l))
	return n
}

func (m *FundAccountProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Recipient)
	if l > 0 {
		n += 1 + l + sovDao(uint64(l))
	}
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovDao(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovDao(uint64(l))
	}
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovDao(uint64(l))
		}
	}
	return n
}

func sovDao(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDao(x uint64) (n int) {
	return sovDao(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FundTreasuryProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDao
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
			return fmt.Errorf("proto: FundTreasuryProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FundTreasuryProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDao
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
				return ErrInvalidLengthDao
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDao
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDao
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
				return ErrInvalidLengthDao
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDao
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDao
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
				return ErrInvalidLengthDao
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDao
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDao
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
				return ErrInvalidLengthDao
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDao
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = append(m.Amount, types.Coin{})
			if err := m.Amount[len(m.Amount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDao(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDao
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
func (m *ExchangeWithTreasuryProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDao
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
			return fmt.Errorf("proto: ExchangeWithTreasuryProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExchangeWithTreasuryProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDao
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
				return ErrInvalidLengthDao
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDao
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDao
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
				return ErrInvalidLengthDao
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDao
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDao
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
				return ErrInvalidLengthDao
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDao
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CoinsPairs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDao
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
				return ErrInvalidLengthDao
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDao
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CoinsPairs = append(m.CoinsPairs, CoinsExchangePair{})
			if err := m.CoinsPairs[len(m.CoinsPairs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDao(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDao
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
func (m *CoinsExchangePair) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDao
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
			return fmt.Errorf("proto: CoinsExchangePair: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CoinsExchangePair: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CoinAsk", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDao
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
				return ErrInvalidLengthDao
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDao
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CoinAsk.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CoinBid", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDao
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
				return ErrInvalidLengthDao
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDao
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CoinBid.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDao(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDao
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
func (m *FundAccountProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDao
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
			return fmt.Errorf("proto: FundAccountProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FundAccountProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Recipient", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDao
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
				return ErrInvalidLengthDao
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDao
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Recipient = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDao
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
				return ErrInvalidLengthDao
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDao
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDao
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
				return ErrInvalidLengthDao
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDao
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDao
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
				return ErrInvalidLengthDao
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDao
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = append(m.Amount, types.Coin{})
			if err := m.Amount[len(m.Amount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDao(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDao
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
func skipDao(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDao
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
					return 0, ErrIntOverflowDao
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
					return 0, ErrIntOverflowDao
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
				return 0, ErrInvalidLengthDao
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDao
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDao
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDao        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDao          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDao = fmt.Errorf("proto: unexpected end of group")
)
