// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calculate.proto

package calculate

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
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

type Number_ struct {
	Num                  int64    `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Number_) Reset()         { *m = Number_{} }
func (m *Number_) String() string { return proto.CompactTextString(m) }
func (*Number_) ProtoMessage()    {}
func (*Number_) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d8b31321fcad8b1, []int{0}
}

func (m *Number_) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Number_.Unmarshal(m, b)
}
func (m *Number_) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Number_.Marshal(b, m, deterministic)
}
func (m *Number_) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Number_.Merge(m, src)
}
func (m *Number_) XXX_Size() int {
	return xxx_messageInfo_Number_.Size(m)
}
func (m *Number_) XXX_DiscardUnknown() {
	xxx_messageInfo_Number_.DiscardUnknown(m)
}

var xxx_messageInfo_Number_ proto.InternalMessageInfo

func (m *Number_) GetNum() int64 {
	if m != nil {
		return m.Num
	}
	return 0
}

type String_ struct {
	Str                  string   `protobuf:"bytes,1,opt,name=str,proto3" json:"str,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *String_) Reset()         { *m = String_{} }
func (m *String_) String() string { return proto.CompactTextString(m) }
func (*String_) ProtoMessage()    {}
func (*String_) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d8b31321fcad8b1, []int{1}
}

func (m *String_) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_String_.Unmarshal(m, b)
}
func (m *String_) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_String_.Marshal(b, m, deterministic)
}
func (m *String_) XXX_Merge(src proto.Message) {
	xxx_messageInfo_String_.Merge(m, src)
}
func (m *String_) XXX_Size() int {
	return xxx_messageInfo_String_.Size(m)
}
func (m *String_) XXX_DiscardUnknown() {
	xxx_messageInfo_String_.DiscardUnknown(m)
}

var xxx_messageInfo_String_ proto.InternalMessageInfo

func (m *String_) GetStr() string {
	if m != nil {
		return m.Str
	}
	return ""
}

type CalculateInt struct {
	A                    *Number_ `protobuf:"bytes,1,opt,name=a,proto3" json:"a,omitempty"`
	B                    *Number_ `protobuf:"bytes,2,opt,name=b,proto3" json:"b,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CalculateInt) Reset()         { *m = CalculateInt{} }
func (m *CalculateInt) String() string { return proto.CompactTextString(m) }
func (*CalculateInt) ProtoMessage()    {}
func (*CalculateInt) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d8b31321fcad8b1, []int{2}
}

func (m *CalculateInt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalculateInt.Unmarshal(m, b)
}
func (m *CalculateInt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalculateInt.Marshal(b, m, deterministic)
}
func (m *CalculateInt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalculateInt.Merge(m, src)
}
func (m *CalculateInt) XXX_Size() int {
	return xxx_messageInfo_CalculateInt.Size(m)
}
func (m *CalculateInt) XXX_DiscardUnknown() {
	xxx_messageInfo_CalculateInt.DiscardUnknown(m)
}

var xxx_messageInfo_CalculateInt proto.InternalMessageInfo

func (m *CalculateInt) GetA() *Number_ {
	if m != nil {
		return m.A
	}
	return nil
}

func (m *CalculateInt) GetB() *Number_ {
	if m != nil {
		return m.B
	}
	return nil
}

type CalculateString struct {
	A                    *String_ `protobuf:"bytes,1,opt,name=a,proto3" json:"a,omitempty"`
	B                    *String_ `protobuf:"bytes,2,opt,name=b,proto3" json:"b,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CalculateString) Reset()         { *m = CalculateString{} }
func (m *CalculateString) String() string { return proto.CompactTextString(m) }
func (*CalculateString) ProtoMessage()    {}
func (*CalculateString) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d8b31321fcad8b1, []int{3}
}

func (m *CalculateString) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalculateString.Unmarshal(m, b)
}
func (m *CalculateString) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalculateString.Marshal(b, m, deterministic)
}
func (m *CalculateString) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalculateString.Merge(m, src)
}
func (m *CalculateString) XXX_Size() int {
	return xxx_messageInfo_CalculateString.Size(m)
}
func (m *CalculateString) XXX_DiscardUnknown() {
	xxx_messageInfo_CalculateString.DiscardUnknown(m)
}

var xxx_messageInfo_CalculateString proto.InternalMessageInfo

func (m *CalculateString) GetA() *String_ {
	if m != nil {
		return m.A
	}
	return nil
}

func (m *CalculateString) GetB() *String_ {
	if m != nil {
		return m.B
	}
	return nil
}

func init() {
	proto.RegisterType((*Number_)(nil), "calculate.Number_")
	proto.RegisterType((*String_)(nil), "calculate.String_")
	proto.RegisterType((*CalculateInt)(nil), "calculate.CalculateInt")
	proto.RegisterType((*CalculateString)(nil), "calculate.CalculateString")
}

func init() { proto.RegisterFile("calculate.proto", fileDescriptor_1d8b31321fcad8b1) }

var fileDescriptor_1d8b31321fcad8b1 = []byte{
	// 207 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0x4e, 0xcc, 0x49,
	0x2e, 0xcd, 0x49, 0x2c, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x84, 0x0b, 0x28,
	0x49, 0x73, 0xb1, 0xfb, 0x95, 0xe6, 0x26, 0xa5, 0x16, 0xc5, 0x0b, 0x09, 0x70, 0x31, 0xe7, 0x95,
	0xe6, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x30, 0x07, 0x81, 0x98, 0x20, 0xc9, 0xe0, 0x92, 0xa2, 0xcc,
	0xbc, 0x74, 0xb0, 0x64, 0x71, 0x49, 0x11, 0x58, 0x92, 0x33, 0x08, 0xc4, 0x54, 0x0a, 0xe2, 0xe2,
	0x71, 0x86, 0x19, 0xe3, 0x99, 0x57, 0x22, 0xa4, 0xc0, 0xc5, 0x98, 0x08, 0x96, 0xe7, 0x36, 0x12,
	0xd2, 0x43, 0xd8, 0x08, 0x35, 0x3d, 0x88, 0x31, 0x11, 0xa4, 0x22, 0x49, 0x82, 0x09, 0xb7, 0x8a,
	0x24, 0xa5, 0x50, 0x2e, 0x7e, 0xb8, 0x99, 0x10, 0x9b, 0x71, 0x19, 0x0b, 0x75, 0x17, 0x1e, 0x63,
	0xe1, 0x2a, 0x92, 0x8c, 0x0e, 0x33, 0x72, 0x71, 0xc1, 0xcd, 0x8d, 0x17, 0xb2, 0xe2, 0xe2, 0xcc,
	0x03, 0xdb, 0xe9, 0x98, 0x92, 0x22, 0x24, 0x8e, 0xa4, 0x05, 0xd9, 0x3f, 0x52, 0x58, 0x9c, 0xa8,
	0xc4, 0x80, 0xd0, 0xeb, 0x5b, 0x9a, 0x43, 0xaa, 0x5e, 0x5b, 0x2e, 0xce, 0x62, 0xb0, 0xa3, 0x40,
	0xf6, 0x4a, 0x61, 0xd3, 0x0b, 0x71, 0xb3, 0x14, 0x16, 0x6f, 0x28, 0x31, 0x24, 0xb1, 0x81, 0x23,
	0xcf, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x84, 0xd4, 0x9e, 0x1c, 0xcf, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Calculate_Client is the client API for Calculate_ service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type Calculate_Client interface {
	NumberAdd(ctx context.Context, in *CalculateInt, opts ...grpc.CallOption) (*Number_, error)
	NumberMul(ctx context.Context, in *CalculateInt, opts ...grpc.CallOption) (*Number_, error)
	StringAdd(ctx context.Context, in *CalculateString, opts ...grpc.CallOption) (*String_, error)
}

type calculate_Client struct {
	cc *grpc.ClientConn
}

func NewCalculate_Client(cc *grpc.ClientConn) Calculate_Client {
	return &calculate_Client{cc}
}

func (c *calculate_Client) NumberAdd(ctx context.Context, in *CalculateInt, opts ...grpc.CallOption) (*Number_, error) {
	out := new(Number_)
	err := c.cc.Invoke(ctx, "/calculate.Calculate_/numberAdd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculate_Client) NumberMul(ctx context.Context, in *CalculateInt, opts ...grpc.CallOption) (*Number_, error) {
	out := new(Number_)
	err := c.cc.Invoke(ctx, "/calculate.Calculate_/numberMul", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculate_Client) StringAdd(ctx context.Context, in *CalculateString, opts ...grpc.CallOption) (*String_, error) {
	out := new(String_)
	err := c.cc.Invoke(ctx, "/calculate.Calculate_/stringAdd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Calculate_Server is the server API for Calculate_ service.
type Calculate_Server interface {
	NumberAdd(context.Context, *CalculateInt) (*Number_, error)
	NumberMul(context.Context, *CalculateInt) (*Number_, error)
	StringAdd(context.Context, *CalculateString) (*String_, error)
}

func RegisterCalculate_Server(s *grpc.Server, srv Calculate_Server) {
	s.RegisterService(&_Calculate__serviceDesc, srv)
}

func _Calculate__NumberAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculateInt)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Calculate_Server).NumberAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculate.Calculate_/NumberAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Calculate_Server).NumberAdd(ctx, req.(*CalculateInt))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calculate__NumberMul_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculateInt)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Calculate_Server).NumberMul(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculate.Calculate_/NumberMul",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Calculate_Server).NumberMul(ctx, req.(*CalculateInt))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calculate__StringAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculateString)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Calculate_Server).StringAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculate.Calculate_/StringAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Calculate_Server).StringAdd(ctx, req.(*CalculateString))
	}
	return interceptor(ctx, in, info, handler)
}

var _Calculate__serviceDesc = grpc.ServiceDesc{
	ServiceName: "calculate.Calculate_",
	HandlerType: (*Calculate_Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "numberAdd",
			Handler:    _Calculate__NumberAdd_Handler,
		},
		{
			MethodName: "numberMul",
			Handler:    _Calculate__NumberMul_Handler,
		},
		{
			MethodName: "stringAdd",
			Handler:    _Calculate__StringAdd_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "calculate.proto",
}
