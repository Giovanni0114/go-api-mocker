package apiconfig

import (
	"go-api-mocker/pkg/schema"
	"strconv"
)

type ApiConfig struct {
	Port   int
}

var instance *ApiConfig = &ApiConfig{}

func LoadConfig() error {
	data, err := strconv.Atoi(schema.GetSchema().Config.Port)
	defer func() {instance.Port = data}()
	return err
}

func GetPortFormatted() string {
	return ":" + strconv.Itoa(instance.Port)
}
