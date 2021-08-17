package registry

import (
	// "github.com/asim/go-micro/v3/logger"
	"github.com/micro/micro/v3/service/logger"
	"github.com/ygpark2/njro/service/board/repository"
	"github.com/ygpark2/njro/shared/database"

	boardPB "github.com/ygpark2/njro/service/board/proto/board"
	board_entities "github.com/ygpark2/njro/service/board/proto/entities"
	configPB "github.com/ygpark2/njro/shared/proto/config"
)

// Container - provide di Container
type Container struct {
	BoardRepository repository.BoardRepository
	BoardHandler    boardPB.BoardServiceHandler
}

func BuildBoardRepository(cfg configPB.Configuration) repository.BoardRepository {
	// db := database.GetDatabaseConnection(*cfg.Database).(*gorm.DB)

	db, err := database.GetDatabaseConnection(*cfg.Database)
	if db != nil {
		logger.Debug("============== Create board entity ==============")
		db.AutoMigrate(&board_entities.BoardORM{})
	} else {
		logger.Errorf("++++++++++++++ DB Connection Error +++++++++++++ %s", err)
	}

	return repository.NewBoardRepository(db)
}
