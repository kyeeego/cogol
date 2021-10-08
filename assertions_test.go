package cogol

import (
	"errors"
	"fmt"
	"testing"
)

func mockKiller(f *failure) {
	f.ctx.test.success = false
}

type example struct {
	inner int
}

func TestAssertion(t *testing.T) {
	cgl := Init(t)

	g := cgl.Group("Assertions")
	{

		g.T("Testing on errors", func(c *Context) {
			var err error
			a := assertion{nil, err, c, mockKiller}
			a.ToBeNil()

			verify(c, true)
		})

		g.T("Testing on errors (incorrect)", func(c *Context) {
			err := errors.New("err")
			a := assertion{nil, err, c, mockKiller}
			a.ToBeNil()

			verify(c, false)
		})

		g.T("Testing ToBe on integers", func(c *Context) {
			a := assertion{nil, 2 + 2, c, mockKiller}
			a.ToBe(4)

			verify(c, true)
		})

		g.T("Testing ToBe on integers (incorrect)", func(c *Context) {
			a := assertion{nil, 2 + 2, c, mockKiller}
			a.ToBe(5)

			verify(c, false)
		})

		g.T("Testing ToBe on different types", func(c *Context) {
			a := assertion{nil, "string", c, mockKiller}
			a.ToBe(42)

			verify(c, false)
		})

		g.T("Testing ToBeTrue and ToBeFalse on booleans", func(c *Context) {
			a := assertion{nil, true, c, mockKiller}
			a2 := assertion{nil, false, c, mockKiller}
			a.ToBeTrue()
			a2.ToBeFalse()

			verify(c, true)
		})

		g.T("Testing ToBeTrue on non-booleans", func(c *Context) {
			a := assertion{nil, "hi", c, mockKiller}
			a.ToBeTrue()

			verify(c, false)
		})

		g.T("Testing ToBeFalse on non-booleans", func(c *Context) {
			a := assertion{nil, "hi", c, mockKiller}
			a.ToBeFalse()

			verify(c, false)
		})

		g.T("Should die", func(c *Context) {
			a := assertion{nil, false, c, mockKiller}
			a.ToBeTrue()

			verify(c, false)
		})

		g.T("Testing on equal structs", func(c *Context) {
			e := example{inner: 2}
			e2 := example{inner: 2}

			a := assertion{nil, e2, c, mockKiller}
			a.ToBe(e)

			verify(c, true)
		})

		g.T("Testing on non-equal structs", func(c *Context) {
			e := example{inner: 2}
			e2 := example{inner: 3}

			a := assertion{nil, e2, c, mockKiller}
			a.ToBe(e)

			verify(c, false)
		})

		g.T("Zero testing on strings", func(c *Context) {
			a := assertion{nil, "", c, mockKiller}
			a2 := assertion{nil, "full", c, mockKiller}
			a.ToBeZero()
			a2.ToBeNotZero()

			verify(c, true)
		})

		g.T("Zero testing on structs", func(c *Context) {
			a := assertion{nil, example{}, c, mockKiller}
			a2 := assertion{nil, example{2}, c, mockKiller}
			a.ToBeZero()
			a2.ToBeNotZero()

			verify(c, true)
		})
	}

	cgl.Process()
}

func verify(ctx *Context, shouldBeSuccessful bool) {
	if ctx.test.success != shouldBeSuccessful {
		ctx.Kill(fmt.Sprintf("test successfullness was supposed to be %v", shouldBeSuccessful))
	}
}
