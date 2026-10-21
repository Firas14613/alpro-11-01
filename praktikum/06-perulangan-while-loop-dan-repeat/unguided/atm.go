package main

import "fmt"

func main() {
	var saldo int =  10000000
	var login bool = false
	var pass, i int

	for i != 3 && login == false {
		fmt.Println("Masukkan Password ATM: ")
		fmt.Scan(&pass)

		if pass == 1234 {
			login = true
		} else {
			fmt.Println("Password Salah. Silahkan coba lagi.")
			i++
		}
	}

	if login == false {
		fmt.Println("ATM terblokir. Program Selesai.")
		return
	}

	fmt.Println("Login Berhasil!")
	
	pilihan := 0
	for pilihan != 3 {
		fmt.Println("1. Lihat Saldo")
		fmt.Println("2. Tarik Tunai")
		fmt.Println("3. Keluar")
		fmt.Print("Masukkan pilihan: ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			fmt.Printf("Saldo Anda: Rp %d\n", saldo)
		} else if pilihan == 2 {
			var tarik int
			fmt.Print("Masukkan jumlah tarik tunai: Rp ")
			fmt.Scan(&tarik)
			if tarik > saldo {
				fmt.Println("Saldo tidak cukup.")
			} else {
				fmt.Println("Anda berhasil menarik uang sebesar: Rp ", tarik)
				saldo -= tarik
				fmt.Println("Saldo Tersisa: Rp ", saldo)
			}
		} else if pilihan == 3 {
			fmt.Println("Keluar dari sistem. Terima kasih!")
		}
	}

}