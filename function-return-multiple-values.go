package main

import "fmt"

// func main() { // sementara di disable bentrok main
func main22() {
	// returning multiple values
	// function tidak hanya dapat mengembaikan satu value tapi juga bisa multiple value
	// untuk memberitahu jika function mengembalikan multiple value kita harus menulis semua tipe data return valuenya di function
	// example
	firstName, middleName, lastName := getFullName()
	fmt.Println(firstName, middleName, lastName)

	// Menghiraukan Return value
	// Multiple return value wajib di tangkap semua valuenya
	// jika kita ingin menghiraukan value tersebut kita bisa menggunakan tanda _ (garis bawah/underscore)
	// example
	newfirstName, _, _ := getFullName()
	fmt.Println(newfirstName)
}

func getFullName() (string, string, string) {
	return "bachtiar", "fir", "daus"
}
