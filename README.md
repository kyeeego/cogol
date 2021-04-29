```go
package example_test

import (
	"github.com/kyeeego/cogol"
	"testing"
)

func TestSomething(t *testing.T) {
	cgl := cogol.Init(t)
	
	g := cgl.Group("So these tests test something")
	{
		g.BeforeEach(func() {
			// Prepare stuff before each suite
		})

		g.T("2 + 2 should be 4", func(c *cogol.Context) {
			c.Expect(2 + 2).ToBe(4)
		})

		g.T("Do some dangerous operation", func(c *cogol.Context) {
			res, err := someDangerousOp()
			c.Expect(err).ToBeNil()
			c.Expect(res.Foo).ToBe("Bar")
		})

		g.TODO("Implement something later")

		g.AfterEach(func() {
			//  Clean stuff up after each suite
		})
	}

	cgl.Process()
}
```
