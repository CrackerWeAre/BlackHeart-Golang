package crud

import (
	"github.com/ssoyyoung.p/BlackHeart-Golang/model"
	"github.com/ssoyyoung.p/BlackHeart-Golang/utils"
)

func GetCartItem(query string) model.Cart {
	DB := ConnectDB()
	defer DB.Close()

	var cart model.Cart

	row := DB.QueryRow(query)
	err := row.Scan(&cart.CID, &cart.UID, &cart.PID, &cart.POption, &cart.CQuantity)
	utils.CheckErr(err)

	return cart
}