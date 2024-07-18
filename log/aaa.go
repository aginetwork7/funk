package log

import (
	"context"

	"github.com/rs/zerolog"
)

const (
	ModuleName = "module"
)

type (
	Map           = map[string]any
	Hook          = zerolog.Hook
	Event         = zerolog.Event
	Level         = zerolog.Level
	EnableChecker = func(ctx context.Context, lvl Level, nameOrMessage string, fields Map) bool
)

type Logger interface {
	WithName(name string) Logger
	WithFields(m Map) Logger
	WithCallerSkip(skip int) Logger
	WithEvent(evt *Event) Logger
	WithLevel(lvl Level) Logger

	Debug(ctx ...context.Context) *Event
	Info(ctx ...context.Context) *Event
	Warn(ctx ...context.Context) *Event
	Error(ctx ...context.Context) *Event
	Err(err error, ctx ...context.Context) *Event
	Panic(ctx ...context.Context) *Event
	Fatal(ctx ...context.Context) *Event
}

type StdLogger interface {
	Printf(format string, v ...interface{})
	Logf(format string, v ...interface{})
	Print(v ...interface{})
	Log(v ...interface{})
	Println(v ...interface{})
}
