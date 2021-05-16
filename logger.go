package cogol

import "github.com/fatih/color"

type logger interface {
	Info(string)
	Infof(string, ...interface{})
}

type defaultLogger struct {
	test *test
}

func newDefaultLogger(test *test) *defaultLogger {
	return &defaultLogger{
		test: test,
	}
}

func (l defaultLogger) Info(text string) {
	l.test.logs = color.YellowString(text)
}

func (l defaultLogger) Infof(format string, a ...interface{}) {
	l.test.logs = color.YellowString(format, a...)
}
