package main

import (
	"fmt"
	"regexp"
)

func main56() {
	// package regexp
	/*
		package regexp adalah utilitas di go-lang untuk melakukan pencarian regular expression
		regular expression di golang menggunkan library c yang di buat google bernama re2
	*/
	// regexp.MustCompile(string) membuat Regexp
	// regexp.MatchString(string) bool mengecek apakah regexp match dengan string
	// regexp.FindAllString(string, max) mencari string yang match dengan maximum jumlah hasil
	var regex = regexp.MustCompile(`d([a-z])([a-z])s`)
	fmt.Println(regex.MatchString("daus"))
	fmt.Println("===================================")
	fmt.Println(regex.FindAllString("daus dias dais sisi lais", 2))
	fmt.Println("===================================")
	fmt.Println(regex.FindAllString("daus dias dais sisi lais", 1))
	fmt.Println("===================================")
	fmt.Println(regex.FindAllString("daus dias dais sisi lais", -1))

}
