package schema

import (
	"encoding/json"
	"os"
)

type Endpoint struct {
	Path string `json:"path"`
	Type string `json:"type"`
}

type Config struct {
	Port string `json:"port"`
	Key  string `json:"key"`
}

type Schema struct {
	Config    Config     `json:"config"`
	Endpoints []Endpoint `json:"endpoints"`
}

var instance *Schema = &Schema{}

func GetSchema() *Schema {
	return instance
}

func LoadSchema(filePath string) error {
	data, err := os.ReadFile(filePath)

	if err != nil {
		return err
	}

	err = json.Unmarshal(data, instance)
	return err
}
