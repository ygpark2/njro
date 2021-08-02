// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: service/account/proto/profile/profile_service.proto

package profile

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/ygpark2/mboard/service/account/proto/entities"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/asim/go-micro/v3/api"
	client "github.com/asim/go-micro/v3/client"
	server "github.com/asim/go-micro/v3/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for ProfileService service

func NewProfileServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for ProfileService service

type ProfileService interface {
	List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error)
	Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error)
}

type profileService struct {
	c    client.Client
	name string
}

func NewProfileService(name string, c client.Client) ProfileService {
	return &profileService{
		c:    c,
		name: name,
	}
}

func (c *profileService) List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error) {
	req := c.c.NewRequest(c.name, "ProfileService.List", in)
	out := new(ListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileService) Get(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error) {
	req := c.c.NewRequest(c.name, "ProfileService.Get", in)
	out := new(GetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileService) Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error) {
	req := c.c.NewRequest(c.name, "ProfileService.Create", in)
	out := new(CreateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ProfileService service

type ProfileServiceHandler interface {
	List(context.Context, *ListRequest, *ListResponse) error
	Get(context.Context, *GetRequest, *GetResponse) error
	Create(context.Context, *CreateRequest, *CreateResponse) error
}

func RegisterProfileServiceHandler(s server.Server, hdlr ProfileServiceHandler, opts ...server.HandlerOption) error {
	type profileService interface {
		List(ctx context.Context, in *ListRequest, out *ListResponse) error
		Get(ctx context.Context, in *GetRequest, out *GetResponse) error
		Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error
	}
	type ProfileService struct {
		profileService
	}
	h := &profileServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&ProfileService{h}, opts...))
}

type profileServiceHandler struct {
	ProfileServiceHandler
}

func (h *profileServiceHandler) List(ctx context.Context, in *ListRequest, out *ListResponse) error {
	return h.ProfileServiceHandler.List(ctx, in, out)
}

func (h *profileServiceHandler) Get(ctx context.Context, in *GetRequest, out *GetResponse) error {
	return h.ProfileServiceHandler.Get(ctx, in, out)
}

func (h *profileServiceHandler) Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error {
	return h.ProfileServiceHandler.Create(ctx, in, out)
}
