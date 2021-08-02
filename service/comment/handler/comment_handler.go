package handler

import (
	"context"

	// "github.com/jinzhu/gorm"
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/auth"
	"github.com/micro/micro/v3/service/errors"
	"github.com/micro/micro/v3/service/logger"
	uuid "github.com/satori/go.uuid"
	"github.com/thoas/go-funk"
	"gorm.io/gorm"

	commentPB "github.com/ygpark2/mboard/service/comment/proto/comment"
	comment_entities "github.com/ygpark2/mboard/service/comment/proto/entities"
	"github.com/ygpark2/mboard/service/comment/repository"
	myErrors "github.com/ygpark2/mboard/shared/errors"
)

// commentHandler struct
type CommentHandler struct {
	commentRepository repository.CommentRepository
	event             *service.Event
}

// NewCommentHandler returns an instance of `CommentServiceHandler`.
func NewCommentHandler(repo repository.CommentRepository, eve *service.Event) commentPB.CommentServiceHandler {
	return &CommentHandler{
		commentRepository: repo,
		event:             eve,
	}
}

func (h *CommentHandler) Exist(ctx context.Context, req *commentPB.ExistRequest, rsp *commentPB.ExistResponse) error {
	logger.Info("Received boardHandler.Exist request")
	/*
			google.protobuf.StringValue title = 7 [(gorm.field).tag = { size: 255 not_null: true }];
		    google.protobuf.StringValue mobile_title = 8 [(gorm.field).tag = { size: 255 not_null: true }];
		    google.protobuf.UInt32Value order = 9 [(gorm.field).tag = { not_null: true }];
		    google.protobuf.BoolValue search = 10 [(gorm.field).tag = { not_null: true }];
	*/
	model := comment_entities.CommentORM{}
	model.Id = uuid.FromStringOrNil(req.Id.GetValue())
	title := req.Title.GetValue()
	model.Title = &title
	mobileTitle := req.MobileTitle.GetValue()
	model.MobileTitle = &mobileTitle
	order := req.Order.GetValue()
	model.Order = &order
	search := req.Search.GetValue()
	model.Search = &search

	exists := h.commentRepository.Exist(&model)
	logger.Info("user exists? %t", exists)
	rsp.Result = exists
	return nil
}

func (h *CommentHandler) List(ctx context.Context, req *commentPB.ListRequest, rsp *commentPB.ListResponse) error {
	logger.Info("Received boardHandler.List request")
	model := comment_entities.CommentORM{}
	title := req.Title.GetValue()
	model.Title = &title
	mobileTitle := req.MobileTitle.GetValue()
	model.MobileTitle = &mobileTitle
	model.Description = req.Description
	model.Notices = req.Notices
	model.Order = &req.Order.Value
	model.Search = &req.Search.Value

	total, boards, err := h.commentRepository.List(int(req.Limit.GetValue()), int(req.Page.GetValue()), req.Sort.GetValue(), &model)
	if err != nil {
		return errors.NotFound("mkit.service.board.list", "Error %v", err.Error())
	}
	rsp.Total = total

	// newBoards := make([]*accountPB.User, len(boards))
	// for index, board := range boards {
	// 	tmpBoard, _ := board.ToPB(ctx)
	// 	newBoards[index] = &tmpBoard
	// 	// *newBoards[index], _ = board.ToPB(ctx) ???
	// }
	newBoards := funk.Map(boards, func(board *comment_entities.CommentORM) *comment_entities.Board {
		tmpBoard, _ := board.ToPB(ctx)
		return &tmpBoard
	}).([]*comment_entities.Comment)

	rsp.Results = newBoards
	return nil
}

func (h *CommentHandler) Get(ctx context.Context, req *commentPB.GetRequest, rsp *commentPB.GetResponse) error {
	logger.Info("Received boardHandler.Get request")

	id := req.Id.GetValue()
	if id == "" {
		return myErrors.ValidationError("mkit.service.comment.get", "validation error: Missing Id")
	}
	comment, err := h.commentRepository.Get(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			rsp.Result = nil
			return nil
		}
		return myErrors.AppError(myErrors.DBE, err)
	}

	tempComment, _ := comment.ToPB(ctx)
	rsp.Result = &tempComment

	return nil
}

func (h *CommentHandler) Create(ctx context.Context, req *commentPB.CreateRequest, rsp *commentPB.CreateResponse) error {
	logger.Info("Received commentHandler.Create request")

	model := comment_entities.CommentORM{}
	title := req.Title.GetValue()
	model.Title = &title
	mobileTitle := req.MobileTitle.GetValue()
	model.MobileTitle = &mobileTitle
	model.Description = req.Description
	model.Notices = req.Notices
	model.Order = &req.Order.Value
	model.Search = &req.Search.Value

	if err := h.commentRepository.Create(&model); err != nil {
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

func (h *CommentHandler) Update(ctx context.Context, req *commentPB.UpdateRequest, rsp *commentPB.UpdateResponse) error {
	logger.Info("Received commentHandler.Update request")
	// Identify the user
	acc, ok := auth.AccountFromContext(ctx)
	if !ok {
		return errors.Unauthorized("mkit.service.comment.update", "A valid auth token is required")
	}
	logger.Info("Caller Account: %v", acc)

	id := req.Id.GetValue()
	if id == "" {
		return myErrors.ValidationError("mkit.service.comment.update", "validation error: Missing Id")
	}

	model := comment_entities.CommentORM{}
	title := req.Title.GetValue()
	model.Title = &title
	mobileTitle := req.MobileTitle.GetValue()
	model.MobileTitle = &mobileTitle
	model.Description = req.Description
	model.Notices = req.Notices
	model.Order = &req.Order.Value
	model.Search = &req.Search.Value

	if err := h.commentRepository.Update(id, &model); err != nil {
		return myErrors.AppError(myErrors.DBE, err)
	}

	return nil
}

func (h *CommentHandler) Delete(ctx context.Context, req *commentPB.DeleteRequest, rsp *commentPB.DeleteResponse) error {
	logger.Info("Received commentHandler.Delete request")

	id := req.Id.GetValue()
	if id == "" {
		return myErrors.ValidationError("mkit.service.comment.update", "validation error: Missing Id")
	}

	model := comment_entities.CommentORM{}
	model.Id = uuid.FromStringOrNil(id)

	if err := h.commentRepository.Delete(&model); err != nil {
		return myErrors.AppError(myErrors.DBE, err)
	}

	return nil
}
