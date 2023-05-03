package product

import (
	"os"

	"github.com/go-kit/log"
	"github.com/gorilla/mux"
)

func NewProductService(logger log.Logger, router *mux.Router) {
	bl := NewBL(logger)
	eps := makeEndpoint(bl, logger)
	err := AddHandlers(router, eps, logger)
	if err != nil {
		_ = logger.Log("Layer", "Inithandler", "Method", "NewProductService", "Error", err.Error())
		os.Exit(1)
	}
}
