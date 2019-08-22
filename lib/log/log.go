package log

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

var (
	Debug bool
)

func INFO(message string, args ...interface{}) {
}

func WARN(message string, args ...interface{}) {
}

func FATAL(message string, args ...interface{}) {
}

func DEBUG(message string, args ...interface{}) {
	fmt.Println(message, args)
}

func TRACE(message string, args ...interface{}) {
}

func DUMP(message string, args ...interface{}) {
	if Debug {
		fmt.Print(message + ":\n")
		spew.Dump(args...)
		fmt.Print("\n")
	}
}
