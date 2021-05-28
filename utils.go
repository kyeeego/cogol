package cogol

import (
	"github.com/fatih/color"
)

// CustomLogger - Example custom logger
type CustomLogger struct {
	test *Test
}

func (c *CustomLogger) ForTest(test *Test) {
	c.test = test
}

func (c *CustomLogger) CurrentTest() *Test {
	return c.test
}

func (c CustomLogger) Info(text string) {
	c.CurrentTest().logs = append(c.CurrentTest().logs, color.HiYellowString(text))
}

func (c CustomLogger) Infof(format string, args ...interface{}) {
	c.CurrentTest().logs = append(c.CurrentTest().logs, color.HiYellowString(format, args...))
}

func (c CustomLogger) Error(text string) {
	c.CurrentTest().logs = append(c.CurrentTest().logs, color.HiYellowString(text))
}

func (c CustomLogger) Errorf(format string, args ...interface{}) {
	c.CurrentTest().logs = append(c.CurrentTest().logs, color.HiYellowString(format, args...))
}

