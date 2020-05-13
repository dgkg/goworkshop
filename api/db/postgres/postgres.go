package postgres

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/dgkg/goworkshop/api/db"
	"github.com/dgkg/goworkshop/api/db/sqlite"
	"github.com/dgkg/goworkshop/api/model"
)

// Postgres is the connector for postgres.
type Postgres = sqlite.SQLite

// DBCredential is use to connect to the database.
type DBCredential struct {
	Host, Port, DBname, User, Password string
}

// New is the constructor for a Postgres database.
func New(c *DBCredential) db.DB {
	login := fmt.Sprintf(
		"host=%v port=%v user=%v dbname=%v password=%v sslmode=disable",
		c.Host, c.Port, c.User, c.DBname, c.Password)
	db, err := gorm.Open("postgres", login)
	if err != nil {
		panic(fmt.Sprintf("postgres: failed to connect database %v", err))
	}
	// Migrate the schema
	db.AutoMigrate(&model.User{}, &model.Vote{})
	return &Postgres{
		DB: db,
	}
}
