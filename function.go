package main

import "fmt"

// func main() { // sementara di disable bentrok main
func main19() {
	// Function
	// Sebelumnya kita sudah mengenal sebuah function yang wajib di buat agar program kita bisa berjalan yaitu function main
	// function adalah sebuah blok kode yang sengaja dibuat dalam program agar kita bisa di guankan berulang ulang
	// cara membuat function sangat sederhana hanya dengan menggunakan kata func lalu diikuti dengan nama functionnya dengna block kode isi functionnya
	// setelah membuat function kita bisa mengeksekusi function tersebut dengan memanggilnya menggunakan kata kunci nama function nya yang diikuti tanda kurung buka kurung tutup
	// example
	sayHello()
	fmt.Println("================")
	for i := 0; i < 10; i++ {
		sayHello()
	}
}
func sayHello() {
	fmt.Println("Hello Word !!")
}
