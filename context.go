package cogol

import "fmt"

const (
	KILLED = true
)

type Context struct {
	test      *Test
	succeeded chan bool
	failed    chan string
}

func newContext(test *Test) *Context {
	return &Context{
		test:      test,
		succeeded: make(chan bool),
		failed:    make(chan string),
	}
}

func (ctx *Context) Kill() {
	ctx.failed <- fmt.Sprintf("Killed '%v'", ctx.test.name)
	//ctx.success = false
}
