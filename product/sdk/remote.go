package sdk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/yamadev11/e-commerce/product/spec"
	"github.com/yamadev11/e-commerce/sdkutil"
)

type ProductService interface {
	List(ctx context.Context) (*spec.ListResponse, error)
	UpdateQuantity(ctx context.Context, productID, quantity int) error
}

type Product struct {
	sdkutil.BaseSDK
}

func NewProduct() ProductService {
	return &Product{
		BaseSDK: sdkutil.BaseSDK{
			HTTPClient: &http.Client{
				Timeout: time.Minute,
			},
		},
	}
}

// List returns the list of products.
func (svc *Product) List(ctx context.Context) (*spec.ListResponse, error) {

	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/products", nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	var response *spec.ListResponse
	if err := svc.BaseSDK.SendRequest(req, &response); err != nil {
		return nil, err
	}

	fmt.Printf("%v", response)
	return response, nil
}

// UpdateQuantity updates the quantity of the given product.
func (svc *Product) UpdateQuantity(ctx context.Context, productID, quantity int) error {

	payload := map[string]interface{}{
		"quantity": quantity,
	}

	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPatch,
		fmt.Sprintf("http://localhost:8080/products/%d/quantity", productID),
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
