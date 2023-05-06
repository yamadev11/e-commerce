package componenttest

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/yamadev11/e-commerce/order/spec"
	"github.com/yamadev11/e-commerce/order/testutil"
)

var testObj *testutil.TestObj

func TestMain(m *testing.M) {
	testObj = testutil.InitTestInfra(&testing.T{})
	m.Run()
}

func Test_CreateOrder(t *testing.T) {

	y, m, d := time.Now().Date()
	orderDate := fmt.Sprintf("%d/%d/%d", d, m, y)

	type test struct {
		name    string
		args    spec.CreateRequest
		setup   func(mock testutil.MockSDK)
		want    *spec.Order
		wantErr bool
	}

	tests := []test{
		{
			name: "Success",
			args: spec.CreateRequest{
				Items: []spec.Item{
					{ID: 1, Quantity: 2},
					{ID: 2, Quantity: 1},
					{ID: 3, Quantity: 1},
				},
			},

			setup: func(mock testutil.MockSDK) {
				mock.Product.EXPECT().List(gomock.Any()).Return(listProductResponse, nil)
				mock.Product.EXPECT().UpdateQuantity(gomock.Any(), 1, -2).Return(nil)
				mock.Product.EXPECT().UpdateQuantity(gomock.Any(), 2, -1).Return(nil)
				mock.Product.EXPECT().UpdateQuantity(gomock.Any(), 3, -1).Return(nil)
			},

			want: &spec.Order{
				ID: CreateOrderID1,
				Items: []spec.Item{
					{ID: 1, Quantity: 2},
					{ID: 2, Quantity: 1},
					{ID: 3, Quantity: 1},
				},
				Amount:      430000,
				Discount:    43000,
				FinalAmount: 387000,
				Status:      1,
				OrderDate:   orderDate,
			},
			wantErr: false,
		},
		{
			name: "Failure-Negative quantity",
			args: spec.CreateRequest{
				Items: []spec.Item{
					{ID: 1, Quantity: -2},
				},
			},

			setup: func(mock testutil.MockSDK) {
				mock.Product.EXPECT().List(gomock.Any()).Return(listProductResponse, nil)
			},

			want:    nil,
			wantErr: true,
		},
		{
			name: "Failure-Quantity beyond limit",
			args: spec.CreateRequest{
				Items: []spec.Item{
					{ID: 1, Quantity: 11},
				},
			},

			setup: func(mock testutil.MockSDK) {
				mock.Product.EXPECT().List(gomock.Any()).Return(listProductResponse, nil)
			},

			want:    nil,
			wantErr: true,
		},
		{
			name: "Failure-invalid productID",
			args: spec.CreateRequest{
				Items: []spec.Item{
					{ID: 101, Quantity: 2},
				},
			},

			setup: func(mock testutil.MockSDK) {
				mock.Product.EXPECT().List(gomock.Any()).Return(listProductResponse, nil)
			},

			want:    nil,
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			test.setup(testObj.MockSDK)

			order, err := testObj.Order.Create(context.TODO(), test.args.Items)
			if (err != nil) != test.wantErr {
				t.Errorf("%s=%s,%s=%s", "Method", "Test_CreateOrder", "Error", err.Error())
			}

			if !reflect.DeepEqual(order, test.want) {
				t.Errorf("%s=%s, %s=%v, %s=%v, %s=%v", "Method", "Test_CreateOrder",
					"Got", order, "Want", test.want, "Diff", cmp.Diff(order, test.want))
			}
		})
	}
}

func Test_GetOrder(t *testing.T) {

	y, m, d := time.Now().Date()
	orderDate := fmt.Sprintf("%d/%d/%d", d, m, y)

	type test struct {
		name    string
		args    spec.GetRequest
		setup   func(mock testutil.MockSDK)
		want    *spec.GetOrderResponse
		wantErr bool
	}

	tests := []test{
		{
			name: "Success",
			args: spec.GetRequest{
				ID: GetOrderID1,
			},

			setup: func(mock testutil.MockSDK) {
				mock.Product.EXPECT().List(gomock.Any()).Return(listProductResponse, nil)
				mock.Product.EXPECT().UpdateQuantity(gomock.Any(), 2, -2).Return(nil)

				mock.Product.EXPECT().List(gomock.Any()).Return(listProductResponse, nil)
			},

			want: &spec.GetOrderResponse{
				ID: GetOrderID1,
				Items: []spec.OrderItem{
					{
						ID:       2,
						Name:     "iPad",
						Quantity: 2,
						Price:    70000,
						Category: "Premium",
					},
				},
				Amount:      140000,
				Discount:    0,
				FinalAmount: 140000,
				Status:      "Placed",
				OrderDate:   orderDate,
			},
			wantErr: false,
		},

		{
			name: "Failure-invalid orderID",
			args: spec.GetRequest{
				ID: 101,
			},

			setup: func(mock testutil.MockSDK) {
				mock.Product.EXPECT().List(gomock.Any()).Return(listProductResponse, nil)
				mock.Product.EXPECT().UpdateQuantity(gomock.Any(), 2, -2).Return(nil)
			},

			want:    nil,
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			test.setup(testObj.MockSDK)

			createOrderRequest := spec.CreateRequest{
				Items: []spec.Item{
					{ID: 2, Quantity: 2},
				},
			}

			_, err := testObj.Order.Create(context.TODO(), createOrderRequest.Items)
			if err != nil {
				t.Errorf("%s=%s,%s=%s", "Method", "Test_GetOrder", "Error", err.Error())
			}

			listResponse, err := testObj.Order.Get(context.TODO(), test.args.ID)
			if (err != nil) != test.wantErr {
				t.Errorf("%s=%s,%s=%s", "Method", "Test_GetOrder", "Error", err.Error())
			}

			if !reflect.DeepEqual(listResponse, test.want) {
				t.Errorf("%s=%s, %s=%v, %s=%v, %s=%v", "Method", "Test_GetOrder",
					"Got", listResponse, "Want", test.want, "Diff", cmp.Diff(listResponse, test.want))
			}
		})
	}
}

func Test_UpdateOrder(t *testing.T) {

	y, m, d := time.Now().Date()
	orderDate := fmt.Sprintf("%d/%d/%d", d, m, y)

	type test struct {
		name    string
		args    spec.UpdateRequest
		setup   func(mock testutil.MockSDK)
		want    *spec.GetOrderResponse
		wantErr bool
	}

	tests := []test{
		{
			name: "Success",
			args: spec.UpdateRequest{
				ID:     UpdateOrderID1,
				Status: 2,
			},

			setup: func(mock testutil.MockSDK) {
				mock.Product.EXPECT().List(gomock.Any()).Return(listProductResponse, nil)
				mock.Product.EXPECT().UpdateQuantity(gomock.Any(), 1, -2).Return(nil)

				mock.Product.EXPECT().List(gomock.Any()).Return(listProductResponse, nil)
			},

			want: &spec.GetOrderResponse{
				ID: UpdateOrderID1,
				Items: []spec.OrderItem{
					{
						ID:       1,
						Name:     "iPhone14",
						Quantity: 2,
						Price:    80000,
						Category: "Premium",
					},
				},
				Amount:       160000,
				Discount:     0,
				FinalAmount:  160000,
				Status:       "Dispatched",
				OrderDate:    orderDate,
				DispatchDate: orderDate,
			},
			wantErr: false,
		},
		{
			name: "Failure-Invalid orderID",
			args: spec.UpdateRequest{
				ID:     101,
				Status: 2,
			},

			setup: func(mock testutil.MockSDK) {
				mock.Product.EXPECT().List(gomock.Any()).Return(listProductResponse, nil)
				mock.Product.EXPECT().UpdateQuantity(gomock.Any(), 1, -2).Return(nil)
			},

			want:    nil,
			wantErr: true,
		},
		{
			name: "Failure-Invalid status",
			args: spec.UpdateRequest{
				ID:     UpdateOrderID3,
				Status: 10,
			},

			setup: func(mock testutil.MockSDK) {
				mock.Product.EXPECT().List(gomock.Any()).Return(listProductResponse, nil)
				mock.Product.EXPECT().UpdateQuantity(gomock.Any(), 1, -2).Return(nil)
			},

			want:    nil,
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			test.setup(testObj.MockSDK)

			createOrderRequest := spec.CreateRequest{
				Items: []spec.Item{
					{ID: 1, Quantity: 2},
				},
			}

			_, err := testObj.Order.Create(context.TODO(), createOrderRequest.Items)
			if err != nil {
				t.Errorf("%s=%s,%s=%s", "Method", "Test_UpdateOrder", "Error", err.Error())
			}

			err = testObj.Order.Update(context.TODO(), test.args.ID, test.args.Status)
			if (err != nil) != test.wantErr {
				t.Errorf("%s=%s,%s=%s", "Method", "Test_UpdateOrder", "Error", err.Error())
			}

			if !test.wantErr {

				listResponse, err := testObj.Order.Get(context.TODO(), test.args.ID)
				if (err != nil) != test.wantErr {
					t.Errorf("%s=%s,%s=%s", "Method", "Test_UpdateOrder", "Error", err.Error())
				}

				if !reflect.DeepEqual(listResponse, test.want) {
					t.Errorf("%s=%s, %s=%v, %s=%v, %s=%v", "Method", "Test_UpdateOrder",
						"Got", listResponse, "Want", test.want, "Diff", cmp.Diff(listResponse, test.want))
				}
			}
		})
	}
}
