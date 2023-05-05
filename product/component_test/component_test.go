package component_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/yamadev11/e-commerce/product"
	"github.com/yamadev11/e-commerce/product/sdk"
	"github.com/yamadev11/e-commerce/product/spec"
	"github.com/yamadev11/e-commerce/product/testutil"
)

func Test_ListProduct(t *testing.T) {

	testutil.InitTestInfra()
	productSDK := sdk.NewProduct(testutil.TestProductPort)

	type test struct {
		name    string
		want    *spec.ListResponse
		wantErr bool
	}

	tests := []test{
		{
			name: "Success",
			want: &spec.ListResponse{
				Products: []spec.Product{
					{
						ID:          1,
						Name:        "iPhone14",
						Price:       80000,
						Category:    product.CategoryIDNameMap[product.Premium],
						AvlQuantity: 15,
					},
					{
						ID:          2,
						Name:        "iPad",
						Price:       70000,
						Category:    product.CategoryIDNameMap[product.Premium],
						AvlQuantity: 12,
					},
					{
						ID:          3,
						Name:        "macBook Pro",
						Price:       200000,
						Category:    product.CategoryIDNameMap[product.Premium],
						AvlQuantity: 6,
					},
					{
						ID:          4,
						Name:        "Thinkpad E470",
						Price:       70000,
						Category:    product.CategoryIDNameMap[product.Regular],
						AvlQuantity: 25,
					},
					{
						ID:          5,
						Name:        "Ideapad 500",
						Price:       50000,
						Category:    product.CategoryIDNameMap[product.Budget],
						AvlQuantity: 11,
					},
				},
			},
			wantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			listResponse, err := productSDK.List(context.TODO())
			if (err != nil) != test.wantErr {
				t.Errorf("%s=%s,%s=%s", "Method", "Test_ListProduct", "Error", err.Error())
				return
			}

			if !reflect.DeepEqual(listResponse, test.want) {
				t.Errorf("%s=%s, %s=%v, %s=%v, %s=%v", "Method", "Test_ListProduct",
					"Got", listResponse, "Want", test.want, "Diff", cmp.Diff(listResponse, test.want))
			}
		})
	}
}

func Test_UpdateQuantity(t *testing.T) {

	testutil.InitTestInfra()
	productSDK := sdk.NewProduct(testutil.TestProductPort)

	type test struct {
		name    string
		args    spec.UpdateRequest
		wantErr bool
	}

	tests := []test{
		{
			name: "Success",
			args: spec.UpdateRequest{
				ID:       1,
				Quantity: 3,
			},
			wantErr: false,
		},
		{
			name: "Failure",
			args: spec.UpdateRequest{
				ID:       101, // invalid product id
				Quantity: 3,
			},
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			err := productSDK.UpdateQuantity(context.TODO(), test.args.ID, test.args.Quantity)
			if (err != nil) != test.wantErr {
				t.Errorf("%s=%s,%s=%s", "Method", "Test_UpdateQuantity", "Error", err.Error())
				return
			}
		})
	}

}
