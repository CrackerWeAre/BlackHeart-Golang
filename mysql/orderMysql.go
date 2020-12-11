package mysql

import (
	"github.com/ssoyyoung.p/BlackHeart-Golang/model"
	"github.com/ssoyyoung.p/BlackHeart-Golang/mysql/crud"
)

// GetOrderList func
func GetOrderList(uID string) []model.Order {
	query := `select o.oID, o.uID, o.oOrderNum, oOrderDate, o.pID, o.pName, o.pOption, o.pPrice, o.pDiscount, o.oQuantity, o.oDestination, o.oPhoneNum, o.oPayment, o.oInvoice, o.oDelivery, o.oStatus, p.pListImage from order_list 
	as o join product_list as p on o.pID=p.pID where o.uID='` + uID + `'`

	orderList := crud.GetOrderList(query)

	return orderList
}
