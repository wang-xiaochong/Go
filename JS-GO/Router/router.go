package router

import (
	"github.com/gin-gonic/gin"
	controller "Example/Controller"
	utils "Example/Utils"
	middleware "Example/Middleware"
)

// InitRouter 加入路由访问路径
func InitRouter(e *gin.Engine){
	e.Use(utils.CORS(utils.Options{Origin: "*"}))

	js:=e.Group("/js")
	{
		js.GET("/userInfo",controller.JsUser)
		js.GET("/lessonInfo",controller.JsLesson)
	}





	example:=e.Group("/example")
	{
		example.GET("/search",controller.Search)
		example.POST("/add",controller.Add)
		example.PUT("/update", controller.Update)
		example.DELETE("/delete", controller.Delete)
	}
	user:=e.Group("/user")
	{
		user.GET("/findAll",controller.FindAllUser)
		user.GET("/findByN",controller.FindAllUserByName)
		user.POST("/login",controller.Login)
		user.POST("/insert", controller.UserInsert)
		user.PUT("/update", controller.UserUpdate)
		user.DELETE("/delete", controller.UserDelete)
	}
	// 中间件拦截示例
	e.GET("/cookie",utils.SetCK)
	e.GET("/home",middleware.AuthMiddleWare(),controller.Home)
	e.GET("/home2",middleware.JWTAuthMiddleware(),middleware.TokenCheck(),controller.Home2)
}

// SetupRouter 初始化路由
func SetupRouter() *gin.Engine {
	r := gin.Default()
	return r
}