// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: service/comment/proto/comment/comment_service.proto

package comment

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/ygpark2/njro/service/comment/proto/entities"
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

// Api Endpoints for CommentService service

func NewCommentServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for CommentService service

type CommentService interface {
	Exist(ctx context.Context, in *ExistRequest, opts ...client.CallOption) (*ExistResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error)
	Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error)
}

type commentService struct {
	c    client.Client
	name string
}

func NewCommentService(name string, c client.Client) CommentService {
	return &commentService{
		c:    c,
		name: name,
	}
}

func (c *commentService) Exist(ctx context.Context, in *ExistRequest, opts ...client.CallOption) (*ExistResponse, error) {
	req := c.c.NewRequest(c.name, "CommentService.Exist", in)
	out := new(ExistResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentService) List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error) {
	req := c.c.NewRequest(c.name, "CommentService.List", in)
	out := new(ListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentService) Get(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error) {
	req := c.c.NewRequest(c.name, "CommentService.Get", in)
	out := new(GetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentService) Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error) {
	req := c.c.NewRequest(c.name, "CommentService.Create", in)
	out := new(CreateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentService) Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error) {
	req := c.c.NewRequest(c.name, "CommentService.Update", in)
	out := new(UpdateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentService) Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error) {
	req := c.c.NewRequest(c.name, "CommentService.Delete", in)
	out := new(DeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CommentService service

type CommentServiceHandler interface {
	Exist(context.Context, *ExistRequest, *ExistResponse) error
	List(context.Context, *ListRequest, *ListResponse) error
	Get(context.Context, *GetRequest, *GetResponse) error
	Create(context.Context, *CreateRequest, *CreateResponse) error
	Update(context.Context, *UpdateRequest, *UpdateResponse) error
	Delete(context.Context, *DeleteRequest, *DeleteResponse) error
}

func RegisterCommentServiceHandler(s server.Server, hdlr CommentServiceHandler, opts ...server.HandlerOption) error {
	type commentService interface {
		Exist(ctx context.Context, in *ExistRequest, out *ExistResponse) error
		List(ctx context.Context, in *ListRequest, out *ListResponse) error
		Get(ctx context.Context, in *GetRequest, out *GetResponse) error
		Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error
		Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error
		Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error
	}
	type CommentService struct {
		commentService
	}
	h := &commentServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&CommentService{h}, opts...))
}

type commentServiceHandler struct {
	CommentServiceHandler
}

func (h *commentServiceHandler) Exist(ctx context.Context, in *ExistRequest, out *ExistResponse) error {
	return h.CommentServiceHandler.Exist(ctx, in, out)
}

func (h *commentServiceHandler) List(ctx context.Context, in *ListRequest, out *ListResponse) error {
	return h.CommentServiceHandler.List(ctx, in, out)
}

func (h *commentServiceHandler) Get(ctx context.Context, in *GetRequest, out *GetResponse) error {
	return h.CommentServiceHandler.Get(ctx, in, out)
}

func (h *commentServiceHandler) Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error {
	return h.CommentServiceHandler.Create(ctx, in, out)
}

func (h *commentServiceHandler) Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error {
	return h.CommentServiceHandler.Update(ctx, in, out)
}

func (h *commentServiceHandler) Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error {
	return h.CommentServiceHandler.Delete(ctx, in, out)
}
