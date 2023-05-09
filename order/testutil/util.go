package testutil

import (
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/go-kit/log"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/yamadev11/e-commerce/order"
	"github.com/yamadev11/e-commerce/order/sdk"
	"github.com/yamadev11/e-commerce/product/sdk/mock_sdk"
)

func defaultTimestampUTC() log.Valuer {
	return func() interface{} {
		return time.Now().UTC().UnixMilli()
	}
}

var (
	TestOrderServicePort int = 8081
)

type MockSDK struct {
	Product *mock_sdk.MockProductService
}

type TestObj struct {
	Order   sdk.OrderService
	MockSDK MockSDK
}

func InitTestInfra(t *testing.T) *TestObj {
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))
	logger = log.With(logger, "TS", defaultTimestampUTC())
	logger = log.With(logger, "Service", "OrderService")
	logger = log.With(logger, "Caller", log.DefaultCaller)
	router := mux.NewRouter()

	ctrl := gomock.NewController(t)
	mockSDK := MockSDK{}
	mockSDK.Product = mock_sdk.NewMockProductService(ctrl)

	order.NewOrderService(logger, router, mockSDK.Product)

	_ = logger.Log("Msg", "Starting Order Service")
	go func() {
		_ = http.ListenAndServe(fmt.Sprintf(":%d", TestOrderServicePort), router)
	}()

	orderSDK := sdk.NewOrderService(TestOrderServicePort)

	return &TestObj{
		Order:   orderSDK,
		MockSDK: mockSDK,
	}
}
