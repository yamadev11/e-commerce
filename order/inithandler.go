package order

import (
	"os"

	"github.com/go-kit/log"
	"github.com/gorilla/mux"
	productsdk "github.com/yamadev11/e-commerce/product/sdk"
)

func NewOrderService(logger log.Logger, router *mux.Router, product productsdk.ProductService) {
	bl := NewBL(logger, product)
	eps := makeEndpoint(bl, logger)
	err := AddHandlers(router, eps, logger)
	if err != nil {
		_ = logger.Log("Layer", "Inithandler", "Method", "NewOrderService", "Error", err.Error())
		os.Exit(1)
	}
}
