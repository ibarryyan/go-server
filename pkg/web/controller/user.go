package controller

import (
	"count_num/pkg/dao/impl"
	"github.com/gin-gonic/gin"
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

}
