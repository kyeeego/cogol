package cogol

import (
	"sync"
	"testing"
)

type Cogol struct {
	t        *testing.T
	children []*G
	reporter Reporter
}

func Init(t *testing.T) *Cogol {
	return &Cogol{t, []*G{}, DefaultReporter{}}
}

// G is a struct that represents a group of tests
type G struct {
	name       string
	children   []*Test
	todo       []string
	beforeEach func()
	beforeAll  func()
	afterEach  func()
	afterAll   func()
	t          *testing.T
}

// Group is a function that creates a new group (G instance)
func (cgl *Cogol) Group(name string) *G {
	g := &G{
		name: name,
		t:    cgl.t,
	}
	cgl.children = append(cgl.children, g)

	return g
}

func (cgl *Cogol) Process() {
	for _, g := range cgl.children {
		cgl.processGroup(g)
		cgl.reporter.Group(g)
	}
}

// T indicates a typical testcase
func (g *G) T(name string, handler Handler) {
	t := &Test{
		name: name,
		handler: func(c *Context) {
			handler(c)
			c.succeeded <- true
		},
	}

	g.children = append(g.children, t)
}

func (g *G) TODO(name string) {
	g.todo = append(g.todo, name)
}

// process runs all the tests in group
func (cgl *Cogol) processGroup(g *G) {
	var wg sync.WaitGroup

	for _, testCase := range g.children {
		wg.Add(1)
		c := cgl.Context(testCase)

		go func(test *Test, wg *sync.WaitGroup) {
			defer wg.Done()

			go test.handler(c)

			select {
			case <-c.succeeded:
				test.success = true

			case <-c.failed:
				test.success = false
			}

		}(testCase, &wg)
	}
	wg.Wait()
}
