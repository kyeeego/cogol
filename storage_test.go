package cogol

import (
	"testing"
)

func TestStorage(t *testing.T) {
	cgl := Init(t)

	g := cgl.Group("Storage")
	{
		propname := "test1property"
		bepropname := "bepropname"

		// Overriding before each
		g.BeforeEach(func(c *Context) {
			c.Storage.Set(bepropname, 727)
		})

		g.T("Should not be accessible from other tests", func(c *Context) {
			c.Storage.Set(propname, 42)
		})

		g.T("Test 2", func(c *Context) {
			c.Expect(c.Storage.Get(propname)).ToBeNil()
		})

		g.T("should be accessible because of beforeeach", func(c *Context) {
			c.Expect(c.Storage.Get(bepropname)).ToBe(727)
			c.Logger.Info("should be accessible because of beforeeach")
		})

		g.T("here as well", func(c *Context) {
			c.Expect(c.Storage.Get(bepropname)).ToBe(727)
			c.Logger.Info("here as well")
		})

		g.T("propname should be overridden", func(c *Context) {
			c.Storage.Set(propname, 4444)

			c.Expect(c.Storage.Get(propname)).ToBe(4444)
		})
	}

	cgl.Process()
}
