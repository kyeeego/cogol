package cogol

import (
	"fmt"
	"github.com/fatih/color"
	"testing"
)

func TestLogger(t *testing.T) {
	cgl := Init(t)

	g := cgl.Group("Logger")
	{
		g.T("Infof", func(c *Context) {
			c.Logger.Infof("%v world", "hello")

			c.Expect(c.test.logs).ToBe(
				fmt.Sprintf(
					"%v %v",
					color.HiBlueString("\t\tINFO"),
					"hello world\n\n",
				),
			)
		})

		g.AfterEach(func(c *Context) {
			c.test.logs = []string{}
		})
	}

	cgl.Process()
}