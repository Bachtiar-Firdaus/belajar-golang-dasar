package main

import "fmt"

// func main() { // sementara di disable bentrok main
func main5() {
	var name string
	name = "Bachtiar Firdaus"
	fmt.Println(name)

	name = "Testing Manipulasi Data Variabel"
	fmt.Println(name)

	var friendName = "Tyo Saputra"
	fmt.Println(friendName)

	var age = 20
	fmt.Println(age)

	// deklarasi nama var dalam variabel di go tidak wajib
	// bisa digantikan dengan := example

	nameNew := "Bachtiar Firdaus"
	fmt.Println(nameNew)

	friendNameNew := "Tyo Saputra"
	fmt.Println(friendNameNew)

	ageNew := 20
	fmt.Println(ageNew)

	// ketika melakukan reusing variabel tidak menggunakan := melainkan =

	nameNew = "Bachtiar Firdaus - Reusing data"
	fmt.Println(nameNew)

	// deklarasi variabel sekaligus banyak
	var (
		brand       = "Vivo"
		secondBrand = "Oppo"
	)
	fmt.Println(brand, secondBrand)
}
