package schema

import (
	"encoding/json"
	"io/ioutil"
)

type Endpoint struct {
	Path string `json:"path"`
	Type string `json:"type"`
}

type Schema struct {
	Endpoints []Endpoint `json:"endpoints"`
}

func LoadSchema(filePath string) (Schema, error) {
	var schema Schema
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return schema, err
	}
	err = json.Unmarshal(data, &schema)
	return schema, err
}
