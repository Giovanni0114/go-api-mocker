package random

import (
	"fmt"
	"go-api-mocker/pkg/schema"
	"log"
	"math/rand"
	"strconv"
	"strings"
)

var allowedOptions = map[string][]string{
	"string": {"characters", "length", "prefix", "suffix", "choise"},
	"number": {"range", "choise"},
	"const":  {"value"},
	"bool":   {},
}

var randomFunctions = map[string]func(options schema.OptionsMap) string{
	"string": RandomString,
	"number": RandomInt,
	"const":  ReturnConst,
	"bool":   RandomBool,
}

func RandomString(options schema.OptionsMap) string {
	fmt.Println("Options for string")

	for key, value := range options {
		fmt.Printf("  %s: %v\n", key, value)
	}

	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, 10)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func ReturnConst(options schema.OptionsMap) string {
	value, ok := options["value"]

	if !ok {
		log.Println("No 'value' option in const endpoint")
		return "undefined const"
	}

	val, ok := value.(string)
	if !ok {
		log.Println("'value' option is not of type string")
		return "undefined const"
	}

	return val
}

func ConvertSlice[E any](in []any) (out []E) {
	out = make([]E, 0, len(in))
	for _, v := range in {
		out = append(out, v.(E))
	}
	return
}

func RandomInt(options schema.OptionsMap) string {
	if choise, isPresent := options["choise"]; isPresent {
		slice, ok := choise.([]interface{})
		if !ok {
			log.Println("'choise' is not correct in `number` type endpoint ")
			return "incorrect configuration"
		}

		if len(options) > 1 {
			log.Println("WARNING: option 'choise' provided while other options is present, this other options have no effect")
		}
		randIdx := rand.Intn(len(slice))

		ans, ok := slice[randIdx].(float64)
		if !ok {
			log.Printf("'choise' is not correct in `number` type endpoint %d", slice[randIdx])
			return "incorrect configuration"
		}
		return strconv.Itoa(int(ans))
	}

	base := 0
	randomRange := 0

	if rangeVal, isPresent := options["range"]; isPresent {
		ran, ok := rangeVal.(string)
		if !ok {
			log.Println("'range' is not correct in `number` type endpoint ")
			return "incorrect configuration"
		}

		splited := strings.Split(ran, "-")

		if len(splited) != 2 {
			log.Println("'range' is not correct in `number` type endpoint ")
			return "incorrect configuration"
		}

		baseCandidate, baseOk := strconv.Atoi(splited[0])
		randomRangeCandidate, rangeOk := strconv.Atoi(splited[1])

		if baseOk == nil {
			base = baseCandidate
		}

		if rangeOk == nil {
			randomRange = randomRangeCandidate
		}
	}

	return strconv.Itoa(rand.Intn(randomRange) + base)
}

func RandomBool(options schema.OptionsMap) string {
	if len(options) > 0 {
		log.Println("options have no effect on bool endpoints")
	}

	return strconv.FormatBool(rand.Intn(2) == 0)
}

func RandomValue(endpoint schema.Endpoint) string {
	function, isPresent := randomFunctions[endpoint.Type]
	if isPresent {
		return function(endpoint.Options)
	}

	log.Printf("Detected unsupported type `%s` \n", endpoint.Type)
	return "unsupported type"
}
