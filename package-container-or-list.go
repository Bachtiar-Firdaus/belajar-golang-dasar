package main

import (
	"container/list"
	"fmt"
)

func main51() {
	// package container/list
	/*
		package container / list adalah implementasi data double linked list di golang
	*/
	data := list.New()
	data.PushBack("bachtiar")
	data.PushBack("firdaus")
	data.PushBack("palas")

	// data dari depan ke blakang
	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	// data dari depan ke blakang
	for d := data.Back(); d != nil; d = d.Prev() {
		fmt.Println(d.Value)
	}
}
