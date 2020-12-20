package mysql

import (
	"github.com/ssoyyoung.p/BlackHeart-Golang/model"
	"github.com/ssoyyoung.p/BlackHeart-Golang/mysql/crud"
)

func GetCartList(uID string) model.Cart{
	cartQuery := ""
	prodQuery := ""

	cartList := crud.GetCartItem(cartQuery)
	prodList := crud.GetProductItem(prodQuery)

	cartList.Product = prodList

	return cartList
}
