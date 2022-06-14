package controller

import (
	"count_num/pkg/dao/impl"
	"count_num/pkg/entity"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"strconv"
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
	numInfo := impl.dao.GetNumInfoByKey(c, key)
	numInfo.InfoNum = numInfo.InfoNum + 1
	isOk := impl.dao.UpdateNumInfoByKey(c, numInfo)
	c.JSON(200, isOk)
}

func (impl NumInfoControllerImpl) FindNumByKey(c *gin.Context) {
	key := c.Param("key")
	numInfo := impl.dao.GetNumInfoByKey(c, key)
	c.JSON(200, numInfo)
}

func (impl NumInfoControllerImpl) SaveNumInfo(c *gin.Context) {
	body := c.Request.Body
	bytes, err := ioutil.ReadAll(body)
	info := entity.NumInfo{}
	json.Unmarshal(bytes, &info)
	if err != nil {
		panic(err)
	}
	isOk := impl.dao.AddNumInfo(c, info)
	c.JSON(200, isOk)
}

func (impl NumInfoControllerImpl) DeleteById(c *gin.Context) {
	id := c.Param("id")
	i, _ := strconv.Atoi(id)
	isOk := impl.dao.DeleteNumInfoById(c, int64(i))
	c.JSON(200, isOk)
}

func (impl NumInfoControllerImpl) FindAll(c *gin.Context) {
	numInfos := impl.dao.FindAllNumInfo(c)
	c.JSON(200, numInfos)
}
