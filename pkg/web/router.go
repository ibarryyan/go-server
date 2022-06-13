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
		appInfoGroup.POST("/add/:key", controller.NewNumInfoControllerImpl().AddNumByKey)
		appInfoGroup.GET("/find/:key", controller.NewNumInfoControllerImpl().FindNumByKey)
		appInfoGroup.POST("/saveInfo", controller.NewNumInfoControllerImpl().SaveNumInfo)
		appInfoGroup.POST("/deleteInfo/:id", controller.NewNumInfoControllerImpl().DeleteById)
		appInfoGroup.GET("/getAll", controller.NewNumInfoControllerImpl().FindAll)
	}
	r.Run("127.0.0.1:8888")
}
