package main

import (
	"time"

	"github.com/google/uuid"
)

func NewDB() *DB {
	var db DB
	db.Users = make(map[string]User)

	u := User{
		FirstName: "Henri",
		LastName:  "Lepic",
	}
	db.AddUser(&u)

	return &db
}

type DB struct {
	Users map[string]User
}

func (db *DB) AddUser(u *User) *User {
	u.DateOfBirth = time.Now()
	u.UUID = uuid.New().String()
	db.Users[u.UUID] = *u
	return u
}
