// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sig.proto

package ftl_mixadp

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type VerifyReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyReq) Reset()         { *m = VerifyReq{} }
func (m *VerifyReq) String() string { return proto.CompactTextString(m) }
func (*VerifyReq) ProtoMessage()    {}
func (*VerifyReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c066fc0129d1698, []int{0}
}

func (m *VerifyReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyReq.Unmarshal(m, b)
}
func (m *VerifyReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyReq.Marshal(b, m, deterministic)
}
func (m *VerifyReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyReq.Merge(m, src)
}
func (m *VerifyReq) XXX_Size() int {
	return xxx_messageInfo_VerifyReq.Size(m)
}
func (m *VerifyReq) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyReq.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyReq proto.InternalMessageInfo

type VerifyRsp struct {
	ErrCode string `protobuf:"bytes,1,opt,name=err_code,json=errCode,proto3" json:"err_code,omitempty"`
	// 下面列出的错误码：
	//    "sig_err": 票据验证不通过
	//    "no-sig": 没有发现有效票据
	ErrMsg               string   `protobuf:"bytes,2,opt,name=err_msg,json=errMsg,proto3" json:"err_msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyRsp) Reset()         { *m = VerifyRsp{} }
func (m *VerifyRsp) String() string { return proto.CompactTextString(m) }
func (*VerifyRsp) ProtoMessage()    {}
func (*VerifyRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c066fc0129d1698, []int{1}
}

func (m *VerifyRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyRsp.Unmarshal(m, b)
}
func (m *VerifyRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyRsp.Marshal(b, m, deterministic)
}
func (m *VerifyRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyRsp.Merge(m, src)
}
func (m *VerifyRsp) XXX_Size() int {
	return xxx_messageInfo_VerifyRsp.Size(m)
}
func (m *VerifyRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyRsp.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyRsp proto.InternalMessageInfo

func (m *VerifyRsp) GetErrCode() string {
	if m != nil {
		return m.ErrCode
	}
	return ""
}

func (m *VerifyRsp) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func init() {
	proto.RegisterType((*VerifyReq)(nil), "ftl_mixadp.VerifyReq")
	proto.RegisterType((*VerifyRsp)(nil), "ftl_mixadp.VerifyRsp")
}

func init() { proto.RegisterFile("sig.proto", fileDescriptor_2c066fc0129d1698) }

var fileDescriptor_2c066fc0129d1698 = []byte{
	// 144 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2c, 0xce, 0x4c, 0xd7,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4a, 0x2b, 0xc9, 0x89, 0xcf, 0xcd, 0xac, 0x48, 0x4c,
	0x29, 0x50, 0xe2, 0xe6, 0xe2, 0x0c, 0x4b, 0x2d, 0xca, 0x4c, 0xab, 0x0c, 0x4a, 0x2d, 0x54, 0xb2,
	0x87, 0x73, 0x8a, 0x0b, 0x84, 0x24, 0xb9, 0x38, 0x52, 0x8b, 0x8a, 0xe2, 0x93, 0xf3, 0x53, 0x52,
	0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xd8, 0x53, 0x8b, 0x8a, 0x9c, 0xf3, 0x53, 0x52, 0x85,
	0xc4, 0xb9, 0x40, 0xcc, 0xf8, 0xdc, 0xe2, 0x74, 0x09, 0x26, 0xb0, 0x0c, 0x5b, 0x6a, 0x51, 0x91,
	0x6f, 0x71, 0xba, 0x91, 0x2d, 0x17, 0x73, 0x70, 0x66, 0xba, 0x90, 0x19, 0x17, 0x1b, 0xc4, 0x1c,
	0x21, 0x51, 0x3d, 0x84, 0x5d, 0x7a, 0x70, 0x8b, 0xa4, 0xb0, 0x09, 0x17, 0x17, 0x24, 0xb1, 0x81,
	0xdd, 0x67, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xdd, 0xa8, 0xc8, 0x3b, 0xac, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SigClient is the client API for Sig service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SigClient interface {
	// 验票
	//注意：需要Mixer把收到的请求的metadata放在这个验票请求里，票据都在metadata里。
	Verify(ctx context.Context, in *VerifyReq, opts ...grpc.CallOption) (*VerifyRsp, error)
}

type sigClient struct {
	cc *grpc.ClientConn
}

func NewSigClient(cc *grpc.ClientConn) SigClient {
	return &sigClient{cc}
}

func (c *sigClient) Verify(ctx context.Context, in *VerifyReq, opts ...grpc.CallOption) (*VerifyRsp, error) {
	out := new(VerifyRsp)
	err := c.cc.Invoke(ctx, "/ftl_mixadp.Sig/Verify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SigServer is the server API for Sig service.
type SigServer interface {
	// 验票
	//注意：需要Mixer把收到的请求的metadata放在这个验票请求里，票据都在metadata里。
	Verify(context.Context, *VerifyReq) (*VerifyRsp, error)
}

// UnimplementedSigServer can be embedded to have forward compatible implementations.
type UnimplementedSigServer struct {
}

func (*UnimplementedSigServer) Verify(ctx context.Context, req *VerifyReq) (*VerifyRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Verify not implemented")
}

func RegisterSigServer(s *grpc.Server, srv SigServer) {
	s.RegisterService(&_Sig_serviceDesc, srv)
}

func _Sig_Verify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SigServer).Verify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ftl_mixadp.Sig/Verify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SigServer).Verify(ctx, req.(*VerifyReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Sig_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ftl_mixadp.Sig",
	HandlerType: (*SigServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Verify",
			Handler:    _Sig_Verify_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sig.proto",
}