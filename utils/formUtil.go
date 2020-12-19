package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ssoyyoung.p/BlackHeart-Golang/model"
)

// ReadUserWithDefault func
func ReadUserWithDefault(c *gin.Context) model.User {
	user := model.User{}
	user.Uemail = c.PostForm("uEmail")
	user.Upw = c.PostForm("uPW")
	user.Uname = c.PostForm("uName")
	user.Ugender = c.DefaultPostForm("uGender", "null")
	user.Uaddr = c.DefaultPostForm("uAddr", "null")
	user.Upost, _ = strconv.Atoi(c.DefaultPostForm("uPost", "0"))
	user.Uphone = c.PostForm("uPhone")
	user.Ubirth = c.DefaultPostForm("uBirth", "2000-01-01")
	user.Uagree, _ = strconv.Atoi(c.DefaultPostForm("uAgree", "0"))

	return user
}

// ReadUser func
func ReadUser(c *gin.Context) model.User {
	user := model.User{}
	user.UID, _ = strconv.Atoi(c.PostForm("uID"))
	user.Uemail = c.PostForm("uEmail")
	user.Uname = c.PostForm("uName")
	user.Ugender = c.PostForm("uGender")
	user.Uaddr = c.PostForm("uAddr")
	user.Upost, _ = strconv.Atoi(c.PostForm("uPost"))
	user.Uphone = c.PostForm("uPhone")
	user.Ubirth = c.PostForm("uBirth")
	user.Uagree, _ = strconv.Atoi(c.PostForm("uAgree"))
	user.Ulevel, _ = strconv.Atoi(c.PostForm("uLevel"))

	return user
}
