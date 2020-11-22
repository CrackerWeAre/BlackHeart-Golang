package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/ssoyyoung.p/BlackHeart-Golang/mysql"
	"github.com/ssoyyoung.p/BlackHeart-Golang/utils"
)

// GetUserList func
func GetUserList(c *gin.Context) {
	page, _ := strconv.Atoi(c.Param("page"))
	maxResult, _ := strconv.Atoi(c.Param("maxResult"))
	allUsers := mysql.GetUserList()

	c.JSON(http.StatusOK, utils.PgSplit(allUsers, page, maxResult))
}
