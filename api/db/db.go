package db

import (
	"time"

	"github.com/google/uuid"

	"github.com/dgkg/goworkshop/api/model"
)

func NewDB() *DB {
	var db DB
	db.Users = make(map[string]model.User)

	u := model.User{
		FirstName: "Henri",
		LastName:  "Lepic",
	}
	db.AddUser(&u)

	return &db
}

type DB struct {
	Users map[string]model.User
}

func (db *DB) AddUser(u *model.User) *model.User {
	u.DateOfBirth = time.Now()
	u.UUID = uuid.New().String()
	db.Users[u.UUID] = *u
	return u
}
