package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	data := ring.New(5)
	for index := 0; index < data.Len(); index++ {
		data.Value = "value " + strconv.FormatInt(int64(index), 10)
		data = data.Next()
	}

	data.Do(func(value interface{}) {
		fmt.Println(value)
	})

}
