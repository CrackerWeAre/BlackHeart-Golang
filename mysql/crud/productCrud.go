package crud

import (
	"github.com/ssoyyoung.p/BlackHeart-Golang/model"
	"github.com/ssoyyoung.p/BlackHeart-Golang/utils"
)

func GetProductItem(query string) model.Product {
	DB := ConnectDB()
	defer DB.Close()

	var product model.Product

	row := DB.QueryRow(query)
	err := row.Scan(&product.PID, &product.PCode, &product.PName, &product.POption, &product.PPrice, &product.PDiscount, &product.PStock, &product.PListImage, &product.PMainImage, &product.PSubImage, &product.PCategory1, &product.PCategory2, &product.PCategory3, &product.PDisplay, &product.PSoldout)
	utils.CheckErr(err)

	return product
}

func GetProductList(query string) []model.Product{
	DB := ConnectDB()
	defer DB.Close()

	var product model.Product
	var AllProducts []model.Product

	rows, err := DB.Query(query)
	utils.CheckErr(err)

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&product.PID, &product.PCode, &product.PName, &product.POption, &product.PPrice, &product.PDiscount, &product.PStock, &product.PListImage, &product.PMainImage, &product.PSubImage, &product.PCategory1, &product.PCategory2, &product.PCategory3, &product.PDisplay, &product.PSoldout)
		utils.CheckErr(err)

		AllProducts = append(AllProducts, product)
	}

	return AllProducts

}

// GetProductByCate func
func GetProductByCate(query string) []model.Product{
	DB := ConnectDB()
	defer DB.Close()

	var product model.Product
	var AllProducts []model.Product

	rows, err := DB.Query(query)
	utils.CheckErr(err)

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&product.PID, &product.PCode, &product.PName, &product.POption, &product.PPrice, &product.PDiscount, &product.PStock, &product.PListImage, &product.PMainImage, &product.PSubImage, &product.PCategory1, &product.PCategory2, &product.PCategory3, &product.PDisplay, &product.PSoldout)
		utils.CheckErr(err)

		AllProducts = append(AllProducts, product)
	}

	return AllProducts

}

