package main

import (
	"net/http"
	"os"
	"time"

	"github.com/go-kit/log"
	"github.com/gorilla/mux"
	"github.com/yamadev11/e-commerce/order"
	"github.com/yamadev11/e-commerce/product/sdk"
)

func defaultTimestampUTC() log.Valuer {
	return func() interface{} {
		return time.Now().UTC().UnixMilli()
	}
}

func main() {
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))
	logger = log.With(logger, "TS", defaultTimestampUTC())
	logger = log.With(logger, "Service", "OrderService")
	logger = log.With(logger, "Caller", log.DefaultCaller)
	router := mux.NewRouter()

	product := sdk.NewProduct()
	order.NewOrderService(logger, router, product)

	_ = logger.Log("Msg", "Starting Order Service")
	_ = http.ListenAndServe(":8090", router)
}
