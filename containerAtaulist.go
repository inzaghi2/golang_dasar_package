package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushBack("rifqi")
	data.PushBack("inzaghi")
	data.PushBack("zaghi")

	for elemnt := data.Front(); elemnt != nil; elemnt = elemnt.Next() {
		fmt.Println(elemnt.Value)
	}

}
