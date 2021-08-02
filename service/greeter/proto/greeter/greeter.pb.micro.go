// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: service/greeter/proto/greeter/greeter.proto

package greeter

import (
	fmt "fmt"
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

// Api Endpoints for GreeterService service

func NewGreeterServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for GreeterService service

type GreeterService interface {
	// Hello is echo method
	Hello(ctx context.Context, in *HelloRequest, opts ...client.CallOption) (*HelloResponse, error)
}

type greeterService struct {
	c    client.Client
	name string
}

func NewGreeterService(name string, c client.Client) GreeterService {
	return &greeterService{
		c:    c,
		name: name,
	}
}

func (c *greeterService) Hello(ctx context.Context, in *HelloRequest, opts ...client.CallOption) (*HelloResponse, error) {
	req := c.c.NewRequest(c.name, "GreeterService.Hello", in)
	out := new(HelloResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GreeterService service

type GreeterServiceHandler interface {
	// Hello is echo method
	Hello(context.Context, *HelloRequest, *HelloResponse) error
}

func RegisterGreeterServiceHandler(s server.Server, hdlr GreeterServiceHandler, opts ...server.HandlerOption) error {
	type greeterService interface {
		Hello(ctx context.Context, in *HelloRequest, out *HelloResponse) error
	}
	type GreeterService struct {
		greeterService
	}
	h := &greeterServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&GreeterService{h}, opts...))
}

type greeterServiceHandler struct {
	GreeterServiceHandler
}

func (h *greeterServiceHandler) Hello(ctx context.Context, in *HelloRequest, out *HelloResponse) error {
	return h.GreeterServiceHandler.Hello(ctx, in, out)
}