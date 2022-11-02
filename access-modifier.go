package main

import (
	"belajar-golang-dasar/helper"
	"fmt"
)

func main44() {
	// Access Modifier
	/*
		dibahasa pemrograman lain biasanya ada kata kunci yang bisa digunakan untuk menentukan access modifier terhadap suatu function atau variabel
		di golang untuk menentukan access modifier cukup dengan menggunakan nama function atau variable
		jika namanya diawali dengan huruf besar maka artinya bisa diakses dari package lain jika dimulai dengan huruf kecil artinya tidak bisa diakses dari package lain
	*/
	// example
	// example 1
	name := "daus"
	result1 := helper.SayHello(name)
	fmt.Println(result1)

	//  disini function yang awalannya kecil tidak bisa di akses
	// result2 := helper.sayGoodBye(name) // error

}
