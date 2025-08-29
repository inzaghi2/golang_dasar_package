package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)

	utc := time.Date(2024, time.November, 8, 11, 0, 0, 0, time.UTC)
	fmt.Println(utc)
	fmt.Println(utc.Local())

	formatter := "2006-01-02 15:04:05"
	value := "2024-11-08 11:00:00"
	valueTime, err := time.Parse(formatter, value)
	if err != nil {
		fmt.Println("error", err.Error())

	} else {
		fmt.Println(valueTime)
	}

}
