package cogol

import (
	"fmt"
	"github.com/fatih/color"
)

type Logger interface {
	forTest(test *test)
	currentTest() *test

	Info(text string)
	Infof(format string, args ...interface{})
	Error(text string)
	Errorf(format string, args ...interface{})
}

type defaultLogger struct {
	test *test
}

func (l defaultLogger) Info(text string) {
	tag := color.HiBlueString("\t\tINFO")
	l.test.logs = append(l.test.logs, fmt.Sprintf("%v %v\n\n", tag, text))
}

func (l defaultLogger) Infof(format string, a ...interface{}) {
	tag := color.HiBlueString("\t\tINFO")
	text := fmt.Sprintf(format, a...)
	l.test.logs = append(l.test.logs, fmt.Sprintf("%v %v\n\n", tag, text))
}

func (l defaultLogger) Error(text string) {
	tag := color.HiRedString("\t\tERROR")
	l.test.logs = append(l.test.logs, fmt.Sprintf("%v %v\n\n", tag, text))
}

func (l defaultLogger) Errorf(format string, a ...interface{}) {
	tag := color.HiRedString("\t\tERROR")
	text := fmt.Sprintf(format, a...)
	l.test.logs = append(l.test.logs, fmt.Sprintf("%v %v\n\n", tag, text))
}

func (l *defaultLogger) currentTest() *test {
	return l.test
}

func (l *defaultLogger) forTest(test *test) {
	l.test = test
}
