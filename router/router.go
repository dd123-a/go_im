package router

import (
	"github.com/gin-gonic/gin"
	"wenqianIm/api"
	"wenqianIm/conf"
	"wenqianIm/service"
)

func NewROuter() *gin.Engine {
	conf.Init()
	r:=gin.Default()
	r.Use(gin.Recovery(),gin.Logger())
	v1:=r.Group("/")
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(200,"SUCCESS")
		})
		v1.POST("user/register",api.UserRegister)
		v1.GET("ws",service.WsHandler)
	}
	return r
}