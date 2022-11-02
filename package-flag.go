package main

import (
	"flag"
	"fmt"
)

func main47() {
	// package flag
	/*
		package flag berisikan fungsionalitas untuk mempersingkat comand line argument
	*/
	host := flag.String("host", "localhost", "put your database host")
	username := flag.String("username", "bachtiar", "put your database username")
	password := flag.String("password", "firdaus", "put your database password")

	flag.Parse()
	fmt.Println(host, username, password)    // karna hasil nya berbentuk pointer wajib menggunakan * ketika memanggilnya
	fmt.Println(*host, *username, *password) // karna hasil nya berbentuk pointer wajib menggunakan * ketika memanggilnya
	// go run package-flag.go -host=localhost -username=Bachtiar -password=firdaus
}
