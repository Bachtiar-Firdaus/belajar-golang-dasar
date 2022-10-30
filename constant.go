package main

import "fmt"

// func main() { // sementara di disable bentrok main
func main6() {
	// Ketika menggunakan constant wajib langsung mengisi data didalam variabelnya
	const firstName string = "Daus"
	const lastName = "Bachtiar"
	const number = 1000

	// Perbedaan Constant dengan VAR apa bila variabel contant telah berisi data dan tidak ada action setelahnya makan go tidak mendeteksi error sedangkan var sebaliknya

	// sama dengan pemrograman lain constant tidak dapat di ubah dan di reusing variabelnya akan menyebabkan ERROR
	// firstName = "TIDAK BISA DI UBAH"
	// firstName = "TIDAK BISA DI UBAH"

	// example ketika di use or tidak varibel tidak di complaint oleh go
	// fmt.Println(firstName, lastName, number)

	// declarasi multiple constant example
	const (
		brand       = "Vivo"
		secondBrand = "Oppo"
	)
	fmt.Println(brand, secondBrand)

}
