package cogol

import (
	"reflect"
)

type assertion struct {
	expected interface{}
	actual   interface{}
	ctx      *Context
	kill     func(ctx *Context)
}

func (ctx *Context) Expect(actual interface{}) *assertion {
	return &assertion{nil, actual, ctx, kill}
}

func kill(ctx *Context) {
	ctx.Kill()
}

func (a *assertion) ToBe(expected interface{}) {
	a.expected = expected
	if reflect.TypeOf(a.expected) != reflect.TypeOf(a.actual) {
		a.kill(a.ctx)
	}

	if !reflect.DeepEqual(a.expected, a.actual) {
		a.kill(a.ctx)
	}
}

func (a *assertion) NotToBe(unexpected interface{}) {
	if reflect.DeepEqual(a.actual, unexpected) {
		a.kill(a.ctx)
	}
}

func (a *assertion) ToBeNil() {
	if a.actual != nil {
		a.kill(a.ctx)
	}
}

func (a *assertion) ToBeNotNil() {
	if a.actual == nil {
		a.kill(a.ctx)
	}
}

func (a *assertion) ToBeZero() {
	if !reflect.ValueOf(a.actual).IsZero() {
		a.kill(a.ctx)
	}
}

func (a *assertion) ToBeNotZero() {
	if reflect.ValueOf(a.actual).IsZero() {
		a.kill(a.ctx)
	}
}

func (a *assertion) ToBeTrue() {
	if !reflect.DeepEqual(a.actual, true) {
		a.kill(a.ctx)
	}
}

func (a *assertion) ToBeFalse() {
	if !reflect.DeepEqual(a.actual, false) {
		a.kill(a.ctx)
	}
}
