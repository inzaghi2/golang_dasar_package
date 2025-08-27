package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(strconv.ParseBool("true"))
	fmt.Println(strconv.ParseFloat("3.14", 64))
	fmt.Println(strconv.ParseInt("123", 10, 64))
	fmt.Println(strconv.FormatBool(true))
	fmt.Println(strconv.FormatFloat(3.14, 'f', 6, 64))
	fmt.Println(strconv.FormatInt(123, 10))

	result, err := strconv.ParseBool("salah")
	if err != nil {
		fmt.Println("error", err.Error())
	} else {
		fmt.Println(result)
	}

	resultFloat, err := strconv.ParseFloat("3.14", 64)
	if err != nil {
		fmt.Println("error", err.Error())
	} else {
		fmt.Println(resultFloat)
	}

	resultInt, err := strconv.Atoi("salah")
	if err == nil {
		fmt.Println(resultInt)
	} else {
		fmt.Println("error", err.Error())

	}

}
