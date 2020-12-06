package mysql

import (
	"github.com/ssoyyoung.p/BlackHeart-Golang/model"
	"github.com/ssoyyoung.p/BlackHeart-Golang/mysql/crud"
)

// GetOrderList func
func GetOrderList(uID string) []model.Order {
	query := "SELECT * FROM order_list where uID='" + uID + "'"

	orderList := crud.GetOrderList(query)

	return orderList
}
