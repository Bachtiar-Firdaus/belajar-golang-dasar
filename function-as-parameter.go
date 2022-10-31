package main

import "fmt"

// kita sudah pernah membuat alias yang berguana mempersingkat deklarasi
// example penerapan
type Filter func(int) int
type Warna func(string) string
type Merek func(string) string
type Ukuran func(string) string

// func main() { // sementara di disable bentrok main
func main26() {
	// function as parameter
	// function tidak hanya bisa kita simpan di dalam variabel sebagai value
	// namun juga kita bisa gunakan sebagai parameter untuk function lain
	// example 1
	sayHelloWithFileter("Bachtiar", spamFilter)
	filter := spamFilter
	sayHelloWithFileter("Anjing", filter)

	// Function Type Declarations
	// kadang jika function terlalu panjang agak ribet untuk menulikannya di dalam parameter
	// type declaration juga bisa digunakan untuk membuat alias function sehingga akan mempermudah kita menggunakan function sebagai parameter

	// example 2 with sort type
	umurFilter(20, masaFilter)
	// pindahkan function ke dalam variabel
	newMasaFilter := masaFilter
	umurFilter(10, newMasaFilter)

	newWarnaFilter := warnaFilter
	bajuFilter("Merah", newWarnaFilter)

	newMerekFIlter := merekFilter
	sendalFilter("adidas", newMerekFIlter)

	mejaFilter("kecil", ukuranFilter)
	newUkuranFilter := ukuranFilter
	mejaFilter("besar", newUkuranFilter)

}

// example 1
func sayHelloWithFileter(name string, filter func(string) string) {
	nameFiltered := filter(name)
	fmt.Println("Hello ", nameFiltered)
}

// example 1
func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

// example 2 with sort type
func umurFilter(umur int, cek Filter) {
	umurFiltered := cek(umur)
	fmt.Println("Umur Kamu ", umurFiltered)
}

// example 2 with sort type
func masaFilter(umur int) int {
	if umur == 20 {
		return 20
	} else {
		return 0
	}
}

func bajuFilter(baju string, pribadi Warna) {
	newPribadi := pribadi(baju)
	fmt.Println("Baju Kamu", newPribadi)
}

func warnaFilter(baju string) string {
	if baju == "putih" {
		return baju
	} else {
		return "bukan putih"
	}
}

func sendalFilter(sendal string, brand Merek) {
	newBrand := brand(sendal)
	fmt.Println("Sendal kamu termasuk : ", newBrand)
}

func merekFilter(sendal string) string {
	if sendal == "adidas" {
		return "barang mahal"
	} else {
		return "barang murahan"
	}
}

func mejaFilter(meja string, kayu Ukuran) {
	size := kayu(meja)
	fmt.Println("Ukuran Kayu Kamu", size)
}

func ukuranFilter(meja string) string {
	if meja == "besar" {
		return "besar"
	} else {
		return "kecil"
	}
}
