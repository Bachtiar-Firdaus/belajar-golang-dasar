package main

import "fmt"

// func main() { // sementara di disable bentrok main
func main21() {
	// function return value
	// funtion bisa mengembalikan data
	// untuk memberitahu bahwa function mengembalikan data kita harus menulis tipe data kembalian dari function tersebut
	// jika function tersebut kita deklarasikan dengan tipe data pengembalian maka wajib di dalam functionnya kita harus mengembalikan data
	// untuk function mengembalikan data dari function kita bisa menggunakan kata kunci return diikuti dengan datanya
	result := getNama("bachtiar")
	newResult := newGetNama("bachtiar")
	fmt.Println(result)
	fmt.Println(newResult)
}

func getNama(nama string) string {
	return "Nama saya " + nama
}

func newGetNama(nama string) string {
	if nama == "bachtiar" {
		return "anda bachtiar"
	} else {
		return "siapa anda ?"
	}
}
