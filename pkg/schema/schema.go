package schema

import (
	"encoding/json"
	"os"
)

type OptionsMap map[string]interface{};

type Endpoint struct {
	Path string `json:"path"`
	Type string `json:"type"`
	Options OptionsMap `json:"options"`
}

type Schema struct {
	Config    Config                 `json:"config"`
	Endpoints []Endpoint             `json:"endpoints"`
}

type Config struct {
	Port string `json:"port"`
}


var instance *Schema = &Schema{}

func GetSchema() *Schema {
	return instance
}

func LoadSchema(schemaFilePath string) error {
	data, err := os.ReadFile(schemaFilePath)

	if err != nil {
		return err
	}

	err = json.Unmarshal(data, instance)
	return err
}
