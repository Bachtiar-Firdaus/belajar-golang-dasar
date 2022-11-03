package main

import (
	"fmt"
	"sort"
)

// example sort
type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func main53() {
	// package sort
	/*
		package sort adalah package yang berisikan utilitas untuk proses pengurutan
		agar data kita bisa diurutkan kita harus mengimplementasikan kontrak di interface sort interface
	*/
	users := []User{
		{"bachtiar", 60},
		{"joko", 40},
		{"budi", 31},
		{"adit", 33},
	}
	fmt.Println("before", users)
	sort.Sort(UserSlice(users))
	fmt.Println("after", users)

}

func (value UserSlice) Len() int {
	return len(value)
}

func (value UserSlice) Less(i, j int) bool {
	return value[i].Age < value[j].Age
}

func (value UserSlice) Swap(i, j int) {
	value[i], value[j] = value[j], value[i]
}
