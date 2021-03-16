package migrations

import (
	"apple/models"

	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

func m1815217031CreateArticlesTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "1815217031",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&models.Article{}).Error
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.DropTable("Apple_API").Error
		},
	}
}
