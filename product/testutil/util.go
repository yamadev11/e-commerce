package testutil

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

var (
	TestProductPort int = 8081
)

func InitTestInfra() {
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))
	logger = log.With(logger, "TS", defaultTimestampUTC())
	logger = log.With(logger, "Service", "ProductService")
	logger = log.With(logger, "Caller", log.DefaultCaller)
	router := mux.NewRouter()
	product.NewProductService(logger, router)

	_ = logger.Log("Msg", "Starting Product Service")
	go func() {
		_ = http.ListenAndServe(fmt.Sprintf(":%d", TestProductPort), router)
	}()
}
