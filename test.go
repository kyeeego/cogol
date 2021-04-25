package cogol

type Handler = func(c *Context)

type Test struct {
	name    string
	handler Handler
	success bool
}
