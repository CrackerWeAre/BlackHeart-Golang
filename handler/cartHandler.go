package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ssoyyoung.p/BlackHeart-Golang/mysql"
	"github.com/ssoyyoung.p/BlackHeart-Golang/utils"
)

func GetCartList(c *gin.Context) {
	uID := c.Query("uID")

	query := "SELECT count(*) FROM cart_list WHERE uID="+uID
	exist := mysql.CheckExist(query)

	if !exist {
		c.JSON(http.StatusOK, utils.JSONReturnMsg(
			false, "해당 유저의 카드 내역이 존재하지 않습니다.",
			))
		return
	}

	cartList := mysql.GetCartList(uID)
	c.JSON(http.StatusOK, utils.JSONReturnResult(
		true, cartList,
		))
}