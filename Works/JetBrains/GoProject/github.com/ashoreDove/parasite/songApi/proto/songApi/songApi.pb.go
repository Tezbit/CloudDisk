// Code generated by protoc-gen-go. DO NOT EDIT.
// source: songApi.proto

/*
Package go_micro_api_songApi is a generated protocol buffer package.

It is generated from these files:
	songApi.proto

It has these top-level messages:
	Pair
	Request
	Response
*/
package go_micro_api_songApi

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Pair struct {
	Key    string   `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Values []string `protobuf:"bytes,2,rep,name=values" json:"values,omitempty"`
}

func (m *Pair) Reset()                    { *m = Pair{} }
func (m *Pair) String() string            { return proto.CompactTextString(m) }
func (*Pair) ProtoMessage()               {}
func (*Pair) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Pair) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Pair) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

type Request struct {
	Method string           `protobuf:"bytes,1,opt,name=method" json:"method,omitempty"`
	Path   string           `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
	Header map[string]*Pair `protobuf:"bytes,3,rep,name=header" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Get    map[string]*Pair `protobuf:"bytes,4,rep,name=get" json:"get,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Post   map[string]*Pair `protobuf:"bytes,5,rep,name=post" json:"post,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Body   string           `protobuf:"bytes,6,opt,name=body" json:"body,omitempty"`
	Url    string           `protobuf:"bytes,7,opt,name=url" json:"url,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Request) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *Request) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Request) GetHeader() map[string]*Pair {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Request) GetGet() map[string]*Pair {
	if m != nil {
		return m.Get
	}
	return nil
}

func (m *Request) GetPost() map[string]*Pair {
	if m != nil {
		return m.Post
	}
	return nil
}

func (m *Request) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *Request) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type Response struct {
	StatusCode int32            `protobuf:"varint,1,opt,name=statusCode" json:"statusCode,omitempty"`
	Header     map[string]*Pair `protobuf:"bytes,2,rep,name=header" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Body       string           `protobuf:"bytes,3,opt,name=body" json:"body,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Response) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *Response) GetHeader() map[string]*Pair {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Response) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func init() {
	proto.RegisterType((*Pair)(nil), "go.micro.api.songApi.Pair")
	proto.RegisterType((*Request)(nil), "go.micro.api.songApi.Request")
	proto.RegisterType((*Response)(nil), "go.micro.api.songApi.Response")
}

func init() { proto.RegisterFile("songApi.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 382 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x93, 0x4f, 0x6a, 0xeb, 0x30,
	0x10, 0xc6, 0x9f, 0xff, 0xc4, 0x49, 0x26, 0x3c, 0x78, 0x88, 0x47, 0x11, 0x86, 0x86, 0xe0, 0x45,
	0x9b, 0x76, 0x61, 0x42, 0xba, 0x09, 0xed, 0x2a, 0x2d, 0x25, 0xed, 0xa2, 0x10, 0x14, 0x72, 0x00,
	0x27, 0x16, 0xb1, 0xa9, 0x13, 0xb9, 0x92, 0x5c, 0xc8, 0x41, 0x7a, 0x82, 0x9e, 0xaa, 0xb7, 0x29,
	0x92, 0x95, 0x34, 0x0b, 0xd7, 0x9b, 0xa6, 0xbb, 0x91, 0x34, 0xbf, 0x8f, 0x99, 0xf9, 0x46, 0xf0,
	0x57, 0xb0, 0xcd, 0x6a, 0x9c, 0xa7, 0x61, 0xce, 0x99, 0x64, 0xe8, 0xff, 0x8a, 0x85, 0xeb, 0x74,
	0xc9, 0x59, 0x18, 0xe5, 0x69, 0x68, 0xde, 0x82, 0x01, 0xb8, 0xd3, 0x28, 0xe5, 0xe8, 0x1f, 0x38,
	0xcf, 0x74, 0x8b, 0xad, 0x9e, 0xd5, 0x6f, 0x13, 0x15, 0xa2, 0x13, 0xf0, 0x5e, 0xa3, 0xac, 0xa0,
	0x02, 0xdb, 0x3d, 0xa7, 0xdf, 0x26, 0xe6, 0x14, 0xbc, 0xb9, 0xd0, 0x24, 0xf4, 0xa5, 0xa0, 0x42,
	0xaa, 0x9c, 0x35, 0x95, 0x09, 0x8b, 0x0d, 0x68, 0x4e, 0x08, 0x81, 0x9b, 0x47, 0x32, 0xc1, 0xb6,
	0xbe, 0xd5, 0x31, 0x1a, 0x83, 0x97, 0xd0, 0x28, 0xa6, 0x1c, 0x3b, 0x3d, 0xa7, 0xdf, 0x19, 0x5e,
	0x84, 0x55, 0x05, 0x85, 0x46, 0x3a, 0x7c, 0xd0, 0xb9, 0xf7, 0x1b, 0xc9, 0xb7, 0xc4, 0x80, 0x68,
	0x04, 0xce, 0x8a, 0x4a, 0xec, 0x6a, 0xfe, 0xac, 0x9e, 0x9f, 0x50, 0x59, 0xc2, 0x0a, 0x41, 0x37,
	0xe0, 0xe6, 0x4c, 0x48, 0xdc, 0xd0, 0xe8, 0x79, 0x3d, 0x3a, 0x65, 0xc2, 0xb0, 0x1a, 0x52, 0xdd,
	0x2c, 0x58, 0xbc, 0xc5, 0x5e, 0xd9, 0x8d, 0x8a, 0xd5, 0xbc, 0x0a, 0x9e, 0xe1, 0x66, 0x39, 0xaf,
	0x82, 0x67, 0xfe, 0x1c, 0x3a, 0x07, 0x35, 0x57, 0x0c, 0x74, 0x00, 0x0d, 0x3d, 0x42, 0x3d, 0x95,
	0xce, 0xd0, 0xaf, 0x2e, 0x42, 0xb9, 0x41, 0xca, 0xc4, 0x6b, 0x7b, 0x64, 0xf9, 0x04, 0x5a, 0xbb,
	0x56, 0x8e, 0xa6, 0x39, 0x83, 0xf6, 0xbe, 0xc7, 0x63, 0x89, 0x06, 0x1f, 0x16, 0xb4, 0x08, 0x15,
	0x39, 0xdb, 0x08, 0x8a, 0xba, 0x00, 0x42, 0x46, 0xb2, 0x10, 0x77, 0x2c, 0xa6, 0x5a, 0xbb, 0x41,
	0x0e, 0x6e, 0xd0, 0xed, 0x7e, 0x19, 0x6c, 0xed, 0xc8, 0xe5, 0x77, 0x8e, 0x94, 0x7a, 0x95, 0xdb,
	0xb0, 0xb3, 0xc5, 0xf9, 0xb2, 0xe5, 0x97, 0x4c, 0x18, 0xbe, 0x5b, 0xd0, 0x9c, 0x95, 0x6f, 0xe8,
	0x11, 0xbc, 0x19, 0x8d, 0xf8, 0x32, 0x41, 0xa7, 0xb5, 0x6b, 0xe4, 0x77, 0xeb, 0x7b, 0x0a, 0xfe,
	0xa0, 0x27, 0x80, 0x09, 0x95, 0x4a, 0x78, 0xce, 0xb3, 0x1f, 0xcb, 0x2d, 0x3c, 0xfd, 0xd1, 0xaf,
	0x3e, 0x03, 0x00, 0x00, 0xff, 0xff, 0x62, 0x49, 0x23, 0xa5, 0xf9, 0x03, 0x00, 0x00,
}
