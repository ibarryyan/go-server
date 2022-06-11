package web

import (
	"count_num/pkg/web/controller"
	"github.com/gin-gonic/gin"
)

func RunHttp() {
	r := gin.Default()
	//路由组
	appInfoGroup := r.Group("/")
	{
		appInfoGroup.GET("/add/:key", controller.AddNumByKey)
		appInfoGroup.GET("/find/:key", controller.FindNumByKey)
	}
	r.Run("127.0.0.1:8888")
}
