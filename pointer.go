package main

import "fmt"

type Address struct {
	daerah1, daerah2, daerah3 string
}

type NewAddress struct {
	daerah1, daerah2, daerah3 string
}

type Umur struct {
	u1, u2, u3 int
}

type Meja struct {
	m1, m2, m3 int
}

func main39() {
	//pointer
	/*
		pass by value
		secara default di golang semua variabel itu di passing by value bukan by refrence
		artinya jika kita mengirimkan sebuah variabel ke dalam function method atau variabel lainsebenernya yang dikirim adalah duplikasi valuenya
	*/
	address1 := Address{"subang", "palas", "kalianda"}
	address2 := &address1 // data tidak by reverence melainkan duplikasi value

	address2.daerah1 = "Bandung"
	fmt.Println("========tanpa pointer========")
	fmt.Println(address1) // address1 tidak berubah
	fmt.Println(address2)

	// pointer
	// pointer adalah kemampuan membuat refrence ke lokasi data di memory yang sama tanpa menduplikasi data yang sudah ada
	// sederhananya dengan kemampuan pointer kita bisa membuat pass by refrence

	// opretor &
	// untuk membuat sebuah variabel dengan nilai pointer ke variabel yang lain  kita bisa menggunakan operator & diikuti dengan nama variabelnya

	// example 1
	newAddress1 := NewAddress{"rusia", "jakarta", "singapura"}
	newAddress2 := &newAddress1
	newAddress2.daerah2 = "amerika"
	fmt.Println("========dengan pointer========")
	fmt.Println(newAddress1) // address1 tidak berubah
	fmt.Println(newAddress2)

	// operator *
	/*
		saat kita mengubah variabel pointer, maka yang berubah hanya variabel tersebut
		semua variabel yang mengacu ke data yang sama tidak akan berubah
		jika kita ingin mengubah seluruh variabel yang mengacu ke data tersebut kita bisa menggunakan operator *
	*/
	// example 2
	*address2 = Address{"subang", "palas", "kalianda"}
	fmt.Println(address1)
	fmt.Println(address2)

	// example 3
	fmt.Println("=======================================")
	umur1 := Umur{1, 2, 3}
	umur2 := &umur1
	umur3 := &umur1
	umur2.u1 = 10
	*umur2 = Umur{11, 22, 33} // disinilah penggunaan pointer untuk merefrence ke 1 memory satu di ubah semua berubah
	fmt.Println(umur1)
	fmt.Println(umur2)
	fmt.Println(umur3)

	// function new
	/*
		sebelumnya untuk membuat pointer dengan menggunakan operator &
		Go-Lang memiliki function new yang bisa digunakan untuk membuat pointer
		namun function new hanya mengembalikan pointer ke data kosong artinya tidak ada data awal
	*/
	fmt.Println("=======================================")
	// example kode program function new
	meja1 := new(Meja)
	meja2 := meja1
	meja2.m1 = 10
	fmt.Println(meja1)
	fmt.Println(meja2)

}
