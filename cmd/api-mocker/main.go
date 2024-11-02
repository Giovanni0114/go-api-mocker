package main

import (
	"flag"
	"go-api-mocker/pkg/api"
	"go-api-mocker/pkg/apiconfig"
	"go-api-mocker/pkg/schema"
	"log"
	"net/http"
)

const defaultSchemaFilePath = "schemas/schema.json"
const defaultConfigFilePath = "schemas/config.json"

func main() {
	var path string

	flag.StringVar(&path, "s", defaultSchemaFilePath, "path to json schema file")
	flag.StringVar(&path, "c", defaultSchemaFilePath, "path to config file")
	flag.Parse()

	if err := schema.LoadSchema(path); err != nil {
		log.Fatalf("Error loading schema: %v", err)
	}

	if err := apiconfig.LoadConfig(); err != nil {
		log.Fatalf("Error: unable to load schema and config: %v", err)
	}

	api.SetupRoutes()

	log.Printf("Starting server on %v \n", apiconfig.GetPortFormatted())
	log.Fatal(http.ListenAndServe(apiconfig.GetPortFormatted(), nil))
}
