package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()

	r.GET("/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "Hello Gin!",
		})
	})

	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
