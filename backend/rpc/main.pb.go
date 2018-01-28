// Code generated by protoc-gen-go. DO NOT EDIT.
// source: main.proto

/*
Package rpc is a generated protocol buffer package.

It is generated from these files:
	main.proto

It has these top-level messages:
	TerminalStateRequest
	TerminalStateResponse
	TerminalBuyRequest
	TerminalBuyResponse
	TerminalAddDepositOrderRequest
	TerminalAddDepositOrderResponse
	TerminalScanRequest
	TerminalScanResponse
	AbortRequest
	AbortResponse
*/
package rpc

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

type TerminalStateRequest struct {
	TerminalID string `protobuf:"bytes,1,opt,name=TerminalID" json:"TerminalID,omitempty"`
}

func (m *TerminalStateRequest) Reset()                    { *m = TerminalStateRequest{} }
func (m *TerminalStateRequest) String() string            { return proto.CompactTextString(m) }
func (*TerminalStateRequest) ProtoMessage()               {}
func (*TerminalStateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TerminalStateRequest) GetTerminalID() string {
	if m != nil {
		return m.TerminalID
	}
	return ""
}

type TerminalStateResponse struct {
	Accounts           []*TerminalStateResponse_Account `protobuf:"bytes,1,rep,name=Accounts" json:"Accounts,omitempty"`
	PendingOrders      []*TerminalStateResponse_Order   `protobuf:"bytes,2,rep,name=PendingOrders" json:"PendingOrders,omitempty"`
	PendingOrdersTotal int32                            `protobuf:"zigzag32,3,opt,name=PendingOrdersTotal" json:"PendingOrdersTotal,omitempty"`
}

func (m *TerminalStateResponse) Reset()                    { *m = TerminalStateResponse{} }
func (m *TerminalStateResponse) String() string            { return proto.CompactTextString(m) }
func (*TerminalStateResponse) ProtoMessage()               {}
func (*TerminalStateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TerminalStateResponse) GetAccounts() []*TerminalStateResponse_Account {
	if m != nil {
		return m.Accounts
	}
	return nil
}

func (m *TerminalStateResponse) GetPendingOrders() []*TerminalStateResponse_Order {
	if m != nil {
		return m.PendingOrders
	}
	return nil
}

func (m *TerminalStateResponse) GetPendingOrdersTotal() int32 {
	if m != nil {
		return m.PendingOrdersTotal
	}
	return 0
}

type TerminalStateResponse_Account struct {
	ID          string `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	DisplayName string `protobuf:"bytes,2,opt,name=DisplayName" json:"DisplayName,omitempty"`
	Balance     int32  `protobuf:"zigzag32,3,opt,name=Balance" json:"Balance,omitempty"`
}

func (m *TerminalStateResponse_Account) Reset()         { *m = TerminalStateResponse_Account{} }
func (m *TerminalStateResponse_Account) String() string { return proto.CompactTextString(m) }
func (*TerminalStateResponse_Account) ProtoMessage()    {}
func (*TerminalStateResponse_Account) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{1, 0}
}

func (m *TerminalStateResponse_Account) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *TerminalStateResponse_Account) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *TerminalStateResponse_Account) GetBalance() int32 {
	if m != nil {
		return m.Balance
	}
	return 0
}

type TerminalStateResponse_Order struct {
	DisplayName string `protobuf:"bytes,1,opt,name=DisplayName" json:"DisplayName,omitempty"`
	Price       int32  `protobuf:"zigzag32,2,opt,name=Price" json:"Price,omitempty"`
	NeedsReview bool   `protobuf:"varint,3,opt,name=NeedsReview" json:"NeedsReview,omitempty"`
}

func (m *TerminalStateResponse_Order) Reset()                    { *m = TerminalStateResponse_Order{} }
func (m *TerminalStateResponse_Order) String() string            { return proto.CompactTextString(m) }
func (*TerminalStateResponse_Order) ProtoMessage()               {}
func (*TerminalStateResponse_Order) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 1} }

func (m *TerminalStateResponse_Order) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *TerminalStateResponse_Order) GetPrice() int32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *TerminalStateResponse_Order) GetNeedsReview() bool {
	if m != nil {
		return m.NeedsReview
	}
	return false
}

type TerminalBuyRequest struct {
	TerminalID string `protobuf:"bytes,1,opt,name=TerminalID" json:"TerminalID,omitempty"`
	AccountID  string `protobuf:"bytes,2,opt,name=AccountID" json:"AccountID,omitempty"`
}

func (m *TerminalBuyRequest) Reset()                    { *m = TerminalBuyRequest{} }
func (m *TerminalBuyRequest) String() string            { return proto.CompactTextString(m) }
func (*TerminalBuyRequest) ProtoMessage()               {}
func (*TerminalBuyRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *TerminalBuyRequest) GetTerminalID() string {
	if m != nil {
		return m.TerminalID
	}
	return ""
}

func (m *TerminalBuyRequest) GetAccountID() string {
	if m != nil {
		return m.AccountID
	}
	return ""
}

type TerminalBuyResponse struct {
	Error string `protobuf:"bytes,1,opt,name=Error" json:"Error,omitempty"`
}

func (m *TerminalBuyResponse) Reset()                    { *m = TerminalBuyResponse{} }
func (m *TerminalBuyResponse) String() string            { return proto.CompactTextString(m) }
func (*TerminalBuyResponse) ProtoMessage()               {}
func (*TerminalBuyResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *TerminalBuyResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type TerminalAddDepositOrderRequest struct {
	TerminalID   string `protobuf:"bytes,1,opt,name=TerminalID" json:"TerminalID,omitempty"`
	CashInAmount int32  `protobuf:"zigzag32,2,opt,name=CashInAmount" json:"CashInAmount,omitempty"`
}

func (m *TerminalAddDepositOrderRequest) Reset()                    { *m = TerminalAddDepositOrderRequest{} }
func (m *TerminalAddDepositOrderRequest) String() string            { return proto.CompactTextString(m) }
func (*TerminalAddDepositOrderRequest) ProtoMessage()               {}
func (*TerminalAddDepositOrderRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *TerminalAddDepositOrderRequest) GetTerminalID() string {
	if m != nil {
		return m.TerminalID
	}
	return ""
}

func (m *TerminalAddDepositOrderRequest) GetCashInAmount() int32 {
	if m != nil {
		return m.CashInAmount
	}
	return 0
}

type TerminalAddDepositOrderResponse struct {
	Error string `protobuf:"bytes,1,opt,name=Error" json:"Error,omitempty"`
}

func (m *TerminalAddDepositOrderResponse) Reset()                    { *m = TerminalAddDepositOrderResponse{} }
func (m *TerminalAddDepositOrderResponse) String() string            { return proto.CompactTextString(m) }
func (*TerminalAddDepositOrderResponse) ProtoMessage()               {}
func (*TerminalAddDepositOrderResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *TerminalAddDepositOrderResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type TerminalScanRequest struct {
	TerminalID string `protobuf:"bytes,1,opt,name=TerminalID" json:"TerminalID,omitempty"`
	ProductID  string `protobuf:"bytes,2,opt,name=ProductID" json:"ProductID,omitempty"`
}

func (m *TerminalScanRequest) Reset()                    { *m = TerminalScanRequest{} }
func (m *TerminalScanRequest) String() string            { return proto.CompactTextString(m) }
func (*TerminalScanRequest) ProtoMessage()               {}
func (*TerminalScanRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *TerminalScanRequest) GetTerminalID() string {
	if m != nil {
		return m.TerminalID
	}
	return ""
}

func (m *TerminalScanRequest) GetProductID() string {
	if m != nil {
		return m.ProductID
	}
	return ""
}

type TerminalScanResponse struct {
	Error string `protobuf:"bytes,1,opt,name=Error" json:"Error,omitempty"`
}

func (m *TerminalScanResponse) Reset()                    { *m = TerminalScanResponse{} }
func (m *TerminalScanResponse) String() string            { return proto.CompactTextString(m) }
func (*TerminalScanResponse) ProtoMessage()               {}
func (*TerminalScanResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *TerminalScanResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type AbortRequest struct {
	TerminalID string `protobuf:"bytes,1,opt,name=TerminalID" json:"TerminalID,omitempty"`
}

func (m *AbortRequest) Reset()                    { *m = AbortRequest{} }
func (m *AbortRequest) String() string            { return proto.CompactTextString(m) }
func (*AbortRequest) ProtoMessage()               {}
func (*AbortRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *AbortRequest) GetTerminalID() string {
	if m != nil {
		return m.TerminalID
	}
	return ""
}

type AbortResponse struct {
	Error string `protobuf:"bytes,1,opt,name=Error" json:"Error,omitempty"`
}

func (m *AbortResponse) Reset()                    { *m = AbortResponse{} }
func (m *AbortResponse) String() string            { return proto.CompactTextString(m) }
func (*AbortResponse) ProtoMessage()               {}
func (*AbortResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *AbortResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*TerminalStateRequest)(nil), "i6getraenkeabrechnungssystem3000.rpc.TerminalStateRequest")
	proto.RegisterType((*TerminalStateResponse)(nil), "i6getraenkeabrechnungssystem3000.rpc.TerminalStateResponse")
	proto.RegisterType((*TerminalStateResponse_Account)(nil), "i6getraenkeabrechnungssystem3000.rpc.TerminalStateResponse.Account")
	proto.RegisterType((*TerminalStateResponse_Order)(nil), "i6getraenkeabrechnungssystem3000.rpc.TerminalStateResponse.Order")
	proto.RegisterType((*TerminalBuyRequest)(nil), "i6getraenkeabrechnungssystem3000.rpc.TerminalBuyRequest")
	proto.RegisterType((*TerminalBuyResponse)(nil), "i6getraenkeabrechnungssystem3000.rpc.TerminalBuyResponse")
	proto.RegisterType((*TerminalAddDepositOrderRequest)(nil), "i6getraenkeabrechnungssystem3000.rpc.TerminalAddDepositOrderRequest")
	proto.RegisterType((*TerminalAddDepositOrderResponse)(nil), "i6getraenkeabrechnungssystem3000.rpc.TerminalAddDepositOrderResponse")
	proto.RegisterType((*TerminalScanRequest)(nil), "i6getraenkeabrechnungssystem3000.rpc.TerminalScanRequest")
	proto.RegisterType((*TerminalScanResponse)(nil), "i6getraenkeabrechnungssystem3000.rpc.TerminalScanResponse")
	proto.RegisterType((*AbortRequest)(nil), "i6getraenkeabrechnungssystem3000.rpc.AbortRequest")
	proto.RegisterType((*AbortResponse)(nil), "i6getraenkeabrechnungssystem3000.rpc.AbortResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TerminalBackend service

type TerminalBackendClient interface {
	GetState(ctx context.Context, in *TerminalStateRequest, opts ...grpc.CallOption) (*TerminalStateResponse, error)
	Buy(ctx context.Context, in *TerminalBuyRequest, opts ...grpc.CallOption) (*TerminalBuyResponse, error)
	AddDepositOrder(ctx context.Context, in *TerminalAddDepositOrderRequest, opts ...grpc.CallOption) (*TerminalAddDepositOrderResponse, error)
	Scan(ctx context.Context, in *TerminalScanRequest, opts ...grpc.CallOption) (*TerminalScanResponse, error)
	Abort(ctx context.Context, in *AbortRequest, opts ...grpc.CallOption) (*AbortResponse, error)
}

type terminalBackendClient struct {
	cc *grpc.ClientConn
}

func NewTerminalBackendClient(cc *grpc.ClientConn) TerminalBackendClient {
	return &terminalBackendClient{cc}
}

func (c *terminalBackendClient) GetState(ctx context.Context, in *TerminalStateRequest, opts ...grpc.CallOption) (*TerminalStateResponse, error) {
	out := new(TerminalStateResponse)
	err := grpc.Invoke(ctx, "/i6getraenkeabrechnungssystem3000.rpc.TerminalBackend/GetState", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *terminalBackendClient) Buy(ctx context.Context, in *TerminalBuyRequest, opts ...grpc.CallOption) (*TerminalBuyResponse, error) {
	out := new(TerminalBuyResponse)
	err := grpc.Invoke(ctx, "/i6getraenkeabrechnungssystem3000.rpc.TerminalBackend/Buy", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *terminalBackendClient) AddDepositOrder(ctx context.Context, in *TerminalAddDepositOrderRequest, opts ...grpc.CallOption) (*TerminalAddDepositOrderResponse, error) {
	out := new(TerminalAddDepositOrderResponse)
	err := grpc.Invoke(ctx, "/i6getraenkeabrechnungssystem3000.rpc.TerminalBackend/AddDepositOrder", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *terminalBackendClient) Scan(ctx context.Context, in *TerminalScanRequest, opts ...grpc.CallOption) (*TerminalScanResponse, error) {
	out := new(TerminalScanResponse)
	err := grpc.Invoke(ctx, "/i6getraenkeabrechnungssystem3000.rpc.TerminalBackend/Scan", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *terminalBackendClient) Abort(ctx context.Context, in *AbortRequest, opts ...grpc.CallOption) (*AbortResponse, error) {
	out := new(AbortResponse)
	err := grpc.Invoke(ctx, "/i6getraenkeabrechnungssystem3000.rpc.TerminalBackend/Abort", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TerminalBackend service

type TerminalBackendServer interface {
	GetState(context.Context, *TerminalStateRequest) (*TerminalStateResponse, error)
	Buy(context.Context, *TerminalBuyRequest) (*TerminalBuyResponse, error)
	AddDepositOrder(context.Context, *TerminalAddDepositOrderRequest) (*TerminalAddDepositOrderResponse, error)
	Scan(context.Context, *TerminalScanRequest) (*TerminalScanResponse, error)
	Abort(context.Context, *AbortRequest) (*AbortResponse, error)
}

func RegisterTerminalBackendServer(s *grpc.Server, srv TerminalBackendServer) {
	s.RegisterService(&_TerminalBackend_serviceDesc, srv)
}

func _TerminalBackend_GetState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TerminalStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TerminalBackendServer).GetState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/i6getraenkeabrechnungssystem3000.rpc.TerminalBackend/GetState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TerminalBackendServer).GetState(ctx, req.(*TerminalStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TerminalBackend_Buy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TerminalBuyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TerminalBackendServer).Buy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/i6getraenkeabrechnungssystem3000.rpc.TerminalBackend/Buy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TerminalBackendServer).Buy(ctx, req.(*TerminalBuyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TerminalBackend_AddDepositOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TerminalAddDepositOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TerminalBackendServer).AddDepositOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/i6getraenkeabrechnungssystem3000.rpc.TerminalBackend/AddDepositOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TerminalBackendServer).AddDepositOrder(ctx, req.(*TerminalAddDepositOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TerminalBackend_Scan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TerminalScanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TerminalBackendServer).Scan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/i6getraenkeabrechnungssystem3000.rpc.TerminalBackend/Scan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TerminalBackendServer).Scan(ctx, req.(*TerminalScanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TerminalBackend_Abort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AbortRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TerminalBackendServer).Abort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/i6getraenkeabrechnungssystem3000.rpc.TerminalBackend/Abort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TerminalBackendServer).Abort(ctx, req.(*AbortRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TerminalBackend_serviceDesc = grpc.ServiceDesc{
	ServiceName: "i6getraenkeabrechnungssystem3000.rpc.TerminalBackend",
	HandlerType: (*TerminalBackendServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetState",
			Handler:    _TerminalBackend_GetState_Handler,
		},
		{
			MethodName: "Buy",
			Handler:    _TerminalBackend_Buy_Handler,
		},
		{
			MethodName: "AddDepositOrder",
			Handler:    _TerminalBackend_AddDepositOrder_Handler,
		},
		{
			MethodName: "Scan",
			Handler:    _TerminalBackend_Scan_Handler,
		},
		{
			MethodName: "Abort",
			Handler:    _TerminalBackend_Abort_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "main.proto",
}

func init() { proto.RegisterFile("main.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 544 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x95, 0xd1, 0x8a, 0xd3, 0x4e,
	0x14, 0xc6, 0xff, 0x69, 0xff, 0x71, 0xbb, 0x67, 0x77, 0x5d, 0x76, 0xac, 0x10, 0x82, 0xac, 0x25,
	0x28, 0x14, 0x94, 0x50, 0x5a, 0x58, 0x75, 0xbd, 0x6a, 0x37, 0x8b, 0xf4, 0x66, 0x2d, 0xe9, 0x7a,
	0xe3, 0x8d, 0x4c, 0x27, 0x87, 0x6e, 0xd8, 0x76, 0x12, 0x67, 0x26, 0x4a, 0x41, 0xf0, 0xca, 0xf7,
	0xf0, 0x0d, 0x7c, 0x10, 0x5f, 0x4a, 0x3a, 0x4d, 0xd3, 0xb4, 0xd6, 0xd2, 0xb4, 0x97, 0x73, 0x92,
	0xf3, 0xfb, 0xbe, 0x39, 0xf9, 0x0e, 0x01, 0x18, 0xd3, 0x90, 0xbb, 0xb1, 0x88, 0x54, 0x44, 0x9e,
	0x85, 0x17, 0x43, 0x54, 0x82, 0x22, 0xbf, 0x47, 0x3a, 0x10, 0xc8, 0xee, 0x78, 0xc2, 0x87, 0x52,
	0x4e, 0xa4, 0xc2, 0x71, 0xab, 0xd1, 0x68, 0xb8, 0x22, 0x66, 0xce, 0x05, 0x54, 0x6f, 0x51, 0x8c,
	0x43, 0x4e, 0x47, 0x7d, 0x45, 0x15, 0xfa, 0xf8, 0x39, 0x41, 0xa9, 0xc8, 0x39, 0xc0, 0xbc, 0xde,
	0xf5, 0x2c, 0xa3, 0x66, 0xd4, 0x0f, 0xfd, 0x5c, 0xc5, 0xf9, 0x5d, 0x86, 0xc7, 0x2b, 0x8d, 0x32,
	0x8e, 0xb8, 0x44, 0xf2, 0x09, 0x2a, 0x6d, 0xc6, 0xa2, 0x84, 0x2b, 0x69, 0x19, 0xb5, 0x72, 0xfd,
	0xa8, 0x79, 0xe5, 0x6e, 0x63, 0xc5, 0x5d, 0x8b, 0x73, 0x53, 0x96, 0x9f, 0x41, 0xc9, 0x10, 0x4e,
	0x7a, 0xc8, 0x83, 0x90, 0x0f, 0xdf, 0x8b, 0x00, 0x85, 0xb4, 0x4a, 0x5a, 0xa5, 0xbd, 0x8f, 0x8a,
	0x26, 0xf9, 0xcb, 0x5c, 0xe2, 0x02, 0x59, 0x2a, 0xdc, 0x46, 0x8a, 0x8e, 0xac, 0x72, 0xcd, 0xa8,
	0x9f, 0xf9, 0x6b, 0x9e, 0xd8, 0x1f, 0xe0, 0x20, 0x35, 0x49, 0x1e, 0x42, 0x29, 0x1b, 0x5b, 0xa9,
	0xeb, 0x91, 0x1a, 0x1c, 0x79, 0xa1, 0x8c, 0x47, 0x74, 0x72, 0x43, 0xc7, 0x68, 0x95, 0xf4, 0x83,
	0x7c, 0x89, 0x58, 0x70, 0xd0, 0xa1, 0x23, 0xca, 0x19, 0xa6, 0x0a, 0xf3, 0xa3, 0x4d, 0xc1, 0xd4,
	0x2a, 0xab, 0x10, 0xe3, 0x6f, 0x48, 0x15, 0xcc, 0x9e, 0x08, 0xd9, 0x4c, 0xe0, 0xcc, 0x9f, 0x1d,
	0xa6, 0x7d, 0x37, 0x88, 0x81, 0xf4, 0xf1, 0x4b, 0x88, 0x5f, 0x35, 0xbe, 0xe2, 0xe7, 0x4b, 0x8e,
	0x0f, 0x64, 0x3e, 0x97, 0x4e, 0x32, 0xd9, 0x32, 0x03, 0xe4, 0x09, 0x1c, 0xa6, 0xf7, 0xed, 0x7a,
	0xe9, 0x95, 0x16, 0x05, 0xe7, 0x05, 0x3c, 0x5a, 0x62, 0xa6, 0xf1, 0xa8, 0x82, 0x79, 0x2d, 0x44,
	0x24, 0x52, 0xde, 0xec, 0xe0, 0x04, 0x70, 0x3e, 0x7f, 0xb9, 0x1d, 0x04, 0x1e, 0xc6, 0x91, 0x0c,
	0xd5, 0xec, 0xa3, 0x6c, 0x69, 0xc6, 0x81, 0xe3, 0x2b, 0x2a, 0xef, 0xba, 0xbc, 0x3d, 0x9e, 0x1a,
	0x48, 0x27, 0xb0, 0x54, 0x73, 0x5e, 0xc1, 0xd3, 0x7f, 0xaa, 0x6c, 0xb4, 0xd7, 0x5f, 0xdc, 0xa5,
	0xcf, 0x28, 0x2f, 0x30, 0xa0, 0x9e, 0x88, 0x82, 0x84, 0xe5, 0x06, 0x94, 0x15, 0x9c, 0x97, 0xb9,
	0xd5, 0xd3, 0xd0, 0x8d, 0x16, 0x5c, 0x38, 0x6e, 0x0f, 0x22, 0xa1, 0xb6, 0x5d, 0xd0, 0xe7, 0x70,
	0x92, 0xbe, 0xbf, 0x09, 0xdb, 0xfc, 0x65, 0xc2, 0x69, 0xf6, 0x99, 0x28, 0xbb, 0x47, 0x1e, 0x90,
	0x1f, 0x06, 0x54, 0xde, 0xa1, 0xd2, 0x1b, 0x42, 0x2e, 0x77, 0x5a, 0x2b, 0xed, 0xd1, 0x7e, 0xbb,
	0xc7, 0x4a, 0x3a, 0xff, 0x91, 0x6f, 0x50, 0xee, 0x24, 0x13, 0xf2, 0xba, 0x18, 0x65, 0x11, 0x60,
	0xfb, 0xcd, 0x0e, 0x9d, 0x99, 0xfa, 0x4f, 0x03, 0x4e, 0x57, 0x52, 0x42, 0xbc, 0x62, 0xc0, 0xf5,
	0x51, 0xb6, 0xaf, 0xf7, 0xa4, 0x64, 0x16, 0xbf, 0xc3, 0xff, 0xd3, 0xe4, 0x90, 0x82, 0xf7, 0xcc,
	0x45, 0xd8, 0xbe, 0xdc, 0xa5, 0x35, 0x33, 0x20, 0xc0, 0xd4, 0x21, 0x23, 0xcd, 0xed, 0x30, 0xf9,
	0x04, 0xdb, 0xad, 0x42, 0x3d, 0x73, 0xcd, 0x8e, 0xf9, 0xb1, 0x2c, 0x62, 0x36, 0x78, 0xa0, 0xff,
	0x72, 0xad, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x08, 0xc2, 0x04, 0x7b, 0xf3, 0x06, 0x00, 0x00,
}
