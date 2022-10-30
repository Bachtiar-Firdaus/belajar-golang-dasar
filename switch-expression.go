package main

import "fmt"

// func main() { // sementara di disable bentrok main
func main16() {
	// sama dengan bahasa pemrograman lain selain if expression, untuk melakukan percabangan kita juga bisa menggunakan switch Expression
	// switch expression sangat sederhana dibandingkan if
	// biasanya switch expression digunakan untuk melakukan pengecekan ke kondisi dalam satu variabel
	name := "daus"
	switch name {
	case "bachtiar":
		fmt.Println("Hello Bachtiar")
	case "firdaus":
		fmt.Println("Hello Firdaus")
	default:
		fmt.Println("Hello Boleh Kenalan?")
	}

	// switch dengan short statement
	// sama dengan golang if, Switch juga mendukung short statement sebelum variabel yang akan di cek kondisinya
	// example
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Hello Bachtiar")
	case false:
		fmt.Println("Perlu kenalan sepertinya?")
	}

	// switch tanpa kondisi
	// kondisi di switch expression tidak wajib
	// jika kita tidak menggunakan kondisi di switch expression, kita bisa menambahkan kondisi tersebut di setiap casenya
	// example
	newLength := len(name)
	switch {
	case newLength > 10:
		fmt.Println("nama anda lebih dari 10 karakter")
	case newLength < 10:
		fmt.Println("nama anda kurang dari 10 karakter")
	default:
		fmt.Println("nama 10 karakter")

	}

}
