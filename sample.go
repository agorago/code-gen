package api

import "context"

type Request struct {
	X string `json:"x"`
}

type Response struct {
	R string `json:"r,omitempty"`
}

type Service interface {
	Request(ctx context.Context, request *Request) (Response, error)
}