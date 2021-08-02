package repository

import (
	"github.com/micro/micro/v3/service/logger"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"

	comment_entities "github.com/ygpark2/mboard/service/comment/proto/entities"
)

// commentRepository interface
type CommentRepository interface {
	Exist(model *comment_entities.CommentORM) bool
	List(limit, page int, sort string, model *comment_entities.CommentORM) (total int64, users []*comment_entities.CommentORM, err error)
	Get(id string) (*comment_entities.CommentORM, error)
	Create(model *comment_entities.CommentORM) error
	Update(id string, model *comment_entities.CommentORM) error
	Delete(model *comment_entities.CommentORM) error
}

// commentRepository struct
type commentRepository struct {
	db *gorm.DB
}

// NewCommentRepository returns an instance of `CommentRepository`.
func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepository{
		db: db,
	}
}

// Exist
func (repo *commentRepository) Exist(model *comment_entities.CommentORM) bool {
	log.Info().Msgf("Received commentRepository.Exist request %v", *model)
	var count int64
	if model.Username != nil && len(*model.Username) > 0 {
		repo.db.Model(&comment_entities.CommentORM{}).Where("username = ?", model.Username).Count(&count)
		if count > 0 {
			return true
		}
	}
	if len(model.Id.String()) > 0 {
		repo.db.Model(&comment_entities.CommentORM{}).Where("id = ?", model.Id.String()).Count(&count)
		if count > 0 {
			return true
		}
	}
	if model.Email != "" {
		repo.db.Model(&comment_entities.CommentORM{}).Where("email = ?", model.Email).Count(&count)
		if count > 0 {
			return true
		}
	}
	return false
}

// List
func (repo *commentRepository) List(limit, page int, sort string, model *comment_entities.CommentORM) (total int64, users []*comment_entities.CommentORM, err error) {
	db := repo.db

	if limit == 0 {
		limit = 10
	}
	var offset int
	if page > 1 {
		offset = (page - 1) * limit
	} else {
		offset = 0
	}
	if sort == "" {
		sort = "created_at desc"
	}

	if model.Username != nil && len(*model.Username) > 0 {
		db = db.Where("username like ?", "%"+*model.Username+"%")
	}
	if model.FirstName != "" {
		db = db.Where("first_name like ?", "%"+model.FirstName+"%")
	}
	if model.LastName != "" {
		db = db.Where("last_name like ?", "%"+model.LastName+"%")
	}
	if model.Email != "" {
		db = db.Where("email like ?", "%"+model.Email+"%")
	}
	// enable auto preloading for `Profile`
	if err = db.Set("gorm:auto_preload", true).Order(sort).Limit(limit).Offset(offset).Find(&users).Count(&total).Error; err != nil {
		log.Error().Err(err).Msg("Error in commentRepository.List")
		return
	}
	return
}

// Find by ID
func (repo *commentRepository) Get(id string) (user *comment_entities.CommentORM, err error) {
	u2, err := uuid.FromString(id)
	if err != nil {
		return
	}
	user = &comment_entities.CommentORM{Id: u2}
	// enable auto preloading for `Profile`
	if err = repo.db.Set("gorm:auto_preload", true).First(user).Error; err != nil && err != gorm.ErrRecordNotFound {
		log.Error().Err(err).Msg("Error in commentRepository.Get")
	}
	return
}

// Create
func (repo *commentRepository) Create(model *comment_entities.CommentORM) error {
	if exist := repo.Exist(model); exist {
		return errors.New("comment already exist")
	}
	// if err := repo.db.Set("gorm:association_autoupdate", false).Create(model).Error; err != nil {
	if err := repo.db.Create(model).Error; err != nil {
		log.Error().Err(err).Msg("Error in commentRepository.Create")
		return err
	}
	return nil
}

// Update TODO: Translation
func (repo *commentRepository) Update(id string, model *comment_entities.CommentORM) error {
	u2, err := uuid.FromString(id)
	if err != nil {
		return err
	}
	user := &comment_entities.CommentORM{
		Id: u2,
	}
	// result := repo.db.Set("gorm:association_autoupdate", false).Save(model)
	result := repo.db.Model(user).Updates(model)
	if err := result.Error; err != nil {
		log.Error().Err(err).Msg("Error in commentRepository.Update")
		return err
	}
	if rowsAffected := result.RowsAffected; rowsAffected == 0 {
		log.Error().Msgf("Error in commentRepository.Update, rowsAffected: %v", rowsAffected)
		return errors.New("no records updated, No match was found")
	}
	return nil
}

// Delete
func (repo *commentRepository) Delete(model *comment_entities.CommentORM) error {
	result := repo.db.Delete(model)
	if err := result.Error; err != nil {
		log.Error().Err(err).Msg("Error in commentRepository.Delete")
		return err
	}
	if rowsAffected := result.RowsAffected; rowsAffected == 0 {
		log.Error().Msgf("Error in commentRepository.Delete, rowsAffected: %v", rowsAffected)
		return errors.New("no records deleted, No match was found")
	}
	return nil
}
