// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: service/account/proto/user/user_service.proto

package user

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/ygpark2/njro/service/account/proto/entities"
	math "math"
)

import (
	context "context"
	api "github.com/micro/micro/v3/service/api"
	client "github.com/micro/micro/v3/service/client"
	server "github.com/micro/micro/v3/service/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for UserService service

func NewUserServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for UserService service

type UserService interface {
	Exist(ctx context.Context, in *ExistRequest, opts ...client.CallOption) (*ExistResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error)
	Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) Exist(ctx context.Context, in *ExistRequest, opts ...client.CallOption) (*ExistResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.Exist", in)
	out := new(ExistResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.List", in)
	out := new(ListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Get(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.Get", in)
	out := new(GetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.Create", in)
	out := new(CreateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.Update", in)
	out := new(UpdateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.Delete", in)
	out := new(DeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceHandler interface {
	Exist(context.Context, *ExistRequest, *ExistResponse) error
	List(context.Context, *ListRequest, *ListResponse) error
	Get(context.Context, *GetRequest, *GetResponse) error
	Create(context.Context, *CreateRequest, *CreateResponse) error
	Update(context.Context, *UpdateRequest, *UpdateResponse) error
	Delete(context.Context, *DeleteRequest, *DeleteResponse) error
}

func RegisterUserServiceHandler(s server.Server, hdlr UserServiceHandler, opts ...server.HandlerOption) error {
	type userService interface {
		Exist(ctx context.Context, in *ExistRequest, out *ExistResponse) error
		List(ctx context.Context, in *ListRequest, out *ListResponse) error
		Get(ctx context.Context, in *GetRequest, out *GetResponse) error
		Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error
		Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error
		Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error
	}
	type UserService struct {
		userService
	}
	h := &userServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&UserService{h}, opts...))
}

type userServiceHandler struct {
	UserServiceHandler
}

func (h *userServiceHandler) Exist(ctx context.Context, in *ExistRequest, out *ExistResponse) error {
	return h.UserServiceHandler.Exist(ctx, in, out)
}

func (h *userServiceHandler) List(ctx context.Context, in *ListRequest, out *ListResponse) error {
	return h.UserServiceHandler.List(ctx, in, out)
}

func (h *userServiceHandler) Get(ctx context.Context, in *GetRequest, out *GetResponse) error {
	return h.UserServiceHandler.Get(ctx, in, out)
}

func (h *userServiceHandler) Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error {
	return h.UserServiceHandler.Create(ctx, in, out)
}

func (h *userServiceHandler) Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error {
	return h.UserServiceHandler.Update(ctx, in, out)
}

func (h *userServiceHandler) Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error {
	return h.UserServiceHandler.Delete(ctx, in, out)
}
