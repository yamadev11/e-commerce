package product

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/pkg/errors"
	"github.com/yamadev11/e-commerce/product/spec"

	kitHttp "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/gorilla/mux"
)

func AddHandlers(router *mux.Router, userEPs spec.Endpoints, logger log.Logger) error {

	router.Methods(http.MethodGet).Path(spec.ListPath).Handler(kitHttp.NewServer(
		userEPs.ListEP,
		decodeListRequest,
		JSONEncodeAPIResponse,
	))

	router.Methods(http.MethodPut).Path(spec.UpdateQuantityPath).Handler(kitHttp.NewServer(
		userEPs.UpdateQuantityEP,
		decodeUpdateRequest,
		JSONEncodeAPIResponse,
	))

	return nil
}

func JSONEncodeAPIResponse(ctx context.Context, w http.ResponseWriter, resp interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(resp)
}

func decodeListRequest(ctx context.Context, r *http.Request) (interface{}, error) {

	return nil, nil
}

func decodeUpdateRequest(ctx context.Context, r *http.Request) (interface{}, error) {

	id, err := getProductID(r)
	if err != nil {
		return nil, err
	}

	updateRequest := spec.UpdateRequest{}

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&updateRequest)
	if err != nil {
		return nil, err
	}

	updateRequest.ID = id
	return updateRequest, nil
}

func getProductID(r *http.Request) (int, error) {
	vars := mux.Vars(r)
	productIDStr, exists := vars["ID"]
	if !exists {
		err := errors.New("key not found")
		err = errors.Wrapf(err, "msg=Failed to get value for key productID")
		return -1, err
	}

	productID, conversionError := strconv.ParseInt(productIDStr, 10, 32)
	if conversionError != nil || productID < 0 {
		err := errors.New("Invalid productID")
		return -1, err
	}

	return int(productID), nil
}
