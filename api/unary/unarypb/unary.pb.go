// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.1
// source: unary/unarypb/unary.proto

package unarypb

import (
	context "context"
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

type Attack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Damage uint32 `protobuf:"varint,1,opt,name=damage,proto3" json:"damage,omitempty"`
}

func (x *Attack) Reset() {
	*x = Attack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_unary_unarypb_unary_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Attack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Attack) ProtoMessage() {}

func (x *Attack) ProtoReflect() protoreflect.Message {
	mi := &file_unary_unarypb_unary_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Attack.ProtoReflect.Descriptor instead.
func (*Attack) Descriptor() ([]byte, []int) {
	return file_unary_unarypb_unary_proto_rawDescGZIP(), []int{0}
}

func (x *Attack) GetDamage() uint32 {
	if x != nil {
		return x.Damage
	}
	return 0
}

type Boss struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Health int32  `protobuf:"varint,2,opt,name=health,proto3" json:"health,omitempty"`
}

func (x *Boss) Reset() {
	*x = Boss{}
	if protoimpl.UnsafeEnabled {
		mi := &file_unary_unarypb_unary_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Boss) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Boss) ProtoMessage() {}

func (x *Boss) ProtoReflect() protoreflect.Message {
	mi := &file_unary_unarypb_unary_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Boss.ProtoReflect.Descriptor instead.
func (*Boss) Descriptor() ([]byte, []int) {
	return file_unary_unarypb_unary_proto_rawDescGZIP(), []int{1}
}

func (x *Boss) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Boss) GetHealth() int32 {
	if x != nil {
		return x.Health
	}
	return 0
}

type AttackBossRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Boss   *Boss   `protobuf:"bytes,1,opt,name=boss,proto3" json:"boss,omitempty"`
	Attack *Attack `protobuf:"bytes,2,opt,name=attack,proto3" json:"attack,omitempty"`
}

func (x *AttackBossRequest) Reset() {
	*x = AttackBossRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_unary_unarypb_unary_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttackBossRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttackBossRequest) ProtoMessage() {}

func (x *AttackBossRequest) ProtoReflect() protoreflect.Message {
	mi := &file_unary_unarypb_unary_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttackBossRequest.ProtoReflect.Descriptor instead.
func (*AttackBossRequest) Descriptor() ([]byte, []int) {
	return file_unary_unarypb_unary_proto_rawDescGZIP(), []int{2}
}

func (x *AttackBossRequest) GetBoss() *Boss {
	if x != nil {
		return x.Boss
	}
	return nil
}

func (x *AttackBossRequest) GetAttack() *Attack {
	if x != nil {
		return x.Attack
	}
	return nil
}

type AttackBossResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Boss *Boss `protobuf:"bytes,1,opt,name=boss,proto3" json:"boss,omitempty"`
}

func (x *AttackBossResponse) Reset() {
	*x = AttackBossResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_unary_unarypb_unary_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttackBossResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttackBossResponse) ProtoMessage() {}

func (x *AttackBossResponse) ProtoReflect() protoreflect.Message {
	mi := &file_unary_unarypb_unary_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttackBossResponse.ProtoReflect.Descriptor instead.
func (*AttackBossResponse) Descriptor() ([]byte, []int) {
	return file_unary_unarypb_unary_proto_rawDescGZIP(), []int{3}
}

func (x *AttackBossResponse) GetBoss() *Boss {
	if x != nil {
		return x.Boss
	}
	return nil
}

var File_unary_unarypb_unary_proto protoreflect.FileDescriptor

var file_unary_unarypb_unary_proto_rawDesc = []byte{
	0x0a, 0x19, 0x75, 0x6e, 0x61, 0x72, 0x79, 0x2f, 0x75, 0x6e, 0x61, 0x72, 0x79, 0x70, 0x62, 0x2f,
	0x75, 0x6e, 0x61, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x75, 0x6e, 0x61,
	0x72, 0x79, 0x22, 0x20, 0x0a, 0x06, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x12, 0x16, 0x0a, 0x06,
	0x64, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x64, 0x61,
	0x6d, 0x61, 0x67, 0x65, 0x22, 0x2e, 0x0a, 0x04, 0x42, 0x6f, 0x73, 0x73, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x68, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x22, 0x5b, 0x0a, 0x11, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x42, 0x6f,
	0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x04, 0x62, 0x6f, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x75, 0x6e, 0x61, 0x72, 0x79, 0x2e,
	0x42, 0x6f, 0x73, 0x73, 0x52, 0x04, 0x62, 0x6f, 0x73, 0x73, 0x12, 0x25, 0x0a, 0x06, 0x61, 0x74,
	0x74, 0x61, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x75, 0x6e, 0x61,
	0x72, 0x79, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x52, 0x06, 0x61, 0x74, 0x74, 0x61, 0x63,
	0x6b, 0x22, 0x35, 0x0a, 0x12, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x42, 0x6f, 0x73, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x62, 0x6f, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x75, 0x6e, 0x61, 0x72, 0x79, 0x2e, 0x42, 0x6f,
	0x73, 0x73, 0x52, 0x04, 0x62, 0x6f, 0x73, 0x73, 0x32, 0x53, 0x0a, 0x0c, 0x55, 0x6e, 0x61, 0x72,
	0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x0a, 0x41, 0x74, 0x74, 0x61,
	0x63, 0x6b, 0x42, 0x6f, 0x73, 0x73, 0x12, 0x18, 0x2e, 0x75, 0x6e, 0x61, 0x72, 0x79, 0x2e, 0x41,
	0x74, 0x74, 0x61, 0x63, 0x6b, 0x42, 0x6f, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x75, 0x6e, 0x61, 0x72, 0x79, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x42,
	0x6f, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x11, 0x5a,
	0x0f, 0x2e, 0x2f, 0x75, 0x6e, 0x61, 0x72, 0x79, 0x2f, 0x75, 0x6e, 0x61, 0x72, 0x79, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_unary_unarypb_unary_proto_rawDescOnce sync.Once
	file_unary_unarypb_unary_proto_rawDescData = file_unary_unarypb_unary_proto_rawDesc
)

func file_unary_unarypb_unary_proto_rawDescGZIP() []byte {
	file_unary_unarypb_unary_proto_rawDescOnce.Do(func() {
		file_unary_unarypb_unary_proto_rawDescData = protoimpl.X.CompressGZIP(file_unary_unarypb_unary_proto_rawDescData)
	})
	return file_unary_unarypb_unary_proto_rawDescData
}

var file_unary_unarypb_unary_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_unary_unarypb_unary_proto_goTypes = []interface{}{
	(*Attack)(nil),             // 0: unary.Attack
	(*Boss)(nil),               // 1: unary.Boss
	(*AttackBossRequest)(nil),  // 2: unary.AttackBossRequest
	(*AttackBossResponse)(nil), // 3: unary.AttackBossResponse
}
var file_unary_unarypb_unary_proto_depIdxs = []int32{
	1, // 0: unary.AttackBossRequest.boss:type_name -> unary.Boss
	0, // 1: unary.AttackBossRequest.attack:type_name -> unary.Attack
	1, // 2: unary.AttackBossResponse.boss:type_name -> unary.Boss
	2, // 3: unary.UnaryService.AttackBoss:input_type -> unary.AttackBossRequest
	3, // 4: unary.UnaryService.AttackBoss:output_type -> unary.AttackBossResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_unary_unarypb_unary_proto_init() }
func file_unary_unarypb_unary_proto_init() {
	if File_unary_unarypb_unary_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_unary_unarypb_unary_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Attack); i {
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
		file_unary_unarypb_unary_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Boss); i {
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
		file_unary_unarypb_unary_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttackBossRequest); i {
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
		file_unary_unarypb_unary_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttackBossResponse); i {
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
			RawDescriptor: file_unary_unarypb_unary_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_unary_unarypb_unary_proto_goTypes,
		DependencyIndexes: file_unary_unarypb_unary_proto_depIdxs,
		MessageInfos:      file_unary_unarypb_unary_proto_msgTypes,
	}.Build()
	File_unary_unarypb_unary_proto = out.File
	file_unary_unarypb_unary_proto_rawDesc = nil
	file_unary_unarypb_unary_proto_goTypes = nil
	file_unary_unarypb_unary_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UnaryServiceClient is the client API for UnaryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UnaryServiceClient interface {
	AttackBoss(ctx context.Context, in *AttackBossRequest, opts ...grpc.CallOption) (*AttackBossResponse, error)
}

type unaryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUnaryServiceClient(cc grpc.ClientConnInterface) UnaryServiceClient {
	return &unaryServiceClient{cc}
}

func (c *unaryServiceClient) AttackBoss(ctx context.Context, in *AttackBossRequest, opts ...grpc.CallOption) (*AttackBossResponse, error) {
	out := new(AttackBossResponse)
	err := c.cc.Invoke(ctx, "/unary.UnaryService/AttackBoss", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UnaryServiceServer is the server API for UnaryService service.
type UnaryServiceServer interface {
	AttackBoss(context.Context, *AttackBossRequest) (*AttackBossResponse, error)
}

// UnimplementedUnaryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUnaryServiceServer struct {
}

func (*UnimplementedUnaryServiceServer) AttackBoss(context.Context, *AttackBossRequest) (*AttackBossResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AttackBoss not implemented")
}

func RegisterUnaryServiceServer(s *grpc.Server, srv UnaryServiceServer) {
	s.RegisterService(&_UnaryService_serviceDesc, srv)
}

func _UnaryService_AttackBoss_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AttackBossRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UnaryServiceServer).AttackBoss(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/unary.UnaryService/AttackBoss",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UnaryServiceServer).AttackBoss(ctx, req.(*AttackBossRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UnaryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "unary.UnaryService",
	HandlerType: (*UnaryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AttackBoss",
			Handler:    _UnaryService_AttackBoss_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "unary/unarypb/unary.proto",
}
