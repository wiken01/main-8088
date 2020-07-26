package main

import (
	"github.com/gin-gonic/gin"
)



func main() {
	r := gin.Default()
	r.GET("/", handler)
	r.Run("0.0.0.0:8088")
}

func handler(context *gin.Context) {
	context.JSON(200, gin.H{
		"hello": "world",
	})
	return
}
