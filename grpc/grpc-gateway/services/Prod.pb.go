// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.3
// source: Prod.proto

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

type ProdAreas int32

const (
	ProdAreas_A ProdAreas = 0
	ProdAreas_B ProdAreas = 1
	ProdAreas_C ProdAreas = 2
)

// Enum value maps for ProdAreas.
var (
	ProdAreas_name = map[int32]string{
		0: "A",
		1: "B",
		2: "C",
	}
	ProdAreas_value = map[string]int32{
		"A": 0,
		"B": 1,
		"C": 2,
	}
)

func (x ProdAreas) Enum() *ProdAreas {
	p := new(ProdAreas)
	*p = x
	return p
}

func (x ProdAreas) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProdAreas) Descriptor() protoreflect.EnumDescriptor {
	return file_Prod_proto_enumTypes[0].Descriptor()
}

func (ProdAreas) Type() protoreflect.EnumType {
	return &file_Prod_proto_enumTypes[0]
}

func (x ProdAreas) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProdAreas.Descriptor instead.
func (ProdAreas) EnumDescriptor() ([]byte, []int) {
	return file_Prod_proto_rawDescGZIP(), []int{0}
}

type ProdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProdId int32 `protobuf:"varint,1,opt,name=prod_id,json=prodId,proto3" json:"prod_id,omitempty"` //传入的商品ID
}

func (x *ProdRequest) Reset() {
	*x = ProdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Prod_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProdRequest) ProtoMessage() {}

func (x *ProdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Prod_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProdRequest.ProtoReflect.Descriptor instead.
func (*ProdRequest) Descriptor() ([]byte, []int) {
	return file_Prod_proto_rawDescGZIP(), []int{0}
}

func (x *ProdRequest) GetProdId() int32 {
	if x != nil {
		return x.ProdId
	}
	return 0
}

type ProdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProdStock int32 `protobuf:"varint,1,opt,name=prod_stock,json=prodStock,proto3" json:"prod_stock,omitempty"` //商品库存
}

func (x *ProdResponse) Reset() {
	*x = ProdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Prod_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProdResponse) ProtoMessage() {}

func (x *ProdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Prod_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProdResponse.ProtoReflect.Descriptor instead.
func (*ProdResponse) Descriptor() ([]byte, []int) {
	return file_Prod_proto_rawDescGZIP(), []int{1}
}

func (x *ProdResponse) GetProdStock() int32 {
	if x != nil {
		return x.ProdStock
	}
	return 0
}

type QuerySize struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Size int32 `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"` //页尺寸
}

func (x *QuerySize) Reset() {
	*x = QuerySize{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Prod_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuerySize) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuerySize) ProtoMessage() {}

func (x *QuerySize) ProtoReflect() protoreflect.Message {
	mi := &file_Prod_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuerySize.ProtoReflect.Descriptor instead.
func (*QuerySize) Descriptor() ([]byte, []int) {
	return file_Prod_proto_rawDescGZIP(), []int{2}
}

func (x *QuerySize) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

//返回 一堆商品库存，使用了repeated修饰符
type ProdResponseList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Prodres []*ProdResponse `protobuf:"bytes,1,rep,name=prodres,proto3" json:"prodres,omitempty"`
}

func (x *ProdResponseList) Reset() {
	*x = ProdResponseList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Prod_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProdResponseList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProdResponseList) ProtoMessage() {}

func (x *ProdResponseList) ProtoReflect() protoreflect.Message {
	mi := &file_Prod_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProdResponseList.ProtoReflect.Descriptor instead.
func (*ProdResponseList) Descriptor() ([]byte, []int) {
	return file_Prod_proto_rawDescGZIP(), []int{3}
}

func (x *ProdResponseList) GetProdres() []*ProdResponse {
	if x != nil {
		return x.Prodres
	}
	return nil
}

var File_Prod_proto protoreflect.FileDescriptor

var file_Prod_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x26, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x70, 0x72, 0x6f, 0x64, 0x49, 0x64, 0x22, 0x2d, 0x0a, 0x0c, 0x50, 0x72,
	0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72,
	0x6f, 0x64, 0x5f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x70, 0x72, 0x6f, 0x64, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x22, 0x1f, 0x0a, 0x09, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x44, 0x0a, 0x10, 0x50, 0x72,
	0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x30,
	0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x72, 0x65, 0x73,
	0x2a, 0x20, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x41, 0x72, 0x65, 0x61, 0x73, 0x12, 0x05, 0x0a,
	0x01, 0x41, 0x10, 0x00, 0x12, 0x05, 0x0a, 0x01, 0x42, 0x10, 0x01, 0x12, 0x05, 0x0a, 0x01, 0x43,
	0x10, 0x02, 0x32, 0xe9, 0x01, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x59, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x53, 0x74, 0x6f,
	0x63, 0x6b, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72,
	0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x70,
	0x72, 0x6f, 0x64, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x64, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x42, 0x0a,
	0x0d, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x13,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53,
	0x69, 0x7a, 0x65, 0x1a, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50,
	0x72, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x22,
	0x00, 0x12, 0x3b, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x22, 0x00, 0x42, 0x0d,
	0x5a, 0x0b, 0x2e, 0x2e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Prod_proto_rawDescOnce sync.Once
	file_Prod_proto_rawDescData = file_Prod_proto_rawDesc
)

func file_Prod_proto_rawDescGZIP() []byte {
	file_Prod_proto_rawDescOnce.Do(func() {
		file_Prod_proto_rawDescData = protoimpl.X.CompressGZIP(file_Prod_proto_rawDescData)
	})
	return file_Prod_proto_rawDescData
}

var file_Prod_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_Prod_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_Prod_proto_goTypes = []interface{}{
	(ProdAreas)(0),           // 0: services.ProdAreas
	(*ProdRequest)(nil),      // 1: services.ProdRequest
	(*ProdResponse)(nil),     // 2: services.ProdResponse
	(*QuerySize)(nil),        // 3: services.QuerySize
	(*ProdResponseList)(nil), // 4: services.ProdResponseList
	(*ProdModel)(nil),        // 5: services.ProdModel
}
var file_Prod_proto_depIdxs = []int32{
	2, // 0: services.ProdResponseList.prodres:type_name -> services.ProdResponse
	1, // 1: services.ProdService.GetProdStock:input_type -> services.ProdRequest
	3, // 2: services.ProdService.GetProdStocks:input_type -> services.QuerySize
	1, // 3: services.ProdService.GetProdInfo:input_type -> services.ProdRequest
	2, // 4: services.ProdService.GetProdStock:output_type -> services.ProdResponse
	4, // 5: services.ProdService.GetProdStocks:output_type -> services.ProdResponseList
	5, // 6: services.ProdService.GetProdInfo:output_type -> services.ProdModel
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_Prod_proto_init() }
func file_Prod_proto_init() {
	if File_Prod_proto != nil {
		return
	}
	file_Models_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_Prod_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProdRequest); i {
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
		file_Prod_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProdResponse); i {
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
		file_Prod_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuerySize); i {
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
		file_Prod_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProdResponseList); i {
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
			RawDescriptor: file_Prod_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Prod_proto_goTypes,
		DependencyIndexes: file_Prod_proto_depIdxs,
		EnumInfos:         file_Prod_proto_enumTypes,
		MessageInfos:      file_Prod_proto_msgTypes,
	}.Build()
	File_Prod_proto = out.File
	file_Prod_proto_rawDesc = nil
	file_Prod_proto_goTypes = nil
	file_Prod_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ProdServiceClient is the client API for ProdService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProdServiceClient interface {
	GetProdStock(ctx context.Context, in *ProdRequest, opts ...grpc.CallOption) (*ProdResponse, error)
	GetProdStocks(ctx context.Context, in *QuerySize, opts ...grpc.CallOption) (*ProdResponseList, error)
	GetProdInfo(ctx context.Context, in *ProdRequest, opts ...grpc.CallOption) (*ProdModel, error)
}

type prodServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProdServiceClient(cc grpc.ClientConnInterface) ProdServiceClient {
	return &prodServiceClient{cc}
}

func (c *prodServiceClient) GetProdStock(ctx context.Context, in *ProdRequest, opts ...grpc.CallOption) (*ProdResponse, error) {
	out := new(ProdResponse)
	err := c.cc.Invoke(ctx, "/services.ProdService/GetProdStock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *prodServiceClient) GetProdStocks(ctx context.Context, in *QuerySize, opts ...grpc.CallOption) (*ProdResponseList, error) {
	out := new(ProdResponseList)
	err := c.cc.Invoke(ctx, "/services.ProdService/GetProdStocks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *prodServiceClient) GetProdInfo(ctx context.Context, in *ProdRequest, opts ...grpc.CallOption) (*ProdModel, error) {
	out := new(ProdModel)
	err := c.cc.Invoke(ctx, "/services.ProdService/GetProdInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProdServiceServer is the server API for ProdService service.
type ProdServiceServer interface {
	GetProdStock(context.Context, *ProdRequest) (*ProdResponse, error)
	GetProdStocks(context.Context, *QuerySize) (*ProdResponseList, error)
	GetProdInfo(context.Context, *ProdRequest) (*ProdModel, error)
}

// UnimplementedProdServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProdServiceServer struct {
}

func (*UnimplementedProdServiceServer) GetProdStock(context.Context, *ProdRequest) (*ProdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProdStock not implemented")
}
func (*UnimplementedProdServiceServer) GetProdStocks(context.Context, *QuerySize) (*ProdResponseList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProdStocks not implemented")
}
func (*UnimplementedProdServiceServer) GetProdInfo(context.Context, *ProdRequest) (*ProdModel, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProdInfo not implemented")
}

func RegisterProdServiceServer(s *grpc.Server, srv ProdServiceServer) {
	s.RegisterService(&_ProdService_serviceDesc, srv)
}

func _ProdService_GetProdStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProdServiceServer).GetProdStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.ProdService/GetProdStock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProdServiceServer).GetProdStock(ctx, req.(*ProdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProdService_GetProdStocks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySize)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProdServiceServer).GetProdStocks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.ProdService/GetProdStocks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProdServiceServer).GetProdStocks(ctx, req.(*QuerySize))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProdService_GetProdInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProdServiceServer).GetProdInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.ProdService/GetProdInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProdServiceServer).GetProdInfo(ctx, req.(*ProdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProdService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "services.ProdService",
	HandlerType: (*ProdServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProdStock",
			Handler:    _ProdService_GetProdStock_Handler,
		},
		{
			MethodName: "GetProdStocks",
			Handler:    _ProdService_GetProdStocks_Handler,
		},
		{
			MethodName: "GetProdInfo",
			Handler:    _ProdService_GetProdInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Prod.proto",
}
