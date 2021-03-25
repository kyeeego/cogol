package cogol

import (
	"fmt"
	"github.com/fatih/color"
)

const (
	tick  = "✓"
	cross = "✗"
	pencil = "✎"
)

func reportSuccess(ctx *Context) {
	c := color.HiGreenString("    %v PASS: %v\n", tick, ctx.test.name)
	fmt.Print(c)
}

func reportFail(ctx *Context) {
	c := color.HiRedString("    %v FAIL: %v\n", cross, ctx.test.name)
	fmt.Print(c)
}

func reportTodo(name string) {
	c := color.HiMagentaString("    %v TODO: %v\n", pencil, name)
	fmt.Print(c)
}
