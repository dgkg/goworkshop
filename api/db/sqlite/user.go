package sqlite

import "github.com/dgkg/goworkshop/api/model"

func (sqlite *SQLite) AddUser(u *model.User) (*model.User, error) {
	sqlite.DB.Create(u)
	return u, nil
}

func (sqlite *SQLite) DeleteUser(uuid string) error {
	var u model.User
	u.UUID = uuid
	return sqlite.DB.Delete(&u).Error
}
func (sqlite *SQLite) GetByIDUser(uuid string) (u *model.User, e error) {
	sqlite.DB.First(u, "uuid = ?", uuid)
	// sqlite.DB.Where(&model.User{
	// 	moUUID: uuid,
	// }).First(&u)
	return
}
func (sqlite *SQLite) GetAllUser() (us []model.User, e error) {
	sqlite.DB.Find(&us)
	return
}
