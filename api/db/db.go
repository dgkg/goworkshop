package db

import (
	"github.com/dgkg/goworkshop/api/model"
)

type DB interface {
	AddUser(u *model.User) (*model.User, error)
	DeleteUser(uuid string) error
	GetByIDUser(uuid string) (*model.User, error)
	GetAllUser() (map[string]model.User, error)
}
