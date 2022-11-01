package main

import "fmt"

// example interface atau kontrak
type HasName interface {
	GetName() string // string disini sebagai nilai return
}

// example stract
type Person struct {
	Name string
}

// example struct 2
type Animal struct {
	Name string
}

func main34() {
	// Interface
	// notes : interface hanya bisa digunakan jika stract telah terbuat dengan sefinisi yang sama
	/*
		Interface adalah tipe data abstract dia tidak memiliki implementasi langsung
		sebuah interface berisikan definisi definisi method
		biasanya interface digunakan sebagai kontrak
	*/

	// Implementasi interface
	/*
		Setiap tipe data yang sesuai dengan kontrak interface secara otomatis dianggap sebagai interface tersebut
		sehingga kita tidak perlu mengimplementasian interface secara manual
		hal ini agak berbeda dengan bahasa pemerograman lain yang ketika membuat interface kita harus menyebutkan secara eksplisit akan menggunakan interface mana
	*/
	// example struct dengan interface
	person := Person{Name: "Bachtiar"}
	SayHello(person)
	// example struct dengan interface 2
	animal := Animal{Name: "Caty"}
	SayHelloV2(animal)

}

// example penggunaan interface
func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

// example penggunaan interface 2
func SayHelloV2(hasName HasName) {
	fmt.Println("Nama Peliharaan ", hasName.GetName())
}

// example penerapan struct
func (person Person) GetName() string {
	return person.Name
}

// example penerapan struct 2
func (animal Animal) GetName() string {
	return animal.Name
}
