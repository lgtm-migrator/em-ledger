// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: em/slashing/v1beta1/slashing.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
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

type Penalties struct {
	Elements []Penalty `protobuf:"bytes,1,rep,name=elements,proto3" json:"elements" yaml:"elements"`
}

func (m *Penalties) Reset()         { *m = Penalties{} }
func (m *Penalties) String() string { return proto.CompactTextString(m) }
func (*Penalties) ProtoMessage()    {}
func (*Penalties) Descriptor() ([]byte, []int) {
	return fileDescriptor_07d98ad1df708b45, []int{0}
}
func (m *Penalties) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Penalties) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Penalties.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Penalties) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Penalties.Merge(m, src)
}
func (m *Penalties) XXX_Size() int {
	return m.Size()
}
func (m *Penalties) XXX_DiscardUnknown() {
	xxx_messageInfo_Penalties.DiscardUnknown(m)
}

var xxx_messageInfo_Penalties proto.InternalMessageInfo

func (m *Penalties) GetElements() []Penalty {
	if m != nil {
		return m.Elements
	}
	return nil
}

type Penalty struct {
	Validator string                                   `protobuf:"bytes,1,opt,name=validator,proto3" json:"validator,omitempty" yaml:"validator"`
	Amounts   github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,4,rep,name=amounts,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"amounts" yaml:"amounts"`
}

func (m *Penalty) Reset()         { *m = Penalty{} }
func (m *Penalty) String() string { return proto.CompactTextString(m) }
func (*Penalty) ProtoMessage()    {}
func (*Penalty) Descriptor() ([]byte, []int) {
	return fileDescriptor_07d98ad1df708b45, []int{1}
}
func (m *Penalty) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Penalty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Penalty.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Penalty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Penalty.Merge(m, src)
}
func (m *Penalty) XXX_Size() int {
	return m.Size()
}
func (m *Penalty) XXX_DiscardUnknown() {
	xxx_messageInfo_Penalty.DiscardUnknown(m)
}

var xxx_messageInfo_Penalty proto.InternalMessageInfo

func (m *Penalty) GetValidator() string {
	if m != nil {
		return m.Validator
	}
	return ""
}

func (m *Penalty) GetAmounts() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Amounts
	}
	return nil
}

func init() {
	proto.RegisterType((*Penalties)(nil), "em.liquidityprovider.v1beta1.Penalties")
	proto.RegisterType((*Penalty)(nil), "em.liquidityprovider.v1beta1.Penalty")
}

func init() {
	proto.RegisterFile("em/slashing/v1beta1/slashing.proto", fileDescriptor_07d98ad1df708b45)
}

var fileDescriptor_07d98ad1df708b45 = []byte{
	// 359 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0x3f, 0x4e, 0xfb, 0x30,
	0x14, 0x4e, 0xf4, 0xfb, 0x89, 0xd2, 0x20, 0x01, 0x8a, 0x2a, 0xd1, 0x56, 0x28, 0xa9, 0x22, 0x21,
	0x75, 0x49, 0xac, 0x96, 0x8d, 0x31, 0x0c, 0xac, 0xa8, 0x63, 0x17, 0xe4, 0x34, 0x0f, 0xd7, 0xc2,
	0x8e, 0x4b, 0xec, 0x14, 0x72, 0x0b, 0xce, 0xc1, 0x01, 0x38, 0x43, 0xc7, 0x8e, 0x4c, 0x05, 0xb5,
	0x37, 0xe8, 0x09, 0x50, 0x63, 0x37, 0x65, 0x62, 0x4a, 0x9e, 0xbf, 0x3f, 0xef, 0xf3, 0x27, 0x3b,
	0x01, 0x70, 0x24, 0x19, 0x96, 0x53, 0x9a, 0x11, 0x34, 0x1f, 0x24, 0xa0, 0xf0, 0xa0, 0x3e, 0x88,
	0x66, 0xb9, 0x50, 0xc2, 0xbd, 0x04, 0x1e, 0x31, 0xfa, 0x5c, 0xd0, 0x94, 0xaa, 0x72, 0x96, 0x8b,
	0x39, 0x4d, 0x21, 0x8f, 0x0c, 0xb9, 0xdb, 0x22, 0x82, 0x88, 0x8a, 0x88, 0x76, 0x7f, 0x5a, 0xd3,
	0xed, 0x4c, 0x84, 0xe4, 0x42, 0x3e, 0x68, 0x40, 0x0f, 0x06, 0xf2, 0xf4, 0x84, 0x12, 0x2c, 0xa1,
	0x5e, 0x39, 0x11, 0x34, 0xdb, 0x4b, 0x89, 0x10, 0x84, 0x01, 0xaa, 0xa6, 0xa4, 0x78, 0x44, 0x38,
	0x2b, 0x35, 0x14, 0x10, 0xa7, 0x79, 0x0f, 0x19, 0x66, 0x8a, 0x82, 0x74, 0xc7, 0xce, 0x31, 0x30,
	0xe0, 0x90, 0x29, 0xd9, 0xb6, 0x7b, 0xff, 0xfa, 0x27, 0xc3, 0xab, 0xe8, 0xaf, 0xa4, 0x91, 0x96,
	0x96, 0xf1, 0xc5, 0x62, 0xe5, 0x5b, 0xdb, 0x95, 0x7f, 0x56, 0x62, 0xce, 0x6e, 0x82, 0xbd, 0x49,
	0x30, 0xaa, 0xfd, 0x82, 0x0f, 0xdb, 0x69, 0x18, 0xba, 0x3b, 0x74, 0x9a, 0x73, 0xcc, 0x68, 0x8a,
	0x95, 0xc8, 0xdb, 0x76, 0xcf, 0xee, 0x37, 0xe3, 0xd6, 0x76, 0xe5, 0x9f, 0x6b, 0x75, 0x0d, 0x05,
	0xa3, 0x03, 0xcd, 0x7d, 0x71, 0x1a, 0x98, 0x8b, 0x62, 0x17, 0xed, 0x7f, 0x15, 0xad, 0x13, 0x99,
	0x0e, 0x76, 0xb7, 0xae, 0x13, 0xdd, 0x0a, 0x9a, 0xc5, 0xb1, 0x89, 0x73, 0xaa, 0x0d, 0x8d, 0x2e,
	0x78, 0xff, 0xf2, 0xfb, 0x84, 0xaa, 0x69, 0x91, 0x44, 0x13, 0xc1, 0x4d, 0x85, 0xe6, 0x13, 0xca,
	0xf4, 0x09, 0xa9, 0x72, 0x06, 0xb2, 0xb2, 0x90, 0xa3, 0xfd, 0xb6, 0xf8, 0x6e, 0xb1, 0xf6, 0xec,
	0xe5, 0xda, 0xb3, 0xbf, 0xd7, 0x9e, 0xfd, 0xb6, 0xf1, 0xac, 0xe5, 0xc6, 0xb3, 0x3e, 0x37, 0x9e,
	0x35, 0x0e, 0x7f, 0x99, 0x41, 0xc8, 0x45, 0x06, 0x25, 0x02, 0x1e, 0x32, 0x48, 0x09, 0xe4, 0xe8,
	0xf5, 0xf0, 0x0a, 0x2a, 0xdf, 0xe4, 0xa8, 0x6a, 0xfc, 0xfa, 0x27, 0x00, 0x00, 0xff, 0xff, 0x81,
	0x3e, 0xe5, 0xed, 0x21, 0x02, 0x00, 0x00,
}

func (m *Penalties) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Penalties) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Penalties) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Elements) > 0 {
		for iNdEx := len(m.Elements) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Elements[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintSlashing(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Penalty) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Penalty) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Penalty) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Amounts) > 0 {
		for iNdEx := len(m.Amounts) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Amounts[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintSlashing(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Validator) > 0 {
		i -= len(m.Validator)
		copy(dAtA[i:], m.Validator)
		i = encodeVarintSlashing(dAtA, i, uint64(len(m.Validator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSlashing(dAtA []byte, offset int, v uint64) int {
	offset -= sovSlashing(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Penalties) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Elements) > 0 {
		for _, e := range m.Elements {
			l = e.Size()
			n += 1 + l + sovSlashing(uint64(l))
		}
	}
	return n
}

func (m *Penalty) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Validator)
	if l > 0 {
		n += 1 + l + sovSlashing(uint64(l))
	}
	if len(m.Amounts) > 0 {
		for _, e := range m.Amounts {
			l = e.Size()
			n += 1 + l + sovSlashing(uint64(l))
		}
	}
	return n
}

func sovSlashing(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSlashing(x uint64) (n int) {
	return sovSlashing(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Penalties) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSlashing
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
			return fmt.Errorf("proto: Penalties: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Penalties: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Elements", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlashing
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
				return ErrInvalidLengthSlashing
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSlashing
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Elements = append(m.Elements, Penalty{})
			if err := m.Elements[len(m.Elements)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSlashing(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSlashing
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

func (m *Penalty) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSlashing
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
			return fmt.Errorf("proto: Penalty: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Penalty: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlashing
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
				return ErrInvalidLengthSlashing
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSlashing
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Validator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amounts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlashing
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
				return ErrInvalidLengthSlashing
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSlashing
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amounts = append(m.Amounts, types.Coin{})
			if err := m.Amounts[len(m.Amounts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSlashing(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSlashing
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
func skipSlashing(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSlashing
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
					return 0, ErrIntOverflowSlashing
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
					return 0, ErrIntOverflowSlashing
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
				return 0, ErrInvalidLengthSlashing
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSlashing
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSlashing
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSlashing        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSlashing          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSlashing = fmt.Errorf("proto: unexpected end of group")
)
