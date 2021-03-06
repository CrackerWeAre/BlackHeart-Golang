package mysql

import (
	"strconv"

	"github.com/ssoyyoung.p/BlackHeart-Golang/model"
	"github.com/ssoyyoung.p/BlackHeart-Golang/mysql/crud"
)

// GetUserList func
func GetUserList() []model.User {
	query := "SELECT * FROM user_list"
	allUsers := crud.MultiRowQueryUserList(query)

	return allUsers
}

// GetUserDetail func
func GetUserDetail(uEmail string) model.User {
	query := "SELECT * FROM user_list where uEmail='" + uEmail + "'"
	userDetail := crud.RowQueryUserList(query)

	return userDetail
}

// DeleteUser func
func DeleteUser(uID string) {
	query := "DELETE FROM user_list Where uID=" + uID
	crud.ExecQuery(query)
}

// DeleteMultiUser func
func DeleteMultiUser(uID string) {
	query := "DELETE FROM user_list Where uID IN (" + uID + ")"
	crud.ExecQuery(query)
}

// UpdateUser func
func UpdateUser(user model.User) {
	query := "UPDATE user_list SET uName='?', uGender='?', uPhone='?', uBirth='?', uAgree='?', uLevel='?' WHERE uID=?"

	crud.UpdateUser(query, user)
}

// InsertNewUser func
func InsertNewUser(user model.User) {
	query := `INSERT INTO user_list (uEmail, uPW, uName, uGender, uPhone, uBirth, uAgree) 
			  value (?,?,?,?,?,?,?)`

	crud.InsertRowUserList(query, user)
}

// CheckExistByEmail func
func CheckExistByEmail(dbName, uEmail string) bool {
	query := "SELECT count(*) FROM "+dbName+" where uEmail='" + uEmail + "'"
	return crud.CheckUserExistInUserList(query)
}

// CheckExistByID func
func CheckExistByID(dbName string, uID int) bool {
	query := "SELECT count(*) FROM "+dbName+" where uID=" + strconv.Itoa(uID)
	return crud.CheckUserExistInUserList(query)
}

func CheckExist(query string) bool {
	return crud.CheckUserExist(query)
}

// LoginUser func
func LoginUser(user model.User) bool {
	query := "SELECT uEmail, uPW from user_list where uEmail=?"

	return crud.GetEmailAndPwUserList(query, user)
}
