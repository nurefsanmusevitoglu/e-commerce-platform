package repository

import (
	"github.com/nurefsanmusevitoglu/e-commerce-platform/internal/model"
)

var orders []model.Order

func CreateOrder(productCode string, quantity int) {
	var order = model.Order{
		ProductCode: productCode,
		Quantity:    quantity,
	}
	orders = append(orders, order)
}
