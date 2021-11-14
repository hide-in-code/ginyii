package site

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

//app主页
func Index(c *gin.Context) {
	c.HTML(200, "site/index", nil)
}

//测试致命错误
func Fetal(c *gin.Context) {
	viewFile := []string{}
	log.Println(viewFile[10])
	c.JSON(
		http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"data":    "affaa",
			"message": "ok",
		},
	)
}

//json-test
func Test(c *gin.Context) {
	c.JSON(
		http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"data":    "aafg",
			"message": "ok",
		},
	)
}

//登录页面
func Login(c *gin.Context) {
	c.HTML(200, "site/login", nil)
}

func Signin(c *gin.Context)  {
	c.HTML(200, "site/signin", nil)
}

func About(c *gin.Context)  {
	c.HTML(200, "site/about", nil)
}

func Contact(c *gin.Context)  {
	c.HTML(200, "site/contact", nil)
}