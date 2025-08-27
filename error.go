package main

import (
	"errors"
	"fmt"
)

var (
	ValidationError = errors.New("validation error")
	NotFoundError   = errors.New("not found error")
)

func GetByID(id string) error {
	if id == "" {
		return ValidationError
	}
	if id != "rifqi" {
		return NotFoundError
	}
	return nil
}

func main() {
	err := GetByID("rifq")
	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Println("validation error:", err)
		} else if errors.Is(err, NotFoundError) {
			fmt.Println("data not found:", err)
		} else {
			fmt.Println("Uknown Error")
		}

	}

}
