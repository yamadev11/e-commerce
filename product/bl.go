package product

import (
	"context"

	"github.com/go-kit/log"
	"github.com/yamadev11/e-commerce/product/spec"
)

type BL struct {
	logger log.Logger
	dl     *DL
}

func NewBL(logger log.Logger) *BL {
	return &BL{
		dl:     NewDL(),
		logger: logger,
	}
}

// List returns the list of products.
func (bl *BL) List(ctx context.Context) (*spec.ListResponse, error) {

	response := &spec.ListResponse{}
	products := bl.dl.List(ctx)
	for _, product := range products {
		response.Products = append(response.Products, product)
	}

	return response, nil
}

// UpdateQuantity updates the quantity of the given product.
func (bl *BL) UpdateQuantity(ctx context.Context, productID, quantity int) (err error) {

	err = bl.dl.UpdateQuantity(ctx, productID, quantity)
	if err != nil {
		_ = bl.logger.Log("Method", "UpdateQuantity", "Error", err.Error())
		return
	}

	_ = bl.logger.Log("Method", "UpdateQuantity", "Message", "Product quantity updated successfully")
	return
}
