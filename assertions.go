package cogol

import (
	"fmt"
	"reflect"
)

const (
	failureMsg = "Expected %v, got %v"
)

// assertion represents a typical assertion
// Includes expected and actual values, as well as the context (ctx) of the the test
type assertion struct {
	expected interface{}
	actual   interface{}
	ctx      *Context
	kill     func(f *failure)
}

// failure represents a failed assertion, with the test's context (ctx)
// and error message (msg)
type failure struct {
	ctx *Context
	msg string
}

// fail is a more beautiful way to create a failure instance
func (a *assertion) fail(msg string) *failure {
	return &failure{ctx: a.ctx, msg: msg}
}

// UseKiller allows users to use custom assertion killers
func (ctx *Context) UseKiller(killer func(f *failure)) {
	ctx.assertionKiller = killer
}

// Expect creates a new assertion with actual value and nil as an expected one
func (ctx *Context) Expect(actual interface{}) *assertion {
	return &assertion{nil, actual, ctx, ctx.assertionKiller}
}

// defaultKiller just calls ctx.Kill method instead of doing some fancy stuff
func defaultKiller(f *failure) {
	f.ctx.Kill(f.msg)
}

// ToBe receives an expected assertion value and then compares it to the actual one, if they do not match, killer-method id called
func (a *assertion) ToBe(expected interface{}) {
	a.expected = expected
	if a.actual == nil {
		a.kill(a.fail(
			fmt.Sprintf(failureMsg,
				a.expected, "nil",
			)),
		)
		return
	}

	if reflect.TypeOf(a.expected) != reflect.TypeOf(a.actual) {
		a.kill(a.fail(
			fmt.Sprintf(failureMsg,
				"type to be "+reflect.TypeOf(a.expected).String(),
				reflect.TypeOf(a.actual).String()),
		))
		return
	}

	if !reflect.DeepEqual(a.expected, a.actual) {
		a.kill(a.fail(
			fmt.Sprintf(failureMsg, a.expected, a.actual),
		))
		return
	}
	a.ctx.test.success = true
}

func (a *assertion) ToBeNot(unexpected interface{}) {
	if reflect.DeepEqual(a.actual, unexpected) {
		a.kill(a.fail(
			fmt.Sprintf("Expected %v to be not %v", a.actual, unexpected),
		))
		return
	}
	a.ctx.test.success = true
}

func (a *assertion) ToBeNil() {
	if a.actual != nil {
		a.kill(a.fail(
			fmt.Sprintf(failureMsg, "nil", a.actual),
		))
		return
	}
	a.ctx.test.success = true
}

func (a *assertion) ToBeNotNil() {
	if a.actual == nil {
		a.kill(a.fail(
			fmt.Sprintf(failureMsg, "not nil", "nil"),
		))
		return
	}
	a.ctx.test.success = true
}

func (a *assertion) ToBeZero() {
	if !reflect.ValueOf(a.actual).IsZero() {
		a.kill(a.fail(
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
		a.kill(a.fail(
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
		a.kill(a.fail(
			fmt.Sprintf(failureMsg, "true", a.actual),
		))
		return
	}
	a.ctx.test.success = true
}

func (a *assertion) ToBeFalse() {
	if !reflect.DeepEqual(a.actual, false) {
		a.kill(a.fail(
			fmt.Sprintf(failureMsg, "false", a.actual),
		))
		return
	}
	a.ctx.test.success = true
}
