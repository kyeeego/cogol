package cogol

import (
	"reflect"
)

type assertion struct {
	expected interface{}
	actual   interface{}
	ctx      *Context
}

func (ctx *Context) Expect(actual interface{}) *assertion {
	return &assertion{nil, actual, ctx}
}

func (a *assertion) ToBe(expected interface{}) {
	a.expected = expected
	if reflect.TypeOf(a.expected) != reflect.TypeOf(a.actual) {
		a.ctx.Kill()
	}

	if !reflect.DeepEqual(a.expected, a.actual) {
		a.ctx.Kill()
	}
}

func (a *assertion) NotToBe(unexpected interface{}) {
	if reflect.DeepEqual(a.actual, unexpected) {
		a.ctx.Kill()
	}
}

func (a *assertion) ToBeNil() {
	if a.actual != nil {
		a.ctx.Kill()
	}
}

func (a *assertion) ToBeNotNil() {
	if a.actual == nil {
		a.ctx.Kill()
	}
}

func (a *assertion) ToBeTrue() {
	if !reflect.DeepEqual(a.actual, true) {
		a.ctx.Kill()
	}
}

func (a *assertion) ToBeFalse() {
	if !reflect.DeepEqual(a.actual, false) {
		a.ctx.Kill()
	}
}
