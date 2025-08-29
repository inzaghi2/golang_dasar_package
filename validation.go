package main

import (
	"fmt"
	"reflect"
)

type Username struct {
	Nama   string `required:"true" max:"10"`
	Alamat string `required:"true" max:"10"`
	Email  string `required:"true" max:"10"`
}

func IsValid(data any) bool {
	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)

	for index := 0; index < t.NumField(); index++ {
		field := t.Field(index)
		if field.Tag.Get("required") == "true" {
			// Assuming the field is a string for this validation
			if v.Field(index).Interface().(string) == "" {
				return false // Found an empty required field, not valid.
			}
		}
	}
	return true // All required fields were checked and are valid.
}

func main() {
	username := Username{
		Nama:   "rifqi inzaghi",
		Alamat: "Semarang",
		Email:  "",
	}
	fmt.Println(IsValid(username))
}
