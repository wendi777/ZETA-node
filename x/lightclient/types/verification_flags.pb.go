// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lightclient/verification_flags.proto

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

// VerificationFlags is a structure containing information of weather a chain is enabled or not for block header verification
type EnabledChain struct {
	ChainId int64 `protobuf:"varint,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	Enabled bool  `protobuf:"varint,2,opt,name=enabled,proto3" json:"enabled,omitempty"`
}

func (m *EnabledChain) Reset()         { *m = EnabledChain{} }
func (m *EnabledChain) String() string { return proto.CompactTextString(m) }
func (*EnabledChain) ProtoMessage()    {}
func (*EnabledChain) Descriptor() ([]byte, []int) {
	return fileDescriptor_86eae6d737b3f8cc, []int{0}
}
func (m *EnabledChain) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EnabledChain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EnabledChain.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EnabledChain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnabledChain.Merge(m, src)
}
func (m *EnabledChain) XXX_Size() int {
	return m.Size()
}
func (m *EnabledChain) XXX_DiscardUnknown() {
	xxx_messageInfo_EnabledChain.DiscardUnknown(m)
}

var xxx_messageInfo_EnabledChain proto.InternalMessageInfo

func (m *EnabledChain) GetChainId() int64 {
	if m != nil {
		return m.ChainId
	}
	return 0
}

func (m *EnabledChain) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

type BlockHeaderVerification struct {
	EnabledChains []EnabledChain `protobuf:"bytes,1,rep,name=enabled_chains,json=enabledChains,proto3" json:"enabled_chains"`
}

func (m *BlockHeaderVerification) Reset()         { *m = BlockHeaderVerification{} }
func (m *BlockHeaderVerification) String() string { return proto.CompactTextString(m) }
func (*BlockHeaderVerification) ProtoMessage()    {}
func (*BlockHeaderVerification) Descriptor() ([]byte, []int) {
	return fileDescriptor_86eae6d737b3f8cc, []int{1}
}
func (m *BlockHeaderVerification) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BlockHeaderVerification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BlockHeaderVerification.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BlockHeaderVerification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockHeaderVerification.Merge(m, src)
}
func (m *BlockHeaderVerification) XXX_Size() int {
	return m.Size()
}
func (m *BlockHeaderVerification) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockHeaderVerification.DiscardUnknown(m)
}

var xxx_messageInfo_BlockHeaderVerification proto.InternalMessageInfo

func (m *BlockHeaderVerification) GetEnabledChains() []EnabledChain {
	if m != nil {
		return m.EnabledChains
	}
	return nil
}

func init() {
	proto.RegisterType((*EnabledChain)(nil), "zetachain.zetacore.lightclient.EnabledChain")
	proto.RegisterType((*BlockHeaderVerification)(nil), "zetachain.zetacore.lightclient.BlockHeaderVerification")
}

func init() {
	proto.RegisterFile("lightclient/verification_flags.proto", fileDescriptor_86eae6d737b3f8cc)
}

var fileDescriptor_86eae6d737b3f8cc = []byte{
	// 266 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xc9, 0xc9, 0x4c, 0xcf,
	0x28, 0x49, 0xce, 0xc9, 0x4c, 0xcd, 0x2b, 0xd1, 0x2f, 0x4b, 0x2d, 0xca, 0x4c, 0xcb, 0x4c, 0x4e,
	0x2c, 0xc9, 0xcc, 0xcf, 0x8b, 0x4f, 0xcb, 0x49, 0x4c, 0x2f, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x92, 0xab, 0x4a, 0x2d, 0x49, 0x4c, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x03, 0xb3, 0xf2, 0x8b,
	0x52, 0xf5, 0x90, 0x34, 0x4a, 0x89, 0xa4, 0xe7, 0xa7, 0xe7, 0x83, 0x95, 0xea, 0x83, 0x58, 0x10,
	0x5d, 0x4a, 0xce, 0x5c, 0x3c, 0xae, 0x79, 0x89, 0x49, 0x39, 0xa9, 0x29, 0xce, 0x20, 0xad, 0x42,
	0x92, 0x5c, 0x1c, 0x60, 0x33, 0xe2, 0x33, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x98, 0x83, 0xd8,
	0xc1, 0x7c, 0xcf, 0x14, 0x21, 0x09, 0x2e, 0xf6, 0x54, 0x88, 0x52, 0x09, 0x26, 0x05, 0x46, 0x0d,
	0x8e, 0x20, 0x18, 0x57, 0xa9, 0x84, 0x4b, 0xdc, 0x29, 0x27, 0x3f, 0x39, 0xdb, 0x23, 0x35, 0x31,
	0x25, 0xb5, 0x28, 0x0c, 0xc9, 0x85, 0x42, 0x91, 0x5c, 0x7c, 0x50, 0x55, 0xf1, 0x60, 0x73, 0x8a,
	0x25, 0x18, 0x15, 0x98, 0x35, 0xb8, 0x8d, 0x74, 0xf4, 0xf0, 0x3b, 0x57, 0x0f, 0xd9, 0x55, 0x4e,
	0x2c, 0x27, 0xee, 0xc9, 0x33, 0x04, 0xf1, 0xa6, 0x22, 0x89, 0x15, 0x3b, 0xf9, 0x9c, 0x78, 0x24,
	0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78,
	0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x51, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e,
	0x72, 0x7e, 0xae, 0x3e, 0xc8, 0x70, 0x5d, 0xb0, 0x3d, 0xfa, 0x30, 0x7b, 0xf4, 0x2b, 0xf4, 0x91,
	0x43, 0xb4, 0xa4, 0xb2, 0x20, 0xb5, 0x38, 0x89, 0x0d, 0x1c, 0x1e, 0xc6, 0x80, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xc2, 0x8d, 0x20, 0x39, 0x6d, 0x01, 0x00, 0x00,
}

func (m *EnabledChain) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EnabledChain) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EnabledChain) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Enabled {
		i--
		if m.Enabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.ChainId != 0 {
		i = encodeVarintVerificationFlags(dAtA, i, uint64(m.ChainId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *BlockHeaderVerification) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockHeaderVerification) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BlockHeaderVerification) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.EnabledChains) > 0 {
		for iNdEx := len(m.EnabledChains) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.EnabledChains[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintVerificationFlags(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintVerificationFlags(dAtA []byte, offset int, v uint64) int {
	offset -= sovVerificationFlags(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EnabledChain) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ChainId != 0 {
		n += 1 + sovVerificationFlags(uint64(m.ChainId))
	}
	if m.Enabled {
		n += 2
	}
	return n
}

func (m *BlockHeaderVerification) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.EnabledChains) > 0 {
		for _, e := range m.EnabledChains {
			l = e.Size()
			n += 1 + l + sovVerificationFlags(uint64(l))
		}
	}
	return n
}

func sovVerificationFlags(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozVerificationFlags(x uint64) (n int) {
	return sovVerificationFlags(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EnabledChain) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVerificationFlags
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
			return fmt.Errorf("proto: EnabledChain: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EnabledChain: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			m.ChainId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerificationFlags
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
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Enabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerificationFlags
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Enabled = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipVerificationFlags(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVerificationFlags
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
func (m *BlockHeaderVerification) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVerificationFlags
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
			return fmt.Errorf("proto: BlockHeaderVerification: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlockHeaderVerification: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnabledChains", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerificationFlags
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
				return ErrInvalidLengthVerificationFlags
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVerificationFlags
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EnabledChains = append(m.EnabledChains, EnabledChain{})
			if err := m.EnabledChains[len(m.EnabledChains)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVerificationFlags(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVerificationFlags
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
func skipVerificationFlags(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVerificationFlags
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
					return 0, ErrIntOverflowVerificationFlags
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
					return 0, ErrIntOverflowVerificationFlags
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
				return 0, ErrInvalidLengthVerificationFlags
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupVerificationFlags
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthVerificationFlags
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthVerificationFlags        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVerificationFlags          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupVerificationFlags = fmt.Errorf("proto: unexpected end of group")
)
