package spec

import "github.com/go-kit/kit/endpoint"

type Endpoints struct {
	GetEP    endpoint.Endpoint
	CreateEP endpoint.Endpoint
	UpdateEP endpoint.Endpoint
}
