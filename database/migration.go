package database

import (
	"gorm.io/gorm"
)

func MigrateUp(db *gorm.DB) {
	CreateTopupTableUp(db)
}
