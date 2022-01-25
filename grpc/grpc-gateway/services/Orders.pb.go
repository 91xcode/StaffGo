// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.3
// source: Orders.proto

package services

import (
	context "context"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderMain *OrderMain `protobuf:"bytes,1,opt,name=order_main,json=orderMain,proto3" json:"order_main,omitempty"`
}

func (x *OrderRequest) Reset() {
	*x = OrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Orders_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderRequest) ProtoMessage() {}

func (x *OrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Orders_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderRequest.ProtoReflect.Descriptor instead.
func (*OrderRequest) Descriptor() ([]byte, []int) {
	return file_Orders_proto_rawDescGZIP(), []int{0}
}

func (x *OrderRequest) GetOrderMain() *OrderMain {
	if x != nil {
		return x.OrderMain
	}
	return nil
}

type OrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *OrderResponse) Reset() {
	*x = OrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Orders_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderResponse) ProtoMessage() {}

func (x *OrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Orders_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderResponse.ProtoReflect.Descriptor instead.
func (*OrderResponse) Descriptor() ([]byte, []int) {
	return file_Orders_proto_rawDescGZIP(), []int{1}
}

func (x *OrderResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *OrderResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_Orders_proto protoreflect.FileDescriptor

var file_Orders_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x42, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x6d, 0x61,
	0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4d, 0x61, 0x69, 0x6e, 0x52, 0x09, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x4d, 0x61, 0x69, 0x6e, 0x22, 0x41, 0x0a, 0x0d, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x6b, 0x0a, 0x0c, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x69, 0x76, 0x63, 0x65, 0x12, 0x5b, 0x0a, 0x08, 0x4e,
	0x65, 0x77, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18,
	0x22, 0x0a, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x3a, 0x0a, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x69, 0x6e, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2e, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Orders_proto_rawDescOnce sync.Once
	file_Orders_proto_rawDescData = file_Orders_proto_rawDesc
)

func file_Orders_proto_rawDescGZIP() []byte {
	file_Orders_proto_rawDescOnce.Do(func() {
		file_Orders_proto_rawDescData = protoimpl.X.CompressGZIP(file_Orders_proto_rawDescData)
	})
	return file_Orders_proto_rawDescData
}

var file_Orders_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_Orders_proto_goTypes = []interface{}{
	(*OrderRequest)(nil),  // 0: services.OrderRequest
	(*OrderResponse)(nil), // 1: services.OrderResponse
	(*OrderMain)(nil),     // 2: services.OrderMain
}
var file_Orders_proto_depIdxs = []int32{
	2, // 0: services.OrderRequest.order_main:type_name -> services.OrderMain
	0, // 1: services.OrderSerivce.NewOrder:input_type -> services.OrderRequest
	1, // 2: services.OrderSerivce.NewOrder:output_type -> services.OrderResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_Orders_proto_init() }
func file_Orders_proto_init() {
	if File_Orders_proto != nil {
		return
	}
	file_Models_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_Orders_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_Orders_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_Orders_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Orders_proto_goTypes,
		DependencyIndexes: file_Orders_proto_depIdxs,
		MessageInfos:      file_Orders_proto_msgTypes,
	}.Build()
	File_Orders_proto = out.File
	file_Orders_proto_rawDesc = nil
	file_Orders_proto_goTypes = nil
	file_Orders_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// OrderSerivceClient is the client API for OrderSerivce service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OrderSerivceClient interface {
	NewOrder(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*OrderResponse, error)
}

type orderSerivceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderSerivceClient(cc grpc.ClientConnInterface) OrderSerivceClient {
	return &orderSerivceClient{cc}
}

func (c *orderSerivceClient) NewOrder(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*OrderResponse, error) {
	out := new(OrderResponse)
	err := c.cc.Invoke(ctx, "/services.OrderSerivce/NewOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderSerivceServer is the server API for OrderSerivce service.
type OrderSerivceServer interface {
	NewOrder(context.Context, *OrderRequest) (*OrderResponse, error)
}

// UnimplementedOrderSerivceServer can be embedded to have forward compatible implementations.
type UnimplementedOrderSerivceServer struct {
}

func (*UnimplementedOrderSerivceServer) NewOrder(context.Context, *OrderRequest) (*OrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewOrder not implemented")
}

func RegisterOrderSerivceServer(s *grpc.Server, srv OrderSerivceServer) {
	s.RegisterService(&_OrderSerivce_serviceDesc, srv)
}

func _OrderSerivce_NewOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderSerivceServer).NewOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.OrderSerivce/NewOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderSerivceServer).NewOrder(ctx, req.(*OrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _OrderSerivce_serviceDesc = grpc.ServiceDesc{
	ServiceName: "services.OrderSerivce",
	HandlerType: (*OrderSerivceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewOrder",
			Handler:    _OrderSerivce_NewOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Orders.proto",
}
