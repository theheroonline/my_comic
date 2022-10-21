package routers

import (
	"fmt"
	"my_comic/controller"
	"net/http"
	"text/template"
	"time"

	"github.com/gin-gonic/gin"
)

// 日期格式转换
func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%d-%d", year, month, day)
}

func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	gin.SetMode("debug")

	// router.GET("/test", func(context *gin.Context) {
	// 	context.JSON(200, gin.H{"message": "test"})
	// })

	// 设置静态资源路由
	// router.Static("/resource", "./public/resource")
	// router.StaticFile("/favicon.ico", "./public/resource/images/favicon.ico")
	router.LoadHTMLGlob("template/*")
	router.Static("/static", "./static")

	//自定义模板方法
	router.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})

	// router.Use(service.HaveToken())

	/* 用户管理 */
	user := router.Group("/comic")
	{
		user.GET("/list", controller.ComicCtl.Page)
		user.GET("/href", controller.ComicCtl.Href)
		user.GET("/detail", controller.ComicCtl.Detail)
		user.GET("/shuaxin", controller.ComicCtl.Shuaxin)
	}

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	return router
}
