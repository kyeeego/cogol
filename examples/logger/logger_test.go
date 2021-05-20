package storage

import (
	"github.com/kyeeego/cogol"
	"testing"
)

func TestStorage(t *testing.T) {
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
			c.Logger.Info("Hello cogol!")
			c.Logger.Infof("Hello cogol %d!", 2)
		})

		g.T("Another test", func(c *cogol.Context) {
			c.Logger.Error("Error occured")
		})
	}

	// Don't forget to run the "Process" method!
	cgl.Process()
}
