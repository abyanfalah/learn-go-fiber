package main

import (
	"fmt"
	"learn-fiber/core/config"
)

func main() {
	fmt.Print(config.GetEnv("APP_ENV"))
}
