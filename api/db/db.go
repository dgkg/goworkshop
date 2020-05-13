package db

import (
	"github.com/dgkg/goworkshop/api/model"
)

type DB interface {
	DBUser
	DBVote
}

type DBUser interface {
	AddUser(u *model.User) (*model.User, error)
	DeleteUser(uuid string) error
	GetByIDUser(uuid string) (*model.User, error)
	GetAllUser() ([]model.User, error)
}

type DBVote interface {
	AddVote(u *model.Vote) (*model.Vote, error)
	DeleteVote(uuid string) error
	GetByIDVote(uuid string) (*model.Vote, error)
	GetAllVote() ([]model.Vote, error)
}
