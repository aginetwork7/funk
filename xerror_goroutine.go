package xerror

import (
	"fmt"
	"github.com/pubgo/xerror/xerror_util"
	"sync"
)

func defaultGoroutineErrHandle(err XErr) {
	if isErrNil(err) {
		return
	}

	fmt.Println(err.Println())
}

var goroutineErrHandleMap sync.Map
var goroutineErrs = make(chan *goroutineErrEvent, 1)

func init() {
	go func() {
		for {
			select {
			case errEvent := <-goroutineErrs:
				goroutineErrHandle := defaultGoroutineErrHandle
				val, ok := goroutineErrHandleMap.Load(errEvent.name)
				if ok {
					goroutineErrHandle = val.(func(err XErr))
				}
				goroutineErrHandle(errEvent.err)
			}
		}
	}()
}

func RangeGoroutineErrHandler(fn func(name string, fn func(err XErr))) {
	goroutineErrHandleMap.Range(func(key, value interface{}) bool {
		fn(key.(string), value.(func(err XErr)))
		return true
	})
}

func SetGoroutineErrHandler(name string, fn func(err XErr)) error {
	if fn == nil {
		return New("the fn parameters should not be nil")
	}

	_, ok := goroutineErrHandleMap.Load(name)
	if ok {
		return Fmt("%s already exists, fn: %s", name, xerror_util.CallerWithFunc(fn))
	}

	goroutineErrHandleMap.Store(name, fn)
	return nil
}

type goroutineErrEvent struct {
	name string
	err  *xerror
}

func RespGoroutine(name ...string) {
	nm := "__xerror"
	if len(name) > 0 {
		nm = name[0]
	}

	var err error
	handleErr(&err, recover())
	if isErrNil(err) {
		return
	}

	goroutineErrs <- &goroutineErrEvent{name: nm, err: handle(err, "")}
}
