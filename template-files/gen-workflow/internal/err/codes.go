package err

import (
	"context"

	bpluse "gitlab.intelligentb.com/devops/bplus/err"
)

// It is recommended that each module define its own error file

func internalMakeBplusError(ctx context.Context, ll bpluse.LogLevel, e BPlusErrorCode, args map[string]interface{}) bpluse.BPlusError {
	return bpluse.MakeErr(ctx, ll, int(e), e.String(), args)
}

// MakeBplusError - returns a customized CAFUError for BPlus
func MakeBplusError(ctx context.Context, e BPlusErrorCode, args map[string]interface{}) bpluse.BPlusError {
	return internalMakeBplusError(ctx, bpluse.Error, e, args)

}

// MakeBplusWarning - returns a customized CAFUError for BPlus
func MakeBplusWarning(ctx context.Context, e BPlusErrorCode, args map[string]interface{}) bpluse.BPlusError {
	return internalMakeBplusError(ctx, bpluse.Warning, e, args)

}

// MakeBplusErrorWithErrorCode - returns a customized CAFUError for BPlus
func MakeBplusErrorWithErrorCode(ctx context.Context, httpErrorCode int, e BPlusErrorCode, args map[string]interface{}) bpluse.BPlusError {
	return internalMakeBplusError(ctx, bpluse.Error, e, args)

}

// MakeBplusWarningWithErrorCode - returns a customized CAFUError for BPlus
func MakeBplusWarningWithErrorCode(ctx context.Context, httpErrorCode int, e BPlusErrorCode, args map[string]interface{}) bpluse.BPlusError {
	return internalMakeBplusError(ctx, bpluse.Warning, e, args)

}

// BPlusErrorCode - A BPlus error code
type BPlusErrorCode int

// enumeration for B Plus Error codes
const (
	CannotInvokeOperation BPlusErrorCode = iota + 200000
	SecurityException
)

//go:generate stringer -type=BPlusErrorCode
