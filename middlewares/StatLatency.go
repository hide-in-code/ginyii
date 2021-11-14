package Middlewares
// 处理App运行时的一些必要事件

import (
	"fmt"
	"ginyii/common/helpers"
	"ginyii/config"
	"github.com/gin-gonic/gin"
	"time"
)

// StatLatency 捕捉接口运行耗时
// 此处使用的Next()请参考文档：https://blog.csdn.net/qq_37767455/article/details/104712028
func StatLatency(ctx *gin.Context) {
	start := float64(time.Now().UnixNano()) / 1000000 // ms
	ctx.Set("stat_start", helpers.GetTimeDate("i.s.ms.ns")) // 设置公共参数。此参数范围是每个请求的公共参数，不是超全局参数，超全局参数请用globalData。

	// 等其他中间件先执行
	ctx.Next()

	// 获取运行耗时，ms
	end := float64(time.Now().UnixNano()) / 1000000
	latency := end - start

	// 设置公共参数（设置ctx每次请求的全局值）
	ctx.Set("stat_latency", latency)
	//fmt.Println("本次运行耗时=", latency, "ms")

	// 进入耗时治理服务
	runtimeStatLatency(ctx)

	// 计时完成，中断所有后续函数调用
	ctx.Abort()
}

func runtimeStatLatency(ctx *gin.Context)  {

	_host, _ := ctx.Get("host")
	host := helpers.ValueInterfaceToString(_host)
	Uri := host + ctx.Request.RequestURI
	fmt.Println("请求uri=", Uri)

	_statLatency, _ := ctx.Get("stat_latency") // 获取ctx每次请求的全局值
	fmt.Println("接口耗时=", _statLatency, "ms")

	//err := ctx.Request.ParseMultipartForm(128)
	//if err != nil {
	//	fmt.Println("Request.Form为空=", err)
	//	//ctx.Abort()
	//} // 保存表单缓存的内存大小128M
	//data := ctx.Request.Form
	//fmt.Println(" 请求参数=", data)

	statLatency := helpers.StringToFloat(helpers.ValueInterfaceToString(_statLatency)) // ms
	if statLatency > 3*1000 { // 超过3s都记录下来
		helpers.Log(helpers.ValueInterfaceToString(_statLatency) + "ms；" + Uri, ctx.ClientIP())
	}

	ctx.Next()
}



// AppData 设置全局参数
func AppData(ctx *gin.Context) {
	serverConfig := config.GetCommonConf()
	host := serverConfig["HOST"] + ":" + serverConfig["PORT"]

	ctx.Set("host", host)

	ctx.Next()
}