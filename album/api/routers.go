package api

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
	r := gin.New()
	//panic 自动recover
	r.Use(gin.Recovery())

	return r
}
