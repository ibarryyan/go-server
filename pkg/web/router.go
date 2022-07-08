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
	num := r.Group("/num")
	{
		num.POST("/add/:key", controller.NewNumInfoControllerImpl().AddNumByKey)
		num.GET("/findByKey/:key", controller.NewNumInfoControllerImpl().FindNumByKey)
		num.GET("/findById/:id", controller.NewNumInfoControllerImpl().FindNumById)
		num.POST("/saveInfo", controller.NewNumInfoControllerImpl().SaveNumInfo)
		num.POST("/deleteInfo/:id", controller.NewNumInfoControllerImpl().DeleteById)
		num.GET("/getAll", controller.NewNumInfoControllerImpl().FindAll)
		num.POST("/update", controller.NewNumInfoControllerImpl().Update)
	}

	roleInfo := r.Group("/role")
	{
		roleInfo.POST("/save", controller.NewRoleControllerImpl().CreateRole)
		roleInfo.GET("/all", controller.NewRoleControllerImpl().GetAll)
	}

	userInfo := r.Group("/user")
	{
		userInfo.POST("/save", controller.NewUserController().CreateUser)
		userInfo.POST("/login", controller.NewUserController().FindUserByLoginNameAndPwd)
		userInfo.POST("/register", controller.NewUserController().Register)
	}

	auth := r.Group("/auth")
	{
		auth.POST("/add", controller.AddPolicy)
	}

	r.Run("127.0.0.1:" + config.PORT)
}
