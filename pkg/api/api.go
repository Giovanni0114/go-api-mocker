package api

import (
	"encoding/json"
	"go-api-mocker/internal/random"
	"go-api-mocker/pkg/schema"
	"net/http"
)

func SetupRoutes(schema schema.Schema) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		welcomeHandler(w, r)
	})

	for _, endpoint := range schema.Endpoints {
		endpoint := endpoint // capture range variable
		http.HandleFunc(endpoint.Path, func(w http.ResponseWriter, r *http.Request) {
			value := random.RandomValue(endpoint.Type)
			response, err := json.Marshal(value)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(response)
		})
	}
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the API!"))
}
