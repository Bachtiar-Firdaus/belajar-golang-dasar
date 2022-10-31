package main

import "fmt"

func main30() {
	// defer
	// defer function adalah function yang bisa kita jadwalkan untuk dieksekusi setelah function selesai di eskalasi
	// defer function akan selalu di eksekusi walaupun terjadi error di function yang di eksekusi
	// noted defer digunakan untuk merunning paksa semua logic meskipun ada salah satu code yang error

	// example defer
	runApplication(10)

	// panic
	// panic function adalah function yang bisa kita gunakan untuk menghentikan program
	// panic function biasanya dipanggil ketika terjadi error pada saat program kita sedang berjalan
	// saat panic function dipanggil, program akan terhenti, namun defer funtion tetap akan di eksekusi
	// exmaple panic
	runApp(true)
	fmt.Println("Lanjut aplikasi")
	// example recover setelah panic berhenti
	// recover waib di masukan didalam defer

	// recover
	// recover adalah funtion yang kita bisa gunakan untuk menangkap data panic
	// dengan recover proses panic akan terhenti sehingga program akan tetap berjalan

}

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication(value int) {
	defer logging()
	fmt.Println("Run Aplication")
	// example code error
	result := 10 / value
	fmt.Println("Result : ", result)
}

func runApp(error bool) {
	// example defer
	defer endApp()
	if error {
		// example panic
		panic("ERROR")
	}

}

func endApp() {
	// example recover
	// setelah menagkap error panic maka panic berhenti dan melanjutkan proses selanjutnya
	message := recover()
	if message != nil {
		fmt.Println("Tearjadi Error", message)
	}
	fmt.Println("End App")
}
