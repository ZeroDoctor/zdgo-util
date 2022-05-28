package zdgoutil

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"time"
)

type RetryOpt struct {
	Ctx      context.Context
	Amount   uint
	Duration time.Duration
}

type RetryOption func(*RetryOpt)

func RetryDurationOpt(d time.Duration) RetryOption {
	return func(ro *RetryOpt) {
		ro.Duration = d
	}
}

func RetryAmountOpt(a uint) RetryOption {
	return func(ro *RetryOpt) {
		ro.Amount = a
	}
}

func RetryContextOpt(c context.Context) RetryOption {
	return func(ro *RetryOpt) {
		ro.Ctx = c
	}
}

// Retry attempts to call a function. If functions fails it will
// wait a certain duration defined by RetryOpt before calling the function again.
// Also, defining RetryOpt is optional; defaults:
//
//   Ctx:      context.Background()
//   Amount:   5
//   Duration: 3 * time.Secound
//
func Retry(fn func() error, opts ...RetryOption) error {
	var err error

	retry := &RetryOpt{
		Ctx:      context.Background(),
		Amount:   5,
		Duration: 3 * time.Second,
	}

	for _, opt := range opts {
		opt(retry)
	}

	var count uint
	for count < retry.Amount {
		err = fn()
		if err != nil {
			select {
			case <-time.After(retry.Duration):
			case <-retry.Ctx.Done():
				return errors.New("retry is cancelled")
			}
		} else {
			return nil
		}

		count++
	}

	funcPath := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
	lastSlash := strings.LastIndex(funcPath, "/")
	funcName := funcPath[lastSlash+1:]

	return fmt.Errorf("[function=%s] failed after [count=%d] retries\n\t[error=%w]", funcName, count, err)

}
