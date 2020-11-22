package mysql

import (
	"github.com/ssoyyoung.p/BlackHeart-Golang/model"
	"github.com/ssoyyoung.p/BlackHeart-Golang/mysql/crud"
)

// GetUserList func
func GetUserList() []model.UserList {
	query := "SELECT * FROM user_list"
	allUsers := crud.MultiRowQueryInUserList(query)

	return allUsers
}
