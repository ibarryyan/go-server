package controller

import (
	"count_num/pkg/dao/impl"
	"count_num/pkg/model"
	"count_num/pkg/utils"
	"count_num/pkg/web/auth"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

type UserControllerImpl struct {
	dao *impl.UserDaoImpl
}

type UserController interface {
	CreateUser(c *gin.Context)
	FindUserByLoginNameAndPwd(c *gin.Context)
	Register(c *gin.Context)
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

func (impl UserControllerImpl) FindUserByLoginNameAndPwd(c *gin.Context) {
	body := c.Request.Body
	bytes, err := ioutil.ReadAll(body)
	user := model.User{}
	json.Unmarshal(bytes, &user)
	if err != nil {
		panic(err)
	}
	userByLoginName := impl.dao.GetUserByLoginName(c, user.LoginName)
	//密码通过
	if userByLoginName.Pwd == utils.GetMd5Str(user.Pwd) {
		setToken := auth.SetToken(c, utils.GetTokenStr(), user)
		c.JSON(200, map[string]interface{}{"code": 0, "msg": setToken, "count": 0, "data": utils.GetTokenStr()})
	} else {
		if userByLoginName.Id == 0 {
			c.JSON(200, map[string]interface{}{"code": 0, "msg": "账号不存在", "count": 0, "data": "-1"})
		} else {
			c.JSON(200, map[string]interface{}{"code": 0, "msg": "密码错误", "count": 0, "data": "-1"})
		}
	}
}

func (impl UserControllerImpl) Register(c *gin.Context) {
	body := c.Request.Body
	bytes, err := ioutil.ReadAll(body)
	user := model.User{}
	json.Unmarshal(bytes, &user)
	if err != nil {
		panic(err)
	}
	user.Role = 1
	res := impl.dao.CreateUser(c, user)
	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": res})
}
