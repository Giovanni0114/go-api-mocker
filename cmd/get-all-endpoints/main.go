package main

import (
	"flag"
	"go-api-mocker/pkg/apiconfig"
	"go-api-mocker/pkg/schema"
	"log"
	"net/http"
    "io"
)

const defaultSchemaFilePath = "schemas/schema.json"
const defaultConfigFilePath = "schemas/config.json"

func makeRequest(url string) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatalf("  * ERROR: http request: %s\n", err)
	}

	bodyBytes, _ := io.ReadAll(res.Body)
	log.Printf("  * got response [%d] => %s", res.StatusCode, string(bodyBytes))
}

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

	var portFormatted = apiconfig.GetPortFormatted()
	baseUrl := "http://localhost" + portFormatted

	for _, endpoint := range schema.GetSchema().Endpoints {
		url := baseUrl + endpoint.Path
		log.Printf("Making request for [%s] => [%s]", url, endpoint.Type)
		makeRequest(url)
	}
}
