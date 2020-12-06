package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ssoyyoung.p/BlackHeart-Golang/mysql"
	"github.com/ssoyyoung.p/BlackHeart-Golang/utils"
)

// GetCommentList func
func GetCommentList(c *gin.Context) {

	uID := c.Query("uID")

	commentList := mysql.GetCommentList(uID)

	c.JSON(http.StatusOK, utils.JSONReturnResult(
		true, commentList,
	))
	return
}
