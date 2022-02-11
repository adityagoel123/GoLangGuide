// pkg/errors example
package main

import (
	"fmt"
	"os"
)

// Config holds configuration
type Config struct {
	// configuration fields go here (redacted)
	Name   string
	Gender string
	Age    int
}

func readConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	cfg := &Config{}
	// Parse file here (redacted)
	return cfg, nil
}

func main() {
	//cfg, err := readConfig("/path/to/config.toml")
	cfg, err := readConfig("/Users/B0218162/Documents/LEARNINGS/MEDIUM-BLOG/GO/config.toml")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	} else {
		fmt.Println("File was read succesfully.")
	}
	// Normal operation (redacted)
	fmt.Println(cfg)
}
