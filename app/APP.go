package app

import (
	"ginyii/config"
	Middlewares "ginyii/middlewares"
	"ginyii/routes"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func Run(HttpServer *gin.Engine)  {
	commonConf := config.GetCommonConf()

	// Gin服务
	HttpServer = gin.Default()

	// 捕捉接口运行耗时（必须排第一）
	HttpServer.Use(Middlewares.StatLatency)

	// 设置全局ctx参数（必须排第二）
	HttpServer.Use(Middlewares.AppData)

	// 拦截应用500报错，使之可视化
	HttpServer.Use(Middlewares.AppError500)

	// Gin运行时：release、debug、test
	gin.SetMode(commonConf["ENV"])

	// 注册必要路由，处理默认路由、静态文件路由、404路由等
	routes.RouteMust(HttpServer)

	// web页面路由
	routes.Web(HttpServer)

	//API 路由
	routes.Api(HttpServer)

	// 初始化定时器（立即运行定时器）
	//Task.TimeInterval(0, 0, "0")

	// 实际访问网址和端口
	_host := "127.0.0.1:" + commonConf["PORT"] // 测试访问IP
	host := commonConf["HOST"] + ":" + commonConf["PORT"] // Docker访问IP

	// 终端提示
	log.Println(
		" \n " +
		"访问地址示例：" + host + " >>> \n " +
		"APP主页：http://" + _host + " \n " +
		"")

	err := HttpServer.Run(host)
	if err != nil {
		log.Println("http服务遇到错误，运行中断，error：", err.Error())
		log.Println("提示：注意端口被占时应该首先更改对外暴露的端口，而不是微服务的端口。")
		os.Exit(200)
	}

	return
}