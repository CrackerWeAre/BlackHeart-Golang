package mysql

import (
	"github.com/ssoyyoung.p/BlackHeart-Golang/model"
	"github.com/ssoyyoung.p/BlackHeart-Golang/mysql/crud"
)

func GetProductItem(pID string) model.Product {
	query := "SELECT * FROM product_list WHERE pID=" + pID

	product := crud.GetProductItem(query)

	return product
}


// GetProductList func
func GetProductList() []model.Product {
	query := "SELECT * FROM product_list"
	allProducts := crud.GetProductList(query)

	return allProducts
}

// GetProductByCate
func GetProductByCate(filter string) []model.Product {
	query := "SELECT * FROM product_list"+filter

	product := crud.GetProductByCate(query)

	return product
}
