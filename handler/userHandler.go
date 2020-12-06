package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
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

// UpdateUser func
func UpdateUser(c *gin.Context) {
	user := utils.ReadUser(c)
	if user.UID == 0 || user.Uemail == "" {
		c.JSON(http.StatusOK, utils.JSONReturnMsg(
			false, "uID, uEmail은 필수 입력사항입니다",
		))
		return
	}

	exist := mysql.CheckExistByID(user.UID)
	if !exist {
		c.JSON(http.StatusOK, utils.JSONReturnMsg(
			false, "해당유저가 존재하지 않습니다",
		))
		return
	}

	mysql.UpdateUser(user)
	c.JSON(http.StatusOK, utils.JSONReturnMsg(
		true, "유저정보가 업데이트 되었습니다",
	))
	return
}

// CreateUser func
func CreateUser(c *gin.Context) {
	user := utils.ReadUserWithDefault(c)

	if user.Uemail == "" || user.Upw == "" || user.Uname == "" {
		c.JSON(http.StatusOK, utils.JSONReturnMsg(
			false, "uEmail, uPW, uName, uAgree 값은 필수 입력사항입니다",
		))
		return
	}
	exist := mysql.CheckExistByEmail(user.Uemail)
	if exist {
		c.JSON(http.StatusOK, utils.JSONReturnMsg(
			false, "해당 아이디의 유저가 이미 존재합니다",
		))
		return
	}
	mysql.InsertNewUser(user)

	c.JSON(http.StatusOK, utils.JSONReturnMsg(
		true, "유저가 생성되었습니다",
	))
	return
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
	exist := mysql.CheckExistByEmail(user.Uemail)
	if !exist {
		c.JSON(http.StatusOK, utils.JSONReturnMsg(
			false, "존재하지 않는 아이디 입니다",
		))
		return
	}

	status := mysql.LoginUser(user)

	if status {
		token := jwt.New(jwt.SigningMethodHS256)

		claims := token.Claims.(jwt.MapClaims)
		claims["userID"] = user.Uemail
		claims["admin"] = false

		expTime := time.Now().Add(time.Hour * 72).Unix()
		claims["exp"] = expTime

		t, err := token.SignedString([]byte("secret"))
		utils.CheckErr(err)

		var loginUser model.LoginUser
		loginUser.Token = t
		loginUser.TokenEXP = expTime
		loginUser.Uemail = user.Uemail

		c.JSON(http.StatusOK, utils.JSONReturnResult(
			true, loginUser,
		))
		return
	}
	c.JSON(http.StatusOK, utils.JSONReturnMsg(
		false, "잘못된 비밀번호 입니다",
	))
	return
}

// CheckExist func
func CheckExist(c *gin.Context) {
	uEmail := c.Query("uEmail")
	exist := mysql.CheckExistByEmail(uEmail)

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
