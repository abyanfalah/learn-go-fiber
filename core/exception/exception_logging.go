package exception

import (
	"bytes"
	"errors"
	"fmt"
	"learn-fiber/core/config"
	"log"
	"runtime"
	"runtime/debug"
)

var callerPositionSkip = 3

type StackError struct {
	Err   error
	Stack []byte
}

func logError(e error) {
	var se *StackError
	if !errors.As(e, &se) {
		e = newStackError(e)
		_ = errors.As(e, &se)
	}

	msg := e
	caller, _ := getCallerAndLocation(callerPositionSkip)
	fmt.Println()
	log.Printf(`an error occured: "%s"`, msg)
	// fmt.Printf("%-14s: %s\n", "Caller", caller)
	fmt.Printf("%-14s: %s\n", "Location", caller)
	fmt.Printf("%-14s: %s\n", "Message", msg)
	fmt.Printf("%-14s: %s\n", "Stacktrace", "↴")
	printRelevantStack(se.Stack)
	fmt.Println()
}

func (e *StackError) Error() string {
	return e.Err.Error()
}

func newStackError(err error) error {
	return &StackError{
		Err:   err,
		Stack: debug.Stack(),
	}
}

func printRelevantStack(stack []byte) {
	lines := bytes.Split(stack, []byte("\n"))
	for i := range len(lines) - 1 {
		line := lines[i]
		next := lines[i+1]

		if isRelevantStacktrace(next) {
			fmt.Printf("%s\n%s\n", line, next)
		}
	}
}

func getCallerAndLocation(skip int) (string, string) {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		fmt.Println("Caller info: unknown")
		return "unknown", "unknown"
	}
	fn := runtime.FuncForPC(pc)
	caller := fmt.Sprintf("%s:%d", file, line)
	location := fmt.Sprintf("%s:%d", fn.Name(), line)

	return caller, location
}

func isRelevantStacktrace(path []byte) bool {
	isRelevant := bytes.Contains(path, config.GetProjectRoot()) && !bytes.Contains(path, config.GetCoreExceptionPath())
	return isRelevant
}
