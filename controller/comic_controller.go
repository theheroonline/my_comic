package controller

import (
	"fmt"
	"go_comic/model"
	"go_comic/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var ComicCtl = new(comicCtl)

type comicCtl struct{}

func (c *comicCtl) Page(ctx *gin.Context) {
	pageno, _ := strconv.Atoi(ctx.Query("currentPage"))
	pagesize, _ := strconv.Atoi(ctx.Query("pageSize"))
	search := ctx.Query("queryString")

	comics, count := service.ComicService.GetComicList(pageno, pagesize, search)
	// ctx.HTML(http.StatusOK, "index.html", model.Result{
	// 	Data: &comics,
	// 	Code: 200,
	// 	Msg:  "成功",
	// })
	m := make(map[string]interface{})
	m["total"] = count
	m["data"] = comics

	ctx.JSON(http.StatusOK, model.Result{
		Data: &m,
		Code: 200,
		Msg:  "成功",
	})
}

func (c *comicCtl) Detail(ctx *gin.Context) {
	pageno, _ := strconv.Atoi(ctx.Query("currentPage"))
	pagesize, _ := strconv.Atoi(ctx.Query("pageSize"))
	id, _ := strconv.Atoi(ctx.Query("id"))
	fmt.Printf("id: %v\n", id)
	fmt.Printf("pageno: %v\n", pageno)
	fmt.Printf("pagesize: %v\n", pagesize)
	pics, count := service.ComicService.GetComicById(pageno, pagesize, id)
	m := make(map[string]interface{})
	m["total"] = count
	m["data"] = pics
	// ctx.HTML(http.StatusOK, "detail.html", model.Result{
	// 	Data: &m,
	// 	Code: 200,
	// 	Msg:  "成功",
	// })
	ctx.JSON(http.StatusOK, model.Result{
		Data: &m,
		Code: 200,
		Msg:  "成功",
	})
}

func (c *comicCtl) Href(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "detail.html", gin.H{
		"id": ctx.Query("id"),
	})
}

func (c *comicCtl) Shuaxin(ctx *gin.Context) {
	service.ComicService.Shuaxin()
	ctx.JSON(http.StatusOK, "刷新成功")
}
