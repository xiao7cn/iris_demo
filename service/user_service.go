package service

import (
	"iris_demo/model"
	"iris_demo/repo"
)
type UserService interface {
	GetUserList () *model.Result
	PostSaveUser(user model.User) (result model.Result)
	GetUserById(id uint) (result model.Result)
	DelUser(id uint) (result model.Result)
}

type userService struct {}

func NewUserService() UserService {
	return &userService{}
}

var userRepo = repo.NewUserRepository()

func (self userService) GetUserList() *model.Result {
	books := userRepo.GetUserList()
	result := new (model.Result)
	result.Data = books
	result.Code = 200
	result.Msg ="SUCCESS"
	return result
}
func (self userService) PostSaveUser(user model.User) (result model.Result) {
	err := userRepo.SaveUser(user)
	if err != nil{
		result.Code = 400
		result.Msg = err.Error()
	}else{
		result.Code = 200
		result.Msg ="SUCCESS"
		user := userRepo.GetUserByName(user.Name)
		result.Data = user
	}
	return
}

func (self userService) GetUserById(id uint) (result model.Result) {
	user,err := userRepo.GetUserById(id)
	if err!= nil{
		result.Code = 400
		result.Msg = err.Error()
	}else{
		result.Data = user
		result.Code = 200
		result.Msg ="SUCCESS"
	}
	return result
}

func (self userService) DelUser(id uint) (result model.Result) {
	err := userRepo.DeleteUser(id)
	if err!= nil{
		result.Code = 400
		result.Msg = err.Error()
	}else{
		result.Code = 200
		result.Msg ="SUCCESS"
		list := userRepo.GetUserList()
		result.Data = list

	}
	return
}