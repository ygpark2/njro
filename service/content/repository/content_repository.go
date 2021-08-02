package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	uuid "github.com/satori/go.uuid"

	content_entities "github.com/ygpark2/mboard/service/content/proto/entities"
)

// contentRepository interface
type ContentRepository interface {
	Exist(model *content_entities.ContentORM) bool
	List(limit, page int, sort string, model *content_entities.ContentORM) (total int64, users []*content_entities.ContentORM, err error)
	Get(id string) (*content_entities.ContentORM, error)
	Create(model *content_entities.ContentORM) error
	Update(id string, model *content_entities.ContentORM) error
	Delete(model *content_entities.ContentORM) error
}

// contentRepository struct
type contentRepository struct {
	db *gorm.DB
}

// NewContentRepository returns an instance of `ContentRepository`.
func NewContentRepository(db *gorm.DB) ContentRepository {
	return &contentRepository{
		db: db,
	}
}

// Exist
func (repo *contentRepository) Exist(model *content_entities.ContentORM) bool {
	log.Info().Msgf("Received contentRepository.Exist request %v", *model)
	var count int64
	if model.Username != nil && len(*model.Username) > 0 {
		repo.db.Model(&content_entities.ContentORM{}).Where("username = ?", model.Username).Count(&count)
		if count > 0 {
			return true
		}
	}
	if len(model.Id.String()) > 0 {
		repo.db.Model(&content_entities.ContentORM{}).Where("id = ?", model.Id.String()).Count(&count)
		if count > 0 {
			return true
		}
	}
	if model.Email != "" {
		repo.db.Model(&content_entities.ContentORM{}).Where("email = ?", model.Email).Count(&count)
		if count > 0 {
			return true
		}
	}
	return false
}

// List
func (repo *contentRepository) List(limit, page int, sort string, model *content_entities.ContentORM) (total int64, users []*content_entities.ContentORM, err error) {
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
		log.Error().Err(err).Msg("Error in contentRepository.List")
		return
	}
	return
}

// Find by ID
func (repo *contentRepository) Get(id string) (user *content_entities.ContentORM, err error) {
	u2, err := uuid.FromString(id)
	if err != nil {
		return
	}
	user = &content_entities.ContentORM{Id: u2}
	// enable auto preloading for `Profile`
	if err = repo.db.Set("gorm:auto_preload", true).First(user).Error; err != nil && err != gorm.ErrRecordNotFound {
		log.Error().Err(err).Msg("Error in contentRepository.Get")
	}
	return
}

// Create
func (repo *contentRepository) Create(model *content_entities.ContentORM) error {
	if exist := repo.Exist(model); exist {
		return errors.New("comment already exist")
	}
	// if err := repo.db.Set("gorm:association_autoupdate", false).Create(model).Error; err != nil {
	if err := repo.db.Create(model).Error; err != nil {
		log.Error().Err(err).Msg("Error in contentRepository.Create")
		return err
	}
	return nil
}

// Update TODO: Translation
func (repo *contentRepository) Update(id string, model *content_entities.ContentORM) error {
	u2, err := uuid.FromString(id)
	if err != nil {
		return err
	}
	user := &content_entities.ContentORM{
		Id: u2,
	}
	// result := repo.db.Set("gorm:association_autoupdate", false).Save(model)
	result := repo.db.Model(user).Updates(model)
	if err := result.Error; err != nil {
		log.Error().Err(err).Msg("Error in contentRepository.Update")
		return err
	}
	if rowsAffected := result.RowsAffected; rowsAffected == 0 {
		log.Error().Msgf("Error in contentRepository.Update, rowsAffected: %v", rowsAffected)
		return errors.New("no records updated, No match was found")
	}
	return nil
}

// Delete
func (repo *contentRepository) Delete(model *content_entities.ContentORM) error {
	result := repo.db.Delete(model)
	if err := result.Error; err != nil {
		log.Error().Err(err).Msg("Error in contentRepository.Delete")
		return err
	}
	if rowsAffected := result.RowsAffected; rowsAffected == 0 {
		log.Error().Msgf("Error in contentRepository.Delete, rowsAffected: %v", rowsAffected)
		return errors.New("no records deleted, No match was found")
	}
	return nil
}
