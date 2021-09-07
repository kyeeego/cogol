package cogol

import (
	"os"
	"sync"
	"testing"
)

type Cogol struct {
	t        *testing.T
	children []*G
	reporter Reporter
	logger   Logger
	storage  Storage
}

// Init creates a new Cogol instance from *testing.T taken from Go's typical test function
func Init(t *testing.T) *Cogol {
	return &Cogol{
		t:        t,
		children: []*G{},
		reporter: &defaultReporter{},
		logger:   &defaultLogger{},
		storage:  newDefaultStorage()}
}

// G is a struct that represents a group of tests
type G struct {
	name       string
	children   []*test
	todo       []string
	beforeEach handler
	beforeAll  func()
	afterEach  handler
	afterAll   func()
	t          *testing.T
	success    bool
}

type handler = func(c *Context)

type test struct {
	name    string
	handler handler
	success bool
	f       *failure
	logs    []string
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

// Process runs all the groups of a Cogol instance, calculating it's success and then reporting
func (cgl *Cogol) Process() {
	var wg sync.WaitGroup
	mu := &sync.Mutex{}

	old := os.Stdout
	_, os.Stdout, _ = os.Pipe()

	for _, g := range cgl.children {
		wg.Add(1)
		go func(wg *sync.WaitGroup, g *G) {
			defer wg.Done()

			cgl.processGroup(g)
			g.calculateSuccess()

		}(&wg, g)
	}
	wg.Wait()

	os.Stdout = old

	for _, g := range cgl.children {
		// Locking IO so group reports in a single cogol.Cogol instance won't
		// be reported at the same time resulting in a complete mess
		mu.Lock()
		cgl.reporter.Group(g)
		mu.Unlock()
	}
}

// calculateSuccess is a simple algorithm to check whether all tests in a group have passed
func (g *G) calculateSuccess() {
	for _, test := range g.children {
		if !test.success {
			g.success = false
			return
		}
	}
	g.success = true
}

// T creates a typical testcase
func (g *G) T(name string, handler handler) {
	t := &test{
		name: name,
		handler: func(c *Context) {
			handler(c)
			c.succeeded <- true
		},
		logs: []string{},
	}

	g.children = append(g.children, t)
}

// TODO adds name of the test to group's to G.todo array, marking test as TODO
// Does not require a handler
func (g *G) TODO(name string) {
	g.todo = append(g.todo, name)
}

// BeforeEach sets handler that has to be executed before each test
func (g *G) BeforeEach(h handler) {
	g.beforeEach = h
}

// AfterEach sets handler that has to be executed after each test
func (g *G) AfterEach(h handler) {
	g.afterEach = h
}

// processGroup runs all the tests in group
func (cgl *Cogol) processGroup(g *G) {
	var wg sync.WaitGroup

	for _, testCase := range g.children {
		wg.Add(1)

		c := cgl.context(testCase, newDefaultLogger(testCase), newDefaultStorage())

		go func(test *test, wg *sync.WaitGroup) {
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
