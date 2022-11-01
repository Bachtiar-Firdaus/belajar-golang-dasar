package main

import "fmt"

func main36() {
	// nill
	/*
		Biasanya didalam bahasa pemrograman lain object yang belum diinisialisasi maka secara otomatis nilainya adalah null atau nil
		Berbeda dengan golang di golang saat kita buat variabel dengan tipe data tertentu maka secara otomatis akan dibuatkan defaul valuenya
		namun di golang data nil yaitu data kosong
		nil sendiri hanya bisa digunakan di beberapa tipe data seperti interface function map slice pointer dan chanel
	*/
	// example 1
	var person map[string]string = nil
	fmt.Println(person)

	// example 2
	newTest := NewMap("daus")
	fmt.Println(newTest)

}

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}
