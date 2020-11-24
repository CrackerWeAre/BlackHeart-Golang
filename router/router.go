package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/ssoyyoung.p/BlackHeart-Golang/handler"
)

// Router func
func Router() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())

	// User API Router
	userAPI := r.Group("/user")
	{
		userAPI.GET("/list", handler.GetUserList)
		userAPI.POST("create", handler.CreateUser)
	}

	return r
}
