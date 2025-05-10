package config

import "os"

var ProjectRoot []byte

// todo: find reliable way to retrieve project root. os.getwd can still go wrong
func init() {
	if wd, err := os.Getwd(); err == nil {
		ProjectRoot = []byte(wd)
	}
}
