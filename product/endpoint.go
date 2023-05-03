package product

import (
	"context"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
	"github.com/yamadev11/e-commerce/product/spec"
)

func makeEndpoint(svc *BL, logger log.Logger) spec.Endpoints {

	var listEP endpoint.Endpoint
	{
		listEP = makeListEndpoint(svc)
		listEP = addEndpointMiddleware(logger, listEP, "ListProduct", 0)
	}

	var updateQuantityEP endpoint.Endpoint
	{
		updateQuantityEP = makeUpdateEndpoint(svc)
		updateQuantityEP = addEndpointMiddleware(logger, updateQuantityEP, "UpdateQuantity", 0)
	}

	return spec.Endpoints{
		ListEP:           listEP,
		UpdateQuantityEP: updateQuantityEP,
	}
}

func makeListEndpoint(svc *BL) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.List(ctx)
	}
}

func makeUpdateEndpoint(svc *BL) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, _ := request.(spec.UpdateRequest)
		return nil, svc.UpdateQuantity(ctx, req.ID, req.Quantity)
	}
}

func addEndpointMiddleware(logger log.Logger, ep endpoint.Endpoint, epName string, elapsed time.Duration) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		defer func(begin time.Time) {
			d := time.Since(begin)
			if (elapsed != 0) && (d < elapsed) {
				return
			}

			_ = logger.Log("Layer", "EndpointLayer", "Endpoint", epName, "ElapsedTime", d)
		}(time.Now())

		return ep(ctx, request)
	}
}
