package main

import (
	"danbeltoken/routing"
	"danbeltoken/variables"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	variables.Init()
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.html", gin.H{})
	})

	r.POST("/login", routing.Login)

	r.GET("/error", func(ctx *gin.Context) {
		err := ctx.Query("err")
		ctx.HTML(http.StatusOK, "error.html", gin.H{"Error": err})
	})

	r.GET("/user-page", routing.UserPage)
	r.POST("/referral-code", routing.SetRefCode)
	r.POST("/nft", routing.CreateNft)

	r.Run("0.0.0.0:1212")
}
