package cogol

import (
	"fmt"
	"github.com/fatih/color"
	"testing"
)

func TestDefaultLogger(t *testing.T) {
	cgl := Init(t)

	g := cgl.Group("Logger")
	{
		g.BeforeEach(func(c *Context) {
			c.Expect(c.Log()).ToBeNotNil()
		})

		g.T("Info", func(c *Context) {
			c.Log().Info("hello world")

			c.Expect(c.test.logs).ToBe(
				[]string{
					fmt.Sprintf(
						"%v %v",
						color.HiBlueString("\t\tINFO"),
						"hello world\n\n",
					),
				},
			)
		})

		g.T("Infof", func(c *Context) {
			c.Log().Infof("%v world", "hello")

			c.Expect(c.test.logs).ToBe(
				[]string{
					fmt.Sprintf(
						"%v %v",
						color.HiBlueString("\t\tINFO"),
						"hello world\n\n",
					),
				},
			)
		})

		g.T("Error", func(c *Context) {
			c.Log().Error("hello world")

			c.Expect(c.test.logs).ToBe(
				[]string{
					fmt.Sprintf(
						"%v %v",
						color.HiRedString("\t\tERROR"),
						"hello world\n\n",
					),
				},
			)
		})

		g.T("Errorf", func(c *Context) {
			c.Log().Errorf("%v world", "hello")

			c.Expect(c.test.logs).ToBe(
				[]string{
					fmt.Sprintf(
						"%v %v",
						color.HiRedString("\t\tERROR"),
						"hello world\n\n",
					),
				},
			)
		})

		g.T("Multiple logger statements", func(c *Context) {
			c.Log().Info("Hello")
			c.Log().Error("Cogol")

			c.Expect(c.test.logs).ToBe(
				[]string{
					fmt.Sprintf(
						"%v %v",
						color.HiBlueString("\t\tINFO"),
						"Hello\n\n",
					),
					fmt.Sprintf(
						"%v %v",
						color.HiRedString("\t\tERROR"),
						"Cogol\n\n",
					),
				},
			)
		})

		g.AfterEach(func(c *Context) {
			c.test.logs = []string{}
		})
	}

	cgl.Process()
}

func TestCustomLogger(t *testing.T) {
	cgl := Init(t)
	cgl.UseLogger(&CustomLogger{})

	g := cgl.Group("Test custom logger")
	{
		g.T("Info", func(c *Context) {
			c.Log().Info("Hello")
			c.Expect(c.test.logs).ToBe(
				[]string{
					color.HiYellowString("Hello"),
				},
			)
		})

		g.T("Infof", func(c *Context) {
			c.Log().Infof("Hello %v", "world")
			c.Expect(c.test.logs).ToBe(
				[]string{
					color.HiYellowString("Hello world"),
				},
			)
		})

		g.T("Error", func(c *Context) {
			c.Log().Info("Hello")
			c.Expect(c.test.logs).ToBe(
				[]string{
					color.HiYellowString("Hello"),
				},
			)
		})

		g.T("Errorf", func(c *Context) {
			c.Log().Errorf("Hello %v", "world")
			c.Expect(c.test.logs).ToBe(
				[]string{
					color.HiYellowString("Hello world"),
				},
			)
		})

		g.T("Multiple logger statements", func(c *Context) {
			c.Log().Errorf("Hello %v", "world")
			c.Log().Info("Hello cogol")

			c.Expect(c.test.logs).ToBe(
				[]string{
					color.HiYellowString("Hello world"),
					color.HiYellowString("Hello cogol"),
				},
			)
		})

		g.AfterEach(func(c *Context) {
			c.test.logs = []string{}
		})
	}

	cgl.Process()
}