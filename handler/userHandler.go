package handler

import (
	"fmt"
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

	mysql.DeleteUser(uID)
}

// CreateUser func
func CreateUser(c *gin.Context) {
	user := model.User{}
	user.Uemail = c.PostForm("uEmail")
	user.Upw = c.PostForm("uPW")
	user.Uname = c.PostForm("uName")
	if user.Uemail == "" || user.Upw == "" || user.Uname == "" {
		c.String(http.StatusOK, "Uemail, Upw, Uname, uAgree 값은 필수 입력사항입니다.")
		return
	}
	user.Ugender = c.DefaultPostForm("uGender", "null")
	user.Uaddr = c.DefaultPostForm("uAddr", "null")
	user.Upost, _ = strconv.Atoi(c.DefaultPostForm("uPost", "0"))
	user.Uphone = c.DefaultPostForm("uPhone", "null")
	user.Ubirth = c.DefaultPostForm("uBirth", "2000-01-01")
	user.Uagree, _ = strconv.Atoi(c.DefaultPostForm("uAgree", "0"))

	mysql.InsertNewUser(user)
}

// LoginUser func
func LoginUser(c *gin.Context) {
	user := model.User{}
	user.Uemail = c.PostForm("uEmail")
	user.Upw = c.PostForm("uPW")

	if user.Uemail == "" || user.Upw == "" {
		c.String(http.StatusOK, "Uemail, Upw 값은 필수 입력사항입니다.")
		return
	}
	exist := mysql.CheckExist(user.Uemail)
	fmt.Println(exist)
	if exist {
		status := mysql.LoginUser(user)

		if status {
			c.JSON(http.StatusOK, gin.H{"login": "Success"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"login": "Fail"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "존재하지 않는 아이디 입니다"})
	return

}

// CheckExist func
func CheckExist(c *gin.Context) {
	uEmail := c.Query("uEmail")
	exist := mysql.CheckExist(uEmail)

	if exist {
		c.JSON(http.StatusOK, gin.H{"msg": "이미 존재하는 아이디 입니다."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "사용가능한 아이디 입니다."})
	return
}
