package main

import "fmt"

// func main() { // sementara di disable bentrok main
func main23() {
	// Named Return Values
	// Biasanya saat kita memberi tahu bahwa sebuah function mengembalikan value, maka kita hanya mendeklarasikan tipe data return value di function
	// Namun kita juga bisa membuat variabel secara langsung di tipe data return functionnya
	// Example 1
	firstName, middleName, lastName := getComplateName()
	fmt.Println(firstName, middleName, lastName)
	// Example 2
	newFirstName, newMiddleName, newLastName := getComplateName()
	fmt.Println(newFirstName, newMiddleName, newLastName)

}

// example 1
func getComplateName() (firstName, middleName, lastName string) {
	firstName = "Bachtiar"
	middleName = "Fir"
	lastName = "Daus"
	return firstName, middleName, lastName
}

// example 2
func getNewComplateName() (newFirstName, newMiddleName, newLastName string) {
	newFirstName = "Bachtiar"
	newMiddleName = "Fir"
	newLastName = "Daus"
	return
}
