package main

import (
	"danbeltoken/routing"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/user", routing.GetUser)

	r.Run("0.0.0.0:8080")
}
