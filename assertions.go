package cogol

import (
	"fmt"
	"reflect"
)

const (
	failureMsg = "Expected %v, got %v"
)

type assertion struct {
	expected interface{}
	actual   interface{}
	ctx      *Context
	kill     func(f *failure)
}

type failure struct {
	ctx *Context
	msg string
}

func (a *assertion) Fail(msg string) *failure {
	return &failure{ctx: a.ctx, msg: msg}
}

func (ctx *Context) Expect(actual interface{}) *assertion {
	return &assertion{nil, actual, ctx, defaultKiller}
}

func defaultKiller(f *failure) {
	f.ctx.test.f = f
	f.ctx.Kill(f)
}

func (a *assertion) ToBe(expected interface{}) {
	a.expected = expected
	if a.actual == nil {
		a.kill(a.Fail(
			fmt.Sprintf(failureMsg,
				a.expected, "nil",
			)),
		)
		return
	}

	if reflect.TypeOf(a.expected) != reflect.TypeOf(a.actual) {
		a.kill(a.Fail(
			fmt.Sprintf(failureMsg,
				"type to be "+reflect.TypeOf(a.expected).String(),
				reflect.TypeOf(a.actual).String()),
		))
		return
	}

	if !reflect.DeepEqual(a.expected, a.actual) {
		a.kill(a.Fail(
			fmt.Sprintf(failureMsg, a.expected, a.actual),
		))
		return
	}
	a.ctx.test.success = true
}

func (a *assertion) ToBeNot(unexpected interface{}) {
	if reflect.DeepEqual(a.actual, unexpected) {
		a.kill(a.Fail(
			fmt.Sprintf("Expected %v to be not %v", a.actual, unexpected),
		))
		return
	}
	a.ctx.test.success = true
}

func (a *assertion) ToBeNil() {
	if a.actual != nil {
		a.kill(a.Fail(
			fmt.Sprintf(failureMsg, "nil", a.actual),
		))
		return
	}
	a.ctx.test.success = true
}

func (a *assertion) ToBeNotNil() {
	if a.actual == nil {
		a.kill(a.Fail(
			fmt.Sprintf(failureMsg, "not nil", "nil"),
		))
		return
	}
	a.ctx.test.success = true
}

func (a *assertion) ToBeZero() {
	if !reflect.ValueOf(a.actual).IsZero() {
		a.kill(a.Fail(
			fmt.Sprintf(failureMsg,
				"zero value for "+reflect.TypeOf(a.actual).String()+" type",
				a.actual),
		))
		return
	}
	a.ctx.test.success = true
}

func (a *assertion) ToBeNotZero() {
	if reflect.ValueOf(a.actual).IsZero() {
		a.kill(a.Fail(
			fmt.Sprintf(failureMsg,
				"not zero value for "+reflect.TypeOf(a.actual).String()+" type",
				"zero value"),
		))
		return
	}
	a.ctx.test.success = true
}

func (a *assertion) ToBeTrue() {
	if !reflect.DeepEqual(a.actual, true) {
		a.kill(a.Fail(
			fmt.Sprintf(failureMsg, "true", a.actual),
		))
		return
	}
	a.ctx.test.success = true
}

func (a *assertion) ToBeFalse() {
	if !reflect.DeepEqual(a.actual, false) {
		a.kill(a.Fail(
			fmt.Sprintf(failureMsg, "false", a.actual),
		))
		return
	}
	a.ctx.test.success = true
}
