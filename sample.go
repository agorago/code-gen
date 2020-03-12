package api

import "context"

type Request struct {
	X string `json:"x" validate:"required"`
	Y string `json:"y"`
}

type Response struct {
	R string `json:"r,omitempty"`
}

type Service interface {
	RequestBodyResponseBody(ctx context.Context, request *Request) (Response, error)
	RequestBodyNoResponseBody(ctx context.Context, request *Request) error
	NoRequestBodyNoResponseBody(ctx context.Context) error
}
