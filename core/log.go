package core

import "github.com/kataras/golog"

var rootLogger *golog.Logger

func GetLogger(name string) *golog.Logger {
	child := rootLogger.Child(name)
	if !child.NewLine {
		child.NewLine = true
	}
	return child
}

func SetRootLogger(logger *golog.Logger) {
	rootLogger = logger
}
