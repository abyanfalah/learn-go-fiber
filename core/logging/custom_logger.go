package logging

import (
	"fmt"
	"runtime"
	"strings"
)

func LogWithDetails(msg string) {
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		fmt.Println("Unable to retrieve caller info")
		return
	}

	fn := runtime.FuncForPC(pc)
	funcName := fn.Name() // full path like: github.com/user/project/pkg.Func

	fmt.Println(funcName)

	// Extract package name (last segment before function)
	parts := strings.Split(funcName, "/")
	last := parts[len(parts)-1]             // e.g., "pkg.Func"
	pkgFunc := strings.SplitN(last, ".", 2) // ["pkg", "Func"]
	pkgName := pkgFunc[0]

	fmt.Printf("[%s] %s:%d ➤ %s\n", pkgName, file, line, msg)
}
