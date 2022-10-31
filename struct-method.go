package main

import "fmt"

type newCustomer struct {
	Name, Address string
	Age           int
}
type newBio struct {
	Name, Address string
	Age           int
}

func main33() {
	// struct method
	/*
		struct adalah tipe data seperti tipe data yang lainnya, dia bisa digunakan sebagai parameter untuk function
		namun jika kita ingin menambahkan method kedalam structs, sehingga seakan akan sebuah struct memiliki function
		method adalah function
	*/
	// exmaple
	rully := newCustomer{Name: "Rully"}
	rully.sayHi()

	var reformat newBio
	reformat.Address = "Palas"
	reformat.Age = 20
	reformat.Name = "Bachtiar Firdaus"
	reformat.sayHiV2("daus")

}

// ini merupakan struct function
// example 1
func (customer newCustomer) sayHi() {
	fmt.Println("Hello, My name is", customer.Name)
}

//example 2
func (bio newBio) sayHiV2(name string) {
	fmt.Println("Hello", name, " My name is", bio.Name)

}
