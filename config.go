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

	Access struct {
		// EnableFrontEnd is a bool that determins if the
		// front-end interface should be viewed or not. By
		// default this is set to true.
		EnableFrontEnd bool
		
		// EnableAPI is a bool that determins if the API
		// should be ran or not. This should pretty much
		// always be true, but it could also be false if
		// someone decides to hault service. By default this
		// is set to true.
		EnableAPI bool
	}
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