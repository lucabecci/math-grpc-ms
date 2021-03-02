package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/lucabecci/math-grpc-ms/service"
)

type Endpoints struct {
	Add endpoint.Endpoint
}

type MathReq struct {
	NumA float32
	NumB float32
}

type MathRes struct {
	Result float32
}

func MakeEndpoints(s service.Service) Endpoints {
	return Endpoints{
		Add: MakeAddEndpoints(s),
	}
}

func MakeAddEndpoints(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(MathReq)
		result, _ := s.Add(ctx, req.NumA, req.NumB)
		return MathRes{Result: result}, nil
	}
}
