package controller

import (
	"context"
	"count_num/pkg/dao/impl"
	"count_num/pkg/entity"
	"github.com/gin-gonic/gin"
)

func AddNumByKey(c *gin.Context) {
	key := c.Param("key")
	daoImpl := impl.NewCountNumDAOImpl()
	info := daoImpl.GetNumInfoByKey(context.Background(), key)
	info = entity.NumInfo{info.Id, info.Name, info.InfoKey, info.InfoNum + 1}
	daoImpl.UpdateNumInfoByKey(context.Background(), info)
	c.JSON(200, key)
}

func FindNumByKey(c *gin.Context) {
	key := c.Param("key")
	daoImpl := impl.NewCountNumDAOImpl()
	infoByKey := daoImpl.GetNumInfoByKey(context.Background(), key)
	c.JSON(200, infoByKey)
}
