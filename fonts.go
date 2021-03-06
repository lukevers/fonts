package main

import (
	"github.com/inhies/go-log"
	"fmt"
	"os"
)

var (
	l *log.Logger
	LogLevel = log.LogLevel(log.INFO)
	LogFlags = log.Ldate | log.Ltime
	LogFile  = os.Stdout
)

var (
	Conf *Config
	err error
)

func main() {
	// Parse the config file
	Conf, err = ReadConfig("fonts.conf")
	if err != nil {
		fmt.Printf("Could not read configuration file: %s", err)
		os.Exit(1)
	}
	
	// Start the logger
	l, err = log.NewLevel(LogLevel, true, LogFile, "", LogFlags)
	if err != nil {
		fmt.Printf("Could not start logger: %s", err)
		os.Exit(1)
	}
	
	// Start HTTP server
	Serve(Conf)
}