package model

import (
	"time"
)

type User struct {
	UUID        string    `json:"uuid"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	DateOfBirth time.Time `json:"birth_date"`
}

func (u User) String() string {
	return u.FirstName + " " + u.LastName
}
