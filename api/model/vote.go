package model

import (
	"time"

	"github.com/google/uuid"
)

type Vote struct {
	DBModel
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
}

func (v *Vote) BeforeCreate() (err error) {
	v.UUID = uuid.New().String()
	return nil
}
