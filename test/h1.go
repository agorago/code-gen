package main

import (
	"context"
)

type SomeFuncRequest struct {
	X string
}

type SomeFuncResponse struct {
	Y string
}

type SomeOtherFuncResponse struct {
	Z string
}

type H1 interface {
	SomeFunc(ctx context.Context, a int, request SomeFuncRequest) (SomeFuncResponse, error)
	SomeOtherFunc(ctx context.Context, b int) (SomeOtherFuncResponse, error)
}
