package repo

import (
	"example/template/internal/adapters/repo/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(models.User{})
	if err != nil {
		panic("failed to migrate database")
	}
}
