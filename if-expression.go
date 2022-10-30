package main

import "fmt"

// func main() { // sementara di disable bentrok main
func main15() {
	// pada dasarnya if expression pada go lang sama dengan pemrograman lainnya
	// if adalah salah satu kata kunci yang di gunakan untuk percabangan
	// Percabangan artinya kita bisa mengeksekusi kode program tertentu ketika suatu kondisi terpenuhi
	// example
	name := "ck"
	if name == "Bachtiar" {
		fmt.Println("Hello Bachtiar")
	} else if name == "ck" {
		fmt.Println("Sepertinya kita sudah penah berkenalan?")
	} else {
		fmt.Println("Hi, Boleh Kenalan?")
	}

	// di golang mendukung shor statement sebelum kondisi
	// hal ini sangat cocok untuk membuat statement yang sederhana melakukan pengecekan terhadap kondisi
	// example
	// dimana bagian ini 	if length := len(name); merupakan sort statement sebelum kondisi terjadi
	if length := len(name); length > 5 {
		fmt.Println("Hello Bachtiar")
	} else {
		fmt.Println("Hi, Boleh Kenalan?")
	}
}
