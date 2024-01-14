package main

import (
	"os"

	log "github.com/sirupsen/logrus"

	// Blank-import the function package so the init() runs
	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	_ "realizr.io/iso20022"
)

func main() {
	// Use PORT environment variable, or default to 8080.
	port := "8080"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}
	log.Info("Running on port: ", port)
	if err := funcframework.Start(port); err != nil {
		log.Error("funcframework.Start: ", err)
	}
}
