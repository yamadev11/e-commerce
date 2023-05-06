package sdk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/yamadev11/e-commerce/order/spec"
	"github.com/yamadev11/e-commerce/sdkutil"
)

type OrderService interface {
	Get(ctx context.Context, orderID int) (*spec.GetOrderResponse, error)
	Create(ctx context.Context, items []spec.Item) (*spec.Order, error)
	Update(ctx context.Context, orderID, status int) error
}

type Order struct {
	Port int
	sdkutil.BaseSDK
}

func NewOrderService(port int) OrderService {
	return &Order{
		Port: port,
		BaseSDK: sdkutil.BaseSDK{
			HTTPClient: &http.Client{
				Timeout: time.Minute,
			},
		},
	}
}

// Get returns the order details for given order id.
func (svc *Order) Get(ctx context.Context, orderID int) (*spec.GetOrderResponse, error) {

	req, err := http.NewRequest(http.MethodGet,
		fmt.Sprintf("http://localhost:%d/orders/%d", svc.Port, orderID),
		nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	var response *spec.GetOrderResponse
	if err := svc.BaseSDK.SendRequest(req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

// Update updates the order status of the given order id.
func (svc *Order) Update(ctx context.Context, orderID, status int) error {

	payload := map[string]interface{}{
		"status": status,
	}

	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPatch,
		fmt.Sprintf("http://localhost:%d/orders/%d/status", svc.Port, orderID),
		bytes.NewBuffer(payloadJSON))
	if err != nil {
		return err
	}

	req = req.WithContext(ctx)

	var response interface{}
	if err := svc.BaseSDK.SendRequest(req, &response); err != nil {
		return err
	}

	return nil
}

// Create creates new order with specified items/products.
func (svc *Order) Create(ctx context.Context, items []spec.Item) (*spec.Order, error) {
	payload := map[string]interface{}{
		"items": items,
	}

	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost,
		fmt.Sprintf("http://localhost:%d/orders", svc.Port),
		bytes.NewBuffer(payloadJSON))
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	var response *spec.Order
	if err := svc.BaseSDK.SendRequest(req, &response); err != nil {
		return nil, err
	}

	return response, nil
}
