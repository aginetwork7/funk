package try

import (
	"github.com/pubgo/funk/errors"
	"github.com/pubgo/funk/generic"
	"github.com/pubgo/funk/result"
	"github.com/pubgo/funk/stack"
)

func WithErr(gErr *error, fn func() error) {
	if fn == nil {
		*gErr = errors.WrapStack(errors.New("[fn] is nil"))
		return
	}

	defer func() {
		if err := errors.Parse(recover()); !generic.IsNil(err) {
			*gErr = errors.WrapStack(err)
		}

		*gErr = errors.WrapKV(*gErr, "fn_stack", stack.CallerWithFunc(fn).String())
	}()

	*gErr = fn()
}

func Try(fn func() error) (gErr error) {
	if fn == nil {
		gErr = errors.WrapStack(errors.New("[fn] is nil"))
		return
	}

	defer func() {
		if err := errors.Parse(recover()); !generic.IsNil(err) {
			gErr = errors.WrapStack(err)
		}

		gErr = errors.WrapKV(gErr, "fn_stack", stack.CallerWithFunc(fn).String())
	}()

	gErr = fn()
	return
}

func Result[T any](fn func() result.Result[T]) (g result.Result[T]) {
	if fn == nil {
		g = g.WithErr(errors.WrapStack(errors.New("[fn] is nil")))
		return
	}

	defer func() {
		if err := errors.Parse(recover()); !generic.IsNil(err) {
			g = g.WithErr(errors.WrapStack(err))
		}

		if g.IsErr() {
			g = g.WithErr(g.Err(func(err error) error {
				return errors.WrapKV(err, "fn_stack", stack.CallerWithFunc(fn))
			}))
		}
	}()

	g = fn()
	return
}
