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

// CreateUser func
func CreateUser(c *gin.Context) {
	user := model.User{}
	user.Uemail = c.PostForm("uEmail")
	user.Upw = c.PostForm("uPW")
	user.Uname = c.PostForm("uName")
	if user.Uemail == "" || user.Upw == "" || user.Uname == "" {
		c.String(http.StatusOK, "Uemail, Upw, Uname 값은 필수 입력사항입니다.")
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
