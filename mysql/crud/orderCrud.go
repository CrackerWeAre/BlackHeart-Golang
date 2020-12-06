package crud

import (
	"github.com/ssoyyoung.p/BlackHeart-Golang/model"
	"github.com/ssoyyoung.p/BlackHeart-Golang/utils"
)

// GetOrderList func
func GetOrderList(query string) []model.Order {
	DB := ConnectDB()
	defer DB.Close()

	var order model.Order
	var allOrders []model.Order

	rows, err := DB.Query(query)
	utils.CheckErr(err)

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&order.OID, &order.UID, &order.OOrderNum, &order.OOrderDate, &order.PID, &order.PName, &order.POption, &order.PPrice, &order.PDistount, &order.OQuentity, &order.ODestination, &order.OPhoneNum, &order.OPayment, &order.OInvoice, &order.ODelivery, &order.OStatus)
		utils.CheckErr(err)

		allOrders = append(allOrders, order)
	}

	return allOrders
}
