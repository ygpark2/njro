package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/micro/micro/v3/service/logger"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"

	"github.com/ygpark2/mboard/shared/database"

	post_entities "github.com/ygpark2/mboard/service/post/proto/entities"
	configPB "github.com/ygpark2/mboard/shared/proto/config"
)

// postRepository interface
type PostRepository interface {
	Exist(model *post_entities.PostORM) bool
	List(limit, page int, sort string, model *post_entities.PostORM) (total int64, boards []*post_entities.PostORM, err error)
	Get(id string) (*post_entities.PostORM, error)
	Create(model *post_entities.PostORM) error
	Update(id string, model *post_entities.PostORM) error
	Delete(model *post_entities.PostORM) error
}

// postRepository struct
type postRepository struct {
	db *gorm.DB
}

func BuildPostRepository(cfg configPB.Configuration) PostRepository {
	// db := database.GetDatabaseConnection(*cfg.Database).(*gorm.DB)

	db, err := database.GetDatabaseConnection(*cfg.Database)
	if err != nil {
		db.AutoMigrate(&post_entities.PostORM{})
	}

	return NewPostRepository(db)
}

// NewPostRepository returns an instance of `PostRepository`.
func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{
		db: db,
	}
}

// Exist
func (repo *postRepository) Exist(model *post_entities.PostORM) bool {
	logger.Infof("Received PostRepository.Exist request %v", *model)
	var count int64
	if model.Title != nil && len(*model.Title) > 0 {
		repo.db.Model(&post_entities.PostORM{}).Where("title = ?", model.Title).Count(&count)
		if count > 0 {
			return true
		}
	}
	if len(model.Id.String()) > 0 {
		repo.db.Model(&post_entities.PostORM{}).Where("id = ?", model.Id.String()).Count(&count)
		if count > 0 {
			return true
		}
	}
	if *model.MobileTitle != "" {
		repo.db.Model(&post_entities.PostORM{}).Where("mobile_title = ?", model.MobileTitle).Count(&count)
		if count > 0 {
			return true
		}
	}
	return false
}

// List
func (repo *postRepository) List(limit, page int, sort string, model *post_entities.PostORM) (total int64, boards []*post_entities.PostORM, err error) {
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

	if model.Title != nil && len(*model.Title) > 0 {
		db = db.Where("title like ?", "%"+*model.Title+"%")
	}
	if *model.MobileTitle != "" {
		db = db.Where("mobile_title like ?", "%"+*model.Title+"%")
	}
	if model.Description != "" {
		db = db.Where("description like ?", "%"+model.Description+"%")
	}
	// enable auto preloading for `Post`
	if err = db.Set("gorm:auto_preload", true).Order(sort).Limit(limit).Offset(offset).Find(&boards).Count(&total).Error; err != nil {
		logger.Error("Error in PostRepository.List")
		return
	}
	return
}

// Find by ID
func (repo *postRepository) Get(id string) (user *post_entities.PostORM, err error) {
	u2, err := uuid.FromString(id)
	if err != nil {
		return
	}
	user = &post_entities.PostORM{Id: u2}
	// enable auto preloading for `Profile`
	if err = repo.db.Set("gorm:auto_preload", true).First(user).Error; err != nil && err != gorm.ErrRecordNotFound {
		logger.Error("Error in PostRepository.Get")
	}
	return
}

// Create
func (repo *postRepository) Create(model *post_entities.PostORM) error {
	if exist := repo.Exist(model); exist {
		return errors.New("board already exist")
	}
	// if err := repo.db.Set("gorm:association_autoupdate", false).Create(model).Error; err != nil {
	if err := repo.db.Create(model).Error; err != nil {
		logger.Error("Error in PostRepository.Create")
		return err
	}
	return nil
}

// Update TODO: Translation
func (repo *postRepository) Update(id string, model *post_entities.PostORM) error {
	u2, err := uuid.FromString(id)
	if err != nil {
		return err
	}
	user := &post_entities.PostORM{
		Id: u2,
	}
	// result := repo.db.Set("gorm:association_autoupdate", false).Save(model)
	result := repo.db.Model(user).Updates(model)
	if err := result.Error; err != nil {
		logger.Error("Error in PostRepository.Update")
		return err
	}
	if rowsAffected := result.RowsAffected; rowsAffected == 0 {
		logger.Errorf("Error in PostRepository.Update, rowsAffected: %v", rowsAffected)
		return errors.New("no records updated, No match was found")
	}
	return nil
}

// Delete
func (repo *postRepository) Delete(model *post_entities.PostORM) error {
	result := repo.db.Delete(model)
	if err := result.Error; err != nil {
		logger.Error("Error in PostRepository.Delete")
		return err
	}
	if rowsAffected := result.RowsAffected; rowsAffected == 0 {
		logger.Errorf("Error in PostRepository.Delete, rowsAffected: %v", rowsAffected)
		return errors.New("no records deleted, No match was found")
	}
	return nil
}
