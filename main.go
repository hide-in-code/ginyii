package main

import (
	"ginyii/app"
	"github.com/gin-gonic/gin"
)

var Engine *gin.Engine
func main()  {
	app.Run(Engine)
}