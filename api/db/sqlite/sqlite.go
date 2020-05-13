package sqlite

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/dgkg/goworkshop/api/db"
	"github.com/dgkg/goworkshop/api/model"
)

type SQLite struct {
	DB *gorm.DB
}

func New() db.DB {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	db.AutoMigrate(&model.User{}, &model.Vote{})
	return &SQLite{
		DB: db,
	}
}
