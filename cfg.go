package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pelletier/go-toml"
)

// Config is configuration
type Config struct {
	Login struct {
		User     string
		Password string
	}
}

func main() {
	file, err := os.Open("config.toml")
	if err != nil {
		log.Fatalf("error: can't open config file - %s", err)
	}
	defer file.Close()

	cfg := &Config{}
	decoder := toml.NewDecoder(file)

	fmt.Printf("%+v\n", cfg)

	errWhileDecodingConfigFile := decoder.Decode(cfg)
	if errWhileDecodingConfigFile != nil {
		log.Fatalf("error: can't decode configuration file - %s", err)
	}

	fmt.Printf("%+v\n", cfg)
}
