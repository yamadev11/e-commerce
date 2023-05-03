package product

import (
	"context"
	"errors"

	"github.com/yamadev11/e-commerce/product/spec"
)

type DL struct{}

func NewDL() *DL {
	return &DL{}
}

func (dl *DL) List(ctx context.Context) map[int]spec.Product {
	return products
}

func (dl *DL) UpdateQuantity(ctx context.Context, productID, quantity int) error {

	if product, ok := products[productID]; ok {
		product.AvlQuantity = product.AvlQuantity + quantity
		products[productID] = product
		return nil
	}

	return errors.New("invalid productID, couldn't update product quantity")
}

// ProductID constants
const (
	iPhone = iota + 1
	iPad
	macBook
	thinkpad
	ideapad
)

// CategoryID constant
const (
	Premium = iota + 1
	Regular
	Budget
)

var (
	categoryIDNameMap = map[int]string{
		Premium: "Premium",
		Regular: "Regular",
		Budget:  "Budget",
	}

	products = map[int]spec.Product{
		iPhone: {
			ID:          iPhone,
			Name:        "iPhone14",
			Price:       80000,
			CategoryID:  Premium,
			Category:    categoryIDNameMap[Premium],
			AvlQuantity: 15,
		},
		iPad: {
			ID:          iPad,
			Name:        "iPad",
			Price:       70000,
			CategoryID:  Premium,
			Category:    categoryIDNameMap[Premium],
			AvlQuantity: 12,
		},
		macBook: {
			ID:          macBook,
			Name:        "macBook Pro",
			Price:       200000,
			CategoryID:  Premium,
			Category:    categoryIDNameMap[Premium],
			AvlQuantity: 6,
		},
		thinkpad: {
			ID:          thinkpad,
			Name:        "Thinkpad E470",
			Price:       70000,
			CategoryID:  Regular,
			Category:    categoryIDNameMap[Regular],
			AvlQuantity: 25,
		},
		ideapad: {
			ID:          ideapad,
			Name:        "Ideapad 500",
			Price:       50000,
			CategoryID:  Budget,
			Category:    categoryIDNameMap[Budget],
			AvlQuantity: 11,
		},
	}
)
