package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ssoyyoung.p/BlackHeart-Golang/mysql"
	"github.com/ssoyyoung.p/BlackHeart-Golang/utils"
)

// GetReviewCommentList func
func GetReviewCommentList(c *gin.Context) {

	uID := c.Query("uID")
	uIDint, _ := strconv.Atoi(uID)
	exist := mysql.CheckExistByID(uIDint)

	if !exist {
		c.JSON(http.StatusOK, utils.JSONReturnMsg(
			false, "해당 유저의 값이 존재하지 않습니다."))
		return
	}

	commentList := mysql.GetReviewCommentList(uID)

	c.JSON(http.StatusOK, utils.JSONReturnResult(
		true, commentList,
	))
	return
}

// GetReviewBoardList func
func GetReviewBoardList(c *gin.Context) {

	uID := c.Query("uID")
	uIDint, _ := strconv.Atoi(uID)
	exist := mysql.CheckExistByID(uIDint)

	if !exist {
		c.JSON(http.StatusOK, utils.JSONReturnMsg(
			false, "해당 유저의 값이 존재하지 않습니다."))
		return
	}

	commentList := mysql.GetReviewBoardList(uID)

	c.JSON(http.StatusOK, utils.JSONReturnResult(
		true, commentList,
	))
	return

}