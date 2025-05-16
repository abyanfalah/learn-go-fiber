package config

import (
	"fmt"
	"log"
	"os"
)

var projectRoot []byte
var coreExceptionPath []byte
var IsRelevantLoggingPath bool

// todo: find reliable way to retrieve project root. os.getwd can still go wrong
func init() {
	setProjectRoot()
	setCoreExceptionPath()
}

func setProjectRoot() {
	if wd, err := os.Getwd(); err == nil {
		projectRoot = []byte(wd)
	} else {
		log.Fatalf("Failed to get project root: %v", err)
	}
}

func setCoreExceptionPath() {
	if wd, err := os.Getwd(); err == nil {
		coreExceptionPath = fmt.Appendf(nil, "%s/%s", wd, "core/exception")
	} else {
		log.Fatalf("Failed to get error logging path: %v", err)
	}
}

func GetProjectRoot() []byte {
	if projectRoot == nil {
		log.Fatalf("Project root not set")
		return nil
	}

	return projectRoot
}
func GetCoreExceptionPath() []byte {
	if coreExceptionPath == nil {
		log.Fatalf("Error logging path not set")
		return nil
	}
	return coreExceptionPath
}
