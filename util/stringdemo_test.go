package util_test

import "context"

// UpperCaseRequest - the payload for Uppercase service
type UpperCaseRequest struct {
	S string `json:"s"`
}

// UpperCaseResponse - the  Uppercase service response
type UpperCaseResponse struct {
	V   string `json:"v"`
}

// CountRequest - the payload for Count service
type CountRequest struct {
	S string `json:"s"`
}

// CountResponse - the  Count service response
type CountResponse struct {
	V int `json:"v"`
}

// AddNumbersResponse - the  AddNumbers service response
type AddNumbersResponse struct {
	Sum int `json:"sum"`
}

// StringDemoService - the interface that is going to be implemented by the string demo service
// This has methods to illustrate features of the BPlus framework
type StringDemoService interface {
	// Uppercase - Converts the input string into upper case
	Uppercase(ctx context.Context, ucr *UpperCaseRequest) (UpperCaseResponse, error)
	// Count - returns the length of the input string
	Count(ctx context.Context, cr *CountRequest) (CountResponse, error)
	// AddNumbers - adds two numbers and returns the result
	// This method illustrates a GET method implementation in BPlus since there is no request payload required
	AddNumbers(ctx context.Context, arg1 int, arg2 int) (AddNumbersResponse, error)
}
