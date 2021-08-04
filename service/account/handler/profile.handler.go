package handler

import (
	"context"
	"time"

	"github.com/thoas/go-funk"
	"gorm.io/gorm"

	ptypes1 "github.com/golang/protobuf/ptypes"
	"github.com/rs/zerolog"
	uuid "github.com/satori/go.uuid"

	account_entities "github.com/ygpark2/njro/service/account/proto/entities"
	profilePB "github.com/ygpark2/njro/service/account/proto/profile"
	"github.com/ygpark2/njro/service/account/repository"
	njErrors "github.com/ygpark2/njro/shared/errors"
)

// ProfileHandler struct
type profileHandler struct {
	profileRepository repository.ProfileRepository
	contextLogger     zerolog.Logger
}

// NewProfileHandler returns an instance of `ProfileServiceHandler`.
func NewProfileHandler(repo repository.ProfileRepository, logger zerolog.Logger) profilePB.ProfileServiceHandler {
	return &profileHandler{
		profileRepository: repo,
		contextLogger:     logger,
	}
}

func (ph *profileHandler) List(ctx context.Context, req *profilePB.ListRequest, rsp *profilePB.ListResponse) error {
	ph.contextLogger.Debug().Msg("Received ProfileHandler.List request")
	preferredTheme := req.PreferredTheme.GetValue()
	model := account_entities.ProfileORM{
		// UserID:     uuid.FromStringOrNil(req.UserId.GetValue()),
		PreferredTheme: &preferredTheme,
		Gender:         account_entities.Profile_GenderType_name[int32(req.Gender)],
	}

	total, profiles, err := ph.profileRepository.List(req.Limit.GetValue(), req.Page.GetValue(), req.Sort.GetValue(), &model)
	if err != nil {
		return njErrors.AppError(njErrors.DBE, err)
	}
	rsp.Total = total
	// newProfiles := make([]*pb.Profile, len(profiles))
	// for index, profile := range profiles {
	// 	tempProfile, _ := profile.ToPB(ctx)
	// 	newProfiles[index] = &tempProfile
	// }
	newProfiles := funk.Map(profiles, func(profile *account_entities.ProfileORM) *account_entities.Profile {
		tempProfile, _ := profile.ToPB(ctx)
		return &tempProfile
	}).([]*account_entities.Profile)

	rsp.Results = newProfiles
	return nil
}

func (ph *profileHandler) Get(ctx context.Context, req *profilePB.GetRequest, rsp *profilePB.GetResponse) error {
	ph.contextLogger.Debug().Msg("Received ProfileHandler.Get request")
	var profile *account_entities.ProfileORM
	var err error
	switch id := req.Id.(type) {
	case *profilePB.GetRequest_UserId:
		println("GetRequest_UserId")
		println(req.GetId())
		profile, err = ph.profileRepository.GetByUserID(id.UserId.GetValue())
	case *profilePB.GetRequest_ProfileId:
		println("GetRequest_ProfileId")
		println(req.GetId())
		profile, err = ph.profileRepository.Get(id.ProfileId.GetValue())
	case nil:
		return njErrors.ValidationError("mkit.service.account.profile.get", "validation error: Missing Id")
	default:
		return njErrors.ValidationError("mkit.service.account.profile.get", "validation error: Profile.Id has unexpected type %T", id)
	}
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			rsp.Result = nil
			return nil
		}
		return njErrors.AppError(njErrors.DBE, err)
	}

	tempProfile, _ := profile.ToPB(ctx)
	rsp.Result = &tempProfile
	return nil
}

func (ph *profileHandler) Create(ctx context.Context, req *profilePB.CreateRequest, rsp *profilePB.CreateResponse) error {
	ph.contextLogger.Debug().Msg("Received ProfileHandler.Create request")
	model := account_entities.ProfileORM{}
	userId := uuid.FromStringOrNil(req.UserId.GetValue())
	model.UserId = &userId
	model.Tz = req.Tz.GetValue()
	model.Gender = account_entities.Profile_GenderType_name[int32(req.Gender)]
	model.Avatar = req.Avatar.GetValue()
	if req.Birthday != nil {
		var t time.Time
		var err error
		if t, err = ptypes1.Timestamp(req.Birthday); err != nil {
			return njErrors.ValidationError("mkit.service.account.profile.rceate", "Invalid birthday: %v", err)
		}
		model.Birthday = &t
	}
	preferredTheme := req.PreferredTheme.GetValue()
	model.PreferredTheme = &preferredTheme

	if err := ph.profileRepository.Create(&model); err != nil {
		return njErrors.AppError(njErrors.DBE, err)
	}
	return nil
}
