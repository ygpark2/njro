// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: service/comment/proto/entities/entities.proto

package entities

import (
	context "context"
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/infobloxopen/protoc-gen-gorm/options"
	_ "github.com/infobloxopen/protoc-gen-gorm/types"
	math "math"
	strings "strings"
	time "time"

	ptypes1 "github.com/golang/protobuf/ptypes"
	gorm2 "github.com/infobloxopen/atlas-app-toolkit/gorm"
	errors1 "github.com/infobloxopen/protoc-gen-gorm/errors"
	types1 "github.com/infobloxopen/protoc-gen-gorm/types"
	gorm1 "github.com/jinzhu/gorm"
	go_uuid1 "github.com/satori/go.uuid"
	field_mask1 "google.golang.org/genproto/protobuf/field_mask"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = math.Inf

type CommentORM struct {
	BoardId       *string       `gorm:"not null;index:idx_board"`
	ContentId     *string       `gorm:"not null;index:idx_content"`
	CreatedAt     *time.Time    `gorm:"not null"`
	DeletedAt     *time.Time    `gorm:"index:idx_comments_deleted_at"`
	DownVoteCount uint32        `gorm:"not null"`
	Email         *string       `gorm:"size:250;not null"`
	Id            go_uuid1.UUID `gorm:"type:uuid;primary_key;unique;not null"`
	Nickname      *string       `gorm:"size:200;not null"`
	Password      *string
	PostId        *string    `gorm:"not null;index:idx_post"`
	UpVoteCount   uint32     `gorm:"not null"`
	UpdatedAt     *time.Time `gorm:"not null"`
	Url           *string
	UseHtml       *bool   `gorm:"not null"`
	UseSecret     *bool   `gorm:"not null"`
	Userid        *string `gorm:"size:100;not null"`
	Username      *string `gorm:"size:200;not null"`
}

// TableName overrides the default tablename generated by GORM
func (CommentORM) TableName() string {
	return "comments"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *Comment) ToORM(ctx context.Context) (CommentORM, error) {
	to := CommentORM{}
	var err error
	if prehook, ok := interface{}(m).(CommentWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	if m.Id != nil {
		to.Id, err = go_uuid1.FromString(m.Id.Value)
		if err != nil {
			return to, err
		}
	} else {
		to.Id = go_uuid1.Nil
	}
	if m.CreatedAt != nil {
		var t time.Time
		if t, err = ptypes1.Timestamp(m.CreatedAt); err != nil {
			return to, err
		}
		to.CreatedAt = &t
	}
	if m.UpdatedAt != nil {
		var t time.Time
		if t, err = ptypes1.Timestamp(m.UpdatedAt); err != nil {
			return to, err
		}
		to.UpdatedAt = &t
	}
	if m.DeletedAt != nil {
		var t time.Time
		if t, err = ptypes1.Timestamp(m.DeletedAt); err != nil {
			return to, err
		}
		to.DeletedAt = &t
	}
	if m.BoardId != nil {
		v := m.BoardId.Value
		to.BoardId = &v
	}
	if m.PostId != nil {
		v := m.PostId.Value
		to.PostId = &v
	}
	if m.ContentId != nil {
		v := m.ContentId.Value
		to.ContentId = &v
	}
	if m.Userid != nil {
		v := m.Userid.Value
		to.Userid = &v
	}
	if m.Username != nil {
		v := m.Username.Value
		to.Username = &v
	}
	if m.Nickname != nil {
		v := m.Nickname.Value
		to.Nickname = &v
	}
	if m.Email != nil {
		v := m.Email.Value
		to.Email = &v
	}
	if m.Password != nil {
		v := m.Password.Value
		to.Password = &v
	}
	if m.Url != nil {
		v := m.Url.Value
		to.Url = &v
	}
	if m.UseHtml != nil {
		v := m.UseHtml.Value
		to.UseHtml = &v
	}
	if m.UseSecret != nil {
		v := m.UseSecret.Value
		to.UseSecret = &v
	}
	to.UpVoteCount = m.UpVoteCount
	to.DownVoteCount = m.DownVoteCount
	if posthook, ok := interface{}(m).(CommentWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *CommentORM) ToPB(ctx context.Context) (Comment, error) {
	to := Comment{}
	var err error
	if prehook, ok := interface{}(m).(CommentWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = &types1.UUID{Value: m.Id.String()}
	if m.CreatedAt != nil {
		if to.CreatedAt, err = ptypes1.TimestampProto(*m.CreatedAt); err != nil {
			return to, err
		}
	}
	if m.UpdatedAt != nil {
		if to.UpdatedAt, err = ptypes1.TimestampProto(*m.UpdatedAt); err != nil {
			return to, err
		}
	}
	if m.DeletedAt != nil {
		if to.DeletedAt, err = ptypes1.TimestampProto(*m.DeletedAt); err != nil {
			return to, err
		}
	}
	if m.BoardId != nil {
		to.BoardId = &wrappers.StringValue{Value: *m.BoardId}
	}
	if m.PostId != nil {
		to.PostId = &wrappers.StringValue{Value: *m.PostId}
	}
	if m.ContentId != nil {
		to.ContentId = &wrappers.StringValue{Value: *m.ContentId}
	}
	if m.Userid != nil {
		to.Userid = &wrappers.StringValue{Value: *m.Userid}
	}
	if m.Username != nil {
		to.Username = &wrappers.StringValue{Value: *m.Username}
	}
	if m.Nickname != nil {
		to.Nickname = &wrappers.StringValue{Value: *m.Nickname}
	}
	if m.Email != nil {
		to.Email = &wrappers.StringValue{Value: *m.Email}
	}
	if m.Password != nil {
		to.Password = &wrappers.StringValue{Value: *m.Password}
	}
	if m.Url != nil {
		to.Url = &wrappers.StringValue{Value: *m.Url}
	}
	if m.UseHtml != nil {
		to.UseHtml = &wrappers.BoolValue{Value: *m.UseHtml}
	}
	if m.UseSecret != nil {
		to.UseSecret = &wrappers.BoolValue{Value: *m.UseSecret}
	}
	to.UpVoteCount = m.UpVoteCount
	to.DownVoteCount = m.DownVoteCount
	if posthook, ok := interface{}(m).(CommentWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type Comment the arg will be the target, the caller the one being converted from

// CommentBeforeToORM called before default ToORM code
type CommentWithBeforeToORM interface {
	BeforeToORM(context.Context, *CommentORM) error
}

// CommentAfterToORM called after default ToORM code
type CommentWithAfterToORM interface {
	AfterToORM(context.Context, *CommentORM) error
}

// CommentBeforeToPB called before default ToPB code
type CommentWithBeforeToPB interface {
	BeforeToPB(context.Context, *Comment) error
}

// CommentAfterToPB called after default ToPB code
type CommentWithAfterToPB interface {
	AfterToPB(context.Context, *Comment) error
}

// DefaultCreateComment executes a basic gorm create call
func DefaultCreateComment(ctx context.Context, in *Comment, db *gorm1.DB) (*Comment, error) {
	if in == nil {
		return nil, errors1.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(CommentORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(CommentORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type CommentORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type CommentORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm1.DB) error
}

// DefaultReadComment executes a basic gorm read call
func DefaultReadComment(ctx context.Context, in *Comment, db *gorm1.DB) (*Comment, error) {
	if in == nil {
		return nil, errors1.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if ormObj.Id == go_uuid1.Nil {
		return nil, errors1.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(CommentORMWithBeforeReadApplyQuery); ok {
		if db, err = hook.BeforeReadApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	if db, err = gorm2.ApplyFieldSelection(ctx, db, nil, &CommentORM{}); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(CommentORMWithBeforeReadFind); ok {
		if db, err = hook.BeforeReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	ormResponse := CommentORM{}
	if err = db.Where(&ormObj).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormResponse).(CommentORMWithAfterReadFind); ok {
		if err = hook.AfterReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

type CommentORMWithBeforeReadApplyQuery interface {
	BeforeReadApplyQuery(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type CommentORMWithBeforeReadFind interface {
	BeforeReadFind(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type CommentORMWithAfterReadFind interface {
	AfterReadFind(context.Context, *gorm1.DB) error
}

func DefaultDeleteComment(ctx context.Context, in *Comment, db *gorm1.DB) error {
	if in == nil {
		return errors1.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return err
	}
	if ormObj.Id == go_uuid1.Nil {
		return errors1.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(CommentORMWithBeforeDelete_); ok {
		if db, err = hook.BeforeDelete_(ctx, db); err != nil {
			return err
		}
	}
	err = db.Where(&ormObj).Delete(&CommentORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := interface{}(&ormObj).(CommentORMWithAfterDelete_); ok {
		err = hook.AfterDelete_(ctx, db)
	}
	return err
}

type CommentORMWithBeforeDelete_ interface {
	BeforeDelete_(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type CommentORMWithAfterDelete_ interface {
	AfterDelete_(context.Context, *gorm1.DB) error
}

func DefaultDeleteCommentSet(ctx context.Context, in []*Comment, db *gorm1.DB) error {
	if in == nil {
		return errors1.NilArgumentError
	}
	var err error
	keys := []go_uuid1.UUID{}
	for _, obj := range in {
		ormObj, err := obj.ToORM(ctx)
		if err != nil {
			return err
		}
		if ormObj.Id == go_uuid1.Nil {
			return errors1.EmptyIdError
		}
		keys = append(keys, ormObj.Id)
	}
	if hook, ok := (interface{}(&CommentORM{})).(CommentORMWithBeforeDeleteSet); ok {
		if db, err = hook.BeforeDeleteSet(ctx, in, db); err != nil {
			return err
		}
	}
	err = db.Where("id in (?)", keys).Delete(&CommentORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := (interface{}(&CommentORM{})).(CommentORMWithAfterDeleteSet); ok {
		err = hook.AfterDeleteSet(ctx, in, db)
	}
	return err
}

type CommentORMWithBeforeDeleteSet interface {
	BeforeDeleteSet(context.Context, []*Comment, *gorm1.DB) (*gorm1.DB, error)
}
type CommentORMWithAfterDeleteSet interface {
	AfterDeleteSet(context.Context, []*Comment, *gorm1.DB) error
}

// DefaultStrictUpdateComment clears / replaces / appends first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateComment(ctx context.Context, in *Comment, db *gorm1.DB) (*Comment, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateComment")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	lockedRow := &CommentORM{}
	db.Model(&ormObj).Set("gorm:query_option", "FOR UPDATE").Where("id=?", ormObj.Id).First(lockedRow)
	if hook, ok := interface{}(&ormObj).(CommentORMWithBeforeStrictUpdateCleanup); ok {
		if db, err = hook.BeforeStrictUpdateCleanup(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(CommentORMWithBeforeStrictUpdateSave); ok {
		if db, err = hook.BeforeStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(CommentORMWithAfterStrictUpdateSave); ok {
		if err = hook.AfterStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	if err != nil {
		return nil, err
	}
	return &pbResponse, err
}

type CommentORMWithBeforeStrictUpdateCleanup interface {
	BeforeStrictUpdateCleanup(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type CommentORMWithBeforeStrictUpdateSave interface {
	BeforeStrictUpdateSave(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type CommentORMWithAfterStrictUpdateSave interface {
	AfterStrictUpdateSave(context.Context, *gorm1.DB) error
}

// DefaultPatchComment executes a basic gorm update call with patch behavior
func DefaultPatchComment(ctx context.Context, in *Comment, updateMask *field_mask1.FieldMask, db *gorm1.DB) (*Comment, error) {
	if in == nil {
		return nil, errors1.NilArgumentError
	}
	var pbObj Comment
	var err error
	if hook, ok := interface{}(&pbObj).(CommentWithBeforePatchRead); ok {
		if db, err = hook.BeforePatchRead(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbReadRes, err := DefaultReadComment(ctx, &Comment{Id: in.GetId()}, db)
	if err != nil {
		return nil, err
	}
	pbObj = *pbReadRes
	if hook, ok := interface{}(&pbObj).(CommentWithBeforePatchApplyFieldMask); ok {
		if db, err = hook.BeforePatchApplyFieldMask(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	if _, err := DefaultApplyFieldMaskComment(ctx, &pbObj, in, updateMask, "", db); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&pbObj).(CommentWithBeforePatchSave); ok {
		if db, err = hook.BeforePatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := DefaultStrictUpdateComment(ctx, &pbObj, db)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(pbResponse).(CommentWithAfterPatchSave); ok {
		if err = hook.AfterPatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	return pbResponse, nil
}

type CommentWithBeforePatchRead interface {
	BeforePatchRead(context.Context, *Comment, *field_mask1.FieldMask, *gorm1.DB) (*gorm1.DB, error)
}
type CommentWithBeforePatchApplyFieldMask interface {
	BeforePatchApplyFieldMask(context.Context, *Comment, *field_mask1.FieldMask, *gorm1.DB) (*gorm1.DB, error)
}
type CommentWithBeforePatchSave interface {
	BeforePatchSave(context.Context, *Comment, *field_mask1.FieldMask, *gorm1.DB) (*gorm1.DB, error)
}
type CommentWithAfterPatchSave interface {
	AfterPatchSave(context.Context, *Comment, *field_mask1.FieldMask, *gorm1.DB) error
}

// DefaultPatchSetComment executes a bulk gorm update call with patch behavior
func DefaultPatchSetComment(ctx context.Context, objects []*Comment, updateMasks []*field_mask1.FieldMask, db *gorm1.DB) ([]*Comment, error) {
	if len(objects) != len(updateMasks) {
		return nil, fmt.Errorf(errors1.BadRepeatedFieldMaskTpl, len(updateMasks), len(objects))
	}

	results := make([]*Comment, 0, len(objects))
	for i, patcher := range objects {
		pbResponse, err := DefaultPatchComment(ctx, patcher, updateMasks[i], db)
		if err != nil {
			return nil, err
		}

		results = append(results, pbResponse)
	}

	return results, nil
}

// DefaultApplyFieldMaskComment patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskComment(ctx context.Context, patchee *Comment, patcher *Comment, updateMask *field_mask1.FieldMask, prefix string, db *gorm1.DB) (*Comment, error) {
	if patcher == nil {
		return nil, nil
	} else if patchee == nil {
		return nil, errors1.NilArgumentError
	}
	var err error
	var updatedCreatedAt bool
	var updatedUpdatedAt bool
	var updatedDeletedAt bool
	var updatedBoardId bool
	var updatedPostId bool
	var updatedContentId bool
	var updatedUserid bool
	var updatedUsername bool
	var updatedNickname bool
	var updatedEmail bool
	var updatedPassword bool
	var updatedUrl bool
	var updatedUseHtml bool
	var updatedUseSecret bool
	for i, f := range updateMask.Paths {
		if f == prefix+"Id" {
			patchee.Id = patcher.Id
			continue
		}
		if !updatedCreatedAt && strings.HasPrefix(f, prefix+"CreatedAt.") {
			if patcher.CreatedAt == nil {
				patchee.CreatedAt = nil
				continue
			}
			if patchee.CreatedAt == nil {
				patchee.CreatedAt = &timestamp.Timestamp{}
			}
			childMask := &field_mask1.FieldMask{}
			for j := i; j < len(updateMask.Paths); j++ {
				if trimPath := strings.TrimPrefix(updateMask.Paths[j], prefix+"CreatedAt."); trimPath != updateMask.Paths[j] {
					childMask.Paths = append(childMask.Paths, trimPath)
				}
			}
			if err := gorm2.MergeWithMask(patcher.CreatedAt, patchee.CreatedAt, childMask); err != nil {
				return nil, nil
			}
		}
		if f == prefix+"CreatedAt" {
			updatedCreatedAt = true
			patchee.CreatedAt = patcher.CreatedAt
			continue
		}
		if !updatedUpdatedAt && strings.HasPrefix(f, prefix+"UpdatedAt.") {
			if patcher.UpdatedAt == nil {
				patchee.UpdatedAt = nil
				continue
			}
			if patchee.UpdatedAt == nil {
				patchee.UpdatedAt = &timestamp.Timestamp{}
			}
			childMask := &field_mask1.FieldMask{}
			for j := i; j < len(updateMask.Paths); j++ {
				if trimPath := strings.TrimPrefix(updateMask.Paths[j], prefix+"UpdatedAt."); trimPath != updateMask.Paths[j] {
					childMask.Paths = append(childMask.Paths, trimPath)
				}
			}
			if err := gorm2.MergeWithMask(patcher.UpdatedAt, patchee.UpdatedAt, childMask); err != nil {
				return nil, nil
			}
		}
		if f == prefix+"UpdatedAt" {
			updatedUpdatedAt = true
			patchee.UpdatedAt = patcher.UpdatedAt
			continue
		}
		if !updatedDeletedAt && strings.HasPrefix(f, prefix+"DeletedAt.") {
			if patcher.DeletedAt == nil {
				patchee.DeletedAt = nil
				continue
			}
			if patchee.DeletedAt == nil {
				patchee.DeletedAt = &timestamp.Timestamp{}
			}
			childMask := &field_mask1.FieldMask{}
			for j := i; j < len(updateMask.Paths); j++ {
				if trimPath := strings.TrimPrefix(updateMask.Paths[j], prefix+"DeletedAt."); trimPath != updateMask.Paths[j] {
					childMask.Paths = append(childMask.Paths, trimPath)
				}
			}
			if err := gorm2.MergeWithMask(patcher.DeletedAt, patchee.DeletedAt, childMask); err != nil {
				return nil, nil
			}
		}
		if f == prefix+"DeletedAt" {
			updatedDeletedAt = true
			patchee.DeletedAt = patcher.DeletedAt
			continue
		}
		if !updatedBoardId && strings.HasPrefix(f, prefix+"BoardId.") {
			if patcher.BoardId == nil {
				patchee.BoardId = nil
				continue
			}
			if patchee.BoardId == nil {
				patchee.BoardId = &wrappers.StringValue{}
			}
			childMask := &field_mask1.FieldMask{}
			for j := i; j < len(updateMask.Paths); j++ {
				if trimPath := strings.TrimPrefix(updateMask.Paths[j], prefix+"BoardId."); trimPath != updateMask.Paths[j] {
					childMask.Paths = append(childMask.Paths, trimPath)
				}
			}
			if err := gorm2.MergeWithMask(patcher.BoardId, patchee.BoardId, childMask); err != nil {
				return nil, nil
			}
		}
		if f == prefix+"BoardId" {
			updatedBoardId = true
			patchee.BoardId = patcher.BoardId
			continue
		}
		if !updatedPostId && strings.HasPrefix(f, prefix+"PostId.") {
			if patcher.PostId == nil {
				patchee.PostId = nil
				continue
			}
			if patchee.PostId == nil {
				patchee.PostId = &wrappers.StringValue{}
			}
			childMask := &field_mask1.FieldMask{}
			for j := i; j < len(updateMask.Paths); j++ {
				if trimPath := strings.TrimPrefix(updateMask.Paths[j], prefix+"PostId."); trimPath != updateMask.Paths[j] {
					childMask.Paths = append(childMask.Paths, trimPath)
				}
			}
			if err := gorm2.MergeWithMask(patcher.PostId, patchee.PostId, childMask); err != nil {
				return nil, nil
			}
		}
		if f == prefix+"PostId" {
			updatedPostId = true
			patchee.PostId = patcher.PostId
			continue
		}
		if !updatedContentId && strings.HasPrefix(f, prefix+"ContentId.") {
			if patcher.ContentId == nil {
				patchee.ContentId = nil
				continue
			}
			if patchee.ContentId == nil {
				patchee.ContentId = &wrappers.StringValue{}
			}
			childMask := &field_mask1.FieldMask{}
			for j := i; j < len(updateMask.Paths); j++ {
				if trimPath := strings.TrimPrefix(updateMask.Paths[j], prefix+"ContentId."); trimPath != updateMask.Paths[j] {
					childMask.Paths = append(childMask.Paths, trimPath)
				}
			}
			if err := gorm2.MergeWithMask(patcher.ContentId, patchee.ContentId, childMask); err != nil {
				return nil, nil
			}
		}
		if f == prefix+"ContentId" {
			updatedContentId = true
			patchee.ContentId = patcher.ContentId
			continue
		}
		if !updatedUserid && strings.HasPrefix(f, prefix+"Userid.") {
			if patcher.Userid == nil {
				patchee.Userid = nil
				continue
			}
			if patchee.Userid == nil {
				patchee.Userid = &wrappers.StringValue{}
			}
			childMask := &field_mask1.FieldMask{}
			for j := i; j < len(updateMask.Paths); j++ {
				if trimPath := strings.TrimPrefix(updateMask.Paths[j], prefix+"Userid."); trimPath != updateMask.Paths[j] {
					childMask.Paths = append(childMask.Paths, trimPath)
				}
			}
			if err := gorm2.MergeWithMask(patcher.Userid, patchee.Userid, childMask); err != nil {
				return nil, nil
			}
		}
		if f == prefix+"Userid" {
			updatedUserid = true
			patchee.Userid = patcher.Userid
			continue
		}
		if !updatedUsername && strings.HasPrefix(f, prefix+"Username.") {
			if patcher.Username == nil {
				patchee.Username = nil
				continue
			}
			if patchee.Username == nil {
				patchee.Username = &wrappers.StringValue{}
			}
			childMask := &field_mask1.FieldMask{}
			for j := i; j < len(updateMask.Paths); j++ {
				if trimPath := strings.TrimPrefix(updateMask.Paths[j], prefix+"Username."); trimPath != updateMask.Paths[j] {
					childMask.Paths = append(childMask.Paths, trimPath)
				}
			}
			if err := gorm2.MergeWithMask(patcher.Username, patchee.Username, childMask); err != nil {
				return nil, nil
			}
		}
		if f == prefix+"Username" {
			updatedUsername = true
			patchee.Username = patcher.Username
			continue
		}
		if !updatedNickname && strings.HasPrefix(f, prefix+"Nickname.") {
			if patcher.Nickname == nil {
				patchee.Nickname = nil
				continue
			}
			if patchee.Nickname == nil {
				patchee.Nickname = &wrappers.StringValue{}
			}
			childMask := &field_mask1.FieldMask{}
			for j := i; j < len(updateMask.Paths); j++ {
				if trimPath := strings.TrimPrefix(updateMask.Paths[j], prefix+"Nickname."); trimPath != updateMask.Paths[j] {
					childMask.Paths = append(childMask.Paths, trimPath)
				}
			}
			if err := gorm2.MergeWithMask(patcher.Nickname, patchee.Nickname, childMask); err != nil {
				return nil, nil
			}
		}
		if f == prefix+"Nickname" {
			updatedNickname = true
			patchee.Nickname = patcher.Nickname
			continue
		}
		if !updatedEmail && strings.HasPrefix(f, prefix+"Email.") {
			if patcher.Email == nil {
				patchee.Email = nil
				continue
			}
			if patchee.Email == nil {
				patchee.Email = &wrappers.StringValue{}
			}
			childMask := &field_mask1.FieldMask{}
			for j := i; j < len(updateMask.Paths); j++ {
				if trimPath := strings.TrimPrefix(updateMask.Paths[j], prefix+"Email."); trimPath != updateMask.Paths[j] {
					childMask.Paths = append(childMask.Paths, trimPath)
				}
			}
			if err := gorm2.MergeWithMask(patcher.Email, patchee.Email, childMask); err != nil {
				return nil, nil
			}
		}
		if f == prefix+"Email" {
			updatedEmail = true
			patchee.Email = patcher.Email
			continue
		}
		if !updatedPassword && strings.HasPrefix(f, prefix+"Password.") {
			if patcher.Password == nil {
				patchee.Password = nil
				continue
			}
			if patchee.Password == nil {
				patchee.Password = &wrappers.StringValue{}
			}
			childMask := &field_mask1.FieldMask{}
			for j := i; j < len(updateMask.Paths); j++ {
				if trimPath := strings.TrimPrefix(updateMask.Paths[j], prefix+"Password."); trimPath != updateMask.Paths[j] {
					childMask.Paths = append(childMask.Paths, trimPath)
				}
			}
			if err := gorm2.MergeWithMask(patcher.Password, patchee.Password, childMask); err != nil {
				return nil, nil
			}
		}
		if f == prefix+"Password" {
			updatedPassword = true
			patchee.Password = patcher.Password
			continue
		}
		if !updatedUrl && strings.HasPrefix(f, prefix+"Url.") {
			if patcher.Url == nil {
				patchee.Url = nil
				continue
			}
			if patchee.Url == nil {
				patchee.Url = &wrappers.StringValue{}
			}
			childMask := &field_mask1.FieldMask{}
			for j := i; j < len(updateMask.Paths); j++ {
				if trimPath := strings.TrimPrefix(updateMask.Paths[j], prefix+"Url."); trimPath != updateMask.Paths[j] {
					childMask.Paths = append(childMask.Paths, trimPath)
				}
			}
			if err := gorm2.MergeWithMask(patcher.Url, patchee.Url, childMask); err != nil {
				return nil, nil
			}
		}
		if f == prefix+"Url" {
			updatedUrl = true
			patchee.Url = patcher.Url
			continue
		}
		if !updatedUseHtml && strings.HasPrefix(f, prefix+"UseHtml.") {
			if patcher.UseHtml == nil {
				patchee.UseHtml = nil
				continue
			}
			if patchee.UseHtml == nil {
				patchee.UseHtml = &wrappers.BoolValue{}
			}
			childMask := &field_mask1.FieldMask{}
			for j := i; j < len(updateMask.Paths); j++ {
				if trimPath := strings.TrimPrefix(updateMask.Paths[j], prefix+"UseHtml."); trimPath != updateMask.Paths[j] {
					childMask.Paths = append(childMask.Paths, trimPath)
				}
			}
			if err := gorm2.MergeWithMask(patcher.UseHtml, patchee.UseHtml, childMask); err != nil {
				return nil, nil
			}
		}
		if f == prefix+"UseHtml" {
			updatedUseHtml = true
			patchee.UseHtml = patcher.UseHtml
			continue
		}
		if !updatedUseSecret && strings.HasPrefix(f, prefix+"UseSecret.") {
			if patcher.UseSecret == nil {
				patchee.UseSecret = nil
				continue
			}
			if patchee.UseSecret == nil {
				patchee.UseSecret = &wrappers.BoolValue{}
			}
			childMask := &field_mask1.FieldMask{}
			for j := i; j < len(updateMask.Paths); j++ {
				if trimPath := strings.TrimPrefix(updateMask.Paths[j], prefix+"UseSecret."); trimPath != updateMask.Paths[j] {
					childMask.Paths = append(childMask.Paths, trimPath)
				}
			}
			if err := gorm2.MergeWithMask(patcher.UseSecret, patchee.UseSecret, childMask); err != nil {
				return nil, nil
			}
		}
		if f == prefix+"UseSecret" {
			updatedUseSecret = true
			patchee.UseSecret = patcher.UseSecret
			continue
		}
		if f == prefix+"UpVoteCount" {
			patchee.UpVoteCount = patcher.UpVoteCount
			continue
		}
		if f == prefix+"DownVoteCount" {
			patchee.DownVoteCount = patcher.DownVoteCount
			continue
		}
	}
	if err != nil {
		return nil, err
	}
	return patchee, nil
}

// DefaultListComment executes a gorm list call
func DefaultListComment(ctx context.Context, db *gorm1.DB) ([]*Comment, error) {
	in := Comment{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(CommentORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	db, err = gorm2.ApplyCollectionOperators(ctx, db, &CommentORM{}, &Comment{}, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(CommentORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("id")
	ormResponse := []CommentORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(CommentORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*Comment{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type CommentORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type CommentORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type CommentORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm1.DB, *[]CommentORM) error
}