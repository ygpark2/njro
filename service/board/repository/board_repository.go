package repository

import (
	"github.com/asim/go-micro/v3/logger"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"

	"github.com/ygpark2/njro/shared/database"

	board_entities "github.com/ygpark2/njro/service/board/proto/entities"
	configPB "github.com/ygpark2/njro/shared/proto/config"
)

// boardRepository interface
type BoardRepository interface {
	Exist(model *board_entities.BoardORM) bool
	List(limit, page int, sort string, model *board_entities.BoardORM) (total int64, boards []*board_entities.BoardORM, err error)
	Get(id string) (*board_entities.BoardORM, error)
	Create(model *board_entities.BoardORM) error
	Update(id string, model *board_entities.BoardORM) error
	Delete(model *board_entities.BoardORM) error
}

// boardRepository struct
type boardRepository struct {
	db *gorm.DB
}

func BuildBoardRepository(cfg configPB.Configuration) BoardRepository {
	// db := database.GetDatabaseConnection(*cfg.Database).(*gorm.DB)

	db, err := database.GetDatabaseConnection(*cfg.Database)
	if err != nil {
		db.AutoMigrate(&board_entities.BoardORM{})
	}

	return NewBoardRepository(db)
}

// NewBoardRepository returns an instance of `BoardRepository`.
func NewBoardRepository(db *gorm.DB) BoardRepository {
	return &boardRepository{
		db: db,
	}
}

// Exist
func (repo *boardRepository) Exist(model *board_entities.BoardORM) bool {
	logger.Infof("Received boardRepository.Exist request %v", *model)
	var count int64
	if model.Title != nil && len(*model.Title) > 0 {
		repo.db.Model(&board_entities.BoardORM{}).Where("title = ?", model.Title).Count(&count)
		if count > 0 {
			return true
		}
	}
	if len(model.Id.String()) > 0 {
		repo.db.Model(&board_entities.BoardORM{}).Where("id = ?", model.Id.String()).Count(&count)
		if count > 0 {
			return true
		}
	}
	if *model.MobileTitle != "" {
		repo.db.Model(&board_entities.BoardORM{}).Where("mobile_title = ?", model.MobileTitle).Count(&count)
		if count > 0 {
			return true
		}
	}
	return false
}

// List
func (repo *boardRepository) List(limit, page int, sort string, model *board_entities.BoardORM) (total int64, boards []*board_entities.BoardORM, err error) {
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
	// enable auto preloading for `Profile`
	if err = db.Set("gorm:auto_preload", true).Order(sort).Limit(limit).Offset(offset).Find(&boards).Count(&total).Error; err != nil {
		logger.Error("Error in boardRepository.List")
		return
	}
	return
}

// Find by ID
func (repo *boardRepository) Get(id string) (user *board_entities.BoardORM, err error) {
	u2, err := uuid.FromString(id)
	if err != nil {
		return
	}
	user = &board_entities.BoardORM{Id: u2}
	// enable auto preloading for `Profile`
	if err = repo.db.Set("gorm:auto_preload", true).First(user).Error; err != nil && err != gorm.ErrRecordNotFound {
		logger.Error("Error in boardRepository.Get")
	}
	return
}

// Create
func (repo *boardRepository) Create(model *board_entities.BoardORM) error {
	if exist := repo.Exist(model); exist {
		return errors.New("board already exist")
	}
	// if err := repo.db.Set("gorm:association_autoupdate", false).Create(model).Error; err != nil {
	if err := repo.db.Create(model).Error; err != nil {
		logger.Error("Error in boardRepository.Create")
		return err
	}
	return nil
}

// Update TODO: Translation
func (repo *boardRepository) Update(id string, model *board_entities.BoardORM) error {
	u2, err := uuid.FromString(id)
	if err != nil {
		return err
	}
	user := &board_entities.BoardORM{
		Id: u2,
	}
	// result := repo.db.Set("gorm:association_autoupdate", false).Save(model)
	result := repo.db.Model(user).Updates(model)
	if err := result.Error; err != nil {
		logger.Error("Error in boardRepository.Update")
		return err
	}
	if rowsAffected := result.RowsAffected; rowsAffected == 0 {
		logger.Errorf("Error in boardRepository.Update, rowsAffected: %v", rowsAffected)
		return errors.New("no records updated, No match was found")
	}
	return nil
}

// Delete
func (repo *boardRepository) Delete(model *board_entities.BoardORM) error {
	result := repo.db.Delete(model)
	if err := result.Error; err != nil {
		logger.Error("Error in boardRepository.Delete")
		return err
	}
	if rowsAffected := result.RowsAffected; rowsAffected == 0 {
		logger.Errorf("Error in boardRepository.Delete, rowsAffected: %v", rowsAffected)
		return errors.New("no records deleted, No match was found")
	}
	return nil
}
