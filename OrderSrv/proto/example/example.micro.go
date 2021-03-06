// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/example/example.proto

/*
Package go_micro_srv_OrderSrv is a generated protocol buffer package.

It is generated from these files:
	proto/example/example.proto

It has these top-level messages:
	Message
	PostOrdersRequest
	PostOrdersResponse
	UserOrderRequest
	UserOrderResponse
	PutOrdersRequest
	PutOrdersResponse
	UserCommentRequest
	UserCommentResponse
*/
package go_micro_srv_OrderSrv

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Example service

type ExampleService interface {
	// 发布订单
	PostOrders(ctx context.Context, in *PostOrdersRequest, opts ...client.CallOption) (*PostOrdersResponse, error)
	// 查看房东/租客订单的Proto
	GetUserOrder(ctx context.Context, in *UserOrderRequest, opts ...client.CallOption) (*UserOrderResponse, error)
	// 房东同意/拒绝订单的
	PutOrders(ctx context.Context, in *PutOrdersRequest, opts ...client.CallOption) (*PutOrdersResponse, error)
	// 用户评价订单
	PutComment(ctx context.Context, in *UserCommentRequest, opts ...client.CallOption) (*UserCommentResponse, error)
}

type exampleService struct {
	c    client.Client
	name string
}

func NewExampleService(name string, c client.Client) ExampleService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.micro.srv.OrderSrv"
	}
	return &exampleService{
		c:    c,
		name: name,
	}
}

func (c *exampleService) PostOrders(ctx context.Context, in *PostOrdersRequest, opts ...client.CallOption) (*PostOrdersResponse, error) {
	req := c.c.NewRequest(c.name, "Example.PostOrders", in)
	out := new(PostOrdersResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleService) GetUserOrder(ctx context.Context, in *UserOrderRequest, opts ...client.CallOption) (*UserOrderResponse, error) {
	req := c.c.NewRequest(c.name, "Example.GetUserOrder", in)
	out := new(UserOrderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleService) PutOrders(ctx context.Context, in *PutOrdersRequest, opts ...client.CallOption) (*PutOrdersResponse, error) {
	req := c.c.NewRequest(c.name, "Example.PutOrders", in)
	out := new(PutOrdersResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleService) PutComment(ctx context.Context, in *UserCommentRequest, opts ...client.CallOption) (*UserCommentResponse, error) {
	req := c.c.NewRequest(c.name, "Example.PutComment", in)
	out := new(UserCommentResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Example service

type ExampleHandler interface {
	// 发布订单
	PostOrders(context.Context, *PostOrdersRequest, *PostOrdersResponse) error
	// 查看房东/租客订单的Proto
	GetUserOrder(context.Context, *UserOrderRequest, *UserOrderResponse) error
	// 房东同意/拒绝订单的
	PutOrders(context.Context, *PutOrdersRequest, *PutOrdersResponse) error
	// 用户评价订单
	PutComment(context.Context, *UserCommentRequest, *UserCommentResponse) error
}

func RegisterExampleHandler(s server.Server, hdlr ExampleHandler, opts ...server.HandlerOption) error {
	type example interface {
		PostOrders(ctx context.Context, in *PostOrdersRequest, out *PostOrdersResponse) error
		GetUserOrder(ctx context.Context, in *UserOrderRequest, out *UserOrderResponse) error
		PutOrders(ctx context.Context, in *PutOrdersRequest, out *PutOrdersResponse) error
		PutComment(ctx context.Context, in *UserCommentRequest, out *UserCommentResponse) error
	}
	type Example struct {
		example
	}
	h := &exampleHandler{hdlr}
	return s.Handle(s.NewHandler(&Example{h}, opts...))
}

type exampleHandler struct {
	ExampleHandler
}

func (h *exampleHandler) PostOrders(ctx context.Context, in *PostOrdersRequest, out *PostOrdersResponse) error {
	return h.ExampleHandler.PostOrders(ctx, in, out)
}

func (h *exampleHandler) GetUserOrder(ctx context.Context, in *UserOrderRequest, out *UserOrderResponse) error {
	return h.ExampleHandler.GetUserOrder(ctx, in, out)
}

func (h *exampleHandler) PutOrders(ctx context.Context, in *PutOrdersRequest, out *PutOrdersResponse) error {
	return h.ExampleHandler.PutOrders(ctx, in, out)
}

func (h *exampleHandler) PutComment(ctx context.Context, in *UserCommentRequest, out *UserCommentResponse) error {
	return h.ExampleHandler.PutComment(ctx, in, out)
}
