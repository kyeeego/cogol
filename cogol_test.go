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
		g.T("This one works", func(c *cogol.Context) {
			_ = 2 + 2
		})

		g.TODO("this one is TODO")

		g.T("this one fails", func(c *cogol.Context) {
			c.Kill()
		})
	}

	g2 := cgl.Group("Second")
	{
		g2.T("SHOuld work fs", func(c *cogol.Context) {

		})

		g2.T("Should work as well", func(c *cogol.Context) {

		})
	}

	cgl.Process()
}
