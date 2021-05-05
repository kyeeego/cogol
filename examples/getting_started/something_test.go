package getting_started

import (
	"github.com/kyeeego/cogol"
	"testing"
)

func Test(t *testing.T) {
	// First, we need to initialize cogol by running the Init function,
	// which gives us an instance of the framework to use it later
	cgl := cogol.Init(t)

	// Next, we can create groups of tests by running Group method
	// of the framework instance that we've created earlier.
	// Method takes one parameter: the name of the group
	//
	// I highly recommend to create an additional block of code
	// and put all your test there because it strongly improves readability
	// of your tests
	g := cgl.Group("Getting started")
	{
		// This is a test example. We can create it by calling the T method of
		// the group, that we want our test to be in.
		//
		// Method takes two parameters: first is our test's name and second one is the handler,
		// where we can write our test's logic.
		// Notice the Context parameter in our handler. It holds all the important data about the test,
		// and allows you to use assertions. You can learn more about assertions in the corresponding example
		g.T("Hello cogol!", func(c *cogol.Context) {
			c.Expect(c.Storage.Get("hello")).ToBe("world")
		})

		// If you have an idea for a test, but don't feel like writing it, or just cannot write it,
		// use the TODO method of the group. You will not forget about it, because every time
		// you run the tests, you will be reminded of all the tests marked as TODO
		g.TODO("I will write this one eventually...")

		// With BeforeEach method of the group, you can run some code before each test case.
		// For example, you can configure the Context for the test.
		// One of the ways you can set up your test is by using Context's Storage.
		// You can learn about it in the corresponding example
		//
		// Now you can see, why the "Hello cogol!" test is actually passing
		g.BeforeEach(func(c *cogol.Context) {
			c.Storage.Set("hello", "world")
		})

		// With AfterEach method of the group, you can run some code after each test,
		// for example logging
		g.AfterEach(func(c *cogol.Context) {
			// Your logs here
		})
	}

	// In the end, we have to run the Process method of the framework instance.
	// This methods runs all the tests in groups, that were stuck to it
	cgl.Process()
}
