package config

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

var once sync.Once

// GetEnv: actual env fetching logic
func GetEnv(key string) string {
	// set the real environment from .env file ONLY ONCE
	once.Do(InitEnv)

	envValue, isExists := os.LookupEnv(key)
	if !isExists {
		log.Fatalf("Required environment variable %q not set", key)
		// TODO throws exception. figure it out later
		return "== unset =="
	}

	return strings.TrimSpace(envValue)
}

// loads environment variable if not in production
func InitEnv() {
	envValue, isExists := os.LookupEnv("APP_ENV")

	if !isExists {
		fmt.Println("Warning: APP_ENV is not initally set.")
		// return
	}

	switch strings.ToLower(strings.TrimSpace(envValue)) {
	case "production", "prod":
		return
	}

	loadDotEnv(".env")
}

// func loadAppEnv(){
// 	file, err := os.Open(".env")
// 	if err != nil {
// 		log.Printf("Warning: could not open .env file (%s): %v", ".env", err.Error)
// 		return
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		line := scanner.Text()

// 		// commented checking
// 		if strings.HasPrefix(line, "#") || strings.TrimSpace(line) == "" {
// 			continue
// 		}

// 		// less than 2 parts checking
// 		parts := strings.SplitN(line, "=", 2)
// 		if (len(parts)) != 2 {
// 			continue
// 		}

// 		// setting each env
// 		key := strings.TrimSpace(parts[0])
// 		value := strings.Trim(strings.TrimSpace(parts[1]), `"'`)
// 		fmt.Println("setting " + key + " : " + value)
// 		os.Setenv(key, value)

// 		if err := scanner.Err(); err != nil {
// 			log.Fatalf("Error reading .env file: %v", err)
// 		}
// 	}
// }

// loading environment variable from .env file
func loadDotEnv(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Printf("Warning: could not open .env file (%s): %v", path, err.Error)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// commented checking
		if strings.HasPrefix(line, "#") || strings.TrimSpace(line) == "" {
			continue
		}

		// less than 2 parts checking
		parts := strings.SplitN(line, "=", 2)
		if (len(parts)) != 2 {
			continue
		}

		// setting each env
		key := strings.TrimSpace(parts[0])
		value := strings.Trim(strings.TrimSpace(parts[1]), `"'`)
		fmt.Println("setting " + key + " : " + value)
		os.Setenv(key, value)

		if err := scanner.Err(); err != nil {
			log.Fatalf("Error reading .env file: %v", err)
		}
	}
}
