package main

import "fmt"

// func main() { // sementara di disable bentrok main
func main20() {
	// Function Parameter
	// saat membuat function kadang kadang kita membutuhkan data dari luar atau kita sebut parameter
	// kita bisa menambahkan parameter di funtion bisa lebih dari satu
	// Parameter tidaklah wajib jadi kita bisa membuat function tanpa parameter seperti sebelumnya yang sida kita buat
	// namun jika kita menambahkan parameter di function maka ketika memanggil function tersebut kita wajib memasukkan data ke parameternya
	person("Bachtiar", "Firdaus")
}

// noted : ketika menambahkan parameter wajib menambahkan tipe data dari masing masing parameternya
// urutan definisi parameter wajib sama dengan deklarasi
func person(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName, "!")
}
