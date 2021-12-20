package repository

import "github.com/nurefsanmusevitoglu/e-commerce-platform/internal/model"

var products map[string]model.Product

func CreateProduct(productCode string, price int, stock int) {
	// TO-DO: handle err if product already exists
	var product = model.Product {
		ProductCode: productCode,
		Price: price,
		Stock: stock,
	}
	products[product.ProductCode] = product
}

func GetProductInfo(productCode string) model.Product {
	// TO-DO: handle err if product does not exist
	return products[productCode]
}