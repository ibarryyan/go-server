package controller

import (
	"count_num/pkg/dao/impl"
	"count_num/pkg/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

type UserControllerImpl struct {
	dao *impl.UserDaoImpl
}

type UserController interface {
	CreateUser(c *gin.Context)
}

func NewUserController() *UserControllerImpl {
	return &UserControllerImpl{dao: impl.NewUserDaoImpl()}
}

func (impl UserControllerImpl) CreateUser(c *gin.Context) {
	body := c.Request.Body
	bytes, err := ioutil.ReadAll(body)
	user := model.User{}
	json.Unmarshal(bytes, &user)
	if err != nil {
		panic(err)
	}
	res := impl.dao.CreateUser(c, user)
	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": res})
}
