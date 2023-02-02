// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: observer/params.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	common "github.com/zeta-chain/zetacore/common"
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

type ObserverParams struct {
	Chain                 *common.Chain                          `protobuf:"bytes,1,opt,name=chain,proto3" json:"chain,omitempty"`
	BallotThreshold       github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=ballot_threshold,json=ballotThreshold,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"ballot_threshold"`
	MinObserverDelegation github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=min_observer_delegation,json=minObserverDelegation,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"min_observer_delegation"`
	IsSupported           bool                                   `protobuf:"varint,5,opt,name=is_supported,json=isSupported,proto3" json:"is_supported,omitempty"`
}

func (m *ObserverParams) Reset()         { *m = ObserverParams{} }
func (m *ObserverParams) String() string { return proto.CompactTextString(m) }
func (*ObserverParams) ProtoMessage()    {}
func (*ObserverParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_4542fa62877488a1, []int{0}
}
func (m *ObserverParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ObserverParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ObserverParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ObserverParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObserverParams.Merge(m, src)
}
func (m *ObserverParams) XXX_Size() int {
	return m.Size()
}
func (m *ObserverParams) XXX_DiscardUnknown() {
	xxx_messageInfo_ObserverParams.DiscardUnknown(m)
}

var xxx_messageInfo_ObserverParams proto.InternalMessageInfo

func (m *ObserverParams) GetChain() *common.Chain {
	if m != nil {
		return m.Chain
	}
	return nil
}

func (m *ObserverParams) GetIsSupported() bool {
	if m != nil {
		return m.IsSupported
	}
	return false
}

// Params defines the parameters for the module.
type Params struct {
	ObserverParams []*ObserverParams `protobuf:"bytes,1,rep,name=observer_params,json=observerParams,proto3" json:"observer_params,omitempty"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_4542fa62877488a1, []int{1}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetObserverParams() []*ObserverParams {
	if m != nil {
		return m.ObserverParams
	}
	return nil
}

func init() {
	proto.RegisterType((*ObserverParams)(nil), "zetachain.zetacore.observer.ObserverParams")
	proto.RegisterType((*Params)(nil), "zetachain.zetacore.observer.Params")
}

func init() { proto.RegisterFile("observer/params.proto", fileDescriptor_4542fa62877488a1) }

var fileDescriptor_4542fa62877488a1 = []byte{
	// 364 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcd, 0x4f, 0x2a, 0x4e,
	0x2d, 0x2a, 0x4b, 0x2d, 0xd2, 0x2f, 0x48, 0x2c, 0x4a, 0xcc, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x92, 0xae, 0x4a, 0x2d, 0x49, 0x4c, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x03, 0xb3, 0xf2,
	0x8b, 0x52, 0xf5, 0x60, 0x2a, 0xa5, 0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0xea, 0xf4, 0x41, 0x2c,
	0x88, 0x16, 0x29, 0x71, 0xb8, 0x49, 0x30, 0x06, 0x54, 0x42, 0x38, 0x39, 0x3f, 0x37, 0x37, 0x3f,
	0x4f, 0x1f, 0x42, 0x41, 0x04, 0x95, 0x66, 0x33, 0x71, 0xf1, 0xf9, 0x43, 0xd5, 0x05, 0x80, 0x6d,
	0x16, 0x52, 0xe6, 0x62, 0x05, 0xdb, 0x28, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x6d, 0xc4, 0xab, 0x07,
	0xd5, 0xe0, 0x0c, 0x12, 0x0c, 0x82, 0xc8, 0x09, 0x45, 0x72, 0x09, 0x24, 0x25, 0xe6, 0xe4, 0xe4,
	0x97, 0xc4, 0x97, 0x64, 0x14, 0xa5, 0x16, 0x67, 0xe4, 0xe7, 0xa4, 0x48, 0x30, 0x2b, 0x30, 0x6a,
	0x70, 0x3a, 0xe9, 0x9d, 0xb8, 0x27, 0xcf, 0x70, 0xeb, 0x9e, 0xbc, 0x5a, 0x7a, 0x66, 0x49, 0x46,
	0x69, 0x12, 0x48, 0xb7, 0x7e, 0x72, 0x7e, 0x71, 0x6e, 0x7e, 0x31, 0x94, 0xd2, 0x2d, 0x4e, 0xc9,
	0xd6, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x73, 0x49, 0x4d, 0x0e, 0xe2, 0x87, 0x98, 0x13, 0x02,
	0x33, 0x46, 0x28, 0x8d, 0x4b, 0x3c, 0x37, 0x33, 0x2f, 0x1e, 0xe6, 0xfa, 0xf8, 0x94, 0xd4, 0x9c,
	0xd4, 0xf4, 0xc4, 0x92, 0xcc, 0xfc, 0x3c, 0x09, 0x16, 0xb2, 0x6c, 0x10, 0xcd, 0xcd, 0xcc, 0x83,
	0xf9, 0xd1, 0x05, 0x6e, 0x98, 0x90, 0x22, 0x17, 0x4f, 0x66, 0x71, 0x7c, 0x71, 0x69, 0x41, 0x41,
	0x7e, 0x51, 0x49, 0x6a, 0x8a, 0x04, 0xab, 0x02, 0xa3, 0x06, 0x47, 0x10, 0x77, 0x66, 0x71, 0x30,
	0x4c, 0x48, 0x29, 0x85, 0x8b, 0x0d, 0x1a, 0x28, 0x21, 0x5c, 0xfc, 0x70, 0x07, 0x41, 0x62, 0x48,
	0x82, 0x51, 0x81, 0x59, 0x83, 0xdb, 0x48, 0x5b, 0x0f, 0x4f, 0x14, 0xe9, 0xa1, 0x06, 0x6d, 0x10,
	0x5f, 0x3e, 0x0a, 0xdf, 0x8a, 0x65, 0xc6, 0x02, 0x79, 0x06, 0x27, 0xcf, 0x13, 0x8f, 0xe4, 0x18,
	0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5,
	0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0xd2, 0x47, 0xf2, 0x21, 0xc8, 0x70, 0x5d, 0xb0, 0x3d, 0xfa,
	0x30, 0x7b, 0xf4, 0x2b, 0xe0, 0x71, 0x0c, 0xf1, 0x6e, 0x12, 0x1b, 0x38, 0x56, 0x8d, 0x01, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xea, 0x58, 0xa1, 0x49, 0x4f, 0x02, 0x00, 0x00,
}

func (m *ObserverParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ObserverParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ObserverParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.IsSupported {
		i--
		if m.IsSupported {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	{
		size := m.MinObserverDelegation.Size()
		i -= size
		if _, err := m.MinObserverDelegation.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.BallotThreshold.Size()
		i -= size
		if _, err := m.BallotThreshold.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.Chain != nil {
		{
			size, err := m.Chain.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintParams(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ObserverParams) > 0 {
		for iNdEx := len(m.ObserverParams) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ObserverParams[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ObserverParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Chain != nil {
		l = m.Chain.Size()
		n += 1 + l + sovParams(uint64(l))
	}
	l = m.BallotThreshold.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.MinObserverDelegation.Size()
	n += 1 + l + sovParams(uint64(l))
	if m.IsSupported {
		n += 2
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ObserverParams) > 0 {
		for _, e := range m.ObserverParams {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ObserverParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: ObserverParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ObserverParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Chain == nil {
				m.Chain = &common.Chain{}
			}
			if err := m.Chain.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BallotThreshold", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BallotThreshold.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinObserverDelegation", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinObserverDelegation.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsSupported", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
			m.IsSupported = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObserverParams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ObserverParams = append(m.ObserverParams, &ObserverParams{})
			if err := m.ObserverParams[len(m.ObserverParams)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
