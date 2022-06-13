package controller

import (
	"count_num/pkg/dao/impl"
	"fmt"
	"github.com/gin-gonic/gin"
)

type NumInfoControllerImpl struct {
	dao *impl.CountNumDAOImpl
}

type NumInfoController interface {
	AddNumByKey(c *gin.Context)
	FindNumByKey(c *gin.Context)
	SaveNumInfo(c *gin.Context)
	DeleteById(c *gin.Context)
	FindAll(c *gin.Context)
}

func NewNumInfoControllerImpl() *NumInfoControllerImpl {
	return &NumInfoControllerImpl{dao: impl.NewCountNumDAOImpl()}
}

func (impl NumInfoControllerImpl) AddNumByKey(c *gin.Context) {
	key := c.Param("key")
	fmt.Println(impl.dao)
	c.JSON(200, key)
}

func (impl NumInfoControllerImpl) FindNumByKey(c *gin.Context) {
	key := c.Param("key")

	c.JSON(200, key)
}

func (impl NumInfoControllerImpl) SaveNumInfo(c *gin.Context) {
	key := c.Param("key")

	c.JSON(200, key)
}

func (impl NumInfoControllerImpl) DeleteById(c *gin.Context) {
	key := c.Param("key")

	c.JSON(200, key)
}

func (impl NumInfoControllerImpl) FindAll(c *gin.Context) {
	key := c.Param("key")

	c.JSON(200, key)
}
