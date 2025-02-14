package config

import (
	"fmt"
	"os"
)

func getEnvWithoutParser(name string, optional bool) string {
	value := os.Getenv(name)
	if value == "" && !optional {
     panic(fmt.Sprintf("env %s variable is not or its empty ",name))
	}
	return value
}
type parseFunc[T any] func(string) (*T, error)

func getEnv[T any](name string, optional bool, parser parseFunc[T]) T {
	value:= getEnvWithoutParser(name, optional)
	parsedValue,err := parser(value)
	if err != nil {
		panic(fmt.Sprintf("[%s] %s",name,err))
	}
	return *parsedValue
}