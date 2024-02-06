package database

import (
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetConnection() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"))
	if err != nil {
		panic(err)
	}
	if err := db.Use(otelgorm.NewPlugin()); err != nil {
		panic(err)
	}
	return db
}
