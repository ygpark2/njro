package registry

import (
	"github.com/asim/go-micro/v3/logger"
	"github.com/ygpark2/njro/service/comment/repository"
	"github.com/ygpark2/njro/shared/database"

	commentPB "github.com/ygpark2/njro/service/comment/proto/comment"
	comment_entities "github.com/ygpark2/njro/service/comment/proto/entities"
	configPB "github.com/ygpark2/njro/shared/proto/config"
)

// Container - provide di Container
type Container struct {
	CommentRepository repository.CommentRepository
	CommentHandler    commentPB.CommentServiceHandler
}

func BuildCommentRepository(cfg configPB.Configuration) repository.CommentRepository {
	// db := database.GetDatabaseConnection(*cfg.Database).(*gorm.DB)

	db, err := database.GetDatabaseConnection(*cfg.Database)
	if db != nil {
		logger.Debug("============== Create board entity ==============")
		db.AutoMigrate(&comment_entities.CommentORM{})
	} else {
		logger.Errorf("++++++++++++++ DB Connection Error +++++++++++++ %s", err)
	}

	return repository.NewCommentRepository(db)
}
