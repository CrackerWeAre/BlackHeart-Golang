package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/ssoyyoung.p/BlackHeart-Golang/mysql"
	"github.com/ssoyyoung.p/BlackHeart-Golang/utils"
)

// GetOrderList func
func GetOrderList(c *gin.Context) {
	uID := c.Query("uID")
	uIDint, _ := strconv.Atoi(uID)

	if uID == "" {
		c.JSON(http.StatusOK, utils.JSONReturnMsg(
			false, "uID는 필수 파라미터 입니다."))
		return
	}

	exist := mysql.CheckExistByID("order_list", uIDint)

	if !exist {
		c.JSON(http.StatusOK, utils.JSONReturnMsg(
			false, "해당 uID의 주문 내역이 존재하지 않습니다"))
		return
	}

	orderList := mysql.GetOrderList(uID)

	for idx, order := range orderList {
		pID := strconv.Itoa(order.PID)
		product := mysql.GetProductItem(pID)

		orderList[idx].Product = product
	}

	c.JSON(http.StatusOK, utils.JSONReturnResult(
		true, orderList,
	))
	return

}
