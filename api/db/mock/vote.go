package mock

import (
	"errors"

	"github.com/dgkg/goworkshop/api/model"
	"github.com/google/uuid"
)

func (db *Mock) AddVote(v *model.Vote) (*model.Vote, error) {
	v.UUID = uuid.New().String()
	db.Votes[v.UUID] = *v
	return v, nil
}

func (db *Mock) DeleteVote(uuid string) error {
	if _, ok := db.Votes[uuid]; !ok {
		return errors.New("db: Vote not found")
	}
	delete(db.Votes, uuid)
	return nil
}

func (db *Mock) GetByIDVote(uuid string) (*model.Vote, error) {
	u, ok := db.Votes[uuid]
	if !ok {
		return nil, errors.New("db: Vote not found")
	}
	return &u, nil
}

func (db *Mock) GetAllVote() ([]model.Vote, error) {
	us := make([]model.Vote, len(db.Votes))
	for k := range db.Votes {
		us = append(us, db.Votes[k])
	}
	return us, nil
}
