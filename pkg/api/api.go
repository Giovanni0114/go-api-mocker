package api

import (
	"encoding/json"
	"go-api-mocker/internal/random"
	"go-api-mocker/pkg/schema"
	"net/http"
)

func SetupRoutes() {
	for _, endpoint := range schema.GetSchema().Endpoints {
		endpoint := endpoint
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
