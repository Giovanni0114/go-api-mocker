package main

import (
	"log"
	"go-api-mocker/pkg/api"
	"go-api-mocker/pkg/schema"
	"net/http"
)

func main() {
	schemaFilePath := "schemas/schema.json"
	schema, err := schema.LoadSchema(schemaFilePath)
	if err != nil {
		log.Fatalf("Error loading schema: %v", err)
	}

	api.SetupRoutes(schema)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
