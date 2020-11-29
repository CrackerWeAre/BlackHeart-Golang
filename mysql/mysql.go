package mysql

import (
	"github.com/ssoyyoung.p/BlackHeart-Golang/model"
	"github.com/ssoyyoung.p/BlackHeart-Golang/mysql/crud"
)

// GetUserList func
func GetUserList() []model.User {
	query := "SELECT * FROM user_list"
	allUsers := crud.MultiRowQueryUserList(query)

	return allUsers
}

// InsertNewUser func
func InsertNewUser(user model.User) {
	query := `INSERT INTO user_list (uEmail, uPW, uName, uGender, uPhone, uBirth, uAgree) 
			  value (?,?,?,?,?,?,?)`

	crud.InsertRowUserList(query, user)
}

// CheckExist func
func CheckExist(uEmail string) bool {
	return crud.CheckUserExistInUserList(uEmail)
}

// LoginUser func
func LoginUser(user model.User) bool {
	query := "SELECT uEmail, uPW from user_list where uEmail=?"

	return crud.GetEmailAndPwUserList(query, user)
}
