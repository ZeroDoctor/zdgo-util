package zdutil

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

// RetryDurationOpt returns a RetryOption that sets the time duration
// between retries. This overrides the default duration of 3 seconds.
func RetryDurationOpt(d time.Duration) RetryOption {
	return func(ro *RetryOpt) {
		ro.Duration = d
	}
}

// RetryAmountOpt returns a RetryOption that sets the amount of retries.
// This overrides the default of 5 retries.
func RetryAmountOpt(a uint) RetryOption {
	return func(ro *RetryOpt) {
		ro.Amount = a
	}
}

// RetryContextOpt returns a RetryOption that sets the context.Context
// to be used for all retry attempts. This overrides the default context
// of context.Background().
func RetryContextOpt(c context.Context) RetryOption {
	return func(ro *RetryOpt) {
		ro.Ctx = c
	}
}

// Retry will run the provided function until it returns nil or
// the retry amount is exceeded. Between each retry, Retry will
// wait for the specified duration before attempting again.
// If the context.Context is cancelled before the retry limit is
// exceeded, Retry will return an error stating that the retry
// was cancelled.
//
// The function will return the last error returned by the
// function, or an error stating that the retry limit was exceeded.
// The returned error will also contain the name of the function
// that was retried.
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
