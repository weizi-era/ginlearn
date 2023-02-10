package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type User struct {
	Id         string            `form:"id"`
	Name       string            `form:"name"`
	Address    []string          `form:"address" binding:"required"`
	AddressMap map[string]string `form:"addressMap"`
}

func main() {

	r := gin.Default()

	r.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"name": "helloworld",
		})
	})

	r.GET("/get", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"type": "get",
		})
	})

	r.POST("/save", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"type": "save",
		})
	})

	r.DELETE("/delete", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"type": "delete",
		})
	})

	r.PUT("/update", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"type": "update",
		})
	})

	/*r.GET("/user/find/:id", func(ctx *gin.Context) {
		ctx.JSON(200, ctx.Param("id"))
	})*/

	/*r.GET("/user/*path", func(ctx *gin.Context) {
		ctx.JSON(200, ctx.Param("path"))
	})*/

	v1 := r.Group("/v1")
	{
		v1.GET("find", func(ctx *gin.Context) {
			ctx.JSON(200, "v1 find")
		})
	}

	v2 := r.Group("/v2")
	{
		v2.GET("find", func(ctx *gin.Context) {
			ctx.JSON(200, "v2 find")
		})
	}

	r.GET("/user/save", func(ctx *gin.Context) {

		//http://localhost:8080/user/save?id=1235&name=weizhishidai
		/*id := ctx.Query("id")
		name := ctx.Query("name")
		address := ctx.DefaultQuery("address", "taiyuan")
		ctx.JSON(200, gin.H{
			"id":      id,
			"name":    name,
			"address": address,
		})*/

		//http://localhost:8080/user/save?id=1235&name=weizhishidai&address=taiyuan
		/*var user User
		err := ctx.ShouldBindQuery(&user)
		if err != nil {
			log.Println(err)
		}
		ctx.JSON(200, user)*/

		//http://localhost:8080/user/save?address=Beijing&address=shanghai
		/*address := ctx.QueryArray("address")
		ctx.JSON(200, address)*/

		//http://localhost:8080/user/save?addressMap[home]=Beijing&addressMap[company]=shanghai
		addressMap := ctx.QueryMap("addressMap")
		ctx.JSON(200, addressMap)
	})

	//http://localhost:8080/user/post
	r.POST("/user/post", func(ctx *gin.Context) {

		var user User
		err := ctx.ShouldBindJSON(&user)
		if err != nil {
			log.Println(err)
		}

		/*id := ctx.PostForm("id")
		name := ctx.PostForm("name")
		address := ctx.PostFormArray("address")*/

		// map 参数无法通过ShouldBind获取
		//addressMap := ctx.PostFormMap("addressMap")
		//user.AddressMap = addressMap
		ctx.JSON(200, user)
	})

	//http://localhost:8080/user/upload  上传文件
	r.POST("/user/upload", func(ctx *gin.Context) {
		form, err := ctx.MultipartForm()
		if err != nil {
			log.Println(err)
		}
		value := form.Value
		files := form.File
		for _, fileArray := range files {
			for _, file := range fileArray {
				err := ctx.SaveUploadedFile(file, "./"+file.Filename)
				if err != nil {
					log.Println(err)
				}
			}
		}
		ctx.JSON(http.StatusOK, value)
	})

	/*r.GET("/user/return", func(ctx *gin.Context) {

		ctx.File("./logo.png")
	})*/

	err := r.Run(":8080")
	if err != nil {
		log.Println(err)
	}
}
