package web

import (
	"count_num/pkg/config"
	"count_num/pkg/web/controller"
	"count_num/pkg/web/interceptor"
	"github.com/gin-gonic/gin"
)

func RunHttp() {
	r := gin.Default()
	//增加拦截器
	r.Use(interceptor.HttpInterceptor())
	//解决跨域
	r.Use(config.CorsConfig())
	//路由组
	appInfoGroup := r.Group("/")
	{
		appInfoGroup.POST("/add/:key", controller.NewNumInfoControllerImpl().AddNumByKey)
		appInfoGroup.GET("/findByKey/:key", controller.NewNumInfoControllerImpl().FindNumByKey)
		appInfoGroup.GET("/findById/:id", controller.NewNumInfoControllerImpl().FindNumById)
		appInfoGroup.POST("/saveInfo", controller.NewNumInfoControllerImpl().SaveNumInfo)
		appInfoGroup.POST("/deleteInfo/:id", controller.NewNumInfoControllerImpl().DeleteById)
		appInfoGroup.GET("/getAll", controller.NewNumInfoControllerImpl().FindAll)
		appInfoGroup.POST("/update", controller.NewNumInfoControllerImpl().Update)
	}
	r.Run("127.0.0.1:" + config.PORT)
}
