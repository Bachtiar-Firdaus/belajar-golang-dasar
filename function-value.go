package main

import "fmt"

// func main() { // sementara di disable bentrok main
func main25() {
	// Function Value atau bisa disebut function yang disimpan didalam variabel
	// Function adalah first class citizen atau bisa dikatakn function di golang tidak di anak tirikan
	// Function juga merupakan tipe data dan bisa disimpan di dalam variabel
	// example 1
	goodbye := getGoodBye
	fmt.Println(goodbye("Bachtiar Firdaus"))
	// example 2

	result := goodbye("Daus")
	fmt.Println(result)

}

func getGoodBye(name string) string {
	return "Good Bye " + name
}
