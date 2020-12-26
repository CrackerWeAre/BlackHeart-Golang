package crud

import (
	"github.com/ssoyyoung.p/BlackHeart-Golang/model"
	"github.com/ssoyyoung.p/BlackHeart-Golang/utils"
)

func GetCartList(query string) []model.Cart {
	DB := ConnectDB()
	defer DB.Close()

	var cart model.Cart
	var allCarts []model.Cart

	rows, err := DB.Query(query)
	utils.CheckErr(err)

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&cart.CID, &cart.UID, &cart.PID, &cart.CQuantity)
		utils.CheckErr(err)

		allCarts = append(allCarts, cart)
	}



	return allCarts
}