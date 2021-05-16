package cogol

import (
	"sync"
	"testing"
)

func Test(t *testing.T) {
	cgl := Init(t)

	mu := &sync.Mutex{}

	public := 0

	g := cgl.Group("Test runability")
	{
		g.TODO("todo 1")
		g.TODO("todo 2")

		g.T("Should be 2 todos in group", func(c *Context) {
			mu.Lock()
			public++
			mu.Unlock()

			c.Expect(len(g.todo)).ToBe(2)
		})

		g.T("Should be 3 tests in group", func(c *Context) {
			mu.Lock()
			public++
			mu.Unlock()

			c.Expect(len(g.children)).ToBe(3)
		})

		g.T("Running public 3rd time", func(c *Context) {
			mu.Lock()
			public++
			mu.Unlock()
		})
	}

	g2 := cgl.Group("Second")
	{
		g2.T("In cgl should be 2 groups", func(c *Context) {
			c.Expect(len(cgl.children)).ToBe(2)
		})

	}

	cgl.Process()

	// All tests should be ran
	if public != 3 {
		t.Fail()
	}
}

func Test2(t *testing.T) {
	cgl := Init(t)

	mu := &sync.Mutex{}

	public := 0

	g := cgl.Group("Before and after each testing")
	{
		g.BeforeEach(func(c *Context) {
			mu.Lock()
			public++
			mu.Unlock()
		})

		g.T("do stuff", func(c *Context) {})
		g.T("do stuff 2", func(c *Context) {})
		g.T("BeforeEach should be ran 3 times", func(c *Context) {})
	}

	g2 := cgl.Group("Testing persistence")
	{
		old := "old"

		// Old BeforeEach
		g2.BeforeEach(func(c *Context) {
			c.Storage.Set(old, 727)
		})

		// Overriding BeforeEach
		g2.BeforeEach(func(c *Context) {
			c.Storage.Set(old, 4444)
		})

		g2.T("Should be 4444", func(c *Context) {
			c.Expect(c.Storage.Get(old)).ToBe(4444)
		})
	}

	cgl.Process()
	if public != 3 {
		t.Fail()
	}
}

func TestContext_Kill(t *testing.T) {
	faket := &testing.T{}
	cgl := Init(faket)

	g := cgl.Group("Context kill testing")
	{
		g.T("Context.Kill should fail the test immediately", func(c *Context) {
			c.Logger.Infof("%v killed", c.test.name)
			c.Kill(&failure{
				ctx: c,
				msg: "If it failed, then the test is passing",
			})
		})

		g.T("ctx.Kill() calls after the first one should be ignored", func(c *Context) {
			c.Kill(&failure{
				ctx: c,
				msg: "If it failed, then the test is passing",
			})

			c.Kill(&failure{
				ctx: c,
				msg: "Should not appear",
			})
		})

		g.AfterEach(func(c *Context) {
			c.Kill(&failure{
				ctx: c,
				msg: "Fuck",
			})
		})
	}

	cgl.Process()

	if !faket.Failed() {
		t.Fail()
	}
}
