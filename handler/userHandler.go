package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/ssoyyoung.p/BlackHeart-Golang/model"
	"github.com/ssoyyoung.p/BlackHeart-Golang/mysql"
	"github.com/ssoyyoung.p/BlackHeart-Golang/utils"
)

// GetUserList func
func GetUserList(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	maxResult, _ := strconv.Atoi(c.Query("maxResult"))

	allUsers := mysql.GetUserList()

	c.JSON(http.StatusOK, utils.PgSplit(allUsers, page, maxResult))
}

// DeleteUser func
func DeleteUser(c *gin.Context) {
	uID := c.Query("uID")
	if uID == "" {
		c.JSON(http.StatusOK, utils.JSONReturnMsg(
			false, "올바른 uID 값을 입력해주세요",
		))
		return
	}

	mysql.DeleteUser(uID)
	c.JSON(http.StatusOK, utils.JSONReturnMsg(
		true, "삭제되었습니다",
	))
}

// CreateUser func
func CreateUser(c *gin.Context) {
	user := model.User{}
	user.Uemail = c.PostForm("uEmail")
	user.Upw = c.PostForm("uPW")
	user.Uname = c.PostForm("uName")
	user.Ugender = c.DefaultPostForm("uGender", "null")
	user.Uaddr = c.DefaultPostForm("uAddr", "null")
	user.Upost, _ = strconv.Atoi(c.DefaultPostForm("uPost", "0"))
	user.Uphone = c.DefaultPostForm("uPhone", "null")
	user.Ubirth = c.DefaultPostForm("uBirth", "2000-01-01")
	user.Uagree, _ = strconv.Atoi(c.DefaultPostForm("uAgree", "0"))

	if user.Uemail == "" || user.Upw == "" || user.Uname == "" {
		c.JSON(http.StatusOK, utils.JSONReturnMsg(
			false, "Uemail, Upw, Uname, uAgree 값은 필수 입력사항입니다",
		))
		return
	}

	mysql.InsertNewUser(user)

	c.JSON(http.StatusOK, utils.JSONReturnMsg(
		true, "유저가 생성되었습니다",
	))
}

// LoginUser func
func LoginUser(c *gin.Context) {
	user := model.User{}
	user.Uemail = c.PostForm("uEmail")
	user.Upw = c.PostForm("uPW")

	if user.Uemail == "" || user.Upw == "" {
		c.JSON(http.StatusOK, utils.JSONReturnMsg(
			false, "uEmail, uPW 값은 필수 입력사항입니다",
		))
		return
	}
	exist := mysql.CheckExist(user.Uemail)
	if exist {
		status := mysql.LoginUser(user)

		if status {
			c.JSON(http.StatusOK, utils.JSONReturnMsg(
				true, "로그인에 성공했습니다",
			))
			return
		}
		c.JSON(http.StatusOK, utils.JSONReturnMsg(
			false, "잘못된 비밀번호 입니다",
		))
		return
	}

	c.JSON(http.StatusOK, utils.JSONReturnMsg(
		false, "존재하지 않는 아이디 입니다",
	))
	return

}

// CheckExist func
func CheckExist(c *gin.Context) {
	uEmail := c.Query("uEmail")
	exist := mysql.CheckExist(uEmail)

	if exist {
		c.JSON(http.StatusOK, utils.JSONReturnMsg(
			false, "이미 존재하는 아이디 입니다",
		))
		return
	}

	c.JSON(http.StatusOK, utils.JSONReturnMsg(
		true, "사용가능한 아이디 입니다",
	))
	return
}
