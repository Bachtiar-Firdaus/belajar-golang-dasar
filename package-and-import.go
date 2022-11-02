package main

import (
	"belajar-golang-dasar/helper"
	"belajar-golang-dasar/other"
	"fmt"
)

func main43() {
	// Error Package dan Importt
	/*
	   Di Go-Lang versi terbaru, kita sudah diwajibkan menggunakan Go-Modules, yang akan dijelaskan pada PART 2 kelas ini
	   Jadi pada video selanjutnya, pasti akan mendapat error
	   Namun jika ingin mencoba fitur lama Go-Lang, sebelum lanjutkan video selanjutnya, silahkan matikan default fitur Go-Modules dengan menggunakan perintah
	   go env -w  GO111MODULE=off
	*/

	// package
	/*
		package adalah tempat yang bisa gigunakan untuk mengorganisasi kan kode progrma yang kita buat
		dengan menggunakan package kita bisa merapikan kode program yang kita buat
		package sendiri sebenernya hanya direktory folder di sistem operasi kita
	*/

	// import
	/*
		secara standar file golang hanya bisa mengakses file golang lainnya yang berada dalam package yang sama
		jika kita mengakses file golang yang berada diluar package maka kita bisa menggunakan import
	*/
	// example 1
	name := "daus"
	result1 := helper.SayHello(name)
	result2 := other.SayHelloV2(name)
	fmt.Println(result1)
	fmt.Println(result2)

}
