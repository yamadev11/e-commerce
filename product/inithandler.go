package product

import (
	"github.com/go-kit/log"
	"github.com/gorilla/mux"
)

func NewProductService(logger log.Logger, router *mux.Router) {
	bl := NewBL(logger)
	eps := makeEndpoint(bl, logger)
	AddHandlers(router, eps, logger)
}
