package err

import (
	"context"

	bpluse "github.com/MenaEnergyVentures/bplus/err"
)

// It is recommended that each module define its own error file

func internalMakeBplusError(ctx context.Context, ll bpluse.LogLevel, e BPlusErrorCode, args ...interface{}) bpluse.BPlusError {
	return bpluse.MakeErr(ctx, ll, e, ErrMessages[e], args...)
}

// MakeBplusError - returns a customized CAFUError for BPlus
func MakeBplusError(ctx context.Context, e BPlusErrorCode, args ...interface{}) bpluse.BPlusError {
	return internalMakeBplusError(ctx, bpluse.Error, e, args...)

}

// MakeBplusWarning - returns a customized CAFUError for BPlus
func MakeBplusWarning(ctx context.Context, e BPlusErrorCode, args ...interface{}) bpluse.BPlusError {
	return internalMakeBplusError(ctx, bpluse.Warning, e, args...)

}

// BPlusErrorCode - A BPlus error code
type BPlusErrorCode = int

// enumeration for B Plus Error codes
const (
	CannotInvokeOperation BPlusErrorCode = iota + {{.BaseErrorCode}}	
)

// ErrMessages - list of all messages corresponding to this code
var ErrMessages = map[BPlusErrorCode]string{
	CannotInvokeOperation:                     "Operation %s:%s could not be invoked. error is %s",
}
