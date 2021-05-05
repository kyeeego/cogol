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
	beforeEach Handler
	beforeAll  func()
	afterEach  Handler
	afterAll   func()
	t          *testing.T
	success    bool
}

type Handler = func(c *Context)

type Test struct {
	name    string
	handler Handler
	success bool
	f       *failure
}

// Group is a function that creates a new group (G instance)
func (cgl *Cogol) Group(name string) *G {
	g := &G{
		name: name,
		t:    cgl.t,

		beforeEach: func(c *Context) {},
		afterEach:  func(c *Context) {},
		beforeAll:  func() {},
		afterAll:   func() {},
	}
	cgl.children = append(cgl.children, g)

	return g
}

func (cgl *Cogol) Process() {
	var wg sync.WaitGroup

	for _, g := range cgl.children {
		wg.Add(1)
		go func(wg *sync.WaitGroup, g *G) {
			defer wg.Done()

			cgl.processGroup(g)
			g.calculateSuccess()
			cgl.reporter.Group(g)
		}(&wg, g)

	}
	wg.Wait()
}

func (g *G) calculateSuccess() {
	for _, test := range g.children {
		if !test.success {
			g.success = false
			return
		}
	}
	g.success = true
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

func (g *G) BeforeEach(h Handler) {
	g.beforeEach = h
}

func (g *G) AfterEach(h Handler) {
	g.afterEach = h
}

// process runs all the tests in group
func (cgl *Cogol) processGroup(g *G) {
	var wg sync.WaitGroup

	for _, testCase := range g.children {
		wg.Add(1)
		c := cgl.context(testCase)

		go func(test *Test, wg *sync.WaitGroup) {
			defer wg.Done()

			g.beforeEach(c)
			go test.handler(c)

			select {
			case <-c.succeeded:
				test.success = true

			case <-c.failed:
				test.success = false
				cgl.t.Fail()
			}

			g.afterEach(c)

		}(testCase, &wg)
	}
	wg.Wait()
}
