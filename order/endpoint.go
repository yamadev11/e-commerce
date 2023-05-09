package order

import (
	"context"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
	"github.com/yamadev11/e-commerce/order/spec"
)

func makeEndpoint(svc *BL, logger log.Logger) spec.Endpoints {

	var getEP endpoint.Endpoint
	{
		getEP = makeGetEndpoint(svc)
		getEP = addEndpointMiddleware(logger, getEP, "GetOrder", 0)
	}

	var createEP endpoint.Endpoint
	{
		createEP = makeCreateEndpoint(svc)
		createEP = addEndpointMiddleware(logger, createEP, "CreateOrder", 0)
	}

	var updateEP endpoint.Endpoint
	{
		updateEP = makeUpdateEndpoint(svc)
		updateEP = addEndpointMiddleware(logger, updateEP, "UpdateOrder", 0)
	}

	return spec.Endpoints{
		GetEP:    getEP,
		CreateEP: createEP,
		UpdateEP: updateEP,
	}
}

func makeGetEndpoint(svc *BL) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(spec.GetRequest)
		return svc.Get(ctx, req.ID)
	}
}

func makeCreateEndpoint(svc *BL) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(spec.CreateRequest)
		return svc.Create(ctx, req.Items)
	}
}

func makeUpdateEndpoint(svc *BL) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(spec.UpdateRequest)
		return nil, svc.Update(ctx, req.ID, req.Status)
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
