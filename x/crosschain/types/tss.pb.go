// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: crosschain/tss.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

type TSS struct {
	TssPubkey           string   `protobuf:"bytes,3,opt,name=tss_pubkey,json=tssPubkey,proto3" json:"tss_pubkey,omitempty"`
	SignerList          []string `protobuf:"bytes,4,rep,name=signer_list,json=signerList,proto3" json:"signer_list,omitempty"`
	FinalizedZetaHeight int64    `protobuf:"varint,5,opt,name=finalizedZetaHeight,proto3" json:"finalizedZetaHeight,omitempty"`
	KeyGenZetaHeight    int64    `protobuf:"varint,6,opt,name=keyGenZetaHeight,proto3" json:"keyGenZetaHeight,omitempty"`
}

func (m *TSS) Reset()         { *m = TSS{} }
func (m *TSS) String() string { return proto.CompactTextString(m) }
func (*TSS) ProtoMessage()    {}
func (*TSS) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba8ccd105b767be6, []int{0}
}
func (m *TSS) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TSS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TSS.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TSS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TSS.Merge(m, src)
}
func (m *TSS) XXX_Size() int {
	return m.Size()
}
func (m *TSS) XXX_DiscardUnknown() {
	xxx_messageInfo_TSS.DiscardUnknown(m)
}

var xxx_messageInfo_TSS proto.InternalMessageInfo

func (m *TSS) GetTssPubkey() string {
	if m != nil {
		return m.TssPubkey
	}
	return ""
}

func (m *TSS) GetSignerList() []string {
	if m != nil {
		return m.SignerList
	}
	return nil
}

func (m *TSS) GetFinalizedZetaHeight() int64 {
	if m != nil {
		return m.FinalizedZetaHeight
	}
	return 0
}

func (m *TSS) GetKeyGenZetaHeight() int64 {
	if m != nil {
		return m.KeyGenZetaHeight
	}
	return 0
}

func init() {
	proto.RegisterType((*TSS)(nil), "zetachain.zetacore.crosschain.TSS")
}

func init() { proto.RegisterFile("crosschain/tss.proto", fileDescriptor_ba8ccd105b767be6) }

var fileDescriptor_ba8ccd105b767be6 = []byte{
	// 253 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x49, 0x2e, 0xca, 0x2f,
	0x2e, 0x4e, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x2f, 0x29, 0x2e, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x92, 0xad, 0x4a, 0x2d, 0x49, 0x04, 0x0b, 0xea, 0x81, 0x59, 0xf9, 0x45, 0xa9, 0x7a, 0x08,
	0x85, 0x52, 0x22, 0xe9, 0xf9, 0xe9, 0xf9, 0x60, 0x95, 0xfa, 0x20, 0x16, 0x44, 0x93, 0xd2, 0x62,
	0x46, 0x2e, 0xe6, 0x90, 0xe0, 0x60, 0x21, 0x59, 0x2e, 0xae, 0x92, 0xe2, 0xe2, 0xf8, 0x82, 0xd2,
	0xa4, 0xec, 0xd4, 0x4a, 0x09, 0x66, 0x05, 0x46, 0x0d, 0xce, 0x20, 0xce, 0x92, 0xe2, 0xe2, 0x00,
	0xb0, 0x80, 0x90, 0x3c, 0x17, 0x77, 0x71, 0x66, 0x7a, 0x5e, 0x6a, 0x51, 0x7c, 0x4e, 0x66, 0x71,
	0x89, 0x04, 0x8b, 0x02, 0xb3, 0x06, 0x67, 0x10, 0x17, 0x44, 0xc8, 0x27, 0xb3, 0xb8, 0x44, 0xc8,
	0x80, 0x4b, 0x38, 0x2d, 0x33, 0x2f, 0x31, 0x27, 0xb3, 0x2a, 0x35, 0x25, 0x2a, 0xb5, 0x24, 0xd1,
	0x23, 0x35, 0x33, 0x3d, 0xa3, 0x44, 0x82, 0x55, 0x81, 0x51, 0x83, 0x39, 0x08, 0x9b, 0x94, 0x90,
	0x16, 0x97, 0x40, 0x76, 0x6a, 0xa5, 0x7b, 0x6a, 0x1e, 0x92, 0x72, 0x36, 0xb0, 0x72, 0x0c, 0x71,
	0x27, 0xef, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2,
	0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0x32, 0x4c, 0xcf, 0x2c,
	0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x07, 0xf9, 0x5a, 0x17, 0x12, 0x2a, 0xb0, 0x00,
	0xd0, 0xaf, 0xd0, 0x47, 0x0e, 0xab, 0xca, 0x82, 0xd4, 0xe2, 0x24, 0x36, 0xb0, 0xcf, 0x8d, 0x01,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x6f, 0xe2, 0x0f, 0x78, 0x46, 0x01, 0x00, 0x00,
}

func (m *TSS) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TSS) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TSS) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.KeyGenZetaHeight != 0 {
		i = encodeVarintTss(dAtA, i, uint64(m.KeyGenZetaHeight))
		i--
		dAtA[i] = 0x30
	}
	if m.FinalizedZetaHeight != 0 {
		i = encodeVarintTss(dAtA, i, uint64(m.FinalizedZetaHeight))
		i--
		dAtA[i] = 0x28
	}
	if len(m.SignerList) > 0 {
		for iNdEx := len(m.SignerList) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.SignerList[iNdEx])
			copy(dAtA[i:], m.SignerList[iNdEx])
			i = encodeVarintTss(dAtA, i, uint64(len(m.SignerList[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.TssPubkey) > 0 {
		i -= len(m.TssPubkey)
		copy(dAtA[i:], m.TssPubkey)
		i = encodeVarintTss(dAtA, i, uint64(len(m.TssPubkey)))
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}

func encodeVarintTss(dAtA []byte, offset int, v uint64) int {
	offset -= sovTss(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TSS) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TssPubkey)
	if l > 0 {
		n += 1 + l + sovTss(uint64(l))
	}
	if len(m.SignerList) > 0 {
		for _, s := range m.SignerList {
			l = len(s)
			n += 1 + l + sovTss(uint64(l))
		}
	}
	if m.FinalizedZetaHeight != 0 {
		n += 1 + sovTss(uint64(m.FinalizedZetaHeight))
	}
	if m.KeyGenZetaHeight != 0 {
		n += 1 + sovTss(uint64(m.KeyGenZetaHeight))
	}
	return n
}

func sovTss(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTss(x uint64) (n int) {
	return sovTss(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TSS) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTss
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
			return fmt.Errorf("proto: TSS: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TSS: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TssPubkey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTss
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
				return ErrInvalidLengthTss
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTss
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TssPubkey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignerList", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTss
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
				return ErrInvalidLengthTss
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTss
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SignerList = append(m.SignerList, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FinalizedZetaHeight", wireType)
			}
			m.FinalizedZetaHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTss
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FinalizedZetaHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyGenZetaHeight", wireType)
			}
			m.KeyGenZetaHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTss
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.KeyGenZetaHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTss(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTss
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
func skipTss(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTss
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
					return 0, ErrIntOverflowTss
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
					return 0, ErrIntOverflowTss
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
				return 0, ErrInvalidLengthTss
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTss
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTss
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTss        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTss          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTss = fmt.Errorf("proto: unexpected end of group")
)
