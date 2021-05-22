package cogol

import "github.com/fatih/color"

type CustomLogger struct {
	test *test
}

func (c *CustomLogger) forTest(test *test) {
	c.test = test
}

func (c *CustomLogger) currentTest() *test {
	return c.test
}

func (c CustomLogger) Info(text string) {
	c.currentTest().logs = append(c.currentTest().logs, color.HiYellowString(text))
}

func (c CustomLogger) Infof(format string, args ...interface{}) {
	c.currentTest().logs = append(c.currentTest().logs, color.HiYellowString(format, args...))
}

func (c CustomLogger) Error(text string) {
	c.currentTest().logs = append(c.currentTest().logs, color.HiYellowString(text))
}

func (c CustomLogger) Errorf(format string, args ...interface{}) {
	c.currentTest().logs = append(c.currentTest().logs, color.HiYellowString(format, args...))
}

