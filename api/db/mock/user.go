package mock

import (
	"errors"
	"time"

	"github.com/dgkg/goworkshop/api/model"
	"github.com/google/uuid"
)

func (db *Mock) AddUser(u *model.User) (*model.User, error) {
	u.BirthDate = time.Now()
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

func (db *Mock) GetAllUser() ([]model.User, error) {
	us := make([]model.User, len(db.Users))
	for k := range db.Users {
		us = append(us, db.Users[k])
	}
	return us, nil
}
