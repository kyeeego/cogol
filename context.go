package cogol

import (
	"fmt"
	"testing"
)

type Context struct {
	Storage   storage
	test      *test
	t         *testing.T
	succeeded chan bool
	failed    chan string
}

func (cgl Cogol) context(test *test) *Context {
	return &Context{
		test: test,
		Storage: &defaultStorage{
			data: make(map[string]interface{}),
		},
		succeeded: make(chan bool),
		failed:    make(chan string),
		t:         cgl.t,
	}
}

func (ctx *Context) Kill(f *failure) {
	ctx.test.f = f
	ctx.failed <- fmt.Sprintf("Killed '%v'", ctx.test.name)
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
