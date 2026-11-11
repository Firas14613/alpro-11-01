package main

import "fmt"

func main() {
	var L, R int
	fmt.Scan(&L, &R)

	jumlahValid := 0
	for i := L; i <= R; i++ {
		angka := i
		digitSebelumnya := 10
		jumlahDigit := 0
		tidakMenurun := true

		for angka > 0 {
			digit := angka % 10
			jumlahDigit += digit
			if digit > digitSebelumnya {
				tidakMenurun = false
			}
			digitSebelumnya = digit
			angka /= 10
		}

		prima := jumlahDigit > 1
		for pembagi := 2; pembagi*pembagi <= jumlahDigit; pembagi++ {
			if jumlahDigit%pembagi == 0 {
				prima = false
				break
			}
		}

		if tidakMenurun && prima {
			jumlahValid++
		}
	}

	fmt.Println(jumlahValid)
}
