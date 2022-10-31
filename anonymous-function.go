package main

import "fmt"

// kita buat alias
type BlackList func(string) bool
type UmurFilter func(int) bool

func main27() {
	// Anonymouse Function di golang
	// Sebelumnya setiap membuat function kita selalu memberikan sebuah nama pada function tersebut
	// Namun kadang ada kalanya lebih mudah membuat function secara langsung di variabel atau parametertanpa harus membuat function terlebih dahulu
	// hal tersebut dinamakan anonymouse function atau function tanpa nama

	// blacklist := karna anonymus kita bisa langsung kemas dalam variabel
	// example 1

	// // example 2
	// registerUser("anjing", func(name string) bool {
	// 	return name == "anjing"
	// })
	// Declaration anonymouse function
	blacklist := func(name string) bool {
		return name == "anjing"
	}
	umurfilter := func(umur int) bool {
		return umur < 17
	}

	// example case 1
	fmt.Println("=============Example case 1================")
	registerUserV1("Bachtiar", func(name string) bool {
		return name == "Bactiar"
	})

	// example case 2
	fmt.Println("=============Example case 2================")
	registerUserV2("bachtiar", 20, blacklist, umurfilter)

}

// example case 1
func registerUserV1(name string, blacklist BlackList) {
	if blacklist(name) {
		fmt.Println("You Are Blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

// example case 2
func registerUserV2(name string, umur int, blacklist BlackList, umurfilter UmurFilter) {
	// blacklist disini merupakan kategory anonymous function
	// functiontidak dibuat diluar main function
	if blacklist(name) {
		fmt.Println("You Are Blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
	if umurfilter(umur) {
		fmt.Println("Anda Dalam Kategori dibawah umur", umur)
	} else {
		fmt.Println("Anda Dalam Kategori Cukup Dewasa", umur)
	}
}
