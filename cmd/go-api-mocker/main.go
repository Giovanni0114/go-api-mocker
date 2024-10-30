package main

import (
	"go-api-mocker/pkg/api"
	"go-api-mocker/pkg/apiconfig"
	"go-api-mocker/pkg/schema"
	"log"
	"net/http"
)

func loadSchema() bool {
	schemaFilePath := "schemas/schema.json"

	if err := schema.LoadSchema(schemaFilePath); err != nil {
		log.Fatalf("Error loading schema: %v", err)
		return false
	}

	return true
}

func main() {
	if !loadSchema() || !apiconfig.LoadConfig() {
		log.Fatal("Error: unable to load schema and config")
	}

	api.SetupRoutes()

	log.Printf("Starting server on %v \n", apiconfig.GetPortFormatted())
	log.Fatal(http.ListenAndServe(apiconfig.GetPortFormatted(), nil))
}

