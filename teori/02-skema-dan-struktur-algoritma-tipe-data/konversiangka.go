package main

import "fmt"

func main() {
	var angka int

	for {
		fmt.Scanln(&angka)

		if angka > 999 {
			fmt.Println("Angka tidak boleh lebih dari 3 digit")
			continue
		}
		var digit1 = angka / 100
		var digit2 = (angka / 10) % 10
		var digit3 = angka % 10

		fmt.Println(digit1, " ", digit2, " ", digit3)
		break
	}
}

// program 3 digit
// kamus
//   angka : integer
//   digit1, digit2, digit3 : integer
// algoritma
//   ulang
//       input(angka)
//       jika angka > 999 maka
//           output("Angka tidak boleh lebih dari 3 digit")
//       jika tidak
//           digit1 <- angka / 100
//           digit2 <- (angka / 10) % 10
//           digit3 <- angka mod 10
//           output(digit1, " ", digit2, " ", digit3)
//           break (atau: keluar pengulangan)
//       selesai jika
//   selesai ulang
// endprogram
