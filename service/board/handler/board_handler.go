package handler

import (
	"context"

	/*
		"github.com/asim/go-micro/v3"
		"github.com/asim/go-micro/v3/auth"
		"github.com/asim/go-micro/v3/errors"
		"github.com/asim/go-micro/v3/logger"
	*/

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/auth"
	"github.com/micro/micro/v3/service/errors"
	"github.com/micro/micro/v3/service/logger"

	uuid "github.com/satori/go.uuid"
	"github.com/thoas/go-funk"
	"gorm.io/gorm"

	boardPB "github.com/ygpark2/njro/service/board/proto/board"
	board_entities "github.com/ygpark2/njro/service/board/proto/entities"
	"github.com/ygpark2/njro/service/board/repository"
	emailerPB "github.com/ygpark2/njro/service/emailer/proto/emailer"
	myErrors "github.com/ygpark2/njro/shared/errors"
)

// boardHandler struct
type BoardHandler struct {
	boardRepository repository.BoardRepository
	event           service.Event
}

// NewBoardHandler returns an instance of `BoardServiceHandler`.
func NewBoardHandler(repo repository.BoardRepository, eve service.Event) boardPB.BoardServiceHandler {
	return &BoardHandler{
		boardRepository: repo,
		event:           eve,
	}
}

func (h *BoardHandler) Exist(ctx context.Context, req *boardPB.ExistRequest, rsp *boardPB.ExistResponse) error {
	logger.Info("Received boardHandler.Exist request")
	/*
			google.protobuf.StringValue title = 7 [(gorm.field).tag = { size: 255 not_null: true }];
		    google.protobuf.StringValue mobile_title = 8 [(gorm.field).tag = { size: 255 not_null: true }];
		    google.protobuf.UInt32Value order = 9 [(gorm.field).tag = { not_null: true }];
		    google.protobuf.BoolValue search = 10 [(gorm.field).tag = { not_null: true }];
	*/
	model := board_entities.BoardORM{}
	model.Id = uuid.FromStringOrNil(req.Id.GetValue())
	title := req.Title.GetValue()
	model.Title = &title
	mobileTitle := req.MobileTitle.GetValue()
	model.MobileTitle = &mobileTitle
	order := req.Order.GetValue()
	model.Order = &order
	search := req.Search.GetValue()
	model.Search = &search

	exists := h.boardRepository.Exist(&model)
	logger.Info("user exists? %t", exists)
	rsp.Result = exists
	return nil
}

func (h *BoardHandler) List(ctx context.Context, req *boardPB.ListRequest, rsp *boardPB.ListResponse) error {
	logger.Info("Received boardHandler.List request")
	model := board_entities.BoardORM{}
	title := req.Title.GetValue()
	model.Title = &title
	mobileTitle := req.MobileTitle.GetValue()
	model.MobileTitle = &mobileTitle
	model.Description = req.Description
	model.Notices = req.Notices
	model.Order = &req.Order.Value
	model.Search = &req.Search.Value

	total, boards, err := h.boardRepository.List(int(req.Limit.GetValue()), int(req.Page.GetValue()), req.Sort.GetValue(), &model)
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
	newBoards := funk.Map(boards, func(board *board_entities.BoardORM) *board_entities.Board {
		tmpBoard, _ := board.ToPB(ctx)
		return &tmpBoard
	}).([]*board_entities.Board)

	rsp.Results = newBoards
	return nil
}

func (h *BoardHandler) Get(ctx context.Context, req *boardPB.GetRequest, rsp *boardPB.GetResponse) error {
	logger.Info("Received boardHandler.Get request")

	id := req.Id.GetValue()
	if id == "" {
		return myErrors.ValidationError("mkit.service.account.user.get", "validation error: Missing Id")
	}
	board, err := h.boardRepository.Get(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			rsp.Result = nil
			return nil
		}
		return myErrors.AppError(myErrors.DBE, err)
	}

	tempBoard, _ := board.ToPB(ctx)
	rsp.Result = &tempBoard

	return nil
}

func (h *BoardHandler) Create(ctx context.Context, req *boardPB.CreateRequest, rsp *boardPB.CreateResponse) error {
	logger.Info("Received boardHandler.Create request")

	model := board_entities.BoardORM{}
	title := req.Title.Value
	model.Title = &title
	mobileTitle := req.MobileTitle.Value
	model.MobileTitle = &mobileTitle
	model.Description = req.Description
	model.Notices = req.Notices
	model.Order = &req.Order.Value
	model.Search = &req.Search.Value

	if err := h.boardRepository.Create(&model); err != nil {
		return myErrors.AppError(myErrors.DBE, err)
	}

	// send email (TODO: async `go h.Event.Publish(...)`)
	if err := h.event.Publish(ctx, &emailerPB.Message{To: req.Title.GetValue()}); err != nil {
		logger.Error("Received Event.Publish request error")
		return myErrors.AppError(myErrors.PSE, err)
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

func (h *BoardHandler) Update(ctx context.Context, req *boardPB.UpdateRequest, rsp *boardPB.UpdateResponse) error {
	logger.Info("Received boardHandler.Update request")
	// Identify the user
	acc, ok := auth.AccountFromContext(ctx)
	if !ok {
		return errors.Unauthorized("mkit.service.account.user.update", "A valid auth token is required")
	}
	logger.Info("Caller Account: %v", acc)

	id := req.Id.GetValue()
	if id == "" {
		return myErrors.ValidationError("mkit.service.account.user.update", "validation error: Missing Id")
	}

	model := board_entities.BoardORM{}
	title := req.Title.GetValue()
	model.Title = &title
	mobileTitle := req.MobileTitle.GetValue()
	model.MobileTitle = &mobileTitle
	model.Description = req.Description
	model.Notices = req.Notices
	model.Order = &req.Order.Value
	model.Search = &req.Search.Value

	if err := h.boardRepository.Update(id, &model); err != nil {
		return myErrors.AppError(myErrors.DBE, err)
	}

	return nil
}

func (h *BoardHandler) Delete(ctx context.Context, req *boardPB.DeleteRequest, rsp *boardPB.DeleteResponse) error {
	logger.Info("Received boardHandler.Delete request")

	id := req.Id.GetValue()
	if id == "" {
		return myErrors.ValidationError("mkit.service.account.user.update", "validation error: Missing Id")
	}

	model := board_entities.BoardORM{}
	model.Id = uuid.FromStringOrNil(id)

	if err := h.boardRepository.Delete(&model); err != nil {
		return myErrors.AppError(myErrors.DBE, err)
	}

	return nil
}
