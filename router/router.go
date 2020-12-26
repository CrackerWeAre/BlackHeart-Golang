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
		userAPI.GET("/detail", handler.GetUserDetail)
		userAPI.DELETE("/delete", handler.DeleteUser)
		userAPI.POST("/create", handler.CreateUser)
		userAPI.POST("/update", handler.UpdateUser)

		userAPI.POST("/login", handler.LoginUser)
		userAPI.GET("/checkExist", handler.CheckExist)

	}

	userAuth := r.Group("/auth")
	{
		userAuth.POST("/login", handler.Login)
		userAuth.GET("/logout", handler.Logout)
	}

	orderAPI := r.Group("/order")
	{
		orderAPI.GET("/list", handler.GetOrderList)
	}

	cartAPI := r.Group("/cart")
	{
		cartAPI.GET("/list", handler.GetCartList)
	}

	commentAPI := r.Group("/review")
	{
		commentAPI.GET("/board", handler.GetReviewBoardList)
		commentAPI.GET("/comment", handler.GetReviewCommentList)
	}

	productAPI := r.Group("/product")
	{
		productAPI.GET("/detail", handler.GetProductItem)
		productAPI.GET("/list", handler.GetProductList)
		productAPI.GET("/cate", handler.GetProductByCate)
	}

	return r
}
