package main

import "fmt"

// func main() { // sementara di disable bentrok main
func main10() {
	// compare operation adalah operasi perbandingan terhadap dua buah data
	// hasilnya dalam bentuk false / true (boolean)
	// pada dasarnya sama dengan bahasa pemrograman lain
	// example > < >= <= == !=
	var name1 = "Daus"
	var name2 = "Daus"
	var name3 = "Bachtiar"
	// example compare
	var result1 = name1 == name2
	var result2 = name1 == name3
	fmt.Println("Perbandingan yang sama : ", result1)
	fmt.Println("Perbandingan yang tidak sama : ", result2)

	// example pada angka
	var (
		n1 = 1
		n2 = 2
	)
	fmt.Println(n1 > n2)
	fmt.Println(n1 < n2)
	fmt.Println(n1 <= n2)
	fmt.Println(n1 >= n2)
	fmt.Println(n1 == n2)
	fmt.Println(n1 != n2)
}
