package order

import (
	"github.com/go-kit/log"
	"github.com/gorilla/mux"
	productsdk "github.com/yamadev11/e-commerce/product/sdk"
)

func NewOrderService(logger log.Logger, router *mux.Router, product productsdk.ProductService) {
	bl := NewBL(logger, product)
	eps := makeEndpoint(bl, logger)
	AddHandlers(router, eps, logger)
}
