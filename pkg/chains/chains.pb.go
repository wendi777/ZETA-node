// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pkg/chains/chains.proto

package chains

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

type ReceiveStatus int32

const (
	ReceiveStatus_Created ReceiveStatus = 0
	ReceiveStatus_Success ReceiveStatus = 1
	ReceiveStatus_Failed  ReceiveStatus = 2
)

var ReceiveStatus_name = map[int32]string{
	0: "Created",
	1: "Success",
	2: "Failed",
}

var ReceiveStatus_value = map[string]int32{
	"Created": 0,
	"Success": 1,
	"Failed":  2,
}

func (x ReceiveStatus) String() string {
	return proto.EnumName(ReceiveStatus_name, int32(x))
}

func (ReceiveStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_37ad35e0488e8bbc, []int{0}
}

type ChainName int32

const (
	ChainName_empty           ChainName = 0
	ChainName_eth_mainnet     ChainName = 1
	ChainName_zeta_mainnet    ChainName = 2
	ChainName_btc_mainnet     ChainName = 3
	ChainName_polygon_mainnet ChainName = 4
	ChainName_bsc_mainnet     ChainName = 5
	ChainName_goerli_testnet  ChainName = 6
	ChainName_mumbai_testnet  ChainName = 7
	ChainName_ganache_testnet ChainName = 8
	ChainName_baobab_testnet  ChainName = 9
	ChainName_bsc_testnet     ChainName = 10
	ChainName_zeta_testnet    ChainName = 11
	ChainName_btc_testnet     ChainName = 12
	ChainName_sepolia_testnet ChainName = 13
	ChainName_goerli_localnet ChainName = 14
	ChainName_btc_regtest     ChainName = 15
	ChainName_amoy_testnet    ChainName = 16
)

var ChainName_name = map[int32]string{
	0:  "empty",
	1:  "eth_mainnet",
	2:  "zeta_mainnet",
	3:  "btc_mainnet",
	4:  "polygon_mainnet",
	5:  "bsc_mainnet",
	6:  "goerli_testnet",
	7:  "mumbai_testnet",
	8:  "ganache_testnet",
	9:  "baobab_testnet",
	10: "bsc_testnet",
	11: "zeta_testnet",
	12: "btc_testnet",
	13: "sepolia_testnet",
	14: "goerli_localnet",
	15: "btc_regtest",
	16: "amoy_testnet",
}

var ChainName_value = map[string]int32{
	"empty":           0,
	"eth_mainnet":     1,
	"zeta_mainnet":    2,
	"btc_mainnet":     3,
	"polygon_mainnet": 4,
	"bsc_mainnet":     5,
	"goerli_testnet":  6,
	"mumbai_testnet":  7,
	"ganache_testnet": 8,
	"baobab_testnet":  9,
	"bsc_testnet":     10,
	"zeta_testnet":    11,
	"btc_testnet":     12,
	"sepolia_testnet": 13,
	"goerli_localnet": 14,
	"btc_regtest":     15,
	"amoy_testnet":    16,
}

func (x ChainName) String() string {
	return proto.EnumName(ChainName_name, int32(x))
}

func (ChainName) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_37ad35e0488e8bbc, []int{1}
}

type Chain struct {
	ChainName ChainName `protobuf:"varint,1,opt,name=chain_name,json=chainName,proto3,enum=chains.ChainName" json:"chain_name,omitempty"`
	ChainId   int64     `protobuf:"varint,2,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
}

func (m *Chain) Reset()         { *m = Chain{} }
func (m *Chain) String() string { return proto.CompactTextString(m) }
func (*Chain) ProtoMessage()    {}
func (*Chain) Descriptor() ([]byte, []int) {
	return fileDescriptor_37ad35e0488e8bbc, []int{0}
}
func (m *Chain) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Chain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Chain.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Chain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Chain.Merge(m, src)
}
func (m *Chain) XXX_Size() int {
	return m.Size()
}
func (m *Chain) XXX_DiscardUnknown() {
	xxx_messageInfo_Chain.DiscardUnknown(m)
}

var xxx_messageInfo_Chain proto.InternalMessageInfo

func (m *Chain) GetChainName() ChainName {
	if m != nil {
		return m.ChainName
	}
	return ChainName_empty
}

func (m *Chain) GetChainId() int64 {
	if m != nil {
		return m.ChainId
	}
	return 0
}

func init() {
	proto.RegisterEnum("chains.ReceiveStatus", ReceiveStatus_name, ReceiveStatus_value)
	proto.RegisterEnum("chains.ChainName", ChainName_name, ChainName_value)
	proto.RegisterType((*Chain)(nil), "chains.Chain")
}

func init() { proto.RegisterFile("pkg/chains/chains.proto", fileDescriptor_37ad35e0488e8bbc) }

var fileDescriptor_37ad35e0488e8bbc = []byte{
	// 396 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x92, 0x4f, 0x6f, 0xd3, 0x30,
	0x14, 0xc0, 0xe3, 0x6e, 0x6d, 0x97, 0xd7, 0xad, 0x35, 0x06, 0x89, 0xb1, 0x43, 0x34, 0x71, 0x1a,
	0x93, 0x68, 0x10, 0x1c, 0xb9, 0x51, 0x09, 0x89, 0x0b, 0x87, 0x8e, 0x13, 0x97, 0xca, 0x71, 0x9e,
	0x1c, 0x8b, 0x38, 0x8e, 0x12, 0x07, 0xa9, 0x7c, 0x0a, 0x3e, 0x04, 0x07, 0x3e, 0xca, 0x8e, 0x3b,
	0x72, 0x44, 0xed, 0x17, 0x99, 0xec, 0x2c, 0xce, 0x29, 0xef, 0xfd, 0xfc, 0x7b, 0x7f, 0x14, 0x1b,
	0x5e, 0xd6, 0x3f, 0x64, 0x2a, 0x0a, 0xae, 0xaa, 0xf6, 0xe9, 0xb3, 0xae, 0x1b, 0x63, 0x0d, 0x9b,
	0xf5, 0xd9, 0xd5, 0x0b, 0x69, 0xa4, 0xf1, 0x28, 0x75, 0x51, 0x7f, 0xfa, 0xfa, 0x1b, 0x4c, 0x37,
	0xee, 0x9c, 0xbd, 0x03, 0xf0, 0xe2, 0xae, 0xe2, 0x1a, 0x2f, 0xc9, 0x35, 0xb9, 0x59, 0xbe, 0x7f,
	0xb6, 0x7e, 0xea, 0xe4, 0x95, 0xaf, 0x5c, 0xe3, 0x36, 0x16, 0x43, 0xc8, 0x5e, 0xc1, 0x59, 0x5f,
	0xa1, 0xf2, 0xcb, 0xc9, 0x35, 0xb9, 0x39, 0xd9, 0xce, 0x7d, 0xfe, 0x25, 0xbf, 0xfd, 0x08, 0x17,
	0x5b, 0x14, 0xa8, 0x7e, 0xe2, 0x9d, 0xe5, 0xb6, 0x6b, 0xd9, 0x02, 0xe6, 0x9b, 0x06, 0xb9, 0xc5,
	0x9c, 0x46, 0x2e, 0xb9, 0xeb, 0x84, 0xc0, 0xb6, 0xa5, 0x84, 0x01, 0xcc, 0x3e, 0x73, 0x55, 0x62,
	0x4e, 0x27, 0x57, 0xa7, 0x7f, 0xff, 0x24, 0xe4, 0xf6, 0x7e, 0x02, 0x71, 0x18, 0xc8, 0x62, 0x98,
	0xa2, 0xae, 0xed, 0x9e, 0x46, 0x6c, 0x05, 0x0b, 0xb4, 0xc5, 0x4e, 0x73, 0x55, 0x55, 0x68, 0x29,
	0x61, 0x14, 0xce, 0x7f, 0xa1, 0xe5, 0x81, 0x4c, 0x9c, 0x92, 0x59, 0x11, 0xc0, 0x09, 0x7b, 0x0e,
	0xab, 0xda, 0x94, 0x7b, 0x69, 0xaa, 0x00, 0x4f, 0xbd, 0xd5, 0x8e, 0xd6, 0x94, 0x31, 0x58, 0x4a,
	0x83, 0x4d, 0xa9, 0x76, 0x16, 0x5b, 0xeb, 0xd8, 0xcc, 0x31, 0xdd, 0xe9, 0x8c, 0x8f, 0x6c, 0xee,
	0xba, 0x49, 0x5e, 0x71, 0x51, 0x60, 0x80, 0x67, 0x4e, 0xcc, 0xb8, 0xc9, 0x78, 0x16, 0x58, 0x3c,
	0x4c, 0x18, 0x00, 0x84, 0x55, 0x07, 0xb2, 0x18, 0x56, 0x1d, 0xc0, 0xb9, 0x6b, 0xde, 0x62, 0x6d,
	0x4a, 0x35, 0x5a, 0x17, 0x7e, 0x62, 0xbf, 0x59, 0x69, 0x04, 0x2f, 0x1d, 0x5c, 0x0e, 0xa5, 0x0d,
	0x4a, 0x27, 0xd2, 0x95, 0xeb, 0xce, 0xb5, 0xd9, 0x87, 0x3a, 0xda, 0xff, 0xca, 0x4f, 0x9b, 0xfb,
	0x43, 0x42, 0x1e, 0x0e, 0x09, 0xf9, 0x7f, 0x48, 0xc8, 0xef, 0x63, 0x12, 0x3d, 0x1c, 0x93, 0xe8,
	0xdf, 0x31, 0x89, 0xbe, 0xbf, 0x91, 0xca, 0x16, 0x5d, 0xb6, 0x16, 0x46, 0xa7, 0x6e, 0xb1, 0xb7,
	0xfe, 0xea, 0x7c, 0x28, 0x4c, 0x83, 0xe9, 0xf8, 0x9a, 0xb2, 0x99, 0x7f, 0x29, 0x1f, 0x1e, 0x03,
	0x00, 0x00, 0xff, 0xff, 0xda, 0x8f, 0xd7, 0xf3, 0x62, 0x02, 0x00, 0x00,
}

func (m *Chain) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Chain) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Chain) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ChainId != 0 {
		i = encodeVarintChains(dAtA, i, uint64(m.ChainId))
		i--
		dAtA[i] = 0x10
	}
	if m.ChainName != 0 {
		i = encodeVarintChains(dAtA, i, uint64(m.ChainName))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintChains(dAtA []byte, offset int, v uint64) int {
	offset -= sovChains(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Chain) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ChainName != 0 {
		n += 1 + sovChains(uint64(m.ChainName))
	}
	if m.ChainId != 0 {
		n += 1 + sovChains(uint64(m.ChainId))
	}
	return n
}

func sovChains(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozChains(x uint64) (n int) {
	return sovChains(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Chain) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowChains
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
			return fmt.Errorf("proto: Chain: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Chain: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainName", wireType)
			}
			m.ChainName = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChains
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChainName |= ChainName(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			m.ChainId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChains
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChainId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipChains(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthChains
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
func skipChains(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowChains
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
					return 0, ErrIntOverflowChains
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
					return 0, ErrIntOverflowChains
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
				return 0, ErrInvalidLengthChains
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupChains
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthChains
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthChains        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowChains          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupChains = fmt.Errorf("proto: unexpected end of group")
)
