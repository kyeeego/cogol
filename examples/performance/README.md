# Achieving the best performance

In this example you'll learn about best practices to make your tests run at the speed of light!

## Put all the groups inside one function

Let's say you have a bunch of slow tests of different functions that need to be run in different groups. The first
advice is: do not divide your tests by ```Go```'s testing functions, like in ```Demo 1```. ```Go```'s testing functions
are being run one after another in the same goroutine, so after running code in ```Demo 1``` we'll see, that the
execution time was around ```6.001s``` *(Because we've been sleeping for 2 and 4 seconds consecutively)*

Instead, you should create all the groups in a single function, under a single ```cogol.Cogol``` instance
*(what's shown in ```Demo 2```)*. All the ```cogol``` groups that are binded to one ```cogol.Cogol```
instance will be run in parallel, highly improving speed of the tests we've written in ```Demo 1```. Execution time
of ```Demo 2``` code will be around ```4.001s``` *(The longest operation in tests takes exactly 4 seconds)*

#### Demo 1

```go
package main

import (
	"testing"
	"time"
	
	"github.com/kyeeego/cogol"
)

func TestStorage(t *testing.T) {
	cgl := cogol.Init(t)

	g := cgl.Group("Example")
	{
		g.T("Example", func(c *cogol.Context) {
			// Sleeping for 2 seconds, imitating some slow operation
			time.Sleep(time.Second * 2)
		})
	}
	
	cgl.Process()
}

func test(t *testing.T) {
	cgl := cogol.Init(t)
	
    g := cgl.Group("Example 2")
	{
		g.T("Example", func(c *cogol.Context) {
			// Sleeping for 4 seconds, imitating even slower operation
			time.Sleep(time.Second * 4)
		})
	}
	
	cgl.Process()
}
```

#### Demo 2

```go
package main

import (
	"testing"
	"time"
	
	"github.com/kyeeego/cogol"
)

func TestStorage(t *testing.T) {
	cgl := cogol.Init(t)

	g := cgl.Group("Example")
	{
		g.T("Example", func(c *cogol.Context) {
			// Sleeping for 2 seconds, imitating some slow operation
			time.Sleep(time.Second * 2)
		})
	}

	g2 := cgl.Group("Example 2")
	{
		g2.T("Example", func(c *cogol.Context) {
			// Sleeping for 4 seconds, imitating even slower operation
			time.Sleep(time.Second * 4)
		})
	}
	
	cgl.Process()
}
```