package main

import (
	"fmt"
	"os"
)

func main46() {
	//	package os
	/*
		golang telah menyediakan banyak sekali package bawaan saah satunya adalah package os
		package os berisikan fungsionaitas untuk mengakses fitur sistem operasi secara independan
		(bisa digunakan disemua sistem operasi)
	*/
	args := os.Args
	fmt.Println("Argument :", args)
	// fmt.Println("Argument :", args[1])
	// fmt.Println("Argument :", args[2])
	// go run package-os.go bachtiar firdaus
	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname :", hostname)
	} else {
		fmt.Println("Error", err.Error())
	}
}
