package database

import (
	"fmt"
	"os"
	"strings"
	"time"

	// "github.com/jinzhu/gorm"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"

	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	configPB "github.com/ygpark2/mboard/shared/proto/config"
)

// GetDatabaseConnection return (gorm.DB or error)
func GetDatabaseConnection(dbConf configPB.DatabaseConfiguration) (db *gorm.DB, err error) {
	var timezoneCommand string

	dsn, err := dbConf.DSN()
	if err != nil {
		return nil, err
	}

	// gLogger := log.With().Str("module", "gorm").Logger()

	zlog := log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC822}).With().Caller().Logger()
	newLogger := logger.New(
		&zlog,
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)

	conf := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_",                              // table name prefix, table for `User` would be `t_users`
			SingularTable: true,                              // use singular table name, table for `User` would be `user` with this option enabled
			NameReplacer:  strings.NewReplacer("CID", "Cid"), // use name replacer to change struct/field name before convert it to db name
		},
		Logger: newLogger.LogMode(logger.Info), // gormlog.NewGormLogger(gLogger),
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
		DryRun:      false,
		PrepareStmt: false,
	}

	switch dbConf.Dialect {
	case configPB.DatabaseDialect_SQLite3:
		db, err = gorm.Open(sqlite.Open(dsn), conf)
	case configPB.DatabaseDialect_Postgre:
		timezoneCommand = "SET timezone = 'UTC'"
		db, err = gorm.Open(postgres.Open(dsn), conf)
	case configPB.DatabaseDialect_MySQL:
		timezoneCommand = "SET time_zone = '+00:00'"
		db, err = gorm.Open(mysql.Open(dsn), conf)
	default:
		return nil, fmt.Errorf("database dialect %s not supported", dbConf.Dialect)
	}

	if err != nil {
		return
	}

	if dbConf.Logging {
		db.Debug()
	}

	sqlDB, err := db.DB()

	sqlDB.SetMaxOpenConns(int(dbConf.MaxOpenConns))
	sqlDB.SetMaxIdleConns(int(dbConf.MaxIdleConns))
	sqlDB.SetConnMaxLifetime(*dbConf.ConnMaxLifetime)

	if dbConf.Utc {
		if _, err = sqlDB.Exec(timezoneCommand); err != nil {
			return nil, fmt.Errorf("error setting UTC timezone: %w", err)
		}
	}

	return
}
