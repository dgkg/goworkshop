package main

import (
	"time"

	"github.com/google/uuid"
)

var us map[string]User

func init() {
	us = make(map[string]User)
	u := User{
		FirstName: "Henri",
		LastName:  "Lepic",
	}
	AddUser(&u)
}

type User struct {
	UUID        string
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	DateOfBirth time.Time `json:"birth_date"`
}

func (u User) String() string {
	return u.FirstName + " " + u.LastName
}

func AddUser(u *User) *User {
	u.DateOfBirth = time.Now()
	u.UUID = uuid.New().String()
	us[u.UUID] = *u
	return u
}
