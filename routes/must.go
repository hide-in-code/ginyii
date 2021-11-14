package routes

import (
	"ginyii/app/controller/site"
	"ginyii/common/helpers"
	Middlewares "ginyii/middlewares"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

const  AppBasePath string = "./app"
const  StaticFilePathPrefix string = AppBasePath + "/static"
const  ViewFilePathPrefix string = AppBasePath + "/views"

//自动添加路由
//创建视图符合渲染， site/index => ./views/site/index.html
//所有的视图文件都必须是html后缀
//所有布局文件都要放到layouts里面
//提前必须 r.LoadHTMLGlob("./views/**/*")，因为模板的继承关系
func createViewsRender() multitemplate.Renderer {
	render := multitemplate.NewRenderer()

	//布局文件
	layouts, err := filepath.Glob(ViewFilePathPrefix + "/layouts/*.html") //todo conf设置
	if err != nil {
		panic(err.Error())
	}

	//views文件
	includes, err := filepath.Glob(ViewFilePathPrefix + "/**/*.html") //todo conf设置
	if err != nil {
		panic(err.Error())
	}

	// 为layouts/和includes/目录生成 templates map
	for _, include := range includes {
		pathInfo := strings.Split(include, "/")
		vlen := len(pathInfo)
		viewName := pathInfo[vlen - 2] + "/" + strings.Replace(pathInfo[vlen - 1], ".html", "", -1)
		if pathInfo[vlen - 2] == "layouts" { //跳过layouts
			continue
		}

		viewFile := []string{include}
		files := append(viewFile, layouts...)
		render.AddFromFiles(viewName, files...)
	}

	return render
}

//加载必须的路由
func RouteMust(route *gin.Engine) {
	route.StaticFS("/js", http.Dir(StaticFilePathPrefix+ "/js"))
	route.StaticFS("/css", http.Dir(StaticFilePathPrefix+ "/css"))
	route.StaticFS("/fonts", http.Dir(StaticFilePathPrefix+ "/fonts"))

	//动态加载模板文件，必须是两级目录 site/index => /views/site/index.html 必须是html结尾
	route.HTMLRender = createViewsRender()
	route.Any("/", site.Index)           //app主页
	route.StaticFile("/favicon.ico", StaticFilePathPrefix + "/favicon.ico")
	route.GET("/site/test", site.Test)   //app主页

	// 404路由
	route.NoRoute(Middlewares.HttpCorsApi, Middlewares.HttpLimiter(2), func (ctx *gin.Context) {
		var url string = ctx.Request.Host + ctx.Request.URL.Path
		var IP string = ctx.ClientIP()
		log.Println("404路由 >>> " + url, IP)
		ctx.JSONP(http.StatusNotFound, gin.H{
			"state": 404,
			"msg": "GINYII：未定义此名称的路由名",
			"content": map[string]interface{}{
				"url": url,
				"time": helpers.GetTimeDate("Y-m-d H:i:s.ms.ns"),
			},
		})
	})


}