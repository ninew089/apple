package migrations

import (
	"apple/config"
	"log"

	"gopkg.in/gormigrate.v1"
)

func Migrate() {
	db := config.GetDB()
	m := gormigrate.New(
		db,
		gormigrate.DefaultOptions,
		[]*gormigrate.Migration{
			m1815217031CreateArticlesTable(),
			m1925217031CreateArticlesTable(),
		},
	)

	if err := m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
}
