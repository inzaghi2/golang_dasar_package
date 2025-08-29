package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string
}
type Person struct {
	Name    string `required:"true" max:"10"`
	Address string `required:"true" max:"10"`
	Email   string `required:"true" max:"10"`
}

type sturctTag struct {
	Name string `required:"true" max:"10"`
}

func readFiled(value any) {
	valueType := reflect.TypeOf(value)
	fmt.Println("Type Name", valueType.Name())
	for index := 0; index < valueType.NumField(); index++ {
		var structField reflect.StructField = valueType.Field(index)
		fmt.Println(structField.Name, "with type", structField.Type)
		fmt.Println(structField.Tag.Get("required"))
		fmt.Println(structField.Tag.Get("max"))
	}
}

func main() {
	readFiled(Sample{"rifqi"})
	readFiled(Person{"inzaghi", " ", " "})

}
