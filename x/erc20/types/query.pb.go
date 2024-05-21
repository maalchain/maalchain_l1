// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: evmos/erc20/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryTokenPairsRequest is the request type for the Query/TokenPairs RPC
// method.
type QueryTokenPairsRequest struct {
	// pagination defines an optional pagination for the request.
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryTokenPairsRequest) Reset()         { *m = QueryTokenPairsRequest{} }
func (m *QueryTokenPairsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryTokenPairsRequest) ProtoMessage()    {}
func (*QueryTokenPairsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fba814bce17cabdf, []int{0}
}
func (m *QueryTokenPairsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryTokenPairsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryTokenPairsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryTokenPairsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryTokenPairsRequest.Merge(m, src)
}
func (m *QueryTokenPairsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryTokenPairsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryTokenPairsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryTokenPairsRequest proto.InternalMessageInfo

func (m *QueryTokenPairsRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// QueryTokenPairsResponse is the response type for the Query/TokenPairs RPC
// method.
type QueryTokenPairsResponse struct {
	// token_pairs is a slice of registered token pairs for the erc20 module
	TokenPairs []TokenPair `protobuf:"bytes,1,rep,name=token_pairs,json=tokenPairs,proto3" json:"token_pairs"`
	// pagination defines the pagination in the response.
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryTokenPairsResponse) Reset()         { *m = QueryTokenPairsResponse{} }
func (m *QueryTokenPairsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryTokenPairsResponse) ProtoMessage()    {}
func (*QueryTokenPairsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fba814bce17cabdf, []int{1}
}
func (m *QueryTokenPairsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryTokenPairsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryTokenPairsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryTokenPairsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryTokenPairsResponse.Merge(m, src)
}
func (m *QueryTokenPairsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryTokenPairsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryTokenPairsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryTokenPairsResponse proto.InternalMessageInfo

func (m *QueryTokenPairsResponse) GetTokenPairs() []TokenPair {
	if m != nil {
		return m.TokenPairs
	}
	return nil
}

func (m *QueryTokenPairsResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// QueryTokenPairRequest is the request type for the Query/TokenPair RPC method.
type QueryTokenPairRequest struct {
	// token identifier can be either the hex contract address of the ERC20 or the
	// Cosmos base denomination
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (m *QueryTokenPairRequest) Reset()         { *m = QueryTokenPairRequest{} }
func (m *QueryTokenPairRequest) String() string { return proto.CompactTextString(m) }
func (*QueryTokenPairRequest) ProtoMessage()    {}
func (*QueryTokenPairRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fba814bce17cabdf, []int{2}
}
func (m *QueryTokenPairRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryTokenPairRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryTokenPairRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryTokenPairRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryTokenPairRequest.Merge(m, src)
}
func (m *QueryTokenPairRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryTokenPairRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryTokenPairRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryTokenPairRequest proto.InternalMessageInfo

func (m *QueryTokenPairRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

// QueryTokenPairResponse is the response type for the Query/TokenPair RPC
// method.
type QueryTokenPairResponse struct {
	// token_pairs returns the info about a registered token pair for the erc20 module
	TokenPair TokenPair `protobuf:"bytes,1,opt,name=token_pair,json=tokenPair,proto3" json:"token_pair"`
}

func (m *QueryTokenPairResponse) Reset()         { *m = QueryTokenPairResponse{} }
func (m *QueryTokenPairResponse) String() string { return proto.CompactTextString(m) }
func (*QueryTokenPairResponse) ProtoMessage()    {}
func (*QueryTokenPairResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fba814bce17cabdf, []int{3}
}
func (m *QueryTokenPairResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryTokenPairResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryTokenPairResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryTokenPairResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryTokenPairResponse.Merge(m, src)
}
func (m *QueryTokenPairResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryTokenPairResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryTokenPairResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryTokenPairResponse proto.InternalMessageInfo

func (m *QueryTokenPairResponse) GetTokenPair() TokenPair {
	if m != nil {
		return m.TokenPair
	}
	return TokenPair{}
}

// QueryParamsRequest is the request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fba814bce17cabdf, []int{4}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is the response type for the Query/Params RPC
// method.
type QueryParamsResponse struct {
	// params are the erc20 module parameters
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fba814bce17cabdf, []int{5}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func init() {
	proto.RegisterType((*QueryTokenPairsRequest)(nil), "evmos.erc20.v1.QueryTokenPairsRequest")
	proto.RegisterType((*QueryTokenPairsResponse)(nil), "evmos.erc20.v1.QueryTokenPairsResponse")
	proto.RegisterType((*QueryTokenPairRequest)(nil), "evmos.erc20.v1.QueryTokenPairRequest")
	proto.RegisterType((*QueryTokenPairResponse)(nil), "evmos.erc20.v1.QueryTokenPairResponse")
	proto.RegisterType((*QueryParamsRequest)(nil), "evmos.erc20.v1.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "evmos.erc20.v1.QueryParamsResponse")
}

func init() { proto.RegisterFile("evmos/erc20/v1/query.proto", fileDescriptor_fba814bce17cabdf) }

var fileDescriptor_fba814bce17cabdf = []byte{
	// 511 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0x3f, 0x6f, 0xd3, 0x40,
	0x14, 0x8f, 0x5b, 0x1a, 0x29, 0x2f, 0x12, 0xc3, 0x11, 0x4a, 0xb1, 0x2a, 0xb7, 0x18, 0x91, 0x94,
	0x4a, 0xdc, 0x91, 0xc0, 0x8c, 0x50, 0x85, 0x60, 0x60, 0x09, 0x11, 0x03, 0x62, 0x81, 0x4b, 0x78,
	0x72, 0x2d, 0x1a, 0x9f, 0xeb, 0xbb, 0x58, 0xad, 0x10, 0x0b, 0x03, 0x0b, 0x0b, 0x12, 0xe2, 0x1b,
	0xf0, 0x61, 0x3a, 0x56, 0x62, 0x61, 0x42, 0x28, 0xe1, 0x83, 0x20, 0xdf, 0x5d, 0x9c, 0xd8, 0x44,
	0xf1, 0x66, 0xbf, 0xf7, 0x7e, 0xff, 0xde, 0xb3, 0xc1, 0xc5, 0x74, 0x2c, 0x24, 0xc3, 0x64, 0xd4,
	0xbb, 0xcf, 0xd2, 0x2e, 0x3b, 0x9d, 0x60, 0x72, 0x4e, 0xe3, 0x44, 0x28, 0x41, 0xae, 0xea, 0x1e,
	0xd5, 0x3d, 0x9a, 0x76, 0xdd, 0xc3, 0x91, 0x90, 0xd9, 0xf0, 0x90, 0x4b, 0x34, 0x83, 0x2c, 0xed,
	0x0e, 0x51, 0xf1, 0x2e, 0x8b, 0x79, 0x10, 0x46, 0x5c, 0x85, 0x22, 0x32, 0x58, 0xb7, 0xcc, 0x6b,
	0x48, 0x4c, 0x6f, 0xb7, 0xd4, 0x0b, 0x30, 0x42, 0x19, 0x4a, 0xdb, 0x6d, 0x05, 0x22, 0x10, 0xfa,
	0x91, 0x65, 0x4f, 0x73, 0x4c, 0x20, 0x44, 0x70, 0x82, 0x8c, 0xc7, 0x21, 0xe3, 0x51, 0x24, 0x94,
	0x16, 0xb3, 0x18, 0xff, 0x2d, 0x6c, 0xbf, 0xc8, 0xfc, 0xbc, 0x14, 0xef, 0x31, 0xea, 0xf3, 0x30,
	0x91, 0x03, 0x3c, 0x9d, 0xa0, 0x54, 0xe4, 0x29, 0xc0, 0xc2, 0xdb, 0x8e, 0xb3, 0xef, 0x1c, 0x34,
	0x7b, 0x6d, 0x6a, 0x82, 0xd0, 0x2c, 0x08, 0x35, 0x89, 0x6d, 0x10, 0xda, 0xe7, 0x01, 0x5a, 0xec,
	0x60, 0x09, 0xe9, 0xff, 0x70, 0xe0, 0xc6, 0x7f, 0x12, 0x32, 0x16, 0x91, 0x44, 0xf2, 0x18, 0x9a,
	0x2a, 0xab, 0xbe, 0x89, 0xb3, 0xf2, 0x8e, 0xb3, 0xbf, 0x79, 0xd0, 0xec, 0xdd, 0xa4, 0xc5, 0xed,
	0xd1, 0x1c, 0x78, 0x74, 0xe5, 0xe2, 0xf7, 0x5e, 0x6d, 0x00, 0x2a, 0x67, 0x22, 0xcf, 0x0a, 0x2e,
	0x37, 0xb4, 0xcb, 0x4e, 0xa5, 0x4b, 0x23, 0x5f, 0xb0, 0x79, 0x0f, 0xae, 0x17, 0x5d, 0xce, 0xf7,
	0xd0, 0x82, 0x2d, 0xad, 0xa7, 0x57, 0xd0, 0x18, 0x98, 0x17, 0xff, 0x55, 0x79, 0x6f, 0x79, 0xa6,
	0x47, 0x00, 0x8b, 0x4c, 0x76, 0x6f, 0x95, 0x91, 0x1a, 0x79, 0x24, 0xbf, 0x05, 0x44, 0x33, 0xf7,
	0x79, 0xc2, 0xc7, 0xf3, 0x6b, 0xf8, 0xcf, 0xe1, 0x5a, 0xa1, 0x6a, 0xc5, 0x1e, 0x42, 0x3d, 0xd6,
	0x15, 0x2b, 0xb4, 0x5d, 0x16, 0x32, 0xf3, 0x56, 0xc5, 0xce, 0xf6, 0xbe, 0x6f, 0xc2, 0x96, 0x66,
	0x23, 0x9f, 0x1d, 0x80, 0xc5, 0x5d, 0x48, 0xbb, 0x0c, 0x5f, 0xfd, 0x6d, 0xb8, 0x9d, 0xca, 0x39,
	0xe3, 0xcf, 0xef, 0x7c, 0xfa, 0xf9, 0xf7, 0xdb, 0xc6, 0x2d, 0xb2, 0xc7, 0x50, 0x1d, 0x63, 0x32,
	0x0e, 0x23, 0xb5, 0xf8, 0x7a, 0x97, 0x4e, 0x4f, 0xbe, 0x38, 0xd0, 0xc8, 0xf1, 0xe4, 0xce, 0x7a,
	0xfe, 0xb9, 0x8d, 0x76, 0xd5, 0x98, 0x75, 0xc1, 0xb4, 0x8b, 0xbb, 0xa4, 0x53, 0xe1, 0x82, 0x7d,
	0xd0, 0x2f, 0x1f, 0xc9, 0x04, 0xea, 0x66, 0x71, 0xc4, 0x5f, 0x29, 0x51, 0xb8, 0x8d, 0x7b, 0x7b,
	0xed, 0x8c, 0xf5, 0xe0, 0x6b, 0x0f, 0xbb, 0xc4, 0x5d, 0xe5, 0xc1, 0xdc, 0xe5, 0xe8, 0xc9, 0xc5,
	0xd4, 0x73, 0x2e, 0xa7, 0x9e, 0xf3, 0x67, 0xea, 0x39, 0x5f, 0x67, 0x5e, 0xed, 0x72, 0xe6, 0xd5,
	0x7e, 0xcd, 0xbc, 0xda, 0xeb, 0xc3, 0x20, 0x54, 0xc7, 0x93, 0x21, 0x1d, 0x89, 0x31, 0x3b, 0x8b,
	0x4f, 0xf8, 0x3b, 0x4c, 0x97, 0x78, 0xce, 0x2c, 0x93, 0x3a, 0x8f, 0x51, 0x0e, 0xeb, 0xfa, 0xcf,
	0x7e, 0xf0, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xd4, 0x41, 0xad, 0xde, 0xa1, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// TokenPairs retrieves registered token pairs
	TokenPairs(ctx context.Context, in *QueryTokenPairsRequest, opts ...grpc.CallOption) (*QueryTokenPairsResponse, error)
	// TokenPair retrieves a registered token pair
	TokenPair(ctx context.Context, in *QueryTokenPairRequest, opts ...grpc.CallOption) (*QueryTokenPairResponse, error)
	// Params retrieves the erc20 module params
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) TokenPairs(ctx context.Context, in *QueryTokenPairsRequest, opts ...grpc.CallOption) (*QueryTokenPairsResponse, error) {
	out := new(QueryTokenPairsResponse)
	err := c.cc.Invoke(ctx, "/evmos.erc20.v1.Query/TokenPairs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) TokenPair(ctx context.Context, in *QueryTokenPairRequest, opts ...grpc.CallOption) (*QueryTokenPairResponse, error) {
	out := new(QueryTokenPairResponse)
	err := c.cc.Invoke(ctx, "/evmos.erc20.v1.Query/TokenPair", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/evmos.erc20.v1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// TokenPairs retrieves registered token pairs
	TokenPairs(context.Context, *QueryTokenPairsRequest) (*QueryTokenPairsResponse, error)
	// TokenPair retrieves a registered token pair
	TokenPair(context.Context, *QueryTokenPairRequest) (*QueryTokenPairResponse, error)
	// Params retrieves the erc20 module params
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) TokenPairs(ctx context.Context, req *QueryTokenPairsRequest) (*QueryTokenPairsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TokenPairs not implemented")
}
func (*UnimplementedQueryServer) TokenPair(ctx context.Context, req *QueryTokenPairRequest) (*QueryTokenPairResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TokenPair not implemented")
}
func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_TokenPairs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTokenPairsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).TokenPairs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/evmos.erc20.v1.Query/TokenPairs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).TokenPairs(ctx, req.(*QueryTokenPairsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_TokenPair_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTokenPairRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).TokenPair(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/evmos.erc20.v1.Query/TokenPair",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).TokenPair(ctx, req.(*QueryTokenPairRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/evmos.erc20.v1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "evmos.erc20.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TokenPairs",
			Handler:    _Query_TokenPairs_Handler,
		},
		{
			MethodName: "TokenPair",
			Handler:    _Query_TokenPair_Handler,
		},
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "evmos/erc20/v1/query.proto",
}

func (m *QueryTokenPairsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryTokenPairsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryTokenPairsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryTokenPairsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryTokenPairsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryTokenPairsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.TokenPairs) > 0 {
		for iNdEx := len(m.TokenPairs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TokenPairs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryTokenPairRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryTokenPairRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryTokenPairRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Token) > 0 {
		i -= len(m.Token)
		copy(dAtA[i:], m.Token)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Token)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryTokenPairResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryTokenPairResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryTokenPairResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.TokenPair.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryTokenPairsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryTokenPairsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.TokenPairs) > 0 {
		for _, e := range m.TokenPairs {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryTokenPairRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryTokenPairResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.TokenPair.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryTokenPairsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryTokenPairsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryTokenPairsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryTokenPairsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryTokenPairsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryTokenPairsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenPairs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenPairs = append(m.TokenPairs, TokenPair{})
			if err := m.TokenPairs[len(m.TokenPairs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryTokenPairRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryTokenPairRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryTokenPairRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryTokenPairResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryTokenPairResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryTokenPairResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenPair", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TokenPair.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
