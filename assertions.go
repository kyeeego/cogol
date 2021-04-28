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

func (a *assertion) ToBeNil() {
	if a.actual != nil {
		a.ctx.Kill()
	}
}
