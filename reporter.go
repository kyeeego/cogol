package cogol

import (
	"fmt"
	"github.com/fatih/color"
	"testing"
)

const (
	tick   = "✓"
	cross  = "✗"
	pencil = "✎"
)

type Reporter interface {
	Group(g *G)
	Report(test *Test, t *testing.T)
	Todo(todo string)
}

type DefaultReporter struct{}

func (r DefaultReporter) Group(g *G) {
	clr := color.New(color.FgBlack)

	if g.success {
		printHeader(*clr, color.BgHiGreen, fmt.Sprintf("\n%v PASS: Group \"%v\"", tick, g.name))
	} else {
		printHeader(*clr, color.BgHiRed, fmt.Sprintf("\n%v FAIL: Group \"%v\"", cross, g.name))
	}

	for _, test := range g.children {
		r.Report(test, g.t)
	}

	for _, todo := range g.todo {
		r.Todo(todo)
	}
	fmt.Println()
}

func (DefaultReporter) Report(test *Test, t *testing.T) {
	var c string

	if test.success {
		c = color.HiGreenString("    %v PASS: %v\n", tick, test.name)
	} else {
		c = color.HiRedString("    %v FAIL: %v\n", cross, test.name)
		t.Fail()
	}

	fmt.Print(c)
}

func (DefaultReporter) Todo(name string) {
	c := color.HiMagentaString("    %v TODO: %v\n", pencil, name)
	fmt.Print(c)
}

func printHeader(clr color.Color, attr color.Attribute, text string) {
	clr.Add(attr)
	_,_ = clr.Print(text)
	fmt.Print("\n\n")
}
