package exception

import (
	"bytes"
	"errors"
	"fmt"
	"learn-fiber/core/config"
	"runtime"
	"runtime/debug"
)

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

	printCaller(2)
	fmt.Println(se.Err)
	printRelevantStack(se.Stack)
}

func (e *StackError) Error() string {
	return e.Err.Error()
}

// func (e *StackError) Unwrap() error {
// 	return e.Err
// }

func newStackError(err error) error {
	return &StackError{
		Err:   err,
		Stack: debug.Stack(),
	}
}

func printRelevantStack(stack []byte) {
	if config.ProjectRoot == nil {
		return
	}

	lines := bytes.Split(stack, []byte("\n"))
	for i := range len(lines) - 1 {
		line := lines[i]
		next := lines[i+1]

		if bytes.Contains(next, config.ProjectRoot) {
			fmt.Printf("%s\n%s\n", line, next)
		}
	}
}

func printCaller(skip int) {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		fmt.Println("Caller info: unknown")
		return
	}
	fn := runtime.FuncForPC(pc)
	fmt.Printf("Caller: %s:%d\nLocation: %s:%d\n", fn.Name(), line, file, line)
}
