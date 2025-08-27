package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("rifqi inzaghi", "rifqi"))
	fmt.Println(strings.Split("rifqi inzaghi", " "))
	fmt.Println(strings.ToLower("Rifqi Inzaghi"))
	fmt.Println(strings.ToUpper("rifqi inzaghi"))
	fmt.Println(strings.Trim("   rifqi inzaghi   ", " "))
	fmt.Println(strings.ReplaceAll("rifqi inzaghi", "zaghi", "inza"))

}
