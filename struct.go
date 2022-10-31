package main

import "fmt"

// example struct
type Customer struct {
	Name, Address string
	Age           int
}

func main32() {
	// struct
	/*
		struc adalah sebuah template data yang dapat di gunakan untuk menggabungkan nol atau lebih tipe data lainnya dalam satu kesatuan
		struc  biasanya representasi data dalam program aplikasi yang kita buat
		data struct disimpan dalam field
		sederhananya struct adalh kumpulan field
	*/
	// membuat data struct
	/*
		struct adalah template data atau prototype data
		struct tidak bisa langsung digubakan
		Namun kita bisa membuat data/object dari struct yang telah kita buat
	*/
	// example penggunaan struct
	var bioData Customer
	bioData.Address = "Palas"
	bioData.Age = 20
	bioData.Name = "Bachtiar Firdaus"

	fmt.Println(bioData)
	fmt.Println(bioData.Address)
	fmt.Println(bioData.Age)
	fmt.Println(bioData.Name)

	// struct literal
	// Sebelumnya kita telah membuat data struct namun sebenarnya ada banyak cara yang bisa kita gunakan untuk membuat data struct
	// example 1 struct literal
	temen := Customer{
		Name:    "Joko",
		Address: "Indonesia",
		Age:     30,
	}
	fmt.Println(temen)
	// example 2 struct literal
	musuh := Customer{"budi", "indonesia", 30}
	fmt.Println(musuh)

}
