package handler

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/auth"
	"github.com/micro/micro/v3/service/errors"
	"github.com/micro/micro/v3/service/logger"
	uuid "github.com/satori/go.uuid"
	"github.com/thoas/go-funk"

	post_entities "github.com/ygpark2/mboard/service/post/proto/entities"
	postPB "github.com/ygpark2/mboard/service/post/proto/post"
	"github.com/ygpark2/mboard/service/post/repository"
	myErrors "github.com/ygpark2/mboard/shared/errors"
)

// Posts struct
type Posts struct {
	postRepository repository.PostRepository
	event          *service.Event
}

// NewPosts returns an instance of `PostServiceHandler`.
func NewPosts(repo repository.PostRepository, eve *service.Event) postPB.PostServiceHandler {
	return &Posts{
		postRepository: repo,
		event:          eve,
	}
}

func (h *Posts) Exist(ctx context.Context, req *postPB.ExistRequest, rsp *postPB.ExistResponse) error {
	logger.Info("Received Posts.Exist request")
	model := post_entities.PostORM{}
	model.Id = uuid.FromStringOrNil(req.Id.GetValue())
	model.BoardId = req.BoardId.GetValue()
	title := req.Title.GetValue()
	model.Title = &title
	mobileTitle := req.MobileTitle.GetValue()
	model.MobileTitle = &mobileTitle
	order := req.Order.GetValue()
	model.Order = &order
	search := req.Search.GetValue()
	model.Search = &search

	exists := h.postRepository.Exist(&model)
	logger.Info("user exists? %t", exists)
	rsp.Result = exists
	return nil
}

func (h *Posts) List(ctx context.Context, req *postPB.ListRequest, rsp *postPB.ListResponse) error {
	logger.Info("Received Posts.List request")
	model := post_entities.PostORM{}
	title := req.Title.GetValue()
	model.Title = &title
	mobileTitle := req.MobileTitle.GetValue()
	model.MobileTitle = &mobileTitle
	model.Description = req.Description
	model.Notices = req.Notices
	model.Order = &req.Order.Value
	model.Search = &req.Search.Value

	total, posts, err := h.postRepository.List(int(req.Limit.GetValue()), int(req.Page.GetValue()), req.Sort.GetValue(), &model)
	if err != nil {
		return errors.NotFound("mkit.service.post.list", "Error %v", err.Error())
	}
	rsp.Total = total

	// newBoards := make([]*accountPB.User, len(boards))
	// for index, board := range boards {
	// 	tmpBoard, _ := board.ToPB(ctx)
	// 	newBoards[index] = &tmpBoard
	// 	// *newBoards[index], _ = board.ToPB(ctx) ???
	// }
	newPosts := funk.Map(posts, func(post *post_entities.PostORM) *post_entities.Post {
		tmpPost, _ := post.ToPB(ctx)
		return &tmpPost
	}).([]*post_entities.Post)

	rsp.Results = newPosts
	return nil
}

func (h *Posts) Get(ctx context.Context, req *postPB.GetRequest, rsp *postPB.GetResponse) error {
	logger.Info("Received Posts.Get request")

	id := req.Id.GetValue()
	if id == "" {
		return myErrors.ValidationError("mkit.service.post.get", "validation error: Missing Id")
	}
	post, err := h.postRepository.Get(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			rsp.Result = nil
			return nil
		}
		return myErrors.AppError(myErrors.DBE, err)
	}

	tempPost, _ := post.ToPB(ctx)
	rsp.Result = &tempPost

	return nil
}

func (h *Posts) Create(ctx context.Context, req *postPB.CreateRequest, rsp *postPB.CreateResponse) error {
	logger.Info("Received Posts.Create request")

	model := post_entities.PostORM{}
	title := req.Title.GetValue()
	model.Title = &title
	mobileTitle := req.MobileTitle.GetValue()
	model.MobileTitle = &mobileTitle
	model.Description = req.Description
	model.Notices = req.Notices
	model.Order = &req.Order.Value
	model.Search = &req.Search.Value

	if err := h.postRepository.Create(&model); err != nil {
		return myErrors.AppError(myErrors.DBE, err)
	}

	// send email (TODO: async `go h.Event.Publish(...)`)
	/*
		if err := events.Publish(ctx, &emailerPB.Message{To: req.Email.GetValue()}); err != nil {
			log.Error().Err(err).Msg("Received Event.Publish request error")
			return myErrors.AppError(myErrors.PSE, err)
		}
	*/

	return nil
}

func (h *Posts) Update(ctx context.Context, req *postPB.UpdateRequest, rsp *postPB.UpdateResponse) error {
	logger.Info("Received Posts.Update request")
	// Identify the user
	acc, ok := auth.AccountFromContext(ctx)
	if !ok {
		return errors.Unauthorized("mkit.service.post.update", "A valid auth token is required")
	}
	logger.Info("Caller Account: %v", acc)

	id := req.Id.GetValue()
	if id == "" {
		return myErrors.ValidationError("mkit.service.post.update", "validation error: Missing Id")
	}

	model := post_entities.PostORM{}
	model.BoardId = req
	title := req.Title.GetValue()
	model.Title = &title
	mobileTitle := req.MobileTitle.GetValue()
	model.MobileTitle = &mobileTitle
	model.Description = req.Description
	model.Notices = req.Notices
	model.Order = &req.Order.Value
	model.Search = &req.Search.Value

	if err := h.postRepository.Update(id, &model); err != nil {
		return myErrors.AppError(myErrors.DBE, err)
	}

	return nil
}

func (h *Posts) Delete(ctx context.Context, req *postPB.DeleteRequest, rsp *postPB.DeleteResponse) error {
	logger.Info("Received Posts.Delete request")

	id := req.Id.GetValue()
	if id == "" {
		return myErrors.ValidationError("mkit.service.post.update", "validation error: Missing Id")
	}

	model := post_entities.PostORM{}
	model.Id = uuid.FromStringOrNil(id)

	if err := h.postRepository.Delete(&model); err != nil {
		return myErrors.AppError(myErrors.DBE, err)
	}

	return nil
}
