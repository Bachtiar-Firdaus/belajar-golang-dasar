package main

import (
	"fmt"
)

// func main() { // sementara di disable bentrok main
func main14() {
	person := map[string]string{
		"name":    "Daus",
		"address": "Palas",
	}
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	// untuk merubah data dari map
	// example
	person["name"] = "Bachtiar"
	fmt.Println(person)

	// sample function Map
	// len(map) Untuk mendapatkan jumlah data di map
	// map[key] Mengambil data di map dengan key
	// map[key] = value mengubah data di map dengan key
	// make(map[TypeKey]TypeValue) membuat map baru
	// delete(map,key) Menghapus data dimap dengan key

	// example
	// cara penulisan dengan model lama
	// var book map[string]string = make(map[string]string)
	// cara penulisan dengan model baru
	book := make(map[string]string)
	book["title"] = "Buku Go-Lang"
	book["author"] = "Bachtiar Firdaus"
	book["wrong"] = "Ups"
	// case penghapusan data pada map
	fmt.Println("Data Sebelum Dihapus", book)
	fmt.Println("Data Sebelum Dihapus", len(book))
	delete(book, "wrong")
	fmt.Println("Data Setelah Dihapus", book)
	fmt.Println("Data Setelah Dihapus", len(book))

}
