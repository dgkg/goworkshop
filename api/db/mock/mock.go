package mock

import (
	"errors"
	"time"

	"github.com/google/uuid"

	"github.com/dgkg/goworkshop/api/db"
	"github.com/dgkg/goworkshop/api/model"
)

func New() db.DB {
	var db Mock
	db.Users = make(map[string]model.User)

	u := model.User{
		FirstName: "Henri",
		LastName:  "Lepic",
	}
	db.AddUser(&u)

	return &db
}

type Mock struct {
	Users map[string]model.User
}

func (db *Mock) AddUser(u *model.User) (*model.User, error) {
	u.DateOfBirth = time.Now()
	u.UUID = uuid.New().String()
	db.Users[u.UUID] = *u
	return u, nil
}

func (db *Mock) DeleteUser(uuid string) error {
	if _, ok := db.Users[uuid]; !ok {
		return errors.New("db: user not found")
	}
	delete(db.Users, uuid)
	return nil
}

func (db *Mock) GetByIDUser(uuid string) (*model.User, error) {
	u, ok := db.Users[uuid]
	if !ok {
		return nil, errors.New("db: user not found")
	}
	return &u, nil
}

func (db *Mock) GetAllUser() (map[string]model.User, error) {
	return db.Users, nil
}
