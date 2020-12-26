package handler

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"github.com/ssoyyoung.p/BlackHeart-Golang/model"
	"github.com/ssoyyoung.p/BlackHeart-Golang/mysql"
	"github.com/ssoyyoung.p/BlackHeart-Golang/utils"
)

type Cookie struct {
	Name  string
	Value string

	Path       string    // optional
	Domain     string    // optional
	Expires    time.Time // optional
	RawExpires string    // for reading cookies only

	// MaxAge=0 means no 'Max-Age' attribute specified.
	// MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
	// MaxAge>0 means Max-Age attribute present and given in seconds
	MaxAge   int
	Secure   bool
	HttpOnly bool
	Raw      string
	Unparsed []string // Raw text of unparsed attribute-value pairs
}

func Login(c *gin.Context) {
	user := model.User{}
	user.Uemail = c.PostForm("uEmail")
	user.Upw = c.PostForm("uPW")

	if user.Uemail == "" || user.Upw == "" {
		c.JSON(http.StatusUnauthorized, utils.JSONReturnMsg(
			false, "uEmail, uPW 값은 필수 입력사항입니다",
		))
		return
	}
	exist := mysql.CheckExistByEmail("user_list",user.Uemail)
	if !exist {
		c.JSON(http.StatusUnauthorized, utils.JSONReturnMsg(
			false, "존재하지 않는 아이디 입니다",
		))
		return
	}

	status := mysql.LoginUser(user)

	if status {
		userDetail := mysql.GetUserDetail(user.Uemail)

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
		loginUser.Uname = userDetail.Uname
		loginUser.UID = userDetail.UID

		c.SetCookie(
			"blackheart", t,
			60*60*24*7,
			"/",
			"sparker.kr",
			true,
			true,
			)

		c.JSON(http.StatusOK, utils.JSONReturnResult(
			true, loginUser,
		))
		return
	}

	c.JSON(http.StatusUnauthorized, utils.JSONReturnMsg(
		false, "잘못된 비밀번호 입니다",
	))
	return
}

func Logout(c *gin.Context) {
	c.SetCookie(
		"blackheart",
		"",
		0,
		"/",
		"sparker.kr",
		true,
		true,
	)

}
