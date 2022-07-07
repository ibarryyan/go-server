package controller

import (
	"bytes"
	"count_num/pkg/dao/impl"
	"count_num/pkg/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
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
	FindNumById(c *gin.Context)
	Update(context *gin.Context)
}

func NewNumInfoControllerImpl() *NumInfoControllerImpl {
	return &NumInfoControllerImpl{dao: impl.NewCountNumDAOImpl()}
}

func (impl NumInfoControllerImpl) AddNumByKey(c *gin.Context) {
	key := c.Param("key")
	numInfo := impl.dao.GetNumInfoByKey(c, key)
	numInfo.InfoNum = numInfo.InfoNum + 1
	isOk := impl.dao.UpdateNumInfoByKey(c, numInfo)
	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": isOk})
}

func (impl NumInfoControllerImpl) FindNumByKey(c *gin.Context) {
	key := c.Param("key")
	numInfo := impl.dao.GetNumInfoByKey(c, key)
	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": numInfo})
}

func (impl NumInfoControllerImpl) SaveNumInfo(c *gin.Context) {
	body := c.Request.Body
	bytes, err := ioutil.ReadAll(body)
	info := model.NumInfo{}
	json.Unmarshal(bytes, &info)
	if err != nil {
		panic(err)
	}
	isOk := impl.dao.AddNumInfo(c, info)
	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": isOk})
}

func (impl NumInfoControllerImpl) DeleteById(c *gin.Context) {
	id := c.Param("id")
	i, _ := strconv.Atoi(id)
	isOk := impl.dao.DeleteNumInfoById(c, int64(i))
	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": isOk})
}

func (impl NumInfoControllerImpl) FindAll(c *gin.Context) {
	page := c.Query("page")
	limit := c.Query("limit")
	numInfos := impl.dao.FindAllNumInfo(c, cast.ToInt(page), cast.ToInt(limit))
	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": len(numInfos), "data": numInfos})
}

func (impl NumInfoControllerImpl) FindNumById(c *gin.Context) {
	id := c.Param("id")
	i, _ := strconv.Atoi(id)
	numInfo := impl.dao.GetNumInfoById(c, int64(i))
	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": numInfo})
}

func (impl NumInfoControllerImpl) Update(c *gin.Context) {
	body := c.Request.Body
	jsonBytes, err := ioutil.ReadAll(body)
	d := json.NewDecoder(bytes.NewReader(jsonBytes))
	d.UseNumber()
	m := make(map[string]string)
	d.Decode(&m)
	if err != nil {
		panic(err)
	}
	info := model.NumInfo{
		Id:      cast.ToInt64(m["id"]),
		Name:    m["name"],
		InfoKey: m["info_key"],
		InfoNum: cast.ToInt64(m["info_num"]),
	}
	isOk := impl.dao.UpdateNumInfoById(c, info)
	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": isOk})
}
