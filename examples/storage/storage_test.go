package storage

import (
	"github.com/kyeeego/cogol"
	"testing"
)

func TestStorage(t *testing.T) {
	cgl := cogol.Init(t)

	g := cgl.Group("Storage")
	{
		g.BeforeEach(func(c *cogol.Context) {
			// If we need to have a certain value in each and every of our tests,
			// we can run the storage Set method in BeforeEach, so that our value
			// will be put in each test's storage
			c.Storage().Set("public", 42)
		})

		g.T("Basic usage", func(c *cogol.Context) {
			// Context.Storage interface has two methods: Get and Set.
			// Default implementation consists of a simple map[string]interface{}
			//
			// With Set method, we can put some values into storage
			c.Storage().Set("hello", "world")

			c.Expect(c.Storage().Get("hello")).ToBe("world")
		})

		g.T("Can we access values from other tests?", func(c *cogol.Context) {
			// The answer is no. One storage is binded to one context, so if we try to get
			// the value of "hello" that we've set in previous test, we'll get nil
			//
			// The following assertion should fail:
			c.Expect(c.Storage().Get("hello")).ToBe("world")
		})

		g.T("Testing public value", func(c *cogol.Context) {
			// We've set "public" in BeforeEach method, so the value of "public" will be
			// accessible from each and every test
			c.Expect(c.Storage().Get("public")).ToBe(42)
		})

		g.T("Testing public value for the second time", func(c *cogol.Context) {
			c.Expect(c.Storage().Get("public")).ToBe(42)
		})
	}

	// Don't forget to run the "Process" method!
	cgl.Process()
}
