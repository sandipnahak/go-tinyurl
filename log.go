package main

import (
	"log"
	"os"
)

func getLogger() *log.Logger {
	logger := log.New(os.Stdout, "HTTPServer ", log.Ltime|log.Ldate|log.Lshortfile)

	return logger

}
