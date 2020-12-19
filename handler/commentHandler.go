package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ssoyyoung.p/BlackHeart-Golang/mysql"
	"github.com/ssoyyoung.p/BlackHeart-Golang/utils"
)

// GetReviewCommentList func
func GetReviewCommentList(c *gin.Context) {

	uID := c.Query("uID")

	commentList := mysql.GetReviewCommentList(uID)

	c.JSON(http.StatusOK, utils.JSONReturnResult(
		true, commentList,
	))
	return
}

// GetReviewBoardList func
func GetReviewBoardList(c *gin.Context) {

	uID := c.Query("uID")

	commentList := mysql.GetReviewBoardList(uID)

	c.JSON(http.StatusOK, utils.JSONReturnResult(
		true, commentList,
	))
	return
}