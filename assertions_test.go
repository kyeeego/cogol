package cogol

import (
	"errors"
	"fmt"
	"testing"
)

func mockKiller(f *failure) {
	f.ctx.test.success = false
}

func TestAssertion(t *testing.T) {
	cgl := Init(t)

	g := cgl.Group("Assertions")
	{
		g.T("SHOuld fail because err != nil", func(c *Context) {
			err := errors.New("cum")
			a := assertion{nil, err, c, mockKiller}
			a.ToBeNil()

			verify(c, false)
		})

		g.T("SHOuld pass", func(c *Context) {
			var err error
			a := assertion{nil, err, c, mockKiller}
			a.ToBeNil()

			verify(c, true)
		})

		g.T("Should work", func(c *Context) {
			a := assertion{nil, 2 + 2, c, mockKiller}
			a.ToBe(4)

			verify(c, true)
		})

		g.T("Should fail with incorrect value", func(c *Context) {
			a := assertion{nil, 2 + 2, c, mockKiller}
			a.ToBe(5)

			verify(c, false)
		})

		g.T("Should fail with incorrect types", func(c *Context) {
			a := assertion{nil, "string", c, mockKiller}
			a.ToBe(42)

			verify(c, false)
		})

		g.T("SHould be good", func(c *Context) {
			a := assertion{nil, true, c, mockKiller}
			a2 := assertion{nil, false, c, mockKiller}
			a.ToBeTrue()
			a2.ToBeFalse()

			verify(c, true)
		})

		g.T("should die", func(c *Context) {
			a := assertion{nil, false, c, mockKiller}
			a.ToBeTrue()

			verify(c, false)
		})

		g.T("Zero testing", func(c *Context) {
			a := assertion{nil, "", c, mockKiller}
			a2 := assertion{nil, "full", c, mockKiller}
			a.ToBeZero()
			a2.ToBeNotZero()

			verify(c, true)
		})
	}

	cgl.Process()
}

func verify(ctx *Context, shouldBeSuccessful bool) {
	if ctx.test.success != shouldBeSuccessful {
		ctx.Kill(fmt.Sprintf("Test successfullness was supposed to be %v", shouldBeSuccessful))
	}
}
