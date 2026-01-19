package errors_test

import (
	"testing"

	"github.com/pubgo/funk/errors"
	"github.com/pubgo/funk/proto/testcodepb"
)

func TestErrCode(t *testing.T) {
	err1 := errors.Wrap(errors.NewCodeErr(testcodepb.TestErrCodeDbConn), "hello")
	rr := errors.Is(err1, errors.NewCodeErr(testcodepb.TestErrCodeDbConn))
	if !rr {
		t.Fatal("not match")
	}

	t.Log(errors.As(err1, testcodepb.TestErrCodeDbConn))
}

func TestNewCodeErrWithMsg(t *testing.T) {
	err1 := errors.NewCodeErrWithMsg(testcodepb.TestErrCodeDbConn, "hello world")
	var err2 *errors.ErrCode
	if !errors.As(err1, &err2) {
		t.Fatal("errors.As should have returned true")
	}
	if err2.Error() != "Hello World" {
		t.Fatal("message not match")
	}
	t.Log(err2.Error())
}
