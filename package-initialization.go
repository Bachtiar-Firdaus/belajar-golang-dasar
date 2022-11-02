package main

import (
	"belajar-golang-dasar/database"
	_ "belajar-golang-dasar/helper" // ini smaple penggunaan blank identifier
	"fmt"
)

func main() {
	// package initialization
	/*
		saat kita membuat package kita bisa membuat sebuah function yang akan diakses ketika package kita diakses
		ini sangat cocok ketika contohnya jka package kita berisi function-function untuk berkomunikasi
		dengan database kita membuat function inisialisasi untuk koneksi database
		untuk membuat function yang diakses secara otomatis ketika package diakses kita cukup membuat function dengan nama init
	*/
	result := database.GetDatabase()
	fmt.Println(result)

	// Blank identifier
	/*
		kadang kita hanya ingin menjalankan init function di package tanpaharus mengeksekusi salah satu function yang ada di package
		secara default go-lang akan komplen ketika ada package yang di import namun tidak digunakan
		untuk menangani hal tersebut kita bisa menggunakan blank identifier(_) sebelum nama package ketika melakukan import
	*/
}
