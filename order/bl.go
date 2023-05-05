package order

import (
	"context"

	"github.com/go-kit/log"
	"github.com/pkg/errors"
	"github.com/yamadev11/e-commerce/order/spec"
	productsdk "github.com/yamadev11/e-commerce/product/sdk"
	productspec "github.com/yamadev11/e-commerce/product/spec"
)

type BL struct {
	logger  log.Logger
	dl      *DL
	product productsdk.ProductService
}

func NewBL(logger log.Logger, product productsdk.ProductService) *BL {
	return &BL{
		dl:      NewDL(),
		logger:  logger,
		product: product,
	}
}

// Get returns the order details for given order id.
func (svc *BL) Get(ctx context.Context, orderID int) (*spec.GetOrderResponse, error) {
	order, err := svc.dl.Get(ctx, orderID)
	if err != nil {
		_ = svc.logger.Log("Method", "Get", "Error", err.Error())
		return nil, err
	}

	// call list product
	listProductResponse, err := svc.product.List(ctx)
	if err != nil {
		_ = svc.logger.Log("Method", "Get", "Error", err.Error())
		return nil, err
	}

	// create a map for better performance
	productMap := map[int]productspec.Product{}
	for _, product := range listProductResponse.Products {
		productMap[product.ID] = product
	}

	// This should be done at the frontend side and
	// microservice should be kept independent as much as possible.
	orderItems := []spec.OrderItem{}
	for _, item := range order.Items {
		product := productMap[item.ID]
		orderItem := spec.OrderItem{
			ID:       item.ID,
			Name:     product.Name,
			Price:    product.Price,
			Category: product.Category,
			Quantity: item.Quantity,
		}

		orderItems = append(orderItems, orderItem)

	}

	response := &spec.GetOrderResponse{
		ID:           order.ID,
		Items:        orderItems,
		Amount:       order.Amount,
		Discount:     order.Discount,
		FinalAmount:  order.FinalAmount,
		Status:       OrderStatusMap[order.Status],
		OrderDate:    order.OrderDate,
		DispatchDate: order.DispatchDate,
	}

	return response, nil
}

// Create places the new order for the specified items/products.
func (bl *BL) Create(ctx context.Context, items []spec.Item) (*spec.Order, error) {

	listProductResponse, err := bl.product.List(ctx)
	if err != nil {
		_ = bl.logger.Log("Method", "Create", "Error", err.Error())
		return nil, err
	}

	// create a map for better performance
	productMap := map[int]productspec.Product{}
	for _, product := range listProductResponse.Products {
		productMap[product.ID] = product
	}

	itemMap := map[int]int{}
	for _, item := range items {
		if item.Quantity <= 0 {
			err = errors.New("Quantity can not be zero or negative")
			_ = bl.logger.Log("Method", "Create", "Error", err.Error())
			return nil, err
		}

		itemMap[item.ID] = itemMap[item.ID] + item.Quantity
		product, ok := productMap[item.ID]
		if !ok {
			err = errors.New("Invalid ProductID")
			_ = bl.logger.Log("Method", "Create", "Error", err.Error())
			return nil, err
		}

		// check whether quantity is within limit or not
		quantity := itemMap[item.ID]
		if quantity > product.AvlQuantity || quantity > 10 {
			err = errors.New("Product quantity is beyond limit")
			_ = bl.logger.Log("Method", "Create", "Error", err.Error())
			return nil, err
		}

	}

	var amount, discount float64
	var premiumProductCount int
	for productID, quantity := range itemMap {
		// update product quantity
		err = bl.product.UpdateQuantity(ctx, productID, -quantity)
		if err != nil {
			_ = bl.logger.Log("Method", "Create", "Error", err.Error())
			return nil, err
		}

		product := productMap[productID]
		amount += float64(quantity) * product.Price
		if product.Category == "Premium" {
			premiumProductCount++
		}
	}

	// give discount if user has purchased 3 or more premium products
	if premiumProductCount >= 3 {
		discount = amount * 0.1
	}

	// create order
	order, err := bl.dl.Create(ctx, items, amount, discount)
	if err != nil {
		_ = bl.logger.Log("Method", "Create", "Error", err.Error())
		return nil, err
	}

	_ = bl.logger.Log("Method", "Create", "Message", "Order placed successfully")

	return order, nil
}

// Update changes the existing order status with the new one for given order id.
func (bl *BL) Update(ctx context.Context, orderID, status int) error {

	order, err := bl.dl.Get(ctx, orderID)
	if err != nil {
		_ = bl.logger.Log("Method", "Update", "Error", err.Error())
		return err
	}

	if order.Status > status || (order.Status == Completed && status == Cancelled) {
		err := errors.New("Invalid order status, couldn't update order status")
		_ = bl.logger.Log("Method", "Update", "Error", err.Error())
		return err
	}

	err = bl.dl.Update(ctx, orderID, status)
	if err != nil {
		_ = bl.logger.Log("Method", "Update", "Error", err.Error())
		return err
	}

	return nil
}
