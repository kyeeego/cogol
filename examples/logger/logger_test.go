package storage

import (
	"github.com/kyeeego/cogol"
	"testing"
)

func TestLogger(t *testing.T) {
	cgl := cogol.Init(t)

	g := cgl.Group("Logger")
	{
		g.T("Example usage", func(c *cogol.Context) {
			// Logger is accessible from test's context by running c.Logger methods
			// Logger is an interface with 4 methods:
			// 		Info(text string),
			//      Infof(format string, args ...interface{}),
			//      Error(text string),
			//      Errorf(format string, args ...interface{}).
			c.Log().Info("Hello cogol!")
			c.Log().Infof("Hello cogol %d!", 2)
		})

		g.T("Another test", func(c *cogol.Context) {
			c.Log().Error("Error occured")
		})
	}

	// Don't forget to run the "Process" method!
	cgl.Process()
}

