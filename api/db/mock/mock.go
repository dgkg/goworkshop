package mock

import (
	"github.com/dgkg/goworkshop/api/db"
	"github.com/dgkg/goworkshop/api/model"
)

type Mock struct {
	Users map[string]model.User
	Votes map[string]model.Vote
}

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
