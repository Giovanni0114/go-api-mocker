package apiconfig

import (
	"go-api-mocker/pkg/schema"
	"log"
	"strconv"
)

type ApiConfig struct {
	Port   int
}

var instance *ApiConfig = &ApiConfig{}

func LoadConfig() bool {
	config := schema.GetSchema().Config
	data, err := strconv.Atoi(config.Port)

	if err != nil {
		log.Fatalf("[ERROR] verifyConfig: port is not a number, err=%s\n", err.Error())
		return false
	}
	instance.Port = data

	return true
}

func GetPortFormatted() string {
	return ":" + strconv.Itoa(instance.Port)
}
