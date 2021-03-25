package cogol_test

import (
	"testing"

	"github.com/kyeeego/cogol"
)

// real example of usage
func TestG_Process(t *testing.T) {
	g := cogol.Group("Main")
	{
		g.T("This one works", func(c *cogol.Context) {
			_ = 2 + 2
		})

		g.TODO("this one is TODO")

		g.T("this one fails", func(c *cogol.Context) {
			c.Kill()
		})
	}

	g.Process()
}
