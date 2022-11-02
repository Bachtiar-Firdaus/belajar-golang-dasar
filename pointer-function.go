package main

import "fmt"

type ck struct {
	daerah1, daerah2, daerah3 string
}

// example pointer function
type newCK struct {
	a, b, c string
}

func main40() {
	// Pointer di function
	/*
		saat kita membuat parameter di function, secara default adalah pass by value artinya data akan di copy lalu dikirim ke funtion tersebut
		oleh karna itu jika kita mengubah data di dalam function data yang aslinya tidak akan pernah berubah
		hal ini membuat variabel menjadi akan karna tidak akan bisa diubah
		namun kadang kita ingin membuat function yang bisa mengubah data asli parameter tersebut
		untuk melakukan ini kita bisa menggunakan pointer function
		untuk menjadikan sebuah parameter sebagai pointer, kita bisa menggunakan operator * di parameternya
	*/
	//	tanpa pointer data tidak berubah
	// example bukan pointer function
	nck := ck{"subang", "jawa barat", ""}
	ChangeAddressToIndonesia(nck)
	fmt.Println(nck) // tidak berubah

	// example function pointer
	test := newCK{"data1", "data2", ""}
	examplePointer(&test)
	fmt.Println(test) // tidak berubah
}

// example bukan pointer function
func ChangeAddressToIndonesia(nck ck) {
	nck.daerah1 = "jakarta"
}

// example function pointer
func examplePointer(test *newCK) {
	test.c = "ini-akan-merubah-master-data"
}
