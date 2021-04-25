package cogol

import (
	"fmt"
	"github.com/fatih/color"
	"testing"
)

const (
	tick  = "✓"
	cross = "✗"
	pencil = "✎"
)

func reportGroup(g *G) {
	text := color.HiWhiteString("Reporting group \"%v\":\n", g.name)
	fmt.Print(text)

	for _, test := range g.children {
		report(test, g.t)
	}

	for _, todo := range g.todo {
		reportTodo(todo)
	}
	fmt.Println()
}

func report(test *Test, t *testing.T) {
	var c string

	if test.success {
		c = color.HiGreenString("    %v PASS: %v\n", tick, test.name)
	} else {
		c = color.HiRedString("    %v FAIL: %v\n", cross, test.name)
		t.Fail()
	}

	fmt.Print(c)
}

func reportTodo(name string) {
	c := color.HiMagentaString("    %v TODO: %v\n", pencil, name)
	fmt.Print(c)
}
