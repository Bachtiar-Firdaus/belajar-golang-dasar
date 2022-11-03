package main

import (
	"belajar-golang-dasar/validator"
	"fmt"
	"reflect"
)

// example 1
type Sample struct {
	Name string
}

// example 2
type SampleV2 struct {
	Name string `required:"true" max:"10"`
}

func main55() {
	// package reflect
	/*
		dalam bahasa pemrograman biasanya ada fitur reflection dimana kita bisa melihat struktur code kita pada saat aplikasi sedang berjalan
		hal ini bisa dilakukan di go-lang dengan menggunakan package reflect
		fitur ini mungkin tidak kita bahas secara lengkap dalam satu vidio
		anda bisa explorasi package reflec ini secara otodidak
		reflection sangat berguna ketika kita ingin membuat library yang general sehingga mudah digunakan
	*/
	sample := Sample{"daus"}
	sampleType := reflect.TypeOf(sample)
	structField := sampleType.Field(0)

	// example 1
	fmt.Println(sampleType.NumField())
	fmt.Println(structField.Name)

	// example 2
	sampleV2 := SampleV2{"daus"}
	sampleTypeV2 := reflect.TypeOf(sampleV2)
	structFieldV2 := sampleTypeV2.Field(0)
	required := structFieldV2.Tag.Get("required")

	// example 2
	fmt.Println(required)

	// example 3
	fmt.Println(validator.IsValid(sample))

}
