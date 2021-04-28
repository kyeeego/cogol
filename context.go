package cogol

import (
	"fmt"
	"testing"
)

const (
	KILLED = true
)

type Context struct {
	test      *Test
	t         *testing.T
	succeeded chan bool
	failed    chan string
}

func (cgl Cogol) Context(test *Test) *Context {
	return &Context{
		test:      test,
		succeeded: make(chan bool),
		failed:    make(chan string),
		t:         cgl.t,
	}
}

func (ctx *Context) Kill() {
	ctx.failed <- fmt.Sprintf("Killed '%v'", ctx.test.name)
	//ctx.success = false
}
