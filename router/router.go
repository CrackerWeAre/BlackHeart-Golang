package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ssoyyoung.p/BlackHeart-Golang/handler"
)

// Router func
func Router() *gin.Engine {
	r := gin.Default()

	// User API Router
	userAPI := r.Group("/user")
	{
		userAPI.GET("/list/:page/:maxResult", handler.GetUserList)
	}

	return r
}
