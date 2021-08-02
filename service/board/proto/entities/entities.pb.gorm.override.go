package entities

import (
	// "context"
	// "time"

	// "github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

/**
 * GORM
**/

// BeforeCreate implements the GORM BeforeCreate interface for the UserORM type.
// you can use this method to generate new UUID for CREATE operation or let database create it with this annotation:
// {type: "uuid", primary_key: true, not_null:true, default: "uuid_generate_v4()"}];
// we prefer First method as it works with both SQLite & PostgreSQL
func (m *BoardORM) BeforeCreate(tx *gorm.DB) (err error) {
	uuid := uuid.NewV4()
	tx.Statement.SetColumn("Id", uuid.String())
	return nil
}
