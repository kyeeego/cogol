package cogol_test

import (
	"testing"

	"github.com/kyeeego/cogol"
)

// real example of usage
func TestG_Process(t *testing.T) {
	cgl := cogol.Init(t)

	g := cgl.Group("Main")
	{

		g.BeforeEach(func(c *cogol.Context) {
			c.Storage.Set("hello", "world")
		})

		g.T("This one works", func(c *cogol.Context) {
			c.Expect(c.Storage.Get("hello")).ToBe("world")
		})

		g.TODO("this one is TODO")

		g.T("this one fails", func(c *cogol.Context) {
			c.Expect(c.Storage.Get("hello")).ToBe("hello")
		})

		g.T("this one is nil", func(c *cogol.Context) {
			c.Expect(c.Storage.Get("invalid")).ToBeNil()
		})
	}

	cgl.Process()
}
