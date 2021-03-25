package cogol

import (
	"sync"
)

// group is a struct that represents a group of tests
type group struct {
	name       string
	children   []*Test
	todo       []string
	beforeEach Handler
	beforeAll  Handler
	afterEach  Handler
	afterAll   Handler
}

// Group is a function that creates a new group (group instance)
func Group(name string) *group {
	return &group{
		name: name,
	}
}

// T indicates a typical testcase
func (g *group) T(name string, handler Handler) {
	t := &Test{
		name:    name,
		handler: func(c *Context) {
			handler(c)
			c.succeeded <- true
		},
	}

	g.children = append(g.children, t)
}

func (g *group) TODO(name string) {
	g.todo = append(g.todo, name)
}

// Process runs all the tests in group
func (g *group) Process() {
	var wg sync.WaitGroup

	for _, testCase := range g.children {
		wg.Add(1)
		c := New(testCase)

		go func(test *Test, wg *sync.WaitGroup) {
			defer wg.Done()

			go test.handler(c)

			select {
			case <-c.succeeded:
				reportSuccess(c)
				return
			case <-c.failed:
				reportFail(c)
				return
			}

		}(testCase, &wg)
	}
	wg.Wait()

	for _, todo := range g.todo {
		reportTodo(todo)
	}
}
