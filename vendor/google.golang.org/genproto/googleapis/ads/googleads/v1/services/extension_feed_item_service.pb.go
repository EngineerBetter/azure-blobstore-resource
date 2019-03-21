// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/extension_feed_item_service.proto

package services // import "google.golang.org/genproto/googleapis/ads/googleads/v1/services"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/wrappers"
import resources "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import field_mask "google.golang.org/genproto/protobuf/field_mask"

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

// Request message for [ExtensionFeedItemService.GetExtensionFeedItem][google.ads.googleads.v1.services.ExtensionFeedItemService.GetExtensionFeedItem].
type GetExtensionFeedItemRequest struct {
	// The resource name of the extension feed item to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetExtensionFeedItemRequest) Reset()         { *m = GetExtensionFeedItemRequest{} }
func (m *GetExtensionFeedItemRequest) String() string { return proto.CompactTextString(m) }
func (*GetExtensionFeedItemRequest) ProtoMessage()    {}
func (*GetExtensionFeedItemRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_extension_feed_item_service_600f2d34df06b74a, []int{0}
}
func (m *GetExtensionFeedItemRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetExtensionFeedItemRequest.Unmarshal(m, b)
}
func (m *GetExtensionFeedItemRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetExtensionFeedItemRequest.Marshal(b, m, deterministic)
}
func (dst *GetExtensionFeedItemRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetExtensionFeedItemRequest.Merge(dst, src)
}
func (m *GetExtensionFeedItemRequest) XXX_Size() int {
	return xxx_messageInfo_GetExtensionFeedItemRequest.Size(m)
}
func (m *GetExtensionFeedItemRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetExtensionFeedItemRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetExtensionFeedItemRequest proto.InternalMessageInfo

func (m *GetExtensionFeedItemRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [ExtensionFeedItemService.MutateExtensionFeedItems][google.ads.googleads.v1.services.ExtensionFeedItemService.MutateExtensionFeedItems].
type MutateExtensionFeedItemsRequest struct {
	// The ID of the customer whose extension feed items are being
	// modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The list of operations to perform on individual extension feed items.
	Operations []*ExtensionFeedItemOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly         bool     `protobuf:"varint,4,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateExtensionFeedItemsRequest) Reset()         { *m = MutateExtensionFeedItemsRequest{} }
func (m *MutateExtensionFeedItemsRequest) String() string { return proto.CompactTextString(m) }
func (*MutateExtensionFeedItemsRequest) ProtoMessage()    {}
func (*MutateExtensionFeedItemsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_extension_feed_item_service_600f2d34df06b74a, []int{1}
}
func (m *MutateExtensionFeedItemsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateExtensionFeedItemsRequest.Unmarshal(m, b)
}
func (m *MutateExtensionFeedItemsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateExtensionFeedItemsRequest.Marshal(b, m, deterministic)
}
func (dst *MutateExtensionFeedItemsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateExtensionFeedItemsRequest.Merge(dst, src)
}
func (m *MutateExtensionFeedItemsRequest) XXX_Size() int {
	return xxx_messageInfo_MutateExtensionFeedItemsRequest.Size(m)
}
func (m *MutateExtensionFeedItemsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateExtensionFeedItemsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateExtensionFeedItemsRequest proto.InternalMessageInfo

func (m *MutateExtensionFeedItemsRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateExtensionFeedItemsRequest) GetOperations() []*ExtensionFeedItemOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *MutateExtensionFeedItemsRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// A single operation (create, update, remove) on an extension feed item.
type ExtensionFeedItemOperation struct {
	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*ExtensionFeedItemOperation_Create
	//	*ExtensionFeedItemOperation_Update
	//	*ExtensionFeedItemOperation_Remove
	Operation            isExtensionFeedItemOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                               `json:"-"`
	XXX_unrecognized     []byte                                 `json:"-"`
	XXX_sizecache        int32                                  `json:"-"`
}

func (m *ExtensionFeedItemOperation) Reset()         { *m = ExtensionFeedItemOperation{} }
func (m *ExtensionFeedItemOperation) String() string { return proto.CompactTextString(m) }
func (*ExtensionFeedItemOperation) ProtoMessage()    {}
func (*ExtensionFeedItemOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_extension_feed_item_service_600f2d34df06b74a, []int{2}
}
func (m *ExtensionFeedItemOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtensionFeedItemOperation.Unmarshal(m, b)
}
func (m *ExtensionFeedItemOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtensionFeedItemOperation.Marshal(b, m, deterministic)
}
func (dst *ExtensionFeedItemOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtensionFeedItemOperation.Merge(dst, src)
}
func (m *ExtensionFeedItemOperation) XXX_Size() int {
	return xxx_messageInfo_ExtensionFeedItemOperation.Size(m)
}
func (m *ExtensionFeedItemOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtensionFeedItemOperation.DiscardUnknown(m)
}

var xxx_messageInfo_ExtensionFeedItemOperation proto.InternalMessageInfo

func (m *ExtensionFeedItemOperation) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type isExtensionFeedItemOperation_Operation interface {
	isExtensionFeedItemOperation_Operation()
}

type ExtensionFeedItemOperation_Create struct {
	Create *resources.ExtensionFeedItem `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type ExtensionFeedItemOperation_Update struct {
	Update *resources.ExtensionFeedItem `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

type ExtensionFeedItemOperation_Remove struct {
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*ExtensionFeedItemOperation_Create) isExtensionFeedItemOperation_Operation() {}

func (*ExtensionFeedItemOperation_Update) isExtensionFeedItemOperation_Operation() {}

func (*ExtensionFeedItemOperation_Remove) isExtensionFeedItemOperation_Operation() {}

func (m *ExtensionFeedItemOperation) GetOperation() isExtensionFeedItemOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *ExtensionFeedItemOperation) GetCreate() *resources.ExtensionFeedItem {
	if x, ok := m.GetOperation().(*ExtensionFeedItemOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *ExtensionFeedItemOperation) GetUpdate() *resources.ExtensionFeedItem {
	if x, ok := m.GetOperation().(*ExtensionFeedItemOperation_Update); ok {
		return x.Update
	}
	return nil
}

func (m *ExtensionFeedItemOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*ExtensionFeedItemOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ExtensionFeedItemOperation) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ExtensionFeedItemOperation_OneofMarshaler, _ExtensionFeedItemOperation_OneofUnmarshaler, _ExtensionFeedItemOperation_OneofSizer, []interface{}{
		(*ExtensionFeedItemOperation_Create)(nil),
		(*ExtensionFeedItemOperation_Update)(nil),
		(*ExtensionFeedItemOperation_Remove)(nil),
	}
}

func _ExtensionFeedItemOperation_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ExtensionFeedItemOperation)
	// operation
	switch x := m.Operation.(type) {
	case *ExtensionFeedItemOperation_Create:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Create); err != nil {
			return err
		}
	case *ExtensionFeedItemOperation_Update:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Update); err != nil {
			return err
		}
	case *ExtensionFeedItemOperation_Remove:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Remove)
	case nil:
	default:
		return fmt.Errorf("ExtensionFeedItemOperation.Operation has unexpected type %T", x)
	}
	return nil
}

func _ExtensionFeedItemOperation_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ExtensionFeedItemOperation)
	switch tag {
	case 1: // operation.create
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(resources.ExtensionFeedItem)
		err := b.DecodeMessage(msg)
		m.Operation = &ExtensionFeedItemOperation_Create{msg}
		return true, err
	case 2: // operation.update
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(resources.ExtensionFeedItem)
		err := b.DecodeMessage(msg)
		m.Operation = &ExtensionFeedItemOperation_Update{msg}
		return true, err
	case 3: // operation.remove
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Operation = &ExtensionFeedItemOperation_Remove{x}
		return true, err
	default:
		return false, nil
	}
}

func _ExtensionFeedItemOperation_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ExtensionFeedItemOperation)
	// operation
	switch x := m.Operation.(type) {
	case *ExtensionFeedItemOperation_Create:
		s := proto.Size(x.Create)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ExtensionFeedItemOperation_Update:
		s := proto.Size(x.Update)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ExtensionFeedItemOperation_Remove:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Remove)))
		n += len(x.Remove)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Response message for an extension feed item mutate.
type MutateExtensionFeedItemsResponse struct {
	// All results for the mutate.
	Results              []*MutateExtensionFeedItemResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *MutateExtensionFeedItemsResponse) Reset()         { *m = MutateExtensionFeedItemsResponse{} }
func (m *MutateExtensionFeedItemsResponse) String() string { return proto.CompactTextString(m) }
func (*MutateExtensionFeedItemsResponse) ProtoMessage()    {}
func (*MutateExtensionFeedItemsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_extension_feed_item_service_600f2d34df06b74a, []int{3}
}
func (m *MutateExtensionFeedItemsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateExtensionFeedItemsResponse.Unmarshal(m, b)
}
func (m *MutateExtensionFeedItemsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateExtensionFeedItemsResponse.Marshal(b, m, deterministic)
}
func (dst *MutateExtensionFeedItemsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateExtensionFeedItemsResponse.Merge(dst, src)
}
func (m *MutateExtensionFeedItemsResponse) XXX_Size() int {
	return xxx_messageInfo_MutateExtensionFeedItemsResponse.Size(m)
}
func (m *MutateExtensionFeedItemsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateExtensionFeedItemsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateExtensionFeedItemsResponse proto.InternalMessageInfo

func (m *MutateExtensionFeedItemsResponse) GetResults() []*MutateExtensionFeedItemResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for the extension feed item mutate.
type MutateExtensionFeedItemResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateExtensionFeedItemResult) Reset()         { *m = MutateExtensionFeedItemResult{} }
func (m *MutateExtensionFeedItemResult) String() string { return proto.CompactTextString(m) }
func (*MutateExtensionFeedItemResult) ProtoMessage()    {}
func (*MutateExtensionFeedItemResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_extension_feed_item_service_600f2d34df06b74a, []int{4}
}
func (m *MutateExtensionFeedItemResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateExtensionFeedItemResult.Unmarshal(m, b)
}
func (m *MutateExtensionFeedItemResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateExtensionFeedItemResult.Marshal(b, m, deterministic)
}
func (dst *MutateExtensionFeedItemResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateExtensionFeedItemResult.Merge(dst, src)
}
func (m *MutateExtensionFeedItemResult) XXX_Size() int {
	return xxx_messageInfo_MutateExtensionFeedItemResult.Size(m)
}
func (m *MutateExtensionFeedItemResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateExtensionFeedItemResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateExtensionFeedItemResult proto.InternalMessageInfo

func (m *MutateExtensionFeedItemResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetExtensionFeedItemRequest)(nil), "google.ads.googleads.v1.services.GetExtensionFeedItemRequest")
	proto.RegisterType((*MutateExtensionFeedItemsRequest)(nil), "google.ads.googleads.v1.services.MutateExtensionFeedItemsRequest")
	proto.RegisterType((*ExtensionFeedItemOperation)(nil), "google.ads.googleads.v1.services.ExtensionFeedItemOperation")
	proto.RegisterType((*MutateExtensionFeedItemsResponse)(nil), "google.ads.googleads.v1.services.MutateExtensionFeedItemsResponse")
	proto.RegisterType((*MutateExtensionFeedItemResult)(nil), "google.ads.googleads.v1.services.MutateExtensionFeedItemResult")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ExtensionFeedItemServiceClient is the client API for ExtensionFeedItemService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExtensionFeedItemServiceClient interface {
	// Returns the requested extension feed item in full detail.
	GetExtensionFeedItem(ctx context.Context, in *GetExtensionFeedItemRequest, opts ...grpc.CallOption) (*resources.ExtensionFeedItem, error)
	// Creates, updates, or removes extension feed items. Operation
	// statuses are returned.
	MutateExtensionFeedItems(ctx context.Context, in *MutateExtensionFeedItemsRequest, opts ...grpc.CallOption) (*MutateExtensionFeedItemsResponse, error)
}

type extensionFeedItemServiceClient struct {
	cc *grpc.ClientConn
}

func NewExtensionFeedItemServiceClient(cc *grpc.ClientConn) ExtensionFeedItemServiceClient {
	return &extensionFeedItemServiceClient{cc}
}

func (c *extensionFeedItemServiceClient) GetExtensionFeedItem(ctx context.Context, in *GetExtensionFeedItemRequest, opts ...grpc.CallOption) (*resources.ExtensionFeedItem, error) {
	out := new(resources.ExtensionFeedItem)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.ExtensionFeedItemService/GetExtensionFeedItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *extensionFeedItemServiceClient) MutateExtensionFeedItems(ctx context.Context, in *MutateExtensionFeedItemsRequest, opts ...grpc.CallOption) (*MutateExtensionFeedItemsResponse, error) {
	out := new(MutateExtensionFeedItemsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.ExtensionFeedItemService/MutateExtensionFeedItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionFeedItemServiceServer is the server API for ExtensionFeedItemService service.
type ExtensionFeedItemServiceServer interface {
	// Returns the requested extension feed item in full detail.
	GetExtensionFeedItem(context.Context, *GetExtensionFeedItemRequest) (*resources.ExtensionFeedItem, error)
	// Creates, updates, or removes extension feed items. Operation
	// statuses are returned.
	MutateExtensionFeedItems(context.Context, *MutateExtensionFeedItemsRequest) (*MutateExtensionFeedItemsResponse, error)
}

func RegisterExtensionFeedItemServiceServer(s *grpc.Server, srv ExtensionFeedItemServiceServer) {
	s.RegisterService(&_ExtensionFeedItemService_serviceDesc, srv)
}

func _ExtensionFeedItemService_GetExtensionFeedItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetExtensionFeedItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExtensionFeedItemServiceServer).GetExtensionFeedItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.ExtensionFeedItemService/GetExtensionFeedItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExtensionFeedItemServiceServer).GetExtensionFeedItem(ctx, req.(*GetExtensionFeedItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExtensionFeedItemService_MutateExtensionFeedItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateExtensionFeedItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExtensionFeedItemServiceServer).MutateExtensionFeedItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.ExtensionFeedItemService/MutateExtensionFeedItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExtensionFeedItemServiceServer).MutateExtensionFeedItems(ctx, req.(*MutateExtensionFeedItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ExtensionFeedItemService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.ExtensionFeedItemService",
	HandlerType: (*ExtensionFeedItemServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetExtensionFeedItem",
			Handler:    _ExtensionFeedItemService_GetExtensionFeedItem_Handler,
		},
		{
			MethodName: "MutateExtensionFeedItems",
			Handler:    _ExtensionFeedItemService_MutateExtensionFeedItems_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/extension_feed_item_service.proto",
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/extension_feed_item_service.proto", fileDescriptor_extension_feed_item_service_600f2d34df06b74a)
}

var fileDescriptor_extension_feed_item_service_600f2d34df06b74a = []byte{
	// 656 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x95, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0xb1, 0x83, 0x0a, 0xdd, 0x94, 0x8b, 0xc5, 0xc1, 0x0a, 0x94, 0x46, 0xa6, 0x87, 0x2a,
	0x07, 0x5b, 0x09, 0x15, 0x48, 0x6e, 0x2b, 0x14, 0x0b, 0xfa, 0x71, 0xe8, 0x87, 0x8c, 0x54, 0x09,
	0x14, 0xc9, 0xda, 0xc6, 0xd3, 0xc8, 0xaa, 0xed, 0x35, 0xbb, 0xeb, 0x40, 0x55, 0xf5, 0x82, 0x78,
	0x03, 0x5e, 0x00, 0x71, 0xe4, 0x3d, 0x38, 0xc0, 0x81, 0x0b, 0xaf, 0x00, 0x17, 0x9e, 0x02, 0xad,
	0xd7, 0xeb, 0x16, 0x52, 0x37, 0xa8, 0xbd, 0x4d, 0x76, 0xc7, 0xbf, 0xd9, 0xff, 0xfc, 0x77, 0x36,
	0xc8, 0x1b, 0x11, 0x32, 0x8a, 0xc1, 0xc1, 0x21, 0x73, 0x64, 0x28, 0xa2, 0x71, 0xd7, 0x61, 0x40,
	0xc7, 0xd1, 0x10, 0x98, 0x03, 0x6f, 0x39, 0xa4, 0x2c, 0x22, 0x69, 0x70, 0x08, 0x10, 0x06, 0x11,
	0x87, 0x24, 0x28, 0x37, 0xed, 0x8c, 0x12, 0x4e, 0x8c, 0xb6, 0xfc, 0xd0, 0xc6, 0x21, 0xb3, 0x2b,
	0x86, 0x3d, 0xee, 0xda, 0x8a, 0xd1, 0x5a, 0xa9, 0xab, 0x42, 0x81, 0x91, 0x9c, 0xd6, 0x94, 0x91,
	0xf8, 0xd6, 0x7d, 0xf5, 0x71, 0x16, 0x39, 0x38, 0x4d, 0x09, 0xc7, 0x3c, 0x22, 0x29, 0x2b, 0x77,
	0xcb, 0xe2, 0x4e, 0xf1, 0xeb, 0x20, 0x3f, 0x74, 0x0e, 0x23, 0x88, 0xc3, 0x20, 0xc1, 0xec, 0xa8,
	0xcc, 0x78, 0xf0, 0x6f, 0xc6, 0x1b, 0x8a, 0xb3, 0x0c, 0x68, 0x49, 0xb0, 0x3c, 0x74, 0x6f, 0x03,
	0xf8, 0x73, 0x55, 0x7f, 0x1d, 0x20, 0xdc, 0xe2, 0x90, 0xf8, 0xf0, 0x3a, 0x07, 0xc6, 0x8d, 0x87,
	0xe8, 0x8e, 0x3a, 0x65, 0x90, 0xe2, 0x04, 0x4c, 0xad, 0xad, 0x2d, 0xcd, 0xfa, 0x73, 0x6a, 0x71,
	0x07, 0x27, 0x60, 0x7d, 0xd1, 0xd0, 0xc2, 0x76, 0xce, 0x31, 0x87, 0x09, 0x0e, 0x53, 0xa0, 0x05,
	0xd4, 0x1c, 0xe6, 0x8c, 0x93, 0x04, 0x68, 0x10, 0x85, 0x25, 0x06, 0xa9, 0xa5, 0xad, 0xd0, 0x18,
	0x20, 0x44, 0x32, 0xa0, 0x52, 0x9e, 0xa9, 0xb7, 0x1b, 0x4b, 0xcd, 0xde, 0xaa, 0x3d, 0xad, 0xb9,
	0xf6, 0x44, 0xc5, 0x5d, 0x05, 0xf1, 0xcf, 0xf1, 0x84, 0x8e, 0x31, 0x8e, 0xa3, 0x10, 0x73, 0x08,
	0x48, 0x1a, 0x1f, 0x9b, 0x37, 0xdb, 0xda, 0xd2, 0x6d, 0x7f, 0x4e, 0x2d, 0xee, 0xa6, 0xf1, 0xb1,
	0xf5, 0x51, 0x47, 0xad, 0x7a, 0x9e, 0xb1, 0x82, 0x9a, 0x79, 0x56, 0x10, 0x44, 0x7f, 0x0b, 0x42,
	0xb3, 0xd7, 0x52, 0x47, 0x54, 0x0d, 0xb6, 0xd7, 0x85, 0x05, 0xdb, 0x98, 0x1d, 0xf9, 0x48, 0xa6,
	0x8b, 0xd8, 0xd8, 0x41, 0x33, 0x43, 0x0a, 0x98, 0xcb, 0x0e, 0x36, 0x7b, 0xcb, 0xb5, 0xd2, 0xaa,
	0x5b, 0x31, 0xa9, 0x6d, 0xf3, 0x86, 0x5f, 0x52, 0x04, 0x4f, 0xd2, 0x4d, 0xfd, 0x7a, 0x3c, 0x49,
	0x31, 0x4c, 0x34, 0x43, 0x21, 0x21, 0x63, 0x30, 0x1b, 0xc2, 0x1a, 0xb1, 0x23, 0x7f, 0x7b, 0x4d,
	0x34, 0x5b, 0x35, 0xd2, 0x3a, 0x45, 0xed, 0x7a, 0xa7, 0x59, 0x46, 0x52, 0x06, 0xc6, 0x4b, 0x74,
	0x8b, 0x02, 0xcb, 0x63, 0xae, 0x6c, 0x7c, 0x3a, 0xdd, 0xc6, 0x1a, 0xa8, 0x5f, 0x70, 0x7c, 0xc5,
	0xb3, 0x9e, 0xa1, 0xf9, 0x4b, 0x33, 0xff, 0xeb, 0xbe, 0xf6, 0xbe, 0x37, 0x90, 0x39, 0x01, 0x78,
	0x21, 0x8f, 0x62, 0x7c, 0xd5, 0xd0, 0xdd, 0x8b, 0x26, 0xc2, 0x58, 0x9b, 0xae, 0xe2, 0x92, 0x49,
	0x6a, 0x5d, 0xc9, 0x20, 0x6b, 0xf5, 0xdd, 0x8f, 0x9f, 0x1f, 0xf4, 0xc7, 0xc6, 0xb2, 0x78, 0x2f,
	0x4e, 0xfe, 0x92, 0xb6, 0xa6, 0x86, 0x87, 0x39, 0x9d, 0xb3, 0x07, 0xa4, 0xb2, 0xc3, 0xe9, 0x9c,
	0x1a, 0xbf, 0x34, 0x64, 0xd6, 0xd9, 0x65, 0xf4, 0xaf, 0xec, 0x8a, 0x1a, 0xea, 0x96, 0x77, 0x1d,
	0x84, 0xbc, 0x2d, 0x96, 0x57, 0x28, 0x5c, 0xb5, 0x9e, 0x08, 0x85, 0x67, 0x92, 0x4e, 0xce, 0xbd,
	0x16, 0x6b, 0x9d, 0xd3, 0x0b, 0x04, 0xba, 0x49, 0x81, 0x76, 0xb5, 0x8e, 0xf7, 0x5e, 0x47, 0x8b,
	0x43, 0x92, 0x4c, 0x3d, 0x8d, 0x37, 0x5f, 0x67, 0xfb, 0x9e, 0x98, 0xde, 0x3d, 0xed, 0xd5, 0x66,
	0x89, 0x18, 0x91, 0x18, 0xa7, 0x23, 0x9b, 0xd0, 0x91, 0x33, 0x82, 0xb4, 0x98, 0x6d, 0xf5, 0x76,
	0x67, 0x11, 0xab, 0xff, 0xc3, 0x58, 0x51, 0xc1, 0x27, 0xbd, 0xb1, 0xd1, 0xef, 0x7f, 0xd6, 0xdb,
	0x1b, 0x12, 0xd8, 0x0f, 0x99, 0x2d, 0x43, 0x11, 0xed, 0x77, 0xed, 0xb2, 0x30, 0xfb, 0xa6, 0x52,
	0x06, 0xfd, 0x90, 0x0d, 0xaa, 0x94, 0xc1, 0x7e, 0x77, 0xa0, 0x52, 0x7e, 0xeb, 0x8b, 0x72, 0xdd,
	0x75, 0xfb, 0x21, 0x73, 0xdd, 0x2a, 0xc9, 0x75, 0xf7, 0xbb, 0xae, 0xab, 0xd2, 0x0e, 0x66, 0x8a,
	0x73, 0x3e, 0xfa, 0x13, 0x00, 0x00, 0xff, 0xff, 0xf1, 0x46, 0x4e, 0x27, 0xd7, 0x06, 0x00, 0x00,
}
