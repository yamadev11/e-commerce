package componenttest

import (
	"github.com/yamadev11/e-commerce/product"
	productspec "github.com/yamadev11/e-commerce/product/spec"
)

var (
	listProductResponse = &productspec.ListResponse{
		Products: []productspec.Product{
			{
				ID:          1,
				Name:        "iPhone14",
				Price:       80000,
				Category:    product.CategoryIDNameMap[product.Premium],
				AvlQuantity: 15,
			},
			{
				ID:          2,
				Name:        "iPad",
				Price:       70000,
				Category:    product.CategoryIDNameMap[product.Premium],
				AvlQuantity: 12,
			},
			{
				ID:          3,
				Name:        "macBook Pro",
				Price:       200000,
				Category:    product.CategoryIDNameMap[product.Premium],
				AvlQuantity: 6,
			},
			{
				ID:          4,
				Name:        "Thinkpad E470",
				Price:       70000,
				Category:    product.CategoryIDNameMap[product.Regular],
				AvlQuantity: 25,
			},
			{
				ID:          5,
				Name:        "Ideapad 500",
				Price:       50000,
				Category:    product.CategoryIDNameMap[product.Budget],
				AvlQuantity: 11,
			},
		},
	}
)

const (
	CreateOrderID1 = iota + 1
	GetOrderID1
	GetOrderID2
	UpdateOrderID1
	UpdateOrderID2
	UpdateOrderID3
)
