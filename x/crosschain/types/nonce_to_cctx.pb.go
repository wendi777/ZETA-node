// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: crosschain/nonce_to_cctx.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

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

// store key is tss+chainid+nonce
type NonceToCctx struct {
	ChainId   int64  `protobuf:"varint,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	Nonce     int64  `protobuf:"varint,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
	CctxIndex string `protobuf:"bytes,3,opt,name=cctxIndex,proto3" json:"cctxIndex,omitempty"`
	Tss       string `protobuf:"bytes,4,opt,name=tss,proto3" json:"tss,omitempty"`
}

func (m *NonceToCctx) Reset()         { *m = NonceToCctx{} }
func (m *NonceToCctx) String() string { return proto.CompactTextString(m) }
func (*NonceToCctx) ProtoMessage()    {}
func (*NonceToCctx) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfa802a2474dc9c1, []int{0}
}
func (m *NonceToCctx) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NonceToCctx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NonceToCctx.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NonceToCctx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NonceToCctx.Merge(m, src)
}
func (m *NonceToCctx) XXX_Size() int {
	return m.Size()
}
func (m *NonceToCctx) XXX_DiscardUnknown() {
	xxx_messageInfo_NonceToCctx.DiscardUnknown(m)
}

var xxx_messageInfo_NonceToCctx proto.InternalMessageInfo

func (m *NonceToCctx) GetChainId() int64 {
	if m != nil {
		return m.ChainId
	}
	return 0
}

func (m *NonceToCctx) GetNonce() int64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *NonceToCctx) GetCctxIndex() string {
	if m != nil {
		return m.CctxIndex
	}
	return ""
}

func (m *NonceToCctx) GetTss() string {
	if m != nil {
		return m.Tss
	}
	return ""
}

// store key is tss+chainid
type PendingNonces struct {
	NonceLow  int64  `protobuf:"varint,1,opt,name=nonce_low,json=nonceLow,proto3" json:"nonce_low,omitempty"`
	NonceHigh int64  `protobuf:"varint,2,opt,name=nonce_high,json=nonceHigh,proto3" json:"nonce_high,omitempty"`
	ChainId   int64  `protobuf:"varint,3,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	Tss       string `protobuf:"bytes,4,opt,name=tss,proto3" json:"tss,omitempty"`
}

func (m *PendingNonces) Reset()         { *m = PendingNonces{} }
func (m *PendingNonces) String() string { return proto.CompactTextString(m) }
func (*PendingNonces) ProtoMessage()    {}
func (*PendingNonces) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfa802a2474dc9c1, []int{1}
}
func (m *PendingNonces) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PendingNonces) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PendingNonces.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PendingNonces) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PendingNonces.Merge(m, src)
}
func (m *PendingNonces) XXX_Size() int {
	return m.Size()
}
func (m *PendingNonces) XXX_DiscardUnknown() {
	xxx_messageInfo_PendingNonces.DiscardUnknown(m)
}

var xxx_messageInfo_PendingNonces proto.InternalMessageInfo

func (m *PendingNonces) GetNonceLow() int64 {
	if m != nil {
		return m.NonceLow
	}
	return 0
}

func (m *PendingNonces) GetNonceHigh() int64 {
	if m != nil {
		return m.NonceHigh
	}
	return 0
}

func (m *PendingNonces) GetChainId() int64 {
	if m != nil {
		return m.ChainId
	}
	return 0
}

func (m *PendingNonces) GetTss() string {
	if m != nil {
		return m.Tss
	}
	return ""
}

func init() {
	proto.RegisterType((*NonceToCctx)(nil), "zetachain.zetacore.crosschain.NonceToCctx")
	proto.RegisterType((*PendingNonces)(nil), "zetachain.zetacore.crosschain.PendingNonces")
}

func init() { proto.RegisterFile("crosschain/nonce_to_cctx.proto", fileDescriptor_cfa802a2474dc9c1) }

var fileDescriptor_cfa802a2474dc9c1 = []byte{
	// 273 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0x2e, 0xca, 0x2f,
	0x2e, 0x4e, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0xcf, 0xcb, 0xcf, 0x4b, 0x4e, 0x8d, 0x2f, 0xc9, 0x8f,
	0x4f, 0x4e, 0x2e, 0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0xad, 0x4a, 0x2d, 0x49,
	0x04, 0x4b, 0xeb, 0x81, 0x59, 0xf9, 0x45, 0xa9, 0x7a, 0x08, 0x2d, 0x4a, 0x79, 0x5c, 0xdc, 0x7e,
	0x20, 0x5d, 0x21, 0xf9, 0xce, 0xc9, 0x25, 0x15, 0x42, 0x92, 0x5c, 0x1c, 0x60, 0xf1, 0xf8, 0xcc,
	0x14, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xe6, 0x20, 0x76, 0x30, 0xdf, 0x33, 0x45, 0x48, 0x84, 0x8b,
	0x15, 0x6c, 0xbe, 0x04, 0x13, 0x58, 0x1c, 0xc2, 0x11, 0x92, 0xe1, 0xe2, 0x04, 0x59, 0xe6, 0x99,
	0x97, 0x92, 0x5a, 0x21, 0xc1, 0xac, 0xc0, 0xa8, 0xc1, 0x19, 0x84, 0x10, 0x10, 0x12, 0xe0, 0x62,
	0x2e, 0x29, 0x2e, 0x96, 0x60, 0x01, 0x8b, 0x83, 0x98, 0x4a, 0x15, 0x5c, 0xbc, 0x01, 0xa9, 0x79,
	0x29, 0x99, 0x79, 0xe9, 0x60, 0x6b, 0x8b, 0x85, 0xa4, 0xb9, 0x38, 0x21, 0xce, 0xce, 0xc9, 0x2f,
	0x87, 0x5a, 0xc9, 0x01, 0x16, 0xf0, 0xc9, 0x2f, 0x17, 0x92, 0xe5, 0xe2, 0x82, 0x48, 0x66, 0x64,
	0xa6, 0x67, 0x40, 0x2d, 0x86, 0x28, 0xf7, 0xc8, 0x4c, 0xcf, 0x40, 0x71, 0x2d, 0x33, 0xaa, 0x6b,
	0x31, 0x6c, 0x76, 0xf2, 0x3e, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4,
	0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0xc3,
	0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x7d, 0x50, 0x18, 0xe9, 0x42, 0x42,
	0x13, 0x16, 0x5c, 0xfa, 0x15, 0xfa, 0x48, 0x61, 0x5c, 0x52, 0x59, 0x90, 0x5a, 0x9c, 0xc4, 0x06,
	0x0e, 0x5c, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb8, 0xcb, 0x3e, 0xac, 0x7e, 0x01, 0x00,
	0x00,
}

func (m *NonceToCctx) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NonceToCctx) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NonceToCctx) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Tss) > 0 {
		i -= len(m.Tss)
		copy(dAtA[i:], m.Tss)
		i = encodeVarintNonceToCctx(dAtA, i, uint64(len(m.Tss)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.CctxIndex) > 0 {
		i -= len(m.CctxIndex)
		copy(dAtA[i:], m.CctxIndex)
		i = encodeVarintNonceToCctx(dAtA, i, uint64(len(m.CctxIndex)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Nonce != 0 {
		i = encodeVarintNonceToCctx(dAtA, i, uint64(m.Nonce))
		i--
		dAtA[i] = 0x10
	}
	if m.ChainId != 0 {
		i = encodeVarintNonceToCctx(dAtA, i, uint64(m.ChainId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *PendingNonces) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PendingNonces) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PendingNonces) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Tss) > 0 {
		i -= len(m.Tss)
		copy(dAtA[i:], m.Tss)
		i = encodeVarintNonceToCctx(dAtA, i, uint64(len(m.Tss)))
		i--
		dAtA[i] = 0x22
	}
	if m.ChainId != 0 {
		i = encodeVarintNonceToCctx(dAtA, i, uint64(m.ChainId))
		i--
		dAtA[i] = 0x18
	}
	if m.NonceHigh != 0 {
		i = encodeVarintNonceToCctx(dAtA, i, uint64(m.NonceHigh))
		i--
		dAtA[i] = 0x10
	}
	if m.NonceLow != 0 {
		i = encodeVarintNonceToCctx(dAtA, i, uint64(m.NonceLow))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintNonceToCctx(dAtA []byte, offset int, v uint64) int {
	offset -= sovNonceToCctx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *NonceToCctx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ChainId != 0 {
		n += 1 + sovNonceToCctx(uint64(m.ChainId))
	}
	if m.Nonce != 0 {
		n += 1 + sovNonceToCctx(uint64(m.Nonce))
	}
	l = len(m.CctxIndex)
	if l > 0 {
		n += 1 + l + sovNonceToCctx(uint64(l))
	}
	l = len(m.Tss)
	if l > 0 {
		n += 1 + l + sovNonceToCctx(uint64(l))
	}
	return n
}

func (m *PendingNonces) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NonceLow != 0 {
		n += 1 + sovNonceToCctx(uint64(m.NonceLow))
	}
	if m.NonceHigh != 0 {
		n += 1 + sovNonceToCctx(uint64(m.NonceHigh))
	}
	if m.ChainId != 0 {
		n += 1 + sovNonceToCctx(uint64(m.ChainId))
	}
	l = len(m.Tss)
	if l > 0 {
		n += 1 + l + sovNonceToCctx(uint64(l))
	}
	return n
}

func sovNonceToCctx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNonceToCctx(x uint64) (n int) {
	return sovNonceToCctx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *NonceToCctx) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNonceToCctx
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
			return fmt.Errorf("proto: NonceToCctx: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NonceToCctx: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			m.ChainId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNonceToCctx
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
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			m.Nonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNonceToCctx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Nonce |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CctxIndex", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNonceToCctx
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
				return ErrInvalidLengthNonceToCctx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNonceToCctx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CctxIndex = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tss", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNonceToCctx
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
				return ErrInvalidLengthNonceToCctx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNonceToCctx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tss = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNonceToCctx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNonceToCctx
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
func (m *PendingNonces) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNonceToCctx
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
			return fmt.Errorf("proto: PendingNonces: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PendingNonces: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NonceLow", wireType)
			}
			m.NonceLow = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNonceToCctx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NonceLow |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NonceHigh", wireType)
			}
			m.NonceHigh = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNonceToCctx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NonceHigh |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			m.ChainId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNonceToCctx
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
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tss", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNonceToCctx
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
				return ErrInvalidLengthNonceToCctx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNonceToCctx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tss = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNonceToCctx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNonceToCctx
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
func skipNonceToCctx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNonceToCctx
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
					return 0, ErrIntOverflowNonceToCctx
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
					return 0, ErrIntOverflowNonceToCctx
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
				return 0, ErrInvalidLengthNonceToCctx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNonceToCctx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNonceToCctx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNonceToCctx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNonceToCctx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNonceToCctx = fmt.Errorf("proto: unexpected end of group")
)
