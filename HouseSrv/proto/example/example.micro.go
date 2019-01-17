// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/example/example.proto

/*
Package go_micro_srv_HouseSrv is a generated protocol buffer package.

It is generated from these files:
	proto/example/example.proto

It has these top-level messages:
	Message
	UserHousesRequest
	UserHousesResponse
	PostHousesRequest
	PostHousesResponse
	HousesImageRequest
	HousesImageResponse
	HouseInfoRequest
	HouseInfoResponse
	SearchHousesRequest
	SearchHousesResponse
*/
package go_micro_srv_HouseSrv

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
	// 获取当前用户发布的房源信息
	GetUserHouses(ctx context.Context, in *UserHousesRequest, opts ...client.CallOption) (*UserHousesResponse, error)
	// 发布房源
	PostHouses(ctx context.Context, in *PostHousesRequest, opts ...client.CallOption) (*PostHousesResponse, error)
	// 上传房屋图片
	PostHousesImage(ctx context.Context, in *HousesImageRequest, opts ...client.CallOption) (*HousesImageResponse, error)
	// 获取房源详细信息
	GetHouseInfo(ctx context.Context, in *HouseInfoRequest, opts ...client.CallOption) (*HouseInfoResponse, error)
	// 搜索房源
	GetSearchHouses(ctx context.Context, in *SearchHousesRequest, opts ...client.CallOption) (*SearchHousesResponse, error)
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
		name = "go.micro.srv.HouseSrv"
	}
	return &exampleService{
		c:    c,
		name: name,
	}
}

func (c *exampleService) GetUserHouses(ctx context.Context, in *UserHousesRequest, opts ...client.CallOption) (*UserHousesResponse, error) {
	req := c.c.NewRequest(c.name, "Example.GetUserHouses", in)
	out := new(UserHousesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleService) PostHouses(ctx context.Context, in *PostHousesRequest, opts ...client.CallOption) (*PostHousesResponse, error) {
	req := c.c.NewRequest(c.name, "Example.PostHouses", in)
	out := new(PostHousesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleService) PostHousesImage(ctx context.Context, in *HousesImageRequest, opts ...client.CallOption) (*HousesImageResponse, error) {
	req := c.c.NewRequest(c.name, "Example.PostHousesImage", in)
	out := new(HousesImageResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleService) GetHouseInfo(ctx context.Context, in *HouseInfoRequest, opts ...client.CallOption) (*HouseInfoResponse, error) {
	req := c.c.NewRequest(c.name, "Example.GetHouseInfo", in)
	out := new(HouseInfoResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleService) GetSearchHouses(ctx context.Context, in *SearchHousesRequest, opts ...client.CallOption) (*SearchHousesResponse, error) {
	req := c.c.NewRequest(c.name, "Example.GetSearchHouses", in)
	out := new(SearchHousesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Example service

type ExampleHandler interface {
	// 获取当前用户发布的房源信息
	GetUserHouses(context.Context, *UserHousesRequest, *UserHousesResponse) error
	// 发布房源
	PostHouses(context.Context, *PostHousesRequest, *PostHousesResponse) error
	// 上传房屋图片
	PostHousesImage(context.Context, *HousesImageRequest, *HousesImageResponse) error
	// 获取房源详细信息
	GetHouseInfo(context.Context, *HouseInfoRequest, *HouseInfoResponse) error
	// 搜索房源
	GetSearchHouses(context.Context, *SearchHousesRequest, *SearchHousesResponse) error
}

func RegisterExampleHandler(s server.Server, hdlr ExampleHandler, opts ...server.HandlerOption) error {
	type example interface {
		GetUserHouses(ctx context.Context, in *UserHousesRequest, out *UserHousesResponse) error
		PostHouses(ctx context.Context, in *PostHousesRequest, out *PostHousesResponse) error
		PostHousesImage(ctx context.Context, in *HousesImageRequest, out *HousesImageResponse) error
		GetHouseInfo(ctx context.Context, in *HouseInfoRequest, out *HouseInfoResponse) error
		GetSearchHouses(ctx context.Context, in *SearchHousesRequest, out *SearchHousesResponse) error
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

func (h *exampleHandler) GetUserHouses(ctx context.Context, in *UserHousesRequest, out *UserHousesResponse) error {
	return h.ExampleHandler.GetUserHouses(ctx, in, out)
}

func (h *exampleHandler) PostHouses(ctx context.Context, in *PostHousesRequest, out *PostHousesResponse) error {
	return h.ExampleHandler.PostHouses(ctx, in, out)
}

func (h *exampleHandler) PostHousesImage(ctx context.Context, in *HousesImageRequest, out *HousesImageResponse) error {
	return h.ExampleHandler.PostHousesImage(ctx, in, out)
}

func (h *exampleHandler) GetHouseInfo(ctx context.Context, in *HouseInfoRequest, out *HouseInfoResponse) error {
	return h.ExampleHandler.GetHouseInfo(ctx, in, out)
}

func (h *exampleHandler) GetSearchHouses(ctx context.Context, in *SearchHousesRequest, out *SearchHousesResponse) error {
	return h.ExampleHandler.GetSearchHouses(ctx, in, out)
}
