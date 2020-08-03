package xerror

import (
	"github.com/pubgo/xerror/internal/wrapper"
	"strings"
)

type xerrorBase struct {
	Code1  string `json:"code,omitempty"`
	Msg    string `json:"msg,omitempty"`
	Caller string `json:"caller,omitempty"`
}

func (t *xerrorBase) Error() string {
	return t.Code1
}

func (t *xerrorBase) As(err interface{}) bool {
	if t == nil || err == nil {
		return false
	}

	switch e := err.(type) {
	case **xerrorBase:
		return strings.HasPrefix(t.Code1, (*e).Code1)
	case *xerrorBase:
		return strings.HasPrefix(t.Code1, e.Code1)
	case *error:
		return strings.HasPrefix(t.Code1, (*e).Error())
	case error:
		return strings.HasPrefix(t.Code1, e.Error())
	case *string:
		return strings.HasPrefix(t.Code1, *e)
	case string:
		return strings.HasPrefix(t.Code1, e)
	default:
		return false
	}
}

func (t *xerrorBase) FamilyAs(err interface{}) bool {
	return t.As(err)
}

func (t *xerrorBase) New(code string, ms ...string) error {
	var msg string
	if len(ms) > 0 {
		msg = ms[0]
	}

	code = t.Code1 + "->" + code
	xw := &xerrorBase{}
	xw.Code1 = code
	xw.Msg = msg
	xw.Caller = callerWithDepth(wrapper.CallDepth())

	return xw
}
