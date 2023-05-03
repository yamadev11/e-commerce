package product

import (
	"context"
	"errors"

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

func (bl *BL) List(ctx context.Context) (*spec.ListResponse, error) {

	response := &spec.ListResponse{}
	products := bl.dl.List(ctx)
	for _, product := range products {
		response.Products = append(response.Products, product)
	}

	return response, nil
}

func (bl *BL) UpdateQuantity(ctx context.Context, productID, quantity int) (err error) {

	// executes when user is placing order i.e for negative quantity
	if quantity < 0 {
		if quantity < -10 || quantity < -products[productID].AvlQuantity {
			err = errors.New("quantity is beyond limit")
			_ = bl.logger.Log("Method", "UpdateQuantity", "Error", err.Error())
			return
		}
	}

	err = bl.dl.UpdateQuantity(ctx, productID, quantity)
	if err != nil {
		_ = bl.logger.Log("Method", "UpdateQuantity", "Error", err.Error())
		return
	}

	_ = bl.logger.Log("Method", "UpdateQuantity", "Message", "Product quantity updated successfully")
	return
}
