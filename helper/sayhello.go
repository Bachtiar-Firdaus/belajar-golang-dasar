package helper

// ini bisa di akses dari luar package
func SayHello(name string) interface{} {
	return "hello " + name
}

// ini tidak bisa diakses dari luar package
func sayGoodBye(name string) interface{} {
	return "hello " + name
}
