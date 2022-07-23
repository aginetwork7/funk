package logx

import (
	"sync/atomic"

	"github.com/go-logr/logr"
)

var _ logr.LogSink = (*sink)(nil)
var _ logr.CallDepthLogSink = (*sink)(nil)

type sink struct {
	level     int
	callDepth int
	prefix    string
	values    []interface{}
	changeNum int32
	cacheLog  logr.Logger
}

func (s sink) WithCallDepth(depth int) logr.LogSink {
	s.callDepth += depth
	return &s
}

func (s *sink) Init(info logr.RuntimeInfo) {
	s.callDepth += info.CallDepth
}

func (s sink) Enabled(level int) bool {
	return defaultLog.GetSink().Enabled(level)
}

func (s *sink) Info(level int, msg string, keysAndValues ...interface{}) {
	if atomic.LoadInt32(&s.changeNum) != changeNum {
		s.cacheLog = defaultLog.V(level).WithCallDepth(s.callDepth).WithName(s.prefix).WithValues(s.values...)
		atomic.StoreInt32(&s.changeNum, changeNum)
	}
	s.cacheLog.Info(msg, keysAndValues...)
}

func (s *sink) Error(err error, msg string, keysAndValues ...interface{}) {
	if atomic.LoadInt32(&s.changeNum) != changeNum {
		s.cacheLog = defaultLog.WithCallDepth(s.callDepth).WithName(s.prefix).WithValues(s.values...)
		atomic.StoreInt32(&s.changeNum, changeNum)
	}
	s.cacheLog.Error(err, msg, keysAndValues...)
}

func (s sink) WithValues(keysAndValues ...interface{}) logr.LogSink {
	s.values = append(s.values, keysAndValues...)
	return &s
}

func (s sink) WithName(name string) logr.LogSink {
	if len(s.prefix) > 0 {
		s.prefix = s.prefix + "."
	}
	s.prefix += name
	return &s
}
