package sqlite

import "github.com/dgkg/goworkshop/api/model"

func (sqlite *SQLite) AddVote(u *model.Vote) (*model.Vote, error) {
	sqlite.DB.Create(u)
	return u, nil
}

func (sqlite *SQLite) DeleteVote(uuid string) error {
	var u model.Vote
	u.UUID = uuid
	return sqlite.DB.Delete(&u).Error
}
func (sqlite *SQLite) GetByIDVote(uuid string) (u *model.Vote, e error) {
	sqlite.DB.First(u, "uuid = ?", uuid)
	return
}
func (sqlite *SQLite) GetAllVote() (us []model.Vote, e error) {
	sqlite.DB.Find(&us)
	return
}
