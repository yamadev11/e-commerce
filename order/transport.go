package order

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/pkg/errors"
	"github.com/yamadev11/e-commerce/order/spec"

	kitHttp "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/gorilla/mux"
)

func AddHandlers(router *mux.Router, eps spec.Endpoints, logger log.Logger) {

	router.Methods(http.MethodGet).Path(spec.GetPath).Handler(kitHttp.NewServer(
		eps.GetEP,
		decodeGetRequest,
		JSONEncodeAPIResponse,
	))

	router.Methods(http.MethodPost).Path(spec.CreatePath).Handler(kitHttp.NewServer(
		eps.CreateEP,
		decodeCreateRequest,
		JSONEncodeAPIResponse,
	))

	router.Methods(http.MethodPatch).Path(spec.UpdatePath).Handler(kitHttp.NewServer(
		eps.UpdateEP,
		decodeUpdateRequest,
		JSONEncodeAPIResponse,
	))

}

func JSONEncodeAPIResponse(ctx context.Context, w http.ResponseWriter, resp interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(resp)
}

func decodeGetRequest(ctx context.Context, r *http.Request) (interface{}, error) {

	id, err := getOrderID(r)
	if err != nil {
		return nil, err
	}

	return spec.GetRequest{ID: id}, nil
}

func decodeUpdateRequest(ctx context.Context, r *http.Request) (interface{}, error) {

	id, err := getOrderID(r)
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

func decodeCreateRequest(ctx context.Context, r *http.Request) (interface{}, error) {

	createRequest := spec.CreateRequest{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&createRequest)
	if err != nil {
		return nil, err
	}

	return createRequest, nil
}

func getOrderID(r *http.Request) (int, error) {
	vars := mux.Vars(r)
	orderIDStr, exists := vars["ID"]
	if !exists {
		err := errors.New("key not found")
		err = errors.Wrapf(err, "msg=Failed to get value for key orderID")
		return -1, err
	}

	orderID, conversionError := strconv.ParseInt(orderIDStr, 10, 32)
	if conversionError != nil || orderID < 0 {
		err := errors.New("Invalid orderID")
		return -1, err
	}

	return int(orderID), nil
}
