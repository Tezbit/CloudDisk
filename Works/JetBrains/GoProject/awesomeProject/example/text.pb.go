// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.10
// source: text.proto

package example

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

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account  string `protobuf:"bytes,1,opt,name=Account,proto3" json:"Account,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_text_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_text_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_text_proto_rawDescGZIP(), []int{0}
}

func (x *LoginRequest) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code  int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Token string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_text_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_text_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_text_proto_rawDescGZIP(), []int{1}
}

func (x *LoginResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *LoginResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type Music struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Singer string `protobuf:"bytes,3,opt,name=singer,proto3" json:"singer,omitempty"`
	Total  string `protobuf:"bytes,4,opt,name=total,proto3" json:"total,omitempty"`
	Url    string `protobuf:"bytes,5,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *Music) Reset() {
	*x = Music{}
	if protoimpl.UnsafeEnabled {
		mi := &file_text_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Music) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Music) ProtoMessage() {}

func (x *Music) ProtoReflect() protoreflect.Message {
	mi := &file_text_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Music.ProtoReflect.Descriptor instead.
func (*Music) Descriptor() ([]byte, []int) {
	return file_text_proto_rawDescGZIP(), []int{2}
}

func (x *Music) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Music) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Music) GetSinger() string {
	if x != nil {
		return x.Singer
	}
	return ""
}

func (x *Music) GetTotal() string {
	if x != nil {
		return x.Total
	}
	return ""
}

func (x *Music) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type PlaylistPub struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SongList []*Music `protobuf:"bytes,1,rep,name=song_list,json=songList,proto3" json:"song_list,omitempty"`
}

func (x *PlaylistPub) Reset() {
	*x = PlaylistPub{}
	if protoimpl.UnsafeEnabled {
		mi := &file_text_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlaylistPub) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlaylistPub) ProtoMessage() {}

func (x *PlaylistPub) ProtoReflect() protoreflect.Message {
	mi := &file_text_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlaylistPub.ProtoReflect.Descriptor instead.
func (*PlaylistPub) Descriptor() ([]byte, []int) {
	return file_text_proto_rawDescGZIP(), []int{3}
}

func (x *PlaylistPub) GetSongList() []*Music {
	if x != nil {
		return x.SongList
	}
	return nil
}

type Sub struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Sub) Reset() {
	*x = Sub{}
	if protoimpl.UnsafeEnabled {
		mi := &file_text_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sub) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sub) ProtoMessage() {}

func (x *Sub) ProtoReflect() protoreflect.Message {
	mi := &file_text_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sub.ProtoReflect.Descriptor instead.
func (*Sub) Descriptor() ([]byte, []int) {
	return file_text_proto_rawDescGZIP(), []int{4}
}

var File_text_proto protoreflect.FileDescriptor

var file_text_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x44, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x39, 0x0a, 0x0d, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x6b, 0x0a, 0x05, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72,
	0x6c, 0x22, 0x38, 0x0a, 0x0b, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x50, 0x75, 0x62,
	0x12, 0x29, 0x0a, 0x09, 0x73, 0x6f, 0x6e, 0x67, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x75, 0x73, 0x69,
	0x63, 0x52, 0x08, 0x73, 0x6f, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x05, 0x0a, 0x03, 0x53,
	0x75, 0x62, 0x32, 0x6d, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x13, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x07, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73,
	0x68, 0x12, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x75, 0x62, 0x1a, 0x12, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x50, 0x75,
	0x62, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2e, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_text_proto_rawDescOnce sync.Once
	file_text_proto_rawDescData = file_text_proto_rawDesc
)

func file_text_proto_rawDescGZIP() []byte {
	file_text_proto_rawDescOnce.Do(func() {
		file_text_proto_rawDescData = protoimpl.X.CompressGZIP(file_text_proto_rawDescData)
	})
	return file_text_proto_rawDescData
}

var file_text_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_text_proto_goTypes = []interface{}{
	(*LoginRequest)(nil),  // 0: proto.LoginRequest
	(*LoginResponse)(nil), // 1: proto.LoginResponse
	(*Music)(nil),         // 2: proto.Music
	(*PlaylistPub)(nil),   // 3: proto.PlaylistPub
	(*Sub)(nil),           // 4: proto.Sub
}
var file_text_proto_depIdxs = []int32{
	2, // 0: proto.PlaylistPub.song_list:type_name -> proto.Music
	0, // 1: proto.LoginService.Login:input_type -> proto.LoginRequest
	4, // 2: proto.LoginService.Publish:input_type -> proto.Sub
	1, // 3: proto.LoginService.Login:output_type -> proto.LoginResponse
	3, // 4: proto.LoginService.Publish:output_type -> proto.PlaylistPub
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_text_proto_init() }
func file_text_proto_init() {
	if File_text_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_text_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_text_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResponse); i {
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
		file_text_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Music); i {
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
		file_text_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlaylistPub); i {
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
		file_text_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sub); i {
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
			RawDescriptor: file_text_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_text_proto_goTypes,
		DependencyIndexes: file_text_proto_depIdxs,
		MessageInfos:      file_text_proto_msgTypes,
	}.Build()
	File_text_proto = out.File
	file_text_proto_rawDesc = nil
	file_text_proto_goTypes = nil
	file_text_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LoginServiceClient is the client API for LoginService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LoginServiceClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Publish(ctx context.Context, in *Sub, opts ...grpc.CallOption) (*PlaylistPub, error)
}

type loginServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLoginServiceClient(cc grpc.ClientConnInterface) LoginServiceClient {
	return &loginServiceClient{cc}
}

func (c *loginServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/proto.LoginService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginServiceClient) Publish(ctx context.Context, in *Sub, opts ...grpc.CallOption) (*PlaylistPub, error) {
	out := new(PlaylistPub)
	err := c.cc.Invoke(ctx, "/proto.LoginService/Publish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoginServiceServer is the server API for LoginService service.
type LoginServiceServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Publish(context.Context, *Sub) (*PlaylistPub, error)
}

// UnimplementedLoginServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLoginServiceServer struct {
}

func (*UnimplementedLoginServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (*UnimplementedLoginServiceServer) Publish(context.Context, *Sub) (*PlaylistPub, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Publish not implemented")
}

func RegisterLoginServiceServer(s *grpc.Server, srv LoginServiceServer) {
	s.RegisterService(&_LoginService_serviceDesc, srv)
}

func _LoginService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.LoginService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoginService_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Sub)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServiceServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.LoginService/Publish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServiceServer).Publish(ctx, req.(*Sub))
	}
	return interceptor(ctx, in, info, handler)
}

var _LoginService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.LoginService",
	HandlerType: (*LoginServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _LoginService_Login_Handler,
		},
		{
			MethodName: "Publish",
			Handler:    _LoginService_Publish_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "text.proto",
}
