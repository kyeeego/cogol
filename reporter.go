package cogol

import (
	"fmt"
	"github.com/fatih/color"
)

const (
	tick   = "✓"
	cross  = "✗"
	pencil = "✎"
)

type Reporter interface {
	Group(g *G)
	Test(test *Test)
	Todo(todo string)
	Error(f *failure)
	Print(format string, values ...interface{})
}

type defaultReporter struct{}

func (r defaultReporter) Group(g *G) {
	clr := color.New(color.FgBlack)

	printHeader(*clr, g.success, g.name)

	for _, test := range g.children {
		r.Test(test)
		if len(test.logs) != 0 {
			for _, l := range test.logs {
				r.Print(l)
			}
		}
	}

	for _, todo := range g.todo {
		r.Todo(todo)
	}
	fmt.Println()
}

func (r defaultReporter) Test(test *Test) {
	if test.success {
		c := color.HiGreenString("\t%v PASS: %v\n", tick, test.name)
		fmt.Print(c)
	} else {
		c := color.HiRedString("\t%v FAIL: %v\n", cross, test.name)
		fmt.Print(c)
		r.Error(test.f)
	}
}

func (defaultReporter) Error(f *failure) {
	c := color.HiBlackString("\t\t%v\n\n", f.msg)
	fmt.Print(c)
}

func (defaultReporter) Todo(name string) {
	c := color.HiMagentaString("\t%v TODO: %v\n", pencil, name)
	fmt.Print(c)
}

func (defaultReporter) Print(format string, values ...interface{}) {
	color.Yellow(format, values...)
}

func printHeader(clr color.Color, success bool, text string) {
	if success {
		clr.Add(color.BgHiGreen)
		_, _ = clr.Printf(" %v PASS ", tick)
	} else {
		clr.Add(color.BgHiRed)
		_, _ = clr.Printf(" %v FAIL ", cross)
	}
	fmt.Printf(" %v\n\n", text)
}
