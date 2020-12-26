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
	pIDint, _ := strconv.Atoi(pID)

	exist := mysql.CheckExistByID("product_list", pIDint)

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
	return
}

// GetProductByCate func
func GetProductByCate(c *gin.Context) {
	pCategory1 := c.Query("pCategory1")
	pCategory2 := c.Query("pCategory2")
	pCategory3 := c.Query("pCategory3")

	if pCategory1 == "" {
		c.JSON(http.StatusOK, utils.JSONReturnMsg(
			false, "pCategory1은 필수 입력 값 입니다.",
		))
		return
	}

	cateList := []string{pCategory1, pCategory2, pCategory3}

	filter := utils.MakeCateFilter(cateList)

	query := "SELECT count(*) FROM product_list"+filter
	exist := mysql.CheckExist(query)

	if !exist {
		c.JSON(http.StatusOK, utils.JSONReturnMsg(
			false, "해당 카테고리에 대한 상품정보가 존재하지 않습니다",
		))
		return
	}
	productItem := mysql.GetProductByCate(filter)
	c.JSON(http.StatusOK, utils.JSONReturnResult(
		true, productItem,
	))
	return
}