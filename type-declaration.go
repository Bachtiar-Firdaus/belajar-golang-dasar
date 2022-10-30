package main

import "fmt"

// func main() { // sementara di disable bentrok main
func main8() {
	// ini merupakan representasi nama alias dari tipe data contoh yang sudah merupakan defaul nama alias existing
	// byte merupakan tipe data alias dari uint8
	// contoh example 1
	type NoKTP string
	var NoKTPDaus NoKTP = "9821379821739128793298137"
	fmt.Println(NoKTPDaus)

	// contoh exapmle 2
	type Age int8
	var AgeDaus Age = 20
	fmt.Println(AgeDaus)

}
