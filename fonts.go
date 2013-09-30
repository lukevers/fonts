package main

import (
	"github.com/inhies/go-log"
	"fmt"
	"os"
	"database/sql"
)

var (
	l *log.Logger
	LogLevel = log.LogLevel(log.INFO)
	LogFlags = log.Ldate | log.Ltime
	LogFile  = os.Stdout
)

func main() {
	// Parse the config file
	config, err := ReadConfig("fonts.conf")
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

	// Connect to the database
	db, err := sql.Open(config.Database.Driver, config.Database.Resource)
	if err != nil {
		l.Fatalf("Could not connect to database: %s", err)
	}
	l.Info("Connected to database")

	
	
	// Start HTTP server
	Serve(config)
}