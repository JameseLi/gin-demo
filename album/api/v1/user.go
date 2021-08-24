package v1

import "github.com/gin-gonic/gin"

type UserInterface interface {
	GetUserInfo(context *gin.Context)
}
