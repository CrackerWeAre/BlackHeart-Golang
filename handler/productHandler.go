package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ssoyyoung.p/BlackHeart-Golang/mysql"

	"github.com/ssoyyoung.p/BlackHeart-Golang/utils"
)

func GetProductItem(c *gin.Context) {
	pID := c.Query("pID")

	productItem := mysql.GetProductItem(pID)
	c.JSON(http.StatusOK, utils.JSONReturnResult(
		true, productItem,
	))
	return

}