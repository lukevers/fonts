package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	// Address contains the interface, and port all in one to save
	// time instead of combinding multiple variables into one. For
	// example: interface:port
	Address string
}

// ReadConfig reads the configuration file from JSON and returns it in
// the form of a *Config.
func ReadConfig(path string) (config *Config, err error) {
	file, err := os.Open(path)
	defer file.Close()
	
	if err != nil {
		return
	}
	
	config = &Config{}
	err = json.NewDecoder(file).Decode(config)
	
	return
}