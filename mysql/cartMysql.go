package mysql

import (
	"github.com/ssoyyoung.p/BlackHeart-Golang/model"
	"github.com/ssoyyoung.p/BlackHeart-Golang/mysql/crud"
)

func GetCartList(uID string) []model.Cart{
	query := "select * from cart_list where uID='" + uID + "'"

	cartList := crud.GetCartList(query)

	return cartList
}
