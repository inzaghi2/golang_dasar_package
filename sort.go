package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (userslice UserSlice) Len() int {
	return len(userslice)
}
func (userslice UserSlice) Less(i, j int) bool {
	return userslice[i].Age < userslice[j].Age
}

func (userslice UserSlice) Swap(i, j int) {
	userslice[i], userslice[j] = userslice[j], userslice[i]
}

type Interface interface {
	//Len is number of elements in the collection
	Len() int
	//Swap swaps the elements with indexes i dan j
	Swap(i, j int)
	//Less reports whether the element with
	//index i should sort before the element with index j
	Less(i, j int) bool
}

func main() {
	user := []User{
		{
			Name: "rifqi",
			Age:  40,
		},
		{
			Name: "inzaghi",
			Age:  20,
		},
		{
			Name: "zaghi",
			Age:  30,
		},
	}
	sort.Sort(UserSlice(user))
	fmt.Println(user)
}
