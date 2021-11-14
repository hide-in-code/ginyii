package routes

import (
	"ginyii/app/controller/site"
	"github.com/gin-gonic/gin"
)

func Web(route *gin.Engine)  {
	route.GET("/site/index", site.Index)   //app主页
	route.GET("/site/login", site.Login)   //登录页
	route.GET("/site/signin", site.Signin)   //登录页
	route.GET("/site/about", site.About)   //登录页
	route.GET("/site/contact", site.Contact)   //登录页
}