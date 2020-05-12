package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	DBModel
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	BirthDate time.Time `json:"birth_date"`
}

func (u User) String() string {
	return u.FirstName + " " + u.LastName
}

func (u *User) BeforeCreate() (err error) {
	u.BirthDate = time.Now()
	u.UUID = uuid.New().String()
	return nil
}

type DBModel struct {
	UUID      string `gorm:"primary_key" json:"uuid"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
