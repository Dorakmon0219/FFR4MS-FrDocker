package router

import (
	"gitee.com/zengtao321/frdocker/settings"
	"gitee.com/zengtao321/frdocker/web/controller/admin"
	"gitee.com/zengtao321/frdocker/web/controller/command"
	"gitee.com/zengtao321/frdocker/web/controller/container"
	"gitee.com/zengtao321/frdocker/web/controller/logs"
	"gitee.com/zengtao321/frdocker/web/controller/perf"
	"gitee.com/zengtao321/frdocker/web/controller/user"
	"gitee.com/zengtao321/frdocker/web/filter"
	"gitee.com/zengtao321/frdocker/web/swagger"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	gin.SetMode(settings.RUNNING_MODE)
	r := gin.Default()
	r.Use(filter.CorsFilter())
	r.Use(filter.UserAuthFilter())
	command.RegisterRouter(r)
	user.RegisterRouter(r)
	container.RegisterRouter(r)
	perf.RegisterRouter(r)
	admin.RegisterRouter(r)
	swagger.RegisterRouter(r)
	logs.RegisterRouter(r)
	return r
}
