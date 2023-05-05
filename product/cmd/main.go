package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-kit/log"
	"github.com/gorilla/mux"
	"github.com/yamadev11/e-commerce/product"
)

func defaultTimestampUTC() log.Valuer {
	return func() interface{} {
		return time.Now().UTC().UnixMilli()
	}
}

const ProductServicePort int = 8080

func main() {
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))
	logger = log.With(logger, "TS", defaultTimestampUTC())
	logger = log.With(logger, "Service", "ProductService")
	logger = log.With(logger, "Caller", log.DefaultCaller)
	router := mux.NewRouter()
	product.NewProductService(logger, router)

	_ = logger.Log("Msg", "Starting Product Service")
	_ = http.ListenAndServe(fmt.Sprintf(":%d", ProductServicePort), router)
}
