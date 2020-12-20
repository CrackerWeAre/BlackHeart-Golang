package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/ssoyyoung.p/BlackHeart-Golang/mysql"

	"github.com/ssoyyoung.p/BlackHeart-Golang/utils"
)

func GetProductItem(c *gin.Context) {
	pID := c.Query("pID")

	query := "SELECT count(*) FROM product_list where pID="+pID
	exist := mysql.CheckExist(query)

	if !exist {
		c.JSON(http.StatusOK, utils.JSONReturnMsg(
			false, "존재하지 않는 pID 입니다",
		))
		return
	}

	productItem := mysql.GetProductItem(pID)
	c.JSON(http.StatusOK, utils.JSONReturnResult(
		true, productItem,
	))
	return

}

// GetProductList func
func GetProductList(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	maxResult, _ := strconv.Atoi(c.Query("maxResult"))

	allProducts := mysql.GetProductList()

	c.JSON(http.StatusOK, utils.PgSplitP(allProducts, page, maxResult))
}