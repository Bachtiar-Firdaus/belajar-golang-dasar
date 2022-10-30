package main

import "fmt"

// func main() { // sementara di disable bentrok main
func main11() {
	// Boolean Operation hanya bisa digunakan pada data boolean
	// example && || !
	// pada dasarnya operation ini sama dengan operation bahasa pemrograman lain

	// example penerapan saya ambil contoh pada AND
	// true && true -> true
	// true && false -> false
	// false && true -> false
	// false && false -> false

	//  example case
	var (
		d1 = true
		d2 = false
	)
	fmt.Println("Base Data")
	fmt.Println(d1, d2)

	fmt.Println("case pada And")
	fmt.Println(d1 && d2)
	fmt.Println(d1 && d2)
	fmt.Println(d1 && d2)
	fmt.Println(d1 && d2)

	fmt.Println("case pada OR")

	fmt.Println(d1 || d2)
	fmt.Println(d1 || d2)
	fmt.Println(d1 || d2)
	fmt.Println(d1 || d2)

	fmt.Println("case pada Inverse")
	fmt.Println(!d1)
	fmt.Println(!d2)

	// example pada penerapan di real code
	var (
		nilaiAkhir = 90
		absensi    = 80
	)
	lulusNilaiAkhir := nilaiAkhir > 80
	lulusAbsensi := absensi > 80

	fmt.Println("Hasil Akhir Nilai Akhir : ", lulusNilaiAkhir)
	fmt.Println("Hasil Akhir Absensi : ", lulusAbsensi)

}
