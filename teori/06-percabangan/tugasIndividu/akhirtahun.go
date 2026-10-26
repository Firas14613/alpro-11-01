package main

import "fmt"

func main() {
	var totalBelanja, total int
	var kartu, cashback bool
	var diskon = false

	fmt.Scan(&totalBelanja, &kartu)
	
	if totalBelanja >= 100000 {
		diskon = true
		total = totalBelanja - totalBelanja/10
	}
	if kartu == true {
		if totalBelanja >= 200000 {
			cashback = true
			total -= 75000
		} else {
			cashback = false
		}
	}

	fmt.Printf("Kartu? %t\nDiskon? %t\nCashback? %t\nRp. %d", kartu, diskon, cashback, total)
}

// Program akhirTahun
// Kamus
//     totalBelanja, total : integer
//     kartu, cashback, diskon : boolean
// Algoritma
//     input(totalBelanja, kartu)
//     if totalBelanja >= 100000 then
//         diskon <- true
//         total <- totalBelanja - totalBelanja/10
//     endif
//     if kartu then
//         if totalBelanja >= 200000 then
//             cashback <- true
//             total <- total - 75000
//         else
//             cashback <- false
//         endif
//     endif
//     output(Kartu?, Diskon?, Cashback?, total)
// endprogram