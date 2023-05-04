package order

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/yamadev11/e-commerce/order/spec"
)

type DL struct{}

func NewDL() *DL {
	return &DL{}
}

func (dl *DL) List(ctx context.Context) map[int]*spec.Order {
	return orders
}

func (dl *DL) Get(ctx context.Context, orderID int) (*spec.Order, error) {

	if _, ok := orders[orderID]; !ok {
		err := errors.New("Invalid orderID, couldn't get order details")
		return nil, err
	}

	return orders[orderID], nil
}

func (bl *DL) Create(ctx context.Context, items []spec.Item, amount, discount float64) (*spec.Order, error) {

	y, m, d := time.Now().Date()
	orderDate := fmt.Sprintf("%d/%d/%d", d, m, y)

	order := spec.Order{
		ID:          currentOrderID,
		Items:       items,
		Amount:      amount,
		Discount:    discount,
		FinalAmount: amount - discount,
		Status:      Placed,
		OrderDate:   orderDate,
	}

	orders[currentOrderID] = &order
	currentOrderID += 1

	return &order, nil
}

func (dl *DL) Update(ctx context.Context, orderID, status int) error {

	order := orders[orderID]
	order.Status = status
	if status == Dispatched {
		y, m, d := time.Now().Date()
		order.DispatchDate = fmt.Sprintf("%d/%d/%d", d, m, y)
	}

	return nil
}

const (
	Placed = iota + 1
	Dispatched
	Completed
	Cancelled
)

var (
	currentOrderID int = 1

	orders = map[int]*spec.Order{}

	OrderStatusMap = map[int]string{
		Placed:     "Placed",
		Dispatched: "Dispatched",
		Completed:  "Completed",
		Cancelled:  "Cancelled",
	}
)
