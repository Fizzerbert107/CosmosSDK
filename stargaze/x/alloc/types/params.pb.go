// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stargaze/alloc/v1beta1/params.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

type WeightedAddress struct {
	Address string                                 `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty" yaml:"address"`
	Weight  github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=weight,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"weight" yaml:"weight"`
}

func (m *WeightedAddress) Reset()         { *m = WeightedAddress{} }
func (m *WeightedAddress) String() string { return proto.CompactTextString(m) }
func (*WeightedAddress) ProtoMessage()    {}
func (*WeightedAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_23171fab5d42f6a5, []int{0}
}
func (m *WeightedAddress) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WeightedAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WeightedAddress.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WeightedAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WeightedAddress.Merge(m, src)
}
func (m *WeightedAddress) XXX_Size() int {
	return m.Size()
}
func (m *WeightedAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_WeightedAddress.DiscardUnknown(m)
}

var xxx_messageInfo_WeightedAddress proto.InternalMessageInfo

func (m *WeightedAddress) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type DistributionProportions struct {
	NftIncentives    github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=nft_incentives,json=nftIncentives,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"nft_incentives" yaml:"nft_incentives"`
	DeveloperRewards github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=developer_rewards,json=developerRewards,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"developer_rewards" yaml:"developer_rewards"`
}

func (m *DistributionProportions) Reset()         { *m = DistributionProportions{} }
func (m *DistributionProportions) String() string { return proto.CompactTextString(m) }
func (*DistributionProportions) ProtoMessage()    {}
func (*DistributionProportions) Descriptor() ([]byte, []int) {
	return fileDescriptor_23171fab5d42f6a5, []int{1}
}
func (m *DistributionProportions) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DistributionProportions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DistributionProportions.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DistributionProportions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DistributionProportions.Merge(m, src)
}
func (m *DistributionProportions) XXX_Size() int {
	return m.Size()
}
func (m *DistributionProportions) XXX_DiscardUnknown() {
	xxx_messageInfo_DistributionProportions.DiscardUnknown(m)
}

var xxx_messageInfo_DistributionProportions proto.InternalMessageInfo

type Params struct {
	// distribution_proportions defines the proportion of the minted denom
	DistributionProportions DistributionProportions `protobuf:"bytes,1,opt,name=distribution_proportions,json=distributionProportions,proto3" json:"distribution_proportions"`
	// address to receive developer rewards
	WeightedDeveloperRewardsReceivers []WeightedAddress `protobuf:"bytes,2,rep,name=weighted_developer_rewards_receivers,json=weightedDeveloperRewardsReceivers,proto3" json:"weighted_developer_rewards_receivers" yaml:"developer_rewards_receiver"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_23171fab5d42f6a5, []int{2}
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

func (m *Params) GetDistributionProportions() DistributionProportions {
	if m != nil {
		return m.DistributionProportions
	}
	return DistributionProportions{}
}

func (m *Params) GetWeightedDeveloperRewardsReceivers() []WeightedAddress {
	if m != nil {
		return m.WeightedDeveloperRewardsReceivers
	}
	return nil
}

func init() {
	proto.RegisterType((*WeightedAddress)(nil), "publicawesome.stargaze.alloc.v1beta1.WeightedAddress")
	proto.RegisterType((*DistributionProportions)(nil), "publicawesome.stargaze.alloc.v1beta1.DistributionProportions")
	proto.RegisterType((*Params)(nil), "publicawesome.stargaze.alloc.v1beta1.Params")
}

func init() {
	proto.RegisterFile("stargaze/alloc/v1beta1/params.proto", fileDescriptor_23171fab5d42f6a5)
}

var fileDescriptor_23171fab5d42f6a5 = []byte{
	// 469 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x4f, 0x6b, 0xd4, 0x40,
	0x14, 0xdf, 0xac, 0xb2, 0xe2, 0x94, 0x56, 0x0d, 0x4a, 0x17, 0x0f, 0x89, 0x8d, 0x45, 0x2a, 0xd8,
	0x84, 0xb6, 0xe8, 0x41, 0x10, 0x71, 0x59, 0x10, 0x05, 0xa1, 0xcc, 0xa5, 0xe0, 0x25, 0x4c, 0x32,
	0xaf, 0xe9, 0x60, 0x92, 0x09, 0x33, 0xb3, 0x89, 0xf5, 0xe0, 0xc1, 0x4f, 0xe0, 0xd1, 0xaf, 0xe1,
	0xdd, 0x0f, 0xd0, 0x63, 0x8f, 0xe2, 0x21, 0xc8, 0xee, 0x37, 0xd8, 0x4f, 0x20, 0x99, 0x4c, 0xd6,
	0xfe, 0x61, 0xa1, 0x3d, 0xe5, 0x65, 0x66, 0xde, 0xef, 0xf7, 0x7b, 0xbf, 0xf7, 0x1e, 0x7a, 0x2c,
	0x15, 0x11, 0x09, 0xf9, 0x02, 0x01, 0x49, 0x53, 0x1e, 0x07, 0xe5, 0x4e, 0x04, 0x8a, 0xec, 0x04,
	0x05, 0x11, 0x24, 0x93, 0x7e, 0x21, 0xb8, 0xe2, 0xf6, 0x66, 0x31, 0x89, 0x52, 0x16, 0x93, 0x0a,
	0x24, 0xcf, 0xc0, 0xef, 0x52, 0x7c, 0x9d, 0xe2, 0x9b, 0x94, 0x87, 0xf7, 0x13, 0x9e, 0x70, 0x9d,
	0x10, 0x34, 0x51, 0x9b, 0xeb, 0xfd, 0xb0, 0xd0, 0x9d, 0x03, 0x60, 0xc9, 0x91, 0x02, 0xfa, 0x86,
	0x52, 0x01, 0x52, 0xda, 0xcf, 0xd0, 0x2d, 0xd2, 0x86, 0x43, 0xeb, 0x91, 0xb5, 0x75, 0x7b, 0x64,
	0xcf, 0x6b, 0x77, 0xed, 0x98, 0x64, 0xe9, 0x4b, 0xcf, 0x5c, 0x78, 0xb8, 0x7b, 0x62, 0x1f, 0xa0,
	0x41, 0xa5, 0x01, 0x86, 0x7d, 0xfd, 0xf8, 0xf5, 0x49, 0xed, 0xf6, 0xfe, 0xd4, 0xee, 0x93, 0x84,
	0xa9, 0xa3, 0x49, 0xe4, 0xc7, 0x3c, 0x0b, 0x62, 0x2e, 0x33, 0x2e, 0xcd, 0x67, 0x5b, 0xd2, 0x4f,
	0x81, 0x3a, 0x2e, 0x40, 0xfa, 0x63, 0x88, 0xe7, 0xb5, 0xbb, 0xda, 0x42, 0xb7, 0x28, 0x1e, 0x36,
	0x70, 0xde, 0xb7, 0x3e, 0x5a, 0x1f, 0x33, 0xa9, 0x04, 0x8b, 0x26, 0x8a, 0xf1, 0x7c, 0x5f, 0xf0,
	0x82, 0x8b, 0x26, 0x92, 0x76, 0x8e, 0xd6, 0xf2, 0x43, 0x15, 0xb2, 0x3c, 0x86, 0x5c, 0xb1, 0x12,
	0x3a, 0xa5, 0x6f, 0xaf, 0x4d, 0xfe, 0xa0, 0x25, 0x3f, 0x8f, 0xe6, 0xe1, 0xd5, 0xfc, 0x50, 0xbd,
	0x5b, 0xfc, 0xdb, 0x15, 0xba, 0x47, 0xa1, 0x84, 0x94, 0x17, 0x20, 0x42, 0x01, 0x15, 0x11, 0x54,
	0x9a, 0x7a, 0xdf, 0x5f, 0x9b, 0x72, 0xd8, 0x52, 0x5e, 0x02, 0xf4, 0xf0, 0xdd, 0xc5, 0x19, 0x36,
	0x47, 0xbf, 0xfa, 0x68, 0xb0, 0xaf, 0x9b, 0x6d, 0x7f, 0x45, 0x43, 0x7a, 0xc6, 0x8e, 0xb0, 0xf8,
	0xef, 0x87, 0xae, 0x7e, 0x65, 0xf7, 0x95, 0x7f, 0x95, 0x49, 0xf0, 0x97, 0x98, 0x3a, 0xba, 0xd9,
	0x54, 0x82, 0xd7, 0xe9, 0x12, 0xcf, 0x7f, 0x5a, 0x68, 0xb3, 0x32, 0xa3, 0x12, 0x5e, 0x12, 0x1f,
	0x0a, 0x88, 0x81, 0x95, 0x20, 0x1a, 0x5f, 0x6e, 0x6c, 0xad, 0xec, 0x3e, 0xbf, 0x9a, 0x98, 0x0b,
	0xc3, 0x37, 0x7a, 0xda, 0x88, 0x98, 0xd7, 0xee, 0xc6, 0x12, 0x93, 0x16, 0x3c, 0x1e, 0xde, 0xe8,
	0xd4, 0x8c, 0x2f, 0xb8, 0x86, 0x3b, 0x29, 0xa3, 0x0f, 0x27, 0x53, 0xc7, 0x3a, 0x9d, 0x3a, 0xd6,
	0xdf, 0xa9, 0x63, 0x7d, 0x9f, 0x39, 0xbd, 0xd3, 0x99, 0xd3, 0xfb, 0x3d, 0x73, 0x7a, 0x1f, 0xf7,
	0xce, 0xb4, 0xab, 0x15, 0xba, 0x6d, 0x94, 0x06, 0x8b, 0x9d, 0x2b, 0x5f, 0x04, 0x9f, 0xcd, 0xe2,
	0xe9, 0xfe, 0x45, 0x03, 0xbd, 0x34, 0x7b, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x6e, 0x42, 0xd4,
	0x1e, 0x97, 0x03, 0x00, 0x00,
}

func (m *WeightedAddress) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WeightedAddress) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WeightedAddress) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Weight.Size()
		i -= size
		if _, err := m.Weight.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintParams(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DistributionProportions) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DistributionProportions) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DistributionProportions) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.DeveloperRewards.Size()
		i -= size
		if _, err := m.DeveloperRewards.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.NftIncentives.Size()
		i -= size
		if _, err := m.NftIncentives.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
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
	if len(m.WeightedDeveloperRewardsReceivers) > 0 {
		for iNdEx := len(m.WeightedDeveloperRewardsReceivers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.WeightedDeveloperRewardsReceivers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.DistributionProportions.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
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
func (m *WeightedAddress) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = m.Weight.Size()
	n += 1 + l + sovParams(uint64(l))
	return n
}

func (m *DistributionProportions) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.NftIncentives.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.DeveloperRewards.Size()
	n += 1 + l + sovParams(uint64(l))
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.DistributionProportions.Size()
	n += 1 + l + sovParams(uint64(l))
	if len(m.WeightedDeveloperRewardsReceivers) > 0 {
		for _, e := range m.WeightedDeveloperRewardsReceivers {
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
func (m *WeightedAddress) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: WeightedAddress: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WeightedAddress: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
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
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Weight", wireType)
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
			if err := m.Weight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *DistributionProportions) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: DistributionProportions: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DistributionProportions: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NftIncentives", wireType)
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
			if err := m.NftIncentives.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeveloperRewards", wireType)
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
			if err := m.DeveloperRewards.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
				return fmt.Errorf("proto: wrong wireType = %d for field DistributionProportions", wireType)
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
			if err := m.DistributionProportions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WeightedDeveloperRewardsReceivers", wireType)
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
			m.WeightedDeveloperRewardsReceivers = append(m.WeightedDeveloperRewardsReceivers, WeightedAddress{})
			if err := m.WeightedDeveloperRewardsReceivers[len(m.WeightedDeveloperRewardsReceivers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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