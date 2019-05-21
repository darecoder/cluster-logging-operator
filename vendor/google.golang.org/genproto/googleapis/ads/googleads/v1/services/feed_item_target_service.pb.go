// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/feed_item_target_service.proto

package services // import "google.golang.org/genproto/googleapis/ads/googleads/v1/services"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import resources "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Request message for [FeedItemTargetService.GetFeedItemTarget][google.ads.googleads.v1.services.FeedItemTargetService.GetFeedItemTarget].
type GetFeedItemTargetRequest struct {
	// The resource name of the feed item targets to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFeedItemTargetRequest) Reset()         { *m = GetFeedItemTargetRequest{} }
func (m *GetFeedItemTargetRequest) String() string { return proto.CompactTextString(m) }
func (*GetFeedItemTargetRequest) ProtoMessage()    {}
func (*GetFeedItemTargetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_feed_item_target_service_7676c6ec99ca33e2, []int{0}
}
func (m *GetFeedItemTargetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFeedItemTargetRequest.Unmarshal(m, b)
}
func (m *GetFeedItemTargetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFeedItemTargetRequest.Marshal(b, m, deterministic)
}
func (dst *GetFeedItemTargetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFeedItemTargetRequest.Merge(dst, src)
}
func (m *GetFeedItemTargetRequest) XXX_Size() int {
	return xxx_messageInfo_GetFeedItemTargetRequest.Size(m)
}
func (m *GetFeedItemTargetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFeedItemTargetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetFeedItemTargetRequest proto.InternalMessageInfo

func (m *GetFeedItemTargetRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [FeedItemTargetService.MutateFeedItemTargets][google.ads.googleads.v1.services.FeedItemTargetService.MutateFeedItemTargets].
type MutateFeedItemTargetsRequest struct {
	// The ID of the customer whose feed item targets are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The list of operations to perform on individual feed item targets.
	Operations           []*FeedItemTargetOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *MutateFeedItemTargetsRequest) Reset()         { *m = MutateFeedItemTargetsRequest{} }
func (m *MutateFeedItemTargetsRequest) String() string { return proto.CompactTextString(m) }
func (*MutateFeedItemTargetsRequest) ProtoMessage()    {}
func (*MutateFeedItemTargetsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_feed_item_target_service_7676c6ec99ca33e2, []int{1}
}
func (m *MutateFeedItemTargetsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateFeedItemTargetsRequest.Unmarshal(m, b)
}
func (m *MutateFeedItemTargetsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateFeedItemTargetsRequest.Marshal(b, m, deterministic)
}
func (dst *MutateFeedItemTargetsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateFeedItemTargetsRequest.Merge(dst, src)
}
func (m *MutateFeedItemTargetsRequest) XXX_Size() int {
	return xxx_messageInfo_MutateFeedItemTargetsRequest.Size(m)
}
func (m *MutateFeedItemTargetsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateFeedItemTargetsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateFeedItemTargetsRequest proto.InternalMessageInfo

func (m *MutateFeedItemTargetsRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateFeedItemTargetsRequest) GetOperations() []*FeedItemTargetOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

// A single operation (create, remove) on an feed item target.
type FeedItemTargetOperation struct {
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*FeedItemTargetOperation_Create
	//	*FeedItemTargetOperation_Remove
	Operation            isFeedItemTargetOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *FeedItemTargetOperation) Reset()         { *m = FeedItemTargetOperation{} }
func (m *FeedItemTargetOperation) String() string { return proto.CompactTextString(m) }
func (*FeedItemTargetOperation) ProtoMessage()    {}
func (*FeedItemTargetOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_feed_item_target_service_7676c6ec99ca33e2, []int{2}
}
func (m *FeedItemTargetOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeedItemTargetOperation.Unmarshal(m, b)
}
func (m *FeedItemTargetOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeedItemTargetOperation.Marshal(b, m, deterministic)
}
func (dst *FeedItemTargetOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedItemTargetOperation.Merge(dst, src)
}
func (m *FeedItemTargetOperation) XXX_Size() int {
	return xxx_messageInfo_FeedItemTargetOperation.Size(m)
}
func (m *FeedItemTargetOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedItemTargetOperation.DiscardUnknown(m)
}

var xxx_messageInfo_FeedItemTargetOperation proto.InternalMessageInfo

type isFeedItemTargetOperation_Operation interface {
	isFeedItemTargetOperation_Operation()
}

type FeedItemTargetOperation_Create struct {
	Create *resources.FeedItemTarget `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type FeedItemTargetOperation_Remove struct {
	Remove string `protobuf:"bytes,2,opt,name=remove,proto3,oneof"`
}

func (*FeedItemTargetOperation_Create) isFeedItemTargetOperation_Operation() {}

func (*FeedItemTargetOperation_Remove) isFeedItemTargetOperation_Operation() {}

func (m *FeedItemTargetOperation) GetOperation() isFeedItemTargetOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *FeedItemTargetOperation) GetCreate() *resources.FeedItemTarget {
	if x, ok := m.GetOperation().(*FeedItemTargetOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *FeedItemTargetOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*FeedItemTargetOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*FeedItemTargetOperation) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _FeedItemTargetOperation_OneofMarshaler, _FeedItemTargetOperation_OneofUnmarshaler, _FeedItemTargetOperation_OneofSizer, []interface{}{
		(*FeedItemTargetOperation_Create)(nil),
		(*FeedItemTargetOperation_Remove)(nil),
	}
}

func _FeedItemTargetOperation_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*FeedItemTargetOperation)
	// operation
	switch x := m.Operation.(type) {
	case *FeedItemTargetOperation_Create:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Create); err != nil {
			return err
		}
	case *FeedItemTargetOperation_Remove:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Remove)
	case nil:
	default:
		return fmt.Errorf("FeedItemTargetOperation.Operation has unexpected type %T", x)
	}
	return nil
}

func _FeedItemTargetOperation_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*FeedItemTargetOperation)
	switch tag {
	case 1: // operation.create
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(resources.FeedItemTarget)
		err := b.DecodeMessage(msg)
		m.Operation = &FeedItemTargetOperation_Create{msg}
		return true, err
	case 2: // operation.remove
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Operation = &FeedItemTargetOperation_Remove{x}
		return true, err
	default:
		return false, nil
	}
}

func _FeedItemTargetOperation_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*FeedItemTargetOperation)
	// operation
	switch x := m.Operation.(type) {
	case *FeedItemTargetOperation_Create:
		s := proto.Size(x.Create)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *FeedItemTargetOperation_Remove:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Remove)))
		n += len(x.Remove)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Response message for an feed item target mutate.
type MutateFeedItemTargetsResponse struct {
	// All results for the mutate.
	Results              []*MutateFeedItemTargetResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *MutateFeedItemTargetsResponse) Reset()         { *m = MutateFeedItemTargetsResponse{} }
func (m *MutateFeedItemTargetsResponse) String() string { return proto.CompactTextString(m) }
func (*MutateFeedItemTargetsResponse) ProtoMessage()    {}
func (*MutateFeedItemTargetsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_feed_item_target_service_7676c6ec99ca33e2, []int{3}
}
func (m *MutateFeedItemTargetsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateFeedItemTargetsResponse.Unmarshal(m, b)
}
func (m *MutateFeedItemTargetsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateFeedItemTargetsResponse.Marshal(b, m, deterministic)
}
func (dst *MutateFeedItemTargetsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateFeedItemTargetsResponse.Merge(dst, src)
}
func (m *MutateFeedItemTargetsResponse) XXX_Size() int {
	return xxx_messageInfo_MutateFeedItemTargetsResponse.Size(m)
}
func (m *MutateFeedItemTargetsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateFeedItemTargetsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateFeedItemTargetsResponse proto.InternalMessageInfo

func (m *MutateFeedItemTargetsResponse) GetResults() []*MutateFeedItemTargetResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for the feed item target mutate.
type MutateFeedItemTargetResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateFeedItemTargetResult) Reset()         { *m = MutateFeedItemTargetResult{} }
func (m *MutateFeedItemTargetResult) String() string { return proto.CompactTextString(m) }
func (*MutateFeedItemTargetResult) ProtoMessage()    {}
func (*MutateFeedItemTargetResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_feed_item_target_service_7676c6ec99ca33e2, []int{4}
}
func (m *MutateFeedItemTargetResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateFeedItemTargetResult.Unmarshal(m, b)
}
func (m *MutateFeedItemTargetResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateFeedItemTargetResult.Marshal(b, m, deterministic)
}
func (dst *MutateFeedItemTargetResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateFeedItemTargetResult.Merge(dst, src)
}
func (m *MutateFeedItemTargetResult) XXX_Size() int {
	return xxx_messageInfo_MutateFeedItemTargetResult.Size(m)
}
func (m *MutateFeedItemTargetResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateFeedItemTargetResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateFeedItemTargetResult proto.InternalMessageInfo

func (m *MutateFeedItemTargetResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetFeedItemTargetRequest)(nil), "google.ads.googleads.v1.services.GetFeedItemTargetRequest")
	proto.RegisterType((*MutateFeedItemTargetsRequest)(nil), "google.ads.googleads.v1.services.MutateFeedItemTargetsRequest")
	proto.RegisterType((*FeedItemTargetOperation)(nil), "google.ads.googleads.v1.services.FeedItemTargetOperation")
	proto.RegisterType((*MutateFeedItemTargetsResponse)(nil), "google.ads.googleads.v1.services.MutateFeedItemTargetsResponse")
	proto.RegisterType((*MutateFeedItemTargetResult)(nil), "google.ads.googleads.v1.services.MutateFeedItemTargetResult")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FeedItemTargetServiceClient is the client API for FeedItemTargetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FeedItemTargetServiceClient interface {
	// Returns the requested feed item targets in full detail.
	GetFeedItemTarget(ctx context.Context, in *GetFeedItemTargetRequest, opts ...grpc.CallOption) (*resources.FeedItemTarget, error)
	// Creates or removes feed item targets. Operation statuses are returned.
	MutateFeedItemTargets(ctx context.Context, in *MutateFeedItemTargetsRequest, opts ...grpc.CallOption) (*MutateFeedItemTargetsResponse, error)
}

type feedItemTargetServiceClient struct {
	cc *grpc.ClientConn
}

func NewFeedItemTargetServiceClient(cc *grpc.ClientConn) FeedItemTargetServiceClient {
	return &feedItemTargetServiceClient{cc}
}

func (c *feedItemTargetServiceClient) GetFeedItemTarget(ctx context.Context, in *GetFeedItemTargetRequest, opts ...grpc.CallOption) (*resources.FeedItemTarget, error) {
	out := new(resources.FeedItemTarget)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.FeedItemTargetService/GetFeedItemTarget", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedItemTargetServiceClient) MutateFeedItemTargets(ctx context.Context, in *MutateFeedItemTargetsRequest, opts ...grpc.CallOption) (*MutateFeedItemTargetsResponse, error) {
	out := new(MutateFeedItemTargetsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.FeedItemTargetService/MutateFeedItemTargets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FeedItemTargetServiceServer is the server API for FeedItemTargetService service.
type FeedItemTargetServiceServer interface {
	// Returns the requested feed item targets in full detail.
	GetFeedItemTarget(context.Context, *GetFeedItemTargetRequest) (*resources.FeedItemTarget, error)
	// Creates or removes feed item targets. Operation statuses are returned.
	MutateFeedItemTargets(context.Context, *MutateFeedItemTargetsRequest) (*MutateFeedItemTargetsResponse, error)
}

func RegisterFeedItemTargetServiceServer(s *grpc.Server, srv FeedItemTargetServiceServer) {
	s.RegisterService(&_FeedItemTargetService_serviceDesc, srv)
}

func _FeedItemTargetService_GetFeedItemTarget_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFeedItemTargetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedItemTargetServiceServer).GetFeedItemTarget(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.FeedItemTargetService/GetFeedItemTarget",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedItemTargetServiceServer).GetFeedItemTarget(ctx, req.(*GetFeedItemTargetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeedItemTargetService_MutateFeedItemTargets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateFeedItemTargetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedItemTargetServiceServer).MutateFeedItemTargets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.FeedItemTargetService/MutateFeedItemTargets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedItemTargetServiceServer).MutateFeedItemTargets(ctx, req.(*MutateFeedItemTargetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FeedItemTargetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.FeedItemTargetService",
	HandlerType: (*FeedItemTargetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFeedItemTarget",
			Handler:    _FeedItemTargetService_GetFeedItemTarget_Handler,
		},
		{
			MethodName: "MutateFeedItemTargets",
			Handler:    _FeedItemTargetService_MutateFeedItemTargets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/feed_item_target_service.proto",
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/feed_item_target_service.proto", fileDescriptor_feed_item_target_service_7676c6ec99ca33e2)
}

var fileDescriptor_feed_item_target_service_7676c6ec99ca33e2 = []byte{
	// 559 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0x3f, 0x6f, 0xd3, 0x40,
	0x14, 0xc0, 0xb1, 0x2b, 0x05, 0xf5, 0x02, 0x03, 0x27, 0x55, 0x44, 0x56, 0x11, 0x91, 0xe9, 0x50,
	0x65, 0x38, 0xcb, 0x29, 0x03, 0x39, 0x28, 0x91, 0x33, 0x90, 0x56, 0x08, 0xa8, 0x0c, 0x8a, 0x04,
	0x8a, 0x14, 0x1d, 0xf1, 0xc3, 0xb2, 0x14, 0xfb, 0xc2, 0xdd, 0x25, 0x0c, 0x55, 0x07, 0xd8, 0x99,
	0xd8, 0x18, 0x19, 0x59, 0x99, 0xf9, 0x02, 0xac, 0x4c, 0xec, 0x4c, 0x7c, 0x0a, 0x64, 0x5f, 0xce,
	0x24, 0x25, 0x56, 0xa0, 0xdb, 0xbb, 0xbb, 0xf7, 0x7e, 0xef, 0xef, 0x3d, 0xd4, 0x8d, 0x39, 0x8f,
	0x27, 0xe0, 0xb1, 0x48, 0x7a, 0x5a, 0xcc, 0xa5, 0xb9, 0xef, 0x49, 0x10, 0xf3, 0x64, 0x0c, 0xd2,
	0x7b, 0x05, 0x10, 0x8d, 0x12, 0x05, 0xe9, 0x48, 0x31, 0x11, 0x83, 0x1a, 0x2d, 0x5e, 0xc8, 0x54,
	0x70, 0xc5, 0x71, 0x53, 0x5b, 0x11, 0x16, 0x49, 0x52, 0x02, 0xc8, 0xdc, 0x27, 0x06, 0xe0, 0xdc,
	0xa9, 0x72, 0x21, 0x40, 0xf2, 0x99, 0x58, 0xe7, 0x43, 0xb3, 0x9d, 0x5d, 0x63, 0x39, 0x4d, 0x3c,
	0x96, 0x65, 0x5c, 0x31, 0x95, 0xf0, 0x4c, 0xea, 0x57, 0xb7, 0x8b, 0x1a, 0x7d, 0x50, 0x0f, 0x00,
	0xa2, 0x63, 0x05, 0xe9, 0xb3, 0xc2, 0x30, 0x84, 0xd7, 0x33, 0x90, 0x0a, 0xdf, 0x42, 0x57, 0x0d,
	0x7d, 0x94, 0xb1, 0x14, 0x1a, 0x56, 0xd3, 0xda, 0xdf, 0x0e, 0xaf, 0x98, 0xcb, 0xc7, 0x2c, 0x05,
	0xf7, 0xa3, 0x85, 0x76, 0x1f, 0xcd, 0x14, 0x53, 0xb0, 0x0a, 0x91, 0x86, 0x72, 0x13, 0xd5, 0xc7,
	0x33, 0xa9, 0x78, 0x0a, 0x62, 0x94, 0x44, 0x0b, 0x06, 0x32, 0x57, 0xc7, 0x11, 0x7e, 0x8e, 0x10,
	0x9f, 0x82, 0xd0, 0x61, 0x35, 0xec, 0xe6, 0xd6, 0x7e, 0xbd, 0xdd, 0x21, 0x9b, 0x2a, 0x42, 0x56,
	0xdd, 0x3d, 0x31, 0x84, 0x70, 0x09, 0xe6, 0xbe, 0xb7, 0xd0, 0xf5, 0x0a, 0x3d, 0xfc, 0x10, 0xd5,
	0xc6, 0x02, 0x98, 0xd2, 0x69, 0xd5, 0xdb, 0x7e, 0xa5, 0xcb, 0xb2, 0xc4, 0xe7, 0x7c, 0x1e, 0x5d,
	0x0a, 0x17, 0x08, 0xdc, 0x40, 0x35, 0x01, 0x29, 0x9f, 0x43, 0xc3, 0xce, 0xf3, 0xcb, 0x5f, 0xf4,
	0xb9, 0x57, 0x47, 0xdb, 0x65, 0x40, 0xee, 0x1b, 0x74, 0xa3, 0xa2, 0x56, 0x72, 0xca, 0x33, 0x09,
	0x78, 0x80, 0x2e, 0x0b, 0x90, 0xb3, 0x89, 0x32, 0x85, 0xb8, 0xb7, 0xb9, 0x10, 0xeb, 0x88, 0x61,
	0x01, 0x09, 0x0d, 0xcc, 0x0d, 0x90, 0x53, 0xad, 0xf6, 0x4f, 0x8d, 0x6e, 0x7f, 0xd9, 0x42, 0x3b,
	0xab, 0xd6, 0x4f, 0x75, 0x04, 0xf8, 0xab, 0x85, 0xae, 0xfd, 0x35, 0x44, 0x98, 0x6e, 0x8e, 0xbc,
	0x6a, 0xf2, 0x9c, 0xff, 0xef, 0x85, 0xdb, 0x79, 0xf7, 0xfd, 0xe7, 0x07, 0xfb, 0x00, 0xfb, 0xf9,
	0xa7, 0x38, 0x5d, 0x49, 0xe7, 0xd0, 0x0c, 0x9b, 0xf4, 0x5a, 0xc5, 0x2f, 0x59, 0xaa, 0xbc, 0xd7,
	0x3a, 0xc3, 0x3f, 0x2c, 0xb4, 0xb3, 0xb6, 0x2d, 0xf8, 0xfe, 0xc5, 0xaa, 0x6f, 0x66, 0xdf, 0xe9,
	0x5e, 0xd8, 0x5e, 0xcf, 0x83, 0xdb, 0x2d, 0xb2, 0xea, 0xb8, 0xb7, 0xf3, 0xac, 0xfe, 0xa4, 0x71,
	0xba, 0xf4, 0xa3, 0x0e, 0x5b, 0x67, 0xe7, 0x93, 0xa2, 0x69, 0x01, 0xa5, 0x56, 0xab, 0xf7, 0xd6,
	0x46, 0x7b, 0x63, 0x9e, 0x6e, 0x8c, 0xa3, 0xe7, 0xac, 0xed, 0xed, 0x49, 0xbe, 0x24, 0x4e, 0xac,
	0x17, 0x47, 0x0b, 0xfb, 0x98, 0x4f, 0x58, 0x16, 0x13, 0x2e, 0x62, 0x2f, 0x86, 0xac, 0x58, 0x21,
	0x66, 0x1d, 0x4d, 0x13, 0x59, 0xbd, 0x00, 0xef, 0x1a, 0xe1, 0x93, 0xbd, 0xd5, 0x0f, 0x82, 0xcf,
	0x76, 0xb3, 0xaf, 0x81, 0x41, 0x24, 0x89, 0x16, 0x73, 0x69, 0xe0, 0x93, 0x85, 0x63, 0xf9, 0xcd,
	0xa8, 0x0c, 0x83, 0x48, 0x0e, 0x4b, 0x95, 0xe1, 0xc0, 0x1f, 0x1a, 0x95, 0x5f, 0xf6, 0x9e, 0xbe,
	0xa7, 0x34, 0x88, 0x24, 0xa5, 0xa5, 0x12, 0xa5, 0x03, 0x9f, 0x52, 0xa3, 0xf6, 0xb2, 0x56, 0xc4,
	0x79, 0xf0, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xd9, 0xad, 0xe7, 0xfc, 0xa7, 0x05, 0x00, 0x00,
}