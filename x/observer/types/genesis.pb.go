// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: observer/genesis.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type GenesisState struct {
	Ballots           []*Ballot             `protobuf:"bytes,1,rep,name=ballots,proto3" json:"ballots,omitempty"`
	Observers         ObserverSet           `protobuf:"bytes,2,opt,name=observers,proto3" json:"observers"`
	NodeAccountList   []*NodeAccount        `protobuf:"bytes,3,rep,name=nodeAccountList,proto3" json:"nodeAccountList,omitempty"`
	CrosschainFlags   *CrosschainFlags      `protobuf:"bytes,4,opt,name=crosschain_flags,json=crosschainFlags,proto3" json:"crosschain_flags,omitempty"`
	Keygen            *Keygen               `protobuf:"bytes,6,opt,name=keygen,proto3" json:"keygen,omitempty"`
	LastObserverCount *LastObserverCount    `protobuf:"bytes,7,opt,name=last_observer_count,json=lastObserverCount,proto3" json:"last_observer_count,omitempty"`
	ChainParamsList   ChainParamsList       `protobuf:"bytes,8,opt,name=chain_params_list,json=chainParamsList,proto3" json:"chain_params_list"`
	Tss               *TSS                  `protobuf:"bytes,9,opt,name=tss,proto3" json:"tss,omitempty"`
	TssHistory        []TSS                 `protobuf:"bytes,10,rep,name=tss_history,json=tssHistory,proto3" json:"tss_history"`
	TssFundMigrators  []TssFundMigratorInfo `protobuf:"bytes,11,rep,name=tss_fund_migrators,json=tssFundMigrators,proto3" json:"tss_fund_migrators"`
	BlameList         []Blame               `protobuf:"bytes,12,rep,name=blame_list,json=blameList,proto3" json:"blame_list"`
	PendingNonces     []PendingNonces       `protobuf:"bytes,13,rep,name=pending_nonces,json=pendingNonces,proto3" json:"pending_nonces"`
	ChainNonces       []ChainNonces         `protobuf:"bytes,14,rep,name=chain_nonces,json=chainNonces,proto3" json:"chain_nonces"`
	NonceToCctx       []NonceToCctx         `protobuf:"bytes,15,rep,name=nonce_to_cctx,json=nonceToCctx,proto3" json:"nonce_to_cctx"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_15ea8c9d44da7399, []int{0}
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

func (m *GenesisState) GetBallots() []*Ballot {
	if m != nil {
		return m.Ballots
	}
	return nil
}

func (m *GenesisState) GetObservers() ObserverSet {
	if m != nil {
		return m.Observers
	}
	return ObserverSet{}
}

func (m *GenesisState) GetNodeAccountList() []*NodeAccount {
	if m != nil {
		return m.NodeAccountList
	}
	return nil
}

func (m *GenesisState) GetCrosschainFlags() *CrosschainFlags {
	if m != nil {
		return m.CrosschainFlags
	}
	return nil
}

func (m *GenesisState) GetKeygen() *Keygen {
	if m != nil {
		return m.Keygen
	}
	return nil
}

func (m *GenesisState) GetLastObserverCount() *LastObserverCount {
	if m != nil {
		return m.LastObserverCount
	}
	return nil
}

func (m *GenesisState) GetChainParamsList() ChainParamsList {
	if m != nil {
		return m.ChainParamsList
	}
	return ChainParamsList{}
}

func (m *GenesisState) GetTss() *TSS {
	if m != nil {
		return m.Tss
	}
	return nil
}

func (m *GenesisState) GetTssHistory() []TSS {
	if m != nil {
		return m.TssHistory
	}
	return nil
}

func (m *GenesisState) GetTssFundMigrators() []TssFundMigratorInfo {
	if m != nil {
		return m.TssFundMigrators
	}
	return nil
}

func (m *GenesisState) GetBlameList() []Blame {
	if m != nil {
		return m.BlameList
	}
	return nil
}

func (m *GenesisState) GetPendingNonces() []PendingNonces {
	if m != nil {
		return m.PendingNonces
	}
	return nil
}

func (m *GenesisState) GetChainNonces() []ChainNonces {
	if m != nil {
		return m.ChainNonces
	}
	return nil
}

func (m *GenesisState) GetNonceToCctx() []NonceToCctx {
	if m != nil {
		return m.NonceToCctx
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "zetachain.zetacore.observer.GenesisState")
}

func init() { proto.RegisterFile("observer/genesis.proto", fileDescriptor_15ea8c9d44da7399) }

var fileDescriptor_15ea8c9d44da7399 = []byte{
	// 627 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xdf, 0x6e, 0xd3, 0x30,
	0x14, 0xc6, 0x5b, 0x36, 0xed, 0x8f, 0xbb, 0xad, 0x9b, 0x19, 0x60, 0x6d, 0x90, 0x8d, 0x71, 0x53,
	0x21, 0x48, 0x50, 0xb9, 0x44, 0x5c, 0xb0, 0x4a, 0x1b, 0x13, 0x65, 0x8c, 0x74, 0x12, 0x12, 0x17,
	0x8b, 0x5c, 0xd7, 0x4d, 0x23, 0x52, 0xbb, 0x8a, 0x5d, 0xd4, 0xf2, 0x14, 0x3c, 0xd6, 0x2e, 0x77,
	0xc9, 0x15, 0x42, 0xed, 0x13, 0xf0, 0x06, 0x28, 0x8e, 0x9d, 0x34, 0xad, 0x14, 0x7a, 0x67, 0x7d,
	0xc7, 0xdf, 0xef, 0x1c, 0x1f, 0x1f, 0x1b, 0x3c, 0xe4, 0x6d, 0x41, 0xa3, 0xef, 0x34, 0x72, 0x7c,
	0xca, 0xa8, 0x08, 0x84, 0x3d, 0x88, 0xb8, 0xe4, 0xf0, 0xf0, 0x07, 0x95, 0x98, 0xf4, 0x70, 0xc0,
	0x6c, 0xb5, 0xe2, 0x11, 0xb5, 0xcd, 0xd6, 0x83, 0x7d, 0x9f, 0xfb, 0x5c, 0xed, 0x73, 0xe2, 0x55,
	0x62, 0x39, 0x78, 0x90, 0xa2, 0xda, 0x38, 0x0c, 0xb9, 0xd4, 0xf2, 0x7e, 0x26, 0x87, 0xb8, 0x4f,
	0xb5, 0x7a, 0x98, 0xaa, 0x2a, 0x89, 0xc7, 0x38, 0x23, 0x54, 0x27, 0x3f, 0x38, 0xca, 0x82, 0x11,
	0x17, 0x22, 0xd9, 0xd1, 0x0d, 0xb1, 0x2f, 0x16, 0x52, 0x7d, 0xa3, 0x63, 0x9f, 0xb2, 0x05, 0x28,
	0xe3, 0x1d, 0xea, 0x61, 0x42, 0xf8, 0x90, 0x99, 0x3a, 0x1e, 0xcf, 0x04, 0x19, 0xa1, 0x9e, 0xe4,
	0x1e, 0x21, 0x72, 0xa4, 0xa3, 0x8f, 0xd2, 0xa8, 0x59, 0x2c, 0xa4, 0x1a, 0xe0, 0x08, 0xf7, 0x4d,
	0x05, 0x4f, 0x32, 0x99, 0xb2, 0x4e, 0xc0, 0xfc, 0xfc, 0x09, 0x60, 0x1a, 0x96, 0xc2, 0x68, 0x4f,
	0x67, 0x35, 0xaf, 0x3b, 0x64, 0x1d, 0xe1, 0xf5, 0x03, 0x3f, 0xc2, 0x92, 0xeb, 0x64, 0x27, 0x7f,
	0x37, 0xc0, 0xd6, 0x79, 0x72, 0x0f, 0x2d, 0x89, 0x25, 0x85, 0x6f, 0xc1, 0x7a, 0xd2, 0x4c, 0x81,
	0xca, 0xc7, 0x2b, 0xb5, 0x4a, 0xfd, 0x99, 0x5d, 0x70, 0x31, 0xf6, 0xa9, 0xda, 0xeb, 0x1a, 0x0f,
	0x6c, 0x82, 0x4d, 0x13, 0x13, 0xe8, 0xde, 0x71, 0xb9, 0x56, 0xa9, 0xd7, 0x0a, 0x01, 0x9f, 0xf4,
	0xa2, 0x45, 0xe5, 0xe9, 0xea, 0xed, 0xef, 0xa3, 0x92, 0x9b, 0x01, 0xa0, 0x0b, 0xaa, 0x71, 0x5f,
	0xdf, 0x25, 0x6d, 0x6d, 0x06, 0x42, 0xa2, 0x15, 0x55, 0x54, 0x31, 0xf3, 0x32, 0xf3, 0xb8, 0xf3,
	0x00, 0xf8, 0x05, 0xec, 0xce, 0xdf, 0x31, 0x5a, 0x55, 0x85, 0xbe, 0x28, 0x84, 0x36, 0x52, 0xd3,
	0x59, 0xec, 0x71, 0xab, 0x24, 0x2f, 0xc0, 0x37, 0x60, 0x2d, 0x99, 0x0d, 0xb4, 0xa6, 0x70, 0xc5,
	0x8d, 0xfb, 0xa0, 0xb6, 0xba, 0xda, 0x02, 0x6f, 0xc0, 0xfd, 0x10, 0x0b, 0xe9, 0x99, 0xb8, 0xa7,
	0x0a, 0x46, 0xeb, 0x8a, 0x64, 0x17, 0x92, 0x9a, 0x58, 0x48, 0xd3, 0xc5, 0x86, 0x3a, 0xf3, 0x5e,
	0x38, 0x2f, 0xc1, 0x1b, 0xb0, 0x97, 0x1c, 0x38, 0x99, 0x29, 0x2f, 0x8c, 0x7b, 0xb9, 0xb1, 0xcc,
	0xb1, 0x63, 0xfd, 0x4a, 0x99, 0xe2, 0xf6, 0xe9, 0x3b, 0xaa, 0x92, 0xbc, 0x0c, 0xeb, 0x60, 0x45,
	0x0a, 0x81, 0x36, 0x15, 0xf1, 0xb8, 0x90, 0x78, 0xdd, 0x6a, 0xb9, 0xf1, 0x66, 0x78, 0x0e, 0x2a,
	0xf1, 0x5c, 0xf6, 0x02, 0x21, 0x79, 0x34, 0x46, 0x40, 0xdd, 0xec, 0x7f, 0xbd, 0xba, 0x02, 0x20,
	0x85, 0x78, 0x9f, 0x38, 0x61, 0x07, 0x40, 0x33, 0xe0, 0xe9, 0x7c, 0x0b, 0x54, 0x51, 0xbc, 0x57,
	0xc5, 0x3c, 0x21, 0xce, 0x86, 0xac, 0xf3, 0x51, 0x9b, 0x2e, 0x58, 0x97, 0x6b, 0xfe, 0xae, 0xcc,
	0x87, 0xe2, 0x72, 0x81, 0xfa, 0x4f, 0x92, 0xde, 0x6d, 0x29, 0xfa, 0x49, 0xf1, 0xe3, 0x88, 0xb7,
	0x9b, 0xa9, 0x56, 0x5e, 0x3d, 0x81, 0x3b, 0xf9, 0x27, 0x8c, 0xb6, 0x15, 0xec, 0x79, 0x21, 0xec,
	0x2a, 0xb1, 0x5c, 0x2a, 0x87, 0x86, 0x6e, 0x0f, 0x66, 0x45, 0xf8, 0x19, 0x6c, 0xcd, 0xfe, 0x6d,
	0x68, 0x67, 0x89, 0xb7, 0xa2, 0xee, 0x37, 0x07, 0xad, 0x90, 0x4c, 0x82, 0x2e, 0xd8, 0xce, 0x7d,
	0x5e, 0xa8, 0xba, 0xd4, 0xfb, 0x63, 0x84, 0x5e, 0xf3, 0x06, 0x91, 0x23, 0xc3, 0x64, 0x33, 0xd2,
	0xc5, 0xed, 0xc4, 0x2a, 0xdf, 0x4d, 0xac, 0xf2, 0x9f, 0x89, 0x55, 0xfe, 0x39, 0xb5, 0x4a, 0x77,
	0x53, 0xab, 0xf4, 0x6b, 0x6a, 0x95, 0xbe, 0x3a, 0x7e, 0x20, 0x7b, 0xc3, 0xb6, 0x4d, 0x78, 0xdf,
	0x89, 0xb1, 0x2f, 0x55, 0x06, 0xc7, 0x64, 0x70, 0x46, 0x4e, 0xf6, 0xa3, 0x8d, 0x07, 0x54, 0xb4,
	0xd7, 0xd4, 0x2f, 0xf6, 0xfa, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x62, 0x95, 0x52, 0x6b, 0x55,
	0x06, 0x00, 0x00,
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
	if len(m.NonceToCctx) > 0 {
		for iNdEx := len(m.NonceToCctx) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.NonceToCctx[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x7a
		}
	}
	if len(m.ChainNonces) > 0 {
		for iNdEx := len(m.ChainNonces) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ChainNonces[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x72
		}
	}
	if len(m.PendingNonces) > 0 {
		for iNdEx := len(m.PendingNonces) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PendingNonces[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x6a
		}
	}
	if len(m.BlameList) > 0 {
		for iNdEx := len(m.BlameList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BlameList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x62
		}
	}
	if len(m.TssFundMigrators) > 0 {
		for iNdEx := len(m.TssFundMigrators) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TssFundMigrators[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x5a
		}
	}
	if len(m.TssHistory) > 0 {
		for iNdEx := len(m.TssHistory) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TssHistory[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x52
		}
	}
	if m.Tss != nil {
		{
			size, err := m.Tss.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x4a
	}
	{
		size, err := m.ChainParamsList.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	if m.LastObserverCount != nil {
		{
			size, err := m.LastObserverCount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if m.Keygen != nil {
		{
			size, err := m.Keygen.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if m.CrosschainFlags != nil {
		{
			size, err := m.CrosschainFlags.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.NodeAccountList) > 0 {
		for iNdEx := len(m.NodeAccountList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.NodeAccountList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	{
		size, err := m.Observers.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Ballots) > 0 {
		for iNdEx := len(m.Ballots) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Ballots[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Ballots) > 0 {
		for _, e := range m.Ballots {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.Observers.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.NodeAccountList) > 0 {
		for _, e := range m.NodeAccountList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.CrosschainFlags != nil {
		l = m.CrosschainFlags.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.Keygen != nil {
		l = m.Keygen.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.LastObserverCount != nil {
		l = m.LastObserverCount.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = m.ChainParamsList.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if m.Tss != nil {
		l = m.Tss.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.TssHistory) > 0 {
		for _, e := range m.TssHistory {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.TssFundMigrators) > 0 {
		for _, e := range m.TssFundMigrators {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.BlameList) > 0 {
		for _, e := range m.BlameList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.PendingNonces) > 0 {
		for _, e := range m.PendingNonces {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ChainNonces) > 0 {
		for _, e := range m.ChainNonces {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.NonceToCctx) > 0 {
		for _, e := range m.NonceToCctx {
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
				return fmt.Errorf("proto: wrong wireType = %d for field Ballots", wireType)
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
			m.Ballots = append(m.Ballots, &Ballot{})
			if err := m.Ballots[len(m.Ballots)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Observers", wireType)
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
			if err := m.Observers.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeAccountList", wireType)
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
			m.NodeAccountList = append(m.NodeAccountList, &NodeAccount{})
			if err := m.NodeAccountList[len(m.NodeAccountList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CrosschainFlags", wireType)
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
			if m.CrosschainFlags == nil {
				m.CrosschainFlags = &CrosschainFlags{}
			}
			if err := m.CrosschainFlags.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Keygen", wireType)
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
			if m.Keygen == nil {
				m.Keygen = &Keygen{}
			}
			if err := m.Keygen.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastObserverCount", wireType)
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
			if m.LastObserverCount == nil {
				m.LastObserverCount = &LastObserverCount{}
			}
			if err := m.LastObserverCount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainParamsList", wireType)
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
			if err := m.ChainParamsList.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tss", wireType)
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
			if m.Tss == nil {
				m.Tss = &TSS{}
			}
			if err := m.Tss.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TssHistory", wireType)
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
			m.TssHistory = append(m.TssHistory, TSS{})
			if err := m.TssHistory[len(m.TssHistory)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TssFundMigrators", wireType)
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
			m.TssFundMigrators = append(m.TssFundMigrators, TssFundMigratorInfo{})
			if err := m.TssFundMigrators[len(m.TssFundMigrators)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlameList", wireType)
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
			m.BlameList = append(m.BlameList, Blame{})
			if err := m.BlameList[len(m.BlameList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PendingNonces", wireType)
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
			m.PendingNonces = append(m.PendingNonces, PendingNonces{})
			if err := m.PendingNonces[len(m.PendingNonces)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainNonces", wireType)
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
			m.ChainNonces = append(m.ChainNonces, ChainNonces{})
			if err := m.ChainNonces[len(m.ChainNonces)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NonceToCctx", wireType)
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
			m.NonceToCctx = append(m.NonceToCctx, NonceToCctx{})
			if err := m.NonceToCctx[len(m.NonceToCctx)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
