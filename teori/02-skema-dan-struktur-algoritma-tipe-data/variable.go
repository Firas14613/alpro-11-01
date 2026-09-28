package main

import "fmt"

func main() {
	// var name string = "Firas Abdurrahman Sandro"

	// fmt.Println(name)

	// name = "Firas"
	// fmt.Println(name)
	// jika langsung inisialisasi variabel tanpa tipe data, maka tipe data akan otomatis menyesuaikan dengan nilai yang diberikan. jadi tidak perlu menuliskan tipe data secara eksplisit.

	// nama := "Firas Abdurrahman Sandro"
	// fmt.Println(nama)

	// nama = "Firas"
	// fmt.Println(nama)
	// gunakan ":=" untuk cara yang lebih simple, jadi "var" tidak wajib untuk di tulis. ":" Ini deklarasi awal, jadi hanya perlu di ketik sekali saja

	// nama := 17
	// fmt.Println(nama)
	// akan error jika di ubah nilainya

	var (
		firstNama = "Firas"
		lastName  = "Abdurrahman Sandro"
	)

	fmt.Println(firstNama, lastName)
	// variable bisa di tulis langsung banyak seperti ini. jadi tidak perlu mengerik satu persatu
}
