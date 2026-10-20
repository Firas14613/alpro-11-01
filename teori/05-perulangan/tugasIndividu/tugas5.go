package main

import "fmt"

func main () {
	var k, barangMasuk, totalBarang int
	fmt.Scanln(&k)

	for i := 1; i <= 5; i++ {
		fmt.Scanln(&barangMasuk)
		totalBarang += barangMasuk
	
		fmt.Println(totalBarang >= k)
	}
}

// program kapasitasGudang
// kamus
// 	k, barangMasuk, totalBarang : integer
// algoritma
// 	input(k)

// 	for i <- 1 to 5 do
// 		input(barangMasuk)
// 		totalBarang <- totalBarang + barangMasuk

// 		output(totalBarang >= k)
// 	endfor
// endprogram