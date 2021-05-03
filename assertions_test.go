package cogol

import (
	"errors"
	"testing"
)

func mockKiller(c *Context) {
	c.test.success = false
}

func TestAssertion(t *testing.T) {
	cgl := Init(t)

	g2 := cgl.Group("Assertions")
	{
		g2.T("SHOuld fail because err != nil", func(c *Context) {
			err := errors.New("cum")
			a := assertion{nil, err, c, mockKiller}
			a.ToBeNil()

			verify(c, false)
		})

		g2.T("SHOuld pass", func(c *Context) {
			var err error
			a := assertion{nil, err, c, mockKiller}
			a.ToBeNil()

			verify(c, true)
		})

		g2.T("Should work", func(c *Context) {
			a := assertion{nil, 2 + 2, c, mockKiller}
			a.ToBe(4)

			verify(c, true)
		})

		g2.T("Should fail with incorrect value", func(c *Context) {
			a := assertion{nil, 2 + 2, c, mockKiller}
			a.ToBe(5)

			verify(c, false)
		})

		g2.T("Should fail with incorrect types", func(c *Context) {
			a := assertion{nil, "string", c, mockKiller}
			a.ToBe(42)

			verify(c, false)
		})

		g2.T("SHould be good", func(c *Context) {
			a := assertion{nil, true, c, mockKiller}
			a2 := assertion{nil, false, c, mockKiller}
			a.ToBeTrue()
			a2.ToBeFalse()

			verify(c, true)
		})

		g2.T("should die", func(c *Context) {
			a := assertion{nil, false, c, mockKiller}
			a.ToBeTrue()

			verify(c, false)
		})

		g2.T("Zero testing", func(c *Context) {
			a := assertion{nil, "", c, mockKiller}
			a2 := assertion{nil, "full", c, mockKiller}
			a.ToBeZero()
			a2.ToBeNotZero()

			verify(c, true)
		})
	}

	cgl.Process()
}

func TestAssertion_Demo(t *testing.T) {
	cgl := Init(t)

	g := cgl.Group("Assertions demo")
	{
		g.T("Works", func(c *Context) {
			c.Expect(2 + 2).ToBe(4)
			c.Expect(true).ToBeTrue()
			c.Expect("").ToBeZero()
		})
	}

	cgl.Process()
}

func verify(ctx *Context, shouldBeSuccessful bool) {
	if ctx.test.success != shouldBeSuccessful {
		ctx.Kill()
	} else {
		ctx.succeeded <- true
	}
}
