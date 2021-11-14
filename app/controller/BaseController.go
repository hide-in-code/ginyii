package controller

import "github.com/gin-gonic/gin"

type BaseController struct {
	Method  string
	Context *gin.Context
}

