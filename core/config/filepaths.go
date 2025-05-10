package config

import (
	"fmt"
	"os"
)

var ProjectRoot []byte
var ErrorLoggingPath []byte
var IsRelevantLoggingPath bool

// todo: find reliable way to retrieve project root. os.getwd can still go wrong
func init() {
	if wd, err := os.Getwd(); err == nil {
		ProjectRoot = []byte(wd)
	}

	ErrorLoggingPath = fmt.Appendf(nil, "%s/%s", ProjectRoot, "core/exception")
}
