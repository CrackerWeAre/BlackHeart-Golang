package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ssoyyoung.p/BlackHeart-Golang/mysql"
	"github.com/ssoyyoung.p/BlackHeart-Golang/utils"
)

// GetOrderList func
func GetOrderList(c *gin.Context) {
	uID := c.Query("uID")
	fmt.Println(uID)

	orderList := mysql.GetOrderList(uID)

	c.JSON(http.StatusOK, utils.JSONReturnResult(
		true, orderList,
	))
	return

}
