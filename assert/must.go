package assert

import (
	"fmt"
	"os"
	"runtime/debug"

	"github.com/pubgo/funk/errors"
	"github.com/pubgo/funk/try"
)

func Must(err error, args ...interface{}) {
	if err == nil {
		return
	}

	panic(errors.WrapStack(errors.Wrap(err, fmt.Sprint(args...))))
}

func MustFn(errFn func() error, args ...interface{}) {
	err := try.Try(errFn)
	if err == nil {
		return
	}

	panic(errors.WrapStack(errors.Wrap(err, fmt.Sprint(args...))))
}

func MustF(err error, msg string, args ...interface{}) {
	if err == nil {
		return
	}

	panic(errors.WrapStack(errors.Wrap(err, fmt.Sprintf(msg, args...))))
}

func Must1[T any](ret T, err error) T {
	if err != nil {
		panic(errors.WrapStack(err))
	}

	return ret
}

func Exit(err error, args ...interface{}) {
	if err == nil {
		return
	}

	errors.Debug(errors.WrapStack(errors.Wrap(err, fmt.Sprint(args...))))
	debug.PrintStack()
	os.Exit(1)
}

func ExitFn(errFn func() error, args ...interface{}) {
	err := try.Try(errFn)
	if err == nil {
		return
	}

	errors.Debug(errors.WrapStack(errors.Wrap(err, fmt.Sprint(args...))))
	debug.PrintStack()
	os.Exit(1)
}

func ExitF(err error, msg string, args ...interface{}) {
	if err == nil {
		return
	}

	errors.Debug(errors.WrapStack(errors.Wrapf(err, msg, args...)))
	debug.PrintStack()
	os.Exit(1)
}

func Exit1[T any](ret T, err error) T {
	if err != nil {
		errors.Debug(errors.WrapStack(err))
		debug.PrintStack()
		os.Exit(1)
	}

	return ret
}
