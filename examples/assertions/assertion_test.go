package assertions

import (
	"errors"
	"github.com/kyeeego/cogol"
	"testing"
)

func Test(t *testing.T) {
	cgl := cogol.Init(t)

	g := cgl.Group("Assertions example")
	{
		g.T("Successfully adds two numbers", func(c *cogol.Context) {
			c.Expect(42 + 24).ToBe(66)
		})

		g.T("True is true, false is false", func(c *cogol.Context) {
			c.Expect(true).ToBeTrue()
			c.Expect(false).ToBeFalse()
		})

		g.T("Success because error is nil", func(c *cogol.Context) {
			var err error
			c.Expect(err).ToBeNil()
		})

		g.T("This one should fail, error is not nil!", func(c *cogol.Context) {
			err := errors.New("oopsie")
			c.Expect(err).ToBeNil()
		})

		g.T("In case we want error to be present, we can use ToBeNotNil method", func(c *cogol.Context) {
			err := errors.New("oopsie")
			c.Expect(err).ToBeNotNil()
		})
	}

	// Don't forget to run the "Process" method!
	cgl.Process()
}
