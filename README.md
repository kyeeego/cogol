```go
package example_test

import (
	"github.com/kyeeego/cogol"
	"testing"
)

func TestSomething(t *testing.T) {
	g := cogol.Group("So these tests test something")
	{
		g.BeforeEach(func() {
			// Prepare stuff before each suite
		})

		g.T("2 + 2 should be 4", func() {
			cogol.Expect(add(2, 2)).ToBe(4)
		})

		g.T("Do some dangerous operation", func() {
			res, err := someDangerousOp()
			cogol.Expect(err).ToBeNil()
			cogol.Expect(res.Foo).ToBe("Bar")
		})

		g.TODO("Implement something later")

		g.AfterEach(func() {
			//  Clean stuff up after each suite
		})
	}

	g.Process()
}
```
