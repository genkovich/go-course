package main

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

func temp() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()

	startedAt := time.Now()

	defer func() {
		fmt.Printf("execution took %v\n", time.Since(startedAt))
	}()

	panic("not implemented")
}

func main() {
	test := &DTONewEan{
		Ean: "123",
		Id:  "123",
	}

	fields, err := GetFields(test)
	if err != nil {
		panic(err)
	}

	fmt.Println(fields)
}

func GetFields(e any) ([]string, error) {
	t := reflect.TypeOf(e)

	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}

	if t == nil {
		return nil, fmt.Errorf("nil type")
	}

	if t.Kind() != reflect.Struct {
		return nil, fmt.Errorf("expected struct, got %v", t.Kind())
	}

	var fields []string

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		tagValue := field.Tag.Get("json")

		if tagValue == "" {
			continue
		}

		jsonFieldName := strings.Split(tagValue, ",")[0]
		fields = append(fields, jsonFieldName)
	}

	return fields, nil

}

type DTONewEan struct {
	Ean string `json:"ean" validate:"required"`
	Id  string `json:"id"`
}
