package main

import (
	"fmt"
)

// func main() { // sementara di disable bentrok main
func main17() {
	// pada dasarnya for pada go sama dengan bahasa pemrograman lain
	// dalam bahasa pemrograman biasanya ada fitur yang bernama perulangan
	// salah satu fitur perulangan adalah for loops
	counter := 1
	for counter <= 10 {
		fmt.Println("Perulangan ke  ", counter)
		counter++
	}

	// for dengan statement
	// dalam for kita bisa menambahkan statement, dimana terdapat 2 state ment yang bisa tambahkan difor
	// init statement yaitu statement sebelum for diekseskusi
	// post statement, yaitu statement yang akan selalu di akhir tiap perulangan
	// example
	for newCounter := 1; newCounter <= 10; newCounter++ {
		fmt.Println("Perulangan dengan base for yang sama dengan pemrograman lain")
		fmt.Println("Perulangan ke", newCounter)
	}

	// for range
	// for bisa di gunakan untuk melakukan iterasi terhadap semua data collection
	// Data collection contohnya Array, Slice dan Map
	// example old
	slice := []string{"bachtiar", "firdaus", "daus"}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	// for range
	// example new
	friends := []string{"bachtiar", "miqdad", "tyo", "ega"}
	for index, friends := range friends {
		fmt.Println("index", index, "=", friends)
	}

	// for range
	// example hanya butuh data tidak membutuhkan index
	teman := []string{"bachtiar", "miqdad", "tyo", "ega"}
	for _, teman := range teman {
		fmt.Println(teman)
	}

	// for dengan menggunakan map
	// example
	mahasiswa := make(map[string]string)
	mahasiswa["name"] = "ega"
	mahasiswa["semester"] = "6"
	for key, value := range mahasiswa {
		fmt.Println(key, "=", value)
	}
	for _, value := range mahasiswa {
		fmt.Println(value)
	}
}
