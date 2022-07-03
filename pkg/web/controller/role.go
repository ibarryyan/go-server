package controller

import (
	"count_num/pkg/dao/impl"
	"count_num/pkg/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

type RoleControllerImpl struct {
	dao *impl.RoleDaoImpl
}

type RoleController interface {
	CreateRole(c *gin.Context)
	GetAll(c *gin.Context)
}

func NewRoleControllerImpl() *RoleControllerImpl {
	return &RoleControllerImpl{dao: impl.NewRoleDaoImpl()}
}

func (impl RoleControllerImpl) CreateRole(c *gin.Context) {
	body := c.Request.Body
	bytes, err := ioutil.ReadAll(body)
	info := model.Role{}
	json.Unmarshal(bytes, &info)
	if err != nil {
		panic(err)
	}
	isOk := impl.dao.CreateRole(c, info)
	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": isOk})
}

func (impl RoleControllerImpl) GetAll(c *gin.Context) {
	roles := impl.dao.GetAll(c)
	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": len(roles), "data": roles})
}
