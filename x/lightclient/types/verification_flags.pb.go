// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: zetachain/zetacore/lightclient/verification_flags.proto

package types

import (
	fmt "fmt"
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

// VerificationFlags is a structure containing information which chain types are
// enabled for block header verification
type VerificationFlags struct {
	EthTypeChainEnabled bool `protobuf:"varint,1,opt,name=ethTypeChainEnabled,proto3" json:"ethTypeChainEnabled,omitempty"`
	BtcTypeChainEnabled bool `protobuf:"varint,2,opt,name=btcTypeChainEnabled,proto3" json:"btcTypeChainEnabled,omitempty"`
}

func (m *VerificationFlags) Reset()         { *m = VerificationFlags{} }
func (m *VerificationFlags) String() string { return proto.CompactTextString(m) }
func (*VerificationFlags) ProtoMessage()    {}
func (*VerificationFlags) Descriptor() ([]byte, []int) {
	return fileDescriptor_bcf482283292221c, []int{0}
}
func (m *VerificationFlags) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VerificationFlags) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VerificationFlags.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VerificationFlags) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerificationFlags.Merge(m, src)
}
func (m *VerificationFlags) XXX_Size() int {
	return m.Size()
}
func (m *VerificationFlags) XXX_DiscardUnknown() {
	xxx_messageInfo_VerificationFlags.DiscardUnknown(m)
}

var xxx_messageInfo_VerificationFlags proto.InternalMessageInfo

func (m *VerificationFlags) GetEthTypeChainEnabled() bool {
	if m != nil {
		return m.EthTypeChainEnabled
	}
	return false
}

func (m *VerificationFlags) GetBtcTypeChainEnabled() bool {
	if m != nil {
		return m.BtcTypeChainEnabled
	}
	return false
}

func init() {
	proto.RegisterType((*VerificationFlags)(nil), "zetachain.zetacore.lightclient.VerificationFlags")
}

func init() {
	proto.RegisterFile("zetachain/zetacore/lightclient/verification_flags.proto", fileDescriptor_bcf482283292221c)
}

var fileDescriptor_bcf482283292221c = []byte{
	// 199 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0xaf, 0x4a, 0x2d, 0x49,
	0x4c, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x07, 0xb3, 0xf2, 0x8b, 0x52, 0xf5, 0x73, 0x32, 0xd3, 0x33,
	0x4a, 0x92, 0x73, 0x32, 0x53, 0xf3, 0x4a, 0xf4, 0xcb, 0x52, 0x8b, 0x32, 0xd3, 0x32, 0x93, 0x13,
	0x4b, 0x32, 0xf3, 0xf3, 0xe2, 0xd3, 0x72, 0x12, 0xd3, 0x8b, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2,
	0x85, 0xe4, 0xe0, 0x1a, 0xf5, 0x60, 0x1a, 0xf5, 0x90, 0x34, 0x2a, 0x95, 0x73, 0x09, 0x86, 0x21,
	0xe9, 0x75, 0x03, 0x69, 0x15, 0x32, 0xe0, 0x12, 0x4e, 0x2d, 0xc9, 0x08, 0xa9, 0x2c, 0x48, 0x75,
	0x06, 0xe9, 0x74, 0xcd, 0x4b, 0x4c, 0xca, 0x49, 0x4d, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x08,
	0xc2, 0x26, 0x05, 0xd2, 0x91, 0x54, 0x92, 0x8c, 0xa1, 0x83, 0x09, 0xa2, 0x03, 0x8b, 0x94, 0x93,
	0xcf, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1,
	0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0x19, 0xa5, 0x67, 0x96, 0x64,
	0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0x82, 0x3d, 0xab, 0x8b, 0xe6, 0xef, 0x0a, 0x14, 0x9f, 0x97,
	0x54, 0x16, 0xa4, 0x16, 0x27, 0xb1, 0x81, 0x7d, 0x6b, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xab,
	0x69, 0x91, 0x02, 0x28, 0x01, 0x00, 0x00,
}

func (m *VerificationFlags) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VerificationFlags) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VerificationFlags) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BtcTypeChainEnabled {
		i--
		if m.BtcTypeChainEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.EthTypeChainEnabled {
		i--
		if m.EthTypeChainEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
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
func (m *VerificationFlags) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EthTypeChainEnabled {
		n += 2
	}
	if m.BtcTypeChainEnabled {
		n += 2
	}
	return n
}

func sovVerificationFlags(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozVerificationFlags(x uint64) (n int) {
	return sovVerificationFlags(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *VerificationFlags) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: VerificationFlags: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VerificationFlags: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EthTypeChainEnabled", wireType)
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
			m.EthTypeChainEnabled = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BtcTypeChainEnabled", wireType)
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
			m.BtcTypeChainEnabled = bool(v != 0)
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
