package repo

import (
	"iris_demo/datasource"
	"iris_demo/model"
)

type UserRepository interface {
	GetUserList() *[]model.User
	SaveUser(book model.User)(err error)
	GetUserById(id uint)(book model.User,err error)
	DeleteUser(id uint)(err error)
	GetUserByName(name string )*[]model.User
}

func NewUserRepository() UserRepository{
	return &userRepository{}
}

var db = datasource.GetDB()

type userRepository struct {}

func (self userRepository) GetUserList() *[]model.User {
	user:= new([]model.User)
	err:=db.Raw(`select * FROM user`).Scan(user).RowsAffected
	if err > 0 {
		return user
	}else{
		return nil
	}
}

func (self userRepository) GetUserByName(name string ) *[]model.User {
	user:= new([]model.User)
	err:=db.Raw(`select * FROM user where user.name = ?`,name).Scan(user).RowsAffected
	if err > 0 {
		return user
	}else{
		return nil
	}
}


func (self userRepository) SaveUser(user model.User) (err error) {
	if user.ID != 0{
		err := db.Save(&user).Error
		return err
	}else {
		err := db.Create(&user).Error
		return err
	}
}
func (self userRepository) GetUserById(id uint) (user model.User,err error) {
	err = db.First(&user,id).Error
	return user,err
}
func (self userRepository) DeleteUser(id uint) (err error) {
	user:= new(model.User)
	user.ID = id
	err = db.Unscoped().Delete(&user).Error
	return err
}
