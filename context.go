package cogol

import (
	"fmt"
	"testing"
)

type Context struct {
	Storage   storage
	logger    Logger
	test      *test
	t         *testing.T
	succeeded chan bool
	failed    chan string
	killed    bool
}

// context creates a new cgl.Context instance
func (cgl Cogol) context(test *test, logger Logger) *Context {
	return &Context{
		test: test,
		Storage: &defaultStorage{
			data: make(map[string]interface{}),
		},
		logger:    logger,
		succeeded: make(chan bool),
		failed:    make(chan string),
		t:         cgl.t,
		killed:    false,
	}
}

// Kill marks test as failed and stops test's processing
func (ctx *Context) Kill(message string) {
	// Can not kill context more than once, so if ctx.Kill() has been called
	// more than once, we can just ingnore it 'cause test's failed anyway
	if ctx.killed {
		return
	}

	f := &failure{
		ctx: ctx,
		msg: message,
	}

	ctx.test.f = f
	ctx.failed <- fmt.Sprintf("Killed '%v'", ctx.test.name)
	ctx.killed = true
}

func (ctx *Context) Log() Logger {
	return ctx.logger
}

type storage interface {
	Get(key string) interface{}
	Set(key string, value interface{})
}

type defaultStorage struct {
	data map[string]interface{}
}

func (s *defaultStorage) Get(key string) interface{} {
	return s.data[key]
}

func (s *defaultStorage) Set(key string, value interface{}) {
	s.data[key] = value
}
