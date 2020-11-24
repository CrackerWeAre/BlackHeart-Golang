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
	user.Ugender = c.PostForm("uGender")
	user.Uphone = c.PostForm("uPhone")
	user.Ubirth = c.PostForm("uBirth")
	user.Uagree, _ = strconv.Atoi(c.PostForm("uAgree"))

	mysql.InsertNewUser(user)
}
