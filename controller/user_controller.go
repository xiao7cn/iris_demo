package controller

import (
	"github.com/kataras/iris/v12/context"
	"github.com/spf13/cast"
	"iris_demo/model"
	"iris_demo/service"
	"log"
)

type UserController struct {
	Service service.UserService
}

func NewUserController() *UserController {
	return &UserController{Service:service.NewUserService()}
}

//查询所有
func (self *UserController) GetList() (result *model.Result) {
	return self.Service.GetUserList()
}

//保存 and 修改
func (self *UserController) PostSaveUser(ctx context.Context) (result model.Result)  {
	user := new(model.User)
	err := ctx.ReadForm(user)
	if err != nil{
		log.Println(err)
		result.Msg = "数据有错误"
		return
	}
	return self.Service.PostSaveUser(*user)
}

//根据id查询
func (self *UserController) GetUserById(ctx context.Context) (result model.Result)  {
	id := ctx.PostValue("id")
	if id == "" {
		result.Code = 400
		result.Msg = "缺少参数id"
		return
	}
	return self.Service.GetUserById(cast.ToUint(id))
}

//根据id删除
func (self *UserController) PostDelUser(ctx context.Context) (result model.Result)  {
	id := ctx.PostValue("id")
	if id == "" {
		result.Code = 400
		result.Msg = "缺少参数id"
		return
	}
	return self.Service.DelUser(cast.ToUint(id))
}
