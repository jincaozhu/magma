// Code generated by protoc-gen-go. DO NOT EDIT.
// source: feg/protos/swx_proxy.proto

package protos // import "magma/feg/cloud/go/protos"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

// SwxErrorCode reflects Experimental-Result values which are 3GPP failures
// to be processed by EPC. Diameter Base Protocol errors are reflected in gRPC status code
type SwxErrorCode int32

const (
	SwxErrorCode_ERROR_UNDEFINED               SwxErrorCode = 0
	SwxErrorCode_IDENTITY_ALREADY_REGISTERED   SwxErrorCode = 5005
	SwxErrorCode_USER_NO_NON_3GPP_SUBSCRIPTION SwxErrorCode = 5450
)

var SwxErrorCode_name = map[int32]string{
	0:    "ERROR_UNDEFINED",
	5005: "IDENTITY_ALREADY_REGISTERED",
	5450: "USER_NO_NON_3GPP_SUBSCRIPTION",
}
var SwxErrorCode_value = map[string]int32{
	"ERROR_UNDEFINED":               0,
	"IDENTITY_ALREADY_REGISTERED":   5005,
	"USER_NO_NON_3GPP_SUBSCRIPTION": 5450,
}

func (x SwxErrorCode) String() string {
	return proto.EnumName(SwxErrorCode_name, int32(x))
}
func (SwxErrorCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_swx_proxy_492c7361489b2d61, []int{0}
}

type AuthenticationScheme int32

const (
	AuthenticationScheme_EAP_AKA       AuthenticationScheme = 0
	AuthenticationScheme_EAP_AKA_PRIME AuthenticationScheme = 1
)

var AuthenticationScheme_name = map[int32]string{
	0: "EAP_AKA",
	1: "EAP_AKA_PRIME",
}
var AuthenticationScheme_value = map[string]int32{
	"EAP_AKA":       0,
	"EAP_AKA_PRIME": 1,
}

func (x AuthenticationScheme) String() string {
	return proto.EnumName(AuthenticationScheme_name, int32(x))
}
func (AuthenticationScheme) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_swx_proxy_492c7361489b2d61, []int{1}
}

// AuthenticationRequest (Section 8.2.2.1)
type AuthenticationRequest struct {
	// Subscriber identifier
	UserName string `protobuf:"bytes,1,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	// Number of authentication vectors requested
	SipNumAuthVectors uint32 `protobuf:"varint,2,opt,name=sip_num_auth_vectors,json=sipNumAuthVectors,proto3" json:"sip_num_auth_vectors,omitempty"`
	// EAP-AKA or EAP-AKA'
	AuthenticationScheme AuthenticationScheme `protobuf:"varint,3,opt,name=authentication_scheme,json=authenticationScheme,proto3,enum=magma.feg.AuthenticationScheme" json:"authentication_scheme,omitempty"`
	// Concatenation of RAND and AUTS in the case of resync
	ResyncInfo           []byte   `protobuf:"bytes,4,opt,name=resync_info,json=resyncInfo,proto3" json:"resync_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthenticationRequest) Reset()         { *m = AuthenticationRequest{} }
func (m *AuthenticationRequest) String() string { return proto.CompactTextString(m) }
func (*AuthenticationRequest) ProtoMessage()    {}
func (*AuthenticationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_swx_proxy_492c7361489b2d61, []int{0}
}
func (m *AuthenticationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthenticationRequest.Unmarshal(m, b)
}
func (m *AuthenticationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthenticationRequest.Marshal(b, m, deterministic)
}
func (dst *AuthenticationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticationRequest.Merge(dst, src)
}
func (m *AuthenticationRequest) XXX_Size() int {
	return xxx_messageInfo_AuthenticationRequest.Size(m)
}
func (m *AuthenticationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticationRequest proto.InternalMessageInfo

func (m *AuthenticationRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *AuthenticationRequest) GetSipNumAuthVectors() uint32 {
	if m != nil {
		return m.SipNumAuthVectors
	}
	return 0
}

func (m *AuthenticationRequest) GetAuthenticationScheme() AuthenticationScheme {
	if m != nil {
		return m.AuthenticationScheme
	}
	return AuthenticationScheme_EAP_AKA
}

func (m *AuthenticationRequest) GetResyncInfo() []byte {
	if m != nil {
		return m.ResyncInfo
	}
	return nil
}

// MultimediaAuthenticationAnswer (Section 8.2.2.1)
type AuthenticationAnswer struct {
	// Subscriber identifier
	UserName string `protobuf:"bytes,1,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	// For details about fields read 3GPP 29.273
	SipAuthVectors       []*AuthenticationAnswer_SIPAuthVector `protobuf:"bytes,2,rep,name=sip_auth_vectors,json=sipAuthVectors,proto3" json:"sip_auth_vectors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *AuthenticationAnswer) Reset()         { *m = AuthenticationAnswer{} }
func (m *AuthenticationAnswer) String() string { return proto.CompactTextString(m) }
func (*AuthenticationAnswer) ProtoMessage()    {}
func (*AuthenticationAnswer) Descriptor() ([]byte, []int) {
	return fileDescriptor_swx_proxy_492c7361489b2d61, []int{1}
}
func (m *AuthenticationAnswer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthenticationAnswer.Unmarshal(m, b)
}
func (m *AuthenticationAnswer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthenticationAnswer.Marshal(b, m, deterministic)
}
func (dst *AuthenticationAnswer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticationAnswer.Merge(dst, src)
}
func (m *AuthenticationAnswer) XXX_Size() int {
	return xxx_messageInfo_AuthenticationAnswer.Size(m)
}
func (m *AuthenticationAnswer) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticationAnswer.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticationAnswer proto.InternalMessageInfo

func (m *AuthenticationAnswer) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *AuthenticationAnswer) GetSipAuthVectors() []*AuthenticationAnswer_SIPAuthVector {
	if m != nil {
		return m.SipAuthVectors
	}
	return nil
}

// Only for EAP-AKA/EAP-AKA'
type AuthenticationAnswer_SIPAuthVector struct {
	// Contains one of EAP-AKA or EAP-AKA'
	AuthenticationScheme AuthenticationScheme `protobuf:"varint,1,opt,name=authentication_scheme,json=authenticationScheme,proto3,enum=magma.feg.AuthenticationScheme" json:"authentication_scheme,omitempty"`
	// Concatenation of challenge RAND and token AUTN
	RandAutn []byte `protobuf:"bytes,2,opt,name=rand_autn,json=randAutn,proto3" json:"rand_autn,omitempty"`
	// Expected response
	Xres []byte `protobuf:"bytes,3,opt,name=xres,proto3" json:"xres,omitempty"`
	// Confidentiality Key
	ConfidentialityKey []byte `protobuf:"bytes,4,opt,name=confidentiality_key,json=confidentialityKey,proto3" json:"confidentiality_key,omitempty"`
	// Integrity Key
	IntegrityKey         []byte   `protobuf:"bytes,5,opt,name=integrity_key,json=integrityKey,proto3" json:"integrity_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthenticationAnswer_SIPAuthVector) Reset()         { *m = AuthenticationAnswer_SIPAuthVector{} }
func (m *AuthenticationAnswer_SIPAuthVector) String() string { return proto.CompactTextString(m) }
func (*AuthenticationAnswer_SIPAuthVector) ProtoMessage()    {}
func (*AuthenticationAnswer_SIPAuthVector) Descriptor() ([]byte, []int) {
	return fileDescriptor_swx_proxy_492c7361489b2d61, []int{1, 0}
}
func (m *AuthenticationAnswer_SIPAuthVector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthenticationAnswer_SIPAuthVector.Unmarshal(m, b)
}
func (m *AuthenticationAnswer_SIPAuthVector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthenticationAnswer_SIPAuthVector.Marshal(b, m, deterministic)
}
func (dst *AuthenticationAnswer_SIPAuthVector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticationAnswer_SIPAuthVector.Merge(dst, src)
}
func (m *AuthenticationAnswer_SIPAuthVector) XXX_Size() int {
	return xxx_messageInfo_AuthenticationAnswer_SIPAuthVector.Size(m)
}
func (m *AuthenticationAnswer_SIPAuthVector) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticationAnswer_SIPAuthVector.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticationAnswer_SIPAuthVector proto.InternalMessageInfo

func (m *AuthenticationAnswer_SIPAuthVector) GetAuthenticationScheme() AuthenticationScheme {
	if m != nil {
		return m.AuthenticationScheme
	}
	return AuthenticationScheme_EAP_AKA
}

func (m *AuthenticationAnswer_SIPAuthVector) GetRandAutn() []byte {
	if m != nil {
		return m.RandAutn
	}
	return nil
}

func (m *AuthenticationAnswer_SIPAuthVector) GetXres() []byte {
	if m != nil {
		return m.Xres
	}
	return nil
}

func (m *AuthenticationAnswer_SIPAuthVector) GetConfidentialityKey() []byte {
	if m != nil {
		return m.ConfidentialityKey
	}
	return nil
}

func (m *AuthenticationAnswer_SIPAuthVector) GetIntegrityKey() []byte {
	if m != nil {
		return m.IntegrityKey
	}
	return nil
}

// RegistrationRequest:
// ServerAssignmentRequest with ServerAssignmentType set to REGISTRATION (Section 8.2.2.3)
type RegistrationRequest struct {
	// Subscriber identifier
	UserName             string   `protobuf:"bytes,1,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegistrationRequest) Reset()         { *m = RegistrationRequest{} }
func (m *RegistrationRequest) String() string { return proto.CompactTextString(m) }
func (*RegistrationRequest) ProtoMessage()    {}
func (*RegistrationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_swx_proxy_492c7361489b2d61, []int{2}
}
func (m *RegistrationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegistrationRequest.Unmarshal(m, b)
}
func (m *RegistrationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegistrationRequest.Marshal(b, m, deterministic)
}
func (dst *RegistrationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegistrationRequest.Merge(dst, src)
}
func (m *RegistrationRequest) XXX_Size() int {
	return xxx_messageInfo_RegistrationRequest.Size(m)
}
func (m *RegistrationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegistrationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegistrationRequest proto.InternalMessageInfo

func (m *RegistrationRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

// RegistrationAnswer:
// ServerAssignmentAnswer with ServerAssignmentType set to REGISTRATION (Section 8.2.2.3)
type RegistrationAnswer struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegistrationAnswer) Reset()         { *m = RegistrationAnswer{} }
func (m *RegistrationAnswer) String() string { return proto.CompactTextString(m) }
func (*RegistrationAnswer) ProtoMessage()    {}
func (*RegistrationAnswer) Descriptor() ([]byte, []int) {
	return fileDescriptor_swx_proxy_492c7361489b2d61, []int{3}
}
func (m *RegistrationAnswer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegistrationAnswer.Unmarshal(m, b)
}
func (m *RegistrationAnswer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegistrationAnswer.Marshal(b, m, deterministic)
}
func (dst *RegistrationAnswer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegistrationAnswer.Merge(dst, src)
}
func (m *RegistrationAnswer) XXX_Size() int {
	return xxx_messageInfo_RegistrationAnswer.Size(m)
}
func (m *RegistrationAnswer) XXX_DiscardUnknown() {
	xxx_messageInfo_RegistrationAnswer.DiscardUnknown(m)
}

var xxx_messageInfo_RegistrationAnswer proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AuthenticationRequest)(nil), "magma.feg.AuthenticationRequest")
	proto.RegisterType((*AuthenticationAnswer)(nil), "magma.feg.AuthenticationAnswer")
	proto.RegisterType((*AuthenticationAnswer_SIPAuthVector)(nil), "magma.feg.AuthenticationAnswer.SIPAuthVector")
	proto.RegisterType((*RegistrationRequest)(nil), "magma.feg.RegistrationRequest")
	proto.RegisterType((*RegistrationAnswer)(nil), "magma.feg.RegistrationAnswer")
	proto.RegisterEnum("magma.feg.SwxErrorCode", SwxErrorCode_name, SwxErrorCode_value)
	proto.RegisterEnum("magma.feg.AuthenticationScheme", AuthenticationScheme_name, AuthenticationScheme_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SwxProxyClient is the client API for SwxProxy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SwxProxyClient interface {
	// Retrieve authentication vectors from the HSS using MAR/MAA
	Authenticate(ctx context.Context, in *AuthenticationRequest, opts ...grpc.CallOption) (*AuthenticationAnswer, error)
	// Register the AAA server serving a user to the HSS using SAR/SAA
	Register(ctx context.Context, in *RegistrationRequest, opts ...grpc.CallOption) (*RegistrationAnswer, error)
}

type swxProxyClient struct {
	cc *grpc.ClientConn
}

func NewSwxProxyClient(cc *grpc.ClientConn) SwxProxyClient {
	return &swxProxyClient{cc}
}

func (c *swxProxyClient) Authenticate(ctx context.Context, in *AuthenticationRequest, opts ...grpc.CallOption) (*AuthenticationAnswer, error) {
	out := new(AuthenticationAnswer)
	err := c.cc.Invoke(ctx, "/magma.feg.SwxProxy/Authenticate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *swxProxyClient) Register(ctx context.Context, in *RegistrationRequest, opts ...grpc.CallOption) (*RegistrationAnswer, error) {
	out := new(RegistrationAnswer)
	err := c.cc.Invoke(ctx, "/magma.feg.SwxProxy/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SwxProxyServer is the server API for SwxProxy service.
type SwxProxyServer interface {
	// Retrieve authentication vectors from the HSS using MAR/MAA
	Authenticate(context.Context, *AuthenticationRequest) (*AuthenticationAnswer, error)
	// Register the AAA server serving a user to the HSS using SAR/SAA
	Register(context.Context, *RegistrationRequest) (*RegistrationAnswer, error)
}

func RegisterSwxProxyServer(s *grpc.Server, srv SwxProxyServer) {
	s.RegisterService(&_SwxProxy_serviceDesc, srv)
}

func _SwxProxy_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SwxProxyServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.feg.SwxProxy/Authenticate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SwxProxyServer).Authenticate(ctx, req.(*AuthenticationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SwxProxy_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SwxProxyServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.feg.SwxProxy/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SwxProxyServer).Register(ctx, req.(*RegistrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SwxProxy_serviceDesc = grpc.ServiceDesc{
	ServiceName: "magma.feg.SwxProxy",
	HandlerType: (*SwxProxyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Authenticate",
			Handler:    _SwxProxy_Authenticate_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _SwxProxy_Register_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "feg/protos/swx_proxy.proto",
}

func init() {
	proto.RegisterFile("feg/protos/swx_proxy.proto", fileDescriptor_swx_proxy_492c7361489b2d61)
}

var fileDescriptor_swx_proxy_492c7361489b2d61 = []byte{
	// 569 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x86, 0x97, 0x6d, 0x40, 0x7b, 0xd6, 0x8e, 0xce, 0xeb, 0xa4, 0xd2, 0x6a, 0x2c, 0x0a, 0x17,
	0x54, 0x93, 0x68, 0xa4, 0x4e, 0xe2, 0x3e, 0x5b, 0xcd, 0x14, 0x15, 0xd2, 0xc8, 0xe9, 0x40, 0xe3,
	0xc6, 0x0a, 0xa9, 0x9b, 0x45, 0x2c, 0x76, 0xb1, 0x13, 0xda, 0x3c, 0x04, 0x6f, 0xc2, 0x93, 0xf0,
	0x0c, 0x3c, 0x02, 0x6f, 0xc0, 0x0d, 0x4a, 0xda, 0x8d, 0x76, 0x5a, 0x01, 0x89, 0xab, 0x24, 0xbf,
	0xff, 0xfc, 0x39, 0xe7, 0xf3, 0x89, 0xa1, 0x39, 0x66, 0xa1, 0x39, 0x91, 0x22, 0x11, 0xca, 0x54,
	0xd3, 0x19, 0x9d, 0x48, 0x31, 0xcb, 0x3a, 0x85, 0x80, 0xca, 0xb1, 0x1f, 0xc6, 0x7e, 0x67, 0xcc,
	0x42, 0xe3, 0xbb, 0x06, 0x07, 0x56, 0x9a, 0x5c, 0x31, 0x9e, 0x44, 0x81, 0x9f, 0x44, 0x82, 0x13,
	0xf6, 0x29, 0x65, 0x2a, 0x41, 0x2d, 0x28, 0xa7, 0x8a, 0x49, 0xca, 0xfd, 0x98, 0x35, 0x34, 0x5d,
	0x6b, 0x97, 0x49, 0x29, 0x17, 0x1c, 0x3f, 0x66, 0xc8, 0x84, 0xba, 0x8a, 0x26, 0x94, 0xa7, 0x31,
	0xf5, 0xd3, 0xe4, 0x8a, 0x7e, 0x66, 0x41, 0x22, 0xa4, 0x6a, 0x6c, 0xea, 0x5a, 0xbb, 0x4a, 0xf6,
	0x54, 0x34, 0x71, 0xd2, 0x38, 0xcf, 0x7d, 0x3b, 0x5f, 0x40, 0x43, 0x38, 0xf0, 0x57, 0x3e, 0x43,
	0x55, 0x70, 0xc5, 0x62, 0xd6, 0xd8, 0xd2, 0xb5, 0xf6, 0x6e, 0xf7, 0xa8, 0x73, 0x5b, 0x52, 0x67,
	0xb5, 0x1c, 0xaf, 0xb0, 0x91, 0xba, 0x7f, 0x8f, 0x8a, 0x8e, 0x60, 0x47, 0x32, 0x95, 0xf1, 0x80,
	0x46, 0x7c, 0x2c, 0x1a, 0xdb, 0xba, 0xd6, 0xae, 0x10, 0x98, 0x4b, 0x36, 0x1f, 0x0b, 0xe3, 0xe7,
	0x26, 0xd4, 0x57, 0xf3, 0x2c, 0xae, 0xa6, 0x4c, 0xfe, 0xb9, 0xbb, 0x77, 0x50, 0xcb, 0xbb, 0xbb,
	0xd3, 0xd9, 0x56, 0x7b, 0xa7, 0xfb, 0x62, 0x6d, 0x9d, 0xf3, 0xdc, 0x8e, 0x67, 0xbb, 0xbf, 0xdb,
	0x26, 0xbb, 0x2a, 0x9a, 0x2c, 0x51, 0x68, 0xfe, 0xd0, 0xa0, 0xba, 0xe2, 0x58, 0xcf, 0x45, 0xfb,
	0x1f, 0x2e, 0x2d, 0x28, 0x4b, 0x9f, 0x8f, 0xf2, 0x0e, 0x78, 0xb1, 0x27, 0x15, 0x52, 0xca, 0x05,
	0x2b, 0x4d, 0x38, 0x42, 0xb0, 0x3d, 0x93, 0x4c, 0x15, 0xe4, 0x2b, 0xa4, 0xb8, 0x47, 0x26, 0xec,
	0x07, 0x82, 0x8f, 0xa3, 0x51, 0x1e, 0xe5, 0x5f, 0x47, 0x49, 0x46, 0x3f, 0xb2, 0x6c, 0x01, 0x14,
	0xdd, 0x59, 0xea, 0xb3, 0x0c, 0x3d, 0x83, 0x6a, 0xc4, 0x13, 0x16, 0xca, 0x1b, 0xeb, 0x83, 0xc2,
	0x5a, 0xb9, 0x15, 0xfb, 0x2c, 0x33, 0xba, 0xb0, 0x4f, 0x58, 0x18, 0xa9, 0x44, 0xfe, 0xf3, 0x64,
	0x19, 0x75, 0x40, 0xcb, 0xef, 0xcc, 0xb1, 0x1e, 0x47, 0x50, 0xf1, 0xa6, 0x33, 0x2c, 0xa5, 0x90,
	0x67, 0x62, 0xc4, 0xd0, 0x3e, 0x3c, 0xc6, 0x84, 0x0c, 0x08, 0xbd, 0x70, 0x7a, 0xf8, 0x95, 0xed,
	0xe0, 0x5e, 0x6d, 0x03, 0xe9, 0xd0, 0xb2, 0x7b, 0xd8, 0x19, 0xda, 0xc3, 0x4b, 0x6a, 0xbd, 0x26,
	0xd8, 0xea, 0x5d, 0x52, 0x82, 0xcf, 0x6d, 0x6f, 0x88, 0x09, 0xee, 0xd5, 0xbe, 0x3c, 0x47, 0x06,
	0x1c, 0x5e, 0x78, 0x98, 0x50, 0x67, 0x40, 0x9d, 0x81, 0x43, 0x4f, 0xce, 0x5d, 0x97, 0x7a, 0x17,
	0xa7, 0xde, 0x19, 0xb1, 0xdd, 0xa1, 0x3d, 0x70, 0x6a, 0xdf, 0x8e, 0x8f, 0x5f, 0xde, 0x9d, 0x98,
	0x05, 0xd3, 0x1d, 0x78, 0x84, 0x2d, 0x97, 0x5a, 0x7d, 0xab, 0xb6, 0x81, 0xf6, 0xa0, 0xba, 0x78,
	0xa0, 0x2e, 0xb1, 0xdf, 0xe0, 0x9a, 0xd6, 0xfd, 0xaa, 0x41, 0xc9, 0x9b, 0xce, 0xdc, 0xfc, 0x3f,
	0x43, 0x1e, 0x54, 0x96, 0x42, 0x18, 0xd2, 0xd7, 0xee, 0xe3, 0x02, 0x4a, 0xf3, 0xe8, 0x2f, 0x93,
	0x65, 0x6c, 0xa0, 0x3e, 0x94, 0xe6, 0x68, 0x98, 0x44, 0x4f, 0x97, 0xec, 0xf7, 0x30, 0x6e, 0x1e,
	0xae, 0x59, 0xbf, 0x09, 0x3b, 0x6d, 0xbd, 0x7f, 0x52, 0x38, 0xcc, 0xfc, 0x9c, 0x08, 0xae, 0x45,
	0x3a, 0x32, 0x43, 0xb1, 0x38, 0x30, 0x3e, 0x3c, 0x2c, 0xae, 0x27, 0xbf, 0x02, 0x00, 0x00, 0xff,
	0xff, 0xa4, 0x33, 0x22, 0x16, 0x45, 0x04, 0x00, 0x00,
}
