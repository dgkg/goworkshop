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
	db.AutoMigrate(&model.User{})
	return &SQLite{
		DB: db,
	}
}

func (sqlite *SQLite) AddUser(u *model.User) (*model.User, error) {
	sqlite.DB.Create(u)
	return u, nil
}

func (sqlite *SQLite) DeleteUser(uuid string) error {
	var u model.User
	u.UUID = uuid
	return sqlite.DB.Delete(&u).Error
}
func (sqlite *SQLite) GetByIDUser(uuid string) (u *model.User, e error) {
	sqlite.DB.First(u, "uuid = ?", uuid)
	return
}
func (sqlite *SQLite) GetAllUser() (us []model.User, e error) {
	sqlite.DB.Find(&us)
	return
}
