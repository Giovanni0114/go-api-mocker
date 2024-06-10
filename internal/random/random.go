package random

import (
	"math/rand"
)

func RandomString() string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, 10)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func RandomBool() bool {
	return rand.Intn(2) == 0
}

func RandomInt() int {
	return rand.Intn(100)
}

func RandomValue(valueType string) interface{} {
	switch valueType {
	case "string":
		return RandomString()
	case "number":
		return RandomInt()
	case "bool":
		return RandomBool()
	default:
		return "unsupported type"
	}
}
