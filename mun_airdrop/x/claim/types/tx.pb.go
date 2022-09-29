// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: claim/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type MsgClaimFor struct {
	Sender  string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Action  Action `protobuf:"varint,3,opt,name=action,proto3,enum=mun.claim.Action" json:"action,omitempty"`
}

func (m *MsgClaimFor) Reset()         { *m = MsgClaimFor{} }
func (m *MsgClaimFor) String() string { return proto.CompactTextString(m) }
func (*MsgClaimFor) ProtoMessage()    {}
func (*MsgClaimFor) Descriptor() ([]byte, []int) {
	return fileDescriptor_5064b28f68e807c3, []int{0}
}
func (m *MsgClaimFor) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgClaimFor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgClaimFor.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgClaimFor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgClaimFor.Merge(m, src)
}
func (m *MsgClaimFor) XXX_Size() int {
	return m.Size()
}
func (m *MsgClaimFor) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgClaimFor.DiscardUnknown(m)
}

var xxx_messageInfo_MsgClaimFor proto.InternalMessageInfo

func (m *MsgClaimFor) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *MsgClaimFor) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *MsgClaimFor) GetAction() Action {
	if m != nil {
		return m.Action
	}
	return ActionInitialClaim
}

type MsgClaimForResponse struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// total initial claimable amount for the user
	ClaimedAmount github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=claimed_amount,json=claimedAmount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"claimed_amount" yaml:"claimed_amount"`
}

func (m *MsgClaimForResponse) Reset()         { *m = MsgClaimForResponse{} }
func (m *MsgClaimForResponse) String() string { return proto.CompactTextString(m) }
func (*MsgClaimForResponse) ProtoMessage()    {}
func (*MsgClaimForResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5064b28f68e807c3, []int{1}
}
func (m *MsgClaimForResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgClaimForResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgClaimForResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgClaimForResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgClaimForResponse.Merge(m, src)
}
func (m *MsgClaimForResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgClaimForResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgClaimForResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgClaimForResponse proto.InternalMessageInfo

func (m *MsgClaimForResponse) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *MsgClaimForResponse) GetClaimedAmount() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.ClaimedAmount
	}
	return nil
}

type MsgInitialClaim struct {
	Sender string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
}

func (m *MsgInitialClaim) Reset()         { *m = MsgInitialClaim{} }
func (m *MsgInitialClaim) String() string { return proto.CompactTextString(m) }
func (*MsgInitialClaim) ProtoMessage()    {}
func (*MsgInitialClaim) Descriptor() ([]byte, []int) {
	return fileDescriptor_5064b28f68e807c3, []int{2}
}
func (m *MsgInitialClaim) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgInitialClaim) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgInitialClaim.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgInitialClaim) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgInitialClaim.Merge(m, src)
}
func (m *MsgInitialClaim) XXX_Size() int {
	return m.Size()
}
func (m *MsgInitialClaim) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgInitialClaim.DiscardUnknown(m)
}

var xxx_messageInfo_MsgInitialClaim proto.InternalMessageInfo

func (m *MsgInitialClaim) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

type MsgInitialClaimResponse struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// total initial claimable amount for the user
	ClaimedAmount github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=claimed_amount,json=claimedAmount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"claimed_amount" yaml:"claimed_amount"`
}

func (m *MsgInitialClaimResponse) Reset()         { *m = MsgInitialClaimResponse{} }
func (m *MsgInitialClaimResponse) String() string { return proto.CompactTextString(m) }
func (*MsgInitialClaimResponse) ProtoMessage()    {}
func (*MsgInitialClaimResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5064b28f68e807c3, []int{3}
}
func (m *MsgInitialClaimResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgInitialClaimResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgInitialClaimResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgInitialClaimResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgInitialClaimResponse.Merge(m, src)
}
func (m *MsgInitialClaimResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgInitialClaimResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgInitialClaimResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgInitialClaimResponse proto.InternalMessageInfo

func (m *MsgInitialClaimResponse) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *MsgInitialClaimResponse) GetClaimedAmount() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.ClaimedAmount
	}
	return nil
}

func init() {
	proto.RegisterType((*MsgClaimFor)(nil), "mun.claim.MsgClaimFor")
	proto.RegisterType((*MsgClaimForResponse)(nil), "mun.claim.MsgClaimForResponse")
	proto.RegisterType((*MsgInitialClaim)(nil), "mun.claim.MsgInitialClaim")
	proto.RegisterType((*MsgInitialClaimResponse)(nil), "mun.claim.MsgInitialClaimResponse")
}

func init() { proto.RegisterFile("claim/tx.proto", fileDescriptor_5064b28f68e807c3) }

var fileDescriptor_5064b28f68e807c3 = []byte{
	// 408 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x53, 0xbd, 0xae, 0xd3, 0x30,
	0x18, 0x8d, 0x6f, 0xa5, 0xc2, 0xf5, 0x85, 0xa2, 0x86, 0xbf, 0x90, 0xc1, 0xad, 0x32, 0xa5, 0x42,
	0xd8, 0x6a, 0xd9, 0xd8, 0xda, 0x4a, 0x48, 0x1d, 0xca, 0x90, 0x91, 0xa5, 0x72, 0x12, 0x2b, 0x04,
	0x1a, 0xbb, 0x8a, 0x1d, 0xd4, 0x3e, 0x03, 0x0b, 0x2b, 0xaf, 0xc0, 0x3b, 0x30, 0xb0, 0x75, 0xec,
	0xc8, 0x54, 0x50, 0xfb, 0x06, 0x3c, 0x01, 0x8a, 0xe3, 0x96, 0x04, 0xd1, 0x07, 0xb8, 0x4b, 0x12,
	0x7f, 0xe7, 0x7c, 0x27, 0x39, 0xe7, 0xfb, 0x02, 0x3b, 0xd1, 0x92, 0xa6, 0x19, 0x51, 0x6b, 0xbc,
	0xca, 0x85, 0x12, 0xf6, 0x75, 0x56, 0x70, 0xac, 0x6b, 0xee, 0xa3, 0x44, 0x24, 0x42, 0x57, 0x49,
	0xf9, 0x54, 0x11, 0x5c, 0x14, 0x09, 0x99, 0x09, 0x49, 0x42, 0x2a, 0x19, 0xf9, 0x38, 0x0c, 0x99,
	0xa2, 0x43, 0x12, 0x89, 0x94, 0x1b, 0xdc, 0xa9, 0x04, 0xf5, 0x75, 0x91, 0xb3, 0x48, 0xe4, 0x71,
	0x85, 0x78, 0xef, 0xe1, 0xcd, 0x5c, 0x26, 0xd3, 0x12, 0x78, 0x2d, 0x72, 0xfb, 0x09, 0x6c, 0x4b,
	0xc6, 0x63, 0x96, 0x3b, 0xa0, 0x0f, 0xfc, 0xeb, 0xc0, 0x9c, 0x6c, 0x07, 0xde, 0xa1, 0x71, 0x9c,
	0x33, 0x29, 0x9d, 0x2b, 0x0d, 0x9c, 0x8e, 0xf6, 0x00, 0xb6, 0x69, 0xa4, 0x52, 0xc1, 0x9d, 0x56,
	0x1f, 0xf8, 0x9d, 0x51, 0x17, 0x9f, 0x3f, 0x16, 0x8f, 0x35, 0x10, 0x18, 0x82, 0xf7, 0x0d, 0xc0,
	0x87, 0xb5, 0x97, 0x05, 0x4c, 0xae, 0x04, 0x97, 0xac, 0x2e, 0x0e, 0x9a, 0xe2, 0x9f, 0x80, 0xc9,
	0x82, 0xc5, 0x0b, 0x9a, 0x89, 0x82, 0x2b, 0xe7, 0xaa, 0xdf, 0xf2, 0x6f, 0x46, 0xcf, 0x70, 0xe5,
	0x18, 0x97, 0x8e, 0xb1, 0x71, 0x8c, 0xa7, 0x22, 0xe5, 0x93, 0xd9, 0x76, 0xdf, 0xb3, 0x7e, 0xef,
	0x7b, 0x8f, 0x37, 0x34, 0x5b, 0xbe, 0xf2, 0x9a, 0xed, 0xde, 0xd7, 0x9f, 0x3d, 0x3f, 0x49, 0xd5,
	0xbb, 0x22, 0xc4, 0x91, 0xc8, 0x88, 0xc9, 0xad, 0xba, 0xbd, 0x90, 0xf1, 0x07, 0xa2, 0x36, 0x2b,
	0x26, 0xb5, 0x92, 0x0c, 0xee, 0x9b, 0xe6, 0x71, 0xd5, 0x3b, 0x80, 0x0f, 0xe6, 0x32, 0x99, 0xf1,
	0x54, 0xa5, 0x74, 0xa9, 0x5d, 0x5c, 0xca, 0xcb, 0xfb, 0x0e, 0xe0, 0xd3, 0x7f, 0xb8, 0xb7, 0xcd,
	0xee, 0xe8, 0x0b, 0x80, 0xad, 0xb9, 0x4c, 0xec, 0x09, 0xbc, 0xfb, 0x77, 0x3f, 0x6a, 0xd3, 0xad,
	0x8d, 0xd2, 0x45, 0xff, 0xaf, 0x9f, 0x3d, 0xbf, 0x81, 0xf7, 0x1a, 0xb9, 0xb9, 0x4d, 0x7e, 0x1d,
	0x73, 0xbd, 0xcb, 0xd8, 0x49, 0x6f, 0xf2, 0x7c, 0x7b, 0x40, 0x60, 0x77, 0x40, 0xe0, 0xd7, 0x01,
	0x81, 0xcf, 0x47, 0x64, 0xed, 0x8e, 0xc8, 0xfa, 0x71, 0x44, 0xd6, 0xdb, 0x6e, 0x56, 0x70, 0xb2,
	0x26, 0xe6, 0x0f, 0x2a, 0xdd, 0x85, 0x6d, 0xbd, 0xea, 0x2f, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff,
	0xb1, 0x80, 0xda, 0xd3, 0x57, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	ClaimFor(ctx context.Context, in *MsgClaimFor, opts ...grpc.CallOption) (*MsgClaimForResponse, error)
	InitialClaim(ctx context.Context, in *MsgInitialClaim, opts ...grpc.CallOption) (*MsgInitialClaimResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) ClaimFor(ctx context.Context, in *MsgClaimFor, opts ...grpc.CallOption) (*MsgClaimForResponse, error) {
	out := new(MsgClaimForResponse)
	err := c.cc.Invoke(ctx, "/mun.claim.Msg/ClaimFor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) InitialClaim(ctx context.Context, in *MsgInitialClaim, opts ...grpc.CallOption) (*MsgInitialClaimResponse, error) {
	out := new(MsgInitialClaimResponse)
	err := c.cc.Invoke(ctx, "/mun.claim.Msg/InitialClaim", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	ClaimFor(context.Context, *MsgClaimFor) (*MsgClaimForResponse, error)
	InitialClaim(context.Context, *MsgInitialClaim) (*MsgInitialClaimResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) ClaimFor(ctx context.Context, req *MsgClaimFor) (*MsgClaimForResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClaimFor not implemented")
}
func (*UnimplementedMsgServer) InitialClaim(ctx context.Context, req *MsgInitialClaim) (*MsgInitialClaimResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitialClaim not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_ClaimFor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgClaimFor)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ClaimFor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mun.claim.Msg/ClaimFor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ClaimFor(ctx, req.(*MsgClaimFor))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_InitialClaim_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgInitialClaim)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).InitialClaim(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mun.claim.Msg/InitialClaim",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).InitialClaim(ctx, req.(*MsgInitialClaim))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mun.claim.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ClaimFor",
			Handler:    _Msg_ClaimFor_Handler,
		},
		{
			MethodName: "InitialClaim",
			Handler:    _Msg_InitialClaim_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "claim/tx.proto",
}

func (m *MsgClaimFor) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgClaimFor) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgClaimFor) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Action != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Action))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgClaimForResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgClaimForResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgClaimForResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ClaimedAmount) > 0 {
		for iNdEx := len(m.ClaimedAmount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ClaimedAmount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgInitialClaim) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgInitialClaim) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgInitialClaim) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgInitialClaimResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgInitialClaimResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgInitialClaimResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ClaimedAmount) > 0 {
		for iNdEx := len(m.ClaimedAmount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ClaimedAmount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgClaimFor) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Action != 0 {
		n += 1 + sovTx(uint64(m.Action))
	}
	return n
}

func (m *MsgClaimForResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.ClaimedAmount) > 0 {
		for _, e := range m.ClaimedAmount {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgInitialClaim) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgInitialClaimResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.ClaimedAmount) > 0 {
		for _, e := range m.ClaimedAmount {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgClaimFor) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgClaimFor: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgClaimFor: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Action", wireType)
			}
			m.Action = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Action |= Action(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgClaimForResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgClaimForResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgClaimForResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimedAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClaimedAmount = append(m.ClaimedAmount, types.Coin{})
			if err := m.ClaimedAmount[len(m.ClaimedAmount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgInitialClaim) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgInitialClaim: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgInitialClaim: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgInitialClaimResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgInitialClaimResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgInitialClaimResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimedAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClaimedAmount = append(m.ClaimedAmount, types.Coin{})
			if err := m.ClaimedAmount[len(m.ClaimedAmount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
