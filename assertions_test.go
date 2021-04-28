package cogol_test

import (
	"errors"
	"github.com/kyeeego/cogol"
	"testing"
)

func TestAssertion(t *testing.T) {
	cgl := cogol.Init(t)

	g2 := cgl.Group("Assertions")
	{
		g2.T("SHOuld fail because err != nil", func(c *cogol.Context) {
			err := errors.New("cum")
			c.Expect(err).ToBeNil()
		})

		g2.T("SHOuld pass", func(c *cogol.Context) {
			var err error
			c.Expect(err).ToBeNil()
		})

		g2.T("Should work", func(c *cogol.Context) {
			c.Expect(2 + 2).ToBe(4)
		})

		g2.T("Should fail with incorrect value", func(c *cogol.Context) {
			c.Expect(2 + 2).ToBe(5)
		})

		g2.T("Should fail with incorrect types", func(c *cogol.Context) {
			c.Expect("string").ToBe(42)
		})
	}

	cgl.Process()
}
