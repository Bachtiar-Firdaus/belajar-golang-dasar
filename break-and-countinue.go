package main

import "fmt"

// func main() { // sementara di disable bentrok main
func main18() {
	// break & countinue
	// break & countinue adalah kata kunci yang bisa digunakan dalam perulangan
	// break digunakan untuk menghentikan seluruh perulangan
	// countinue adalah digunakan untuk menghentikan perulangan yang berjalan dan langsung melanjutkanke perulangan selanjutnya
	// example break
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("Perulangan menggunakan break ke", i)
	}
	fmt.Println("===================================")

	// example countinue
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println("Perulangan menggunakan break ke", i)
	}
}
