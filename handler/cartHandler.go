package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/ssoyyoung.p/BlackHeart-Golang/mysql"
	"github.com/ssoyyoung.p/BlackHeart-Golang/utils"
)


func GetCartList(c *gin.Context) {
	uID := c.Query("uID")
	uIDint, _ := strconv.Atoi(uID)

	if uID == "" {
		c.JSON(http.StatusOK, utils.JSONReturnMsg(
			false, "uID는 필수 파라미터 입니다."))
		return
	}

	exist := mysql.CheckExistByID("cart_list", uIDint)

	if !exist {
		c.JSON(http.StatusOK, utils.JSONReturnMsg(
			false, "해당 유저의 카드 내역이 존재하지 않습니다.",
			))
		return
	}

	cartList := mysql.GetCartList(uID)

	for idx, cart := range cartList {
		pID := strconv.Itoa(cart.PID)
		product := mysql.GetProductItem(pID)

		cartList[idx].Product = product
	}

	c.JSON(http.StatusOK, utils.JSONReturnResult(
		true, cartList,
		))
	return
}