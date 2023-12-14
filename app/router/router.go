package router

import (
	"github.com/gin-gonic/gin"
	"ldap-server/app/controller"
)

func InitRouter() *gin.Engine {
	//设置模式
	gin.SetMode(gin.ReleaseMode)
	// 创建带有默认中间件的路由:
	// 日志与恢复中间件
	r := gin.Default()
	r.GET("/", controller.GetList)
	return r
}
