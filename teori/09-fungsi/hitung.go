package main

import "fmt"

func hitung(x, y int) int {
	return 2 * x + y - 5
}

func ambangAtas() int {
	return 60
}

func lulus(nilai int) bool {
	return nilai >= ambangAtas()
}

func statusLulus(nilai int) string {
	if lulus(nilai) {
		return "Lulus"
	} else {
		return "Tidak Lulus"
	}
}

func main() {
	fmt.Println(lulus(50))
	fmt.Println(statusLulus(80))
	fmt.Println(ambangAtas())
	hasil := hitung(10, 3)
	fmt.Println(hasil)
	fmt.Println(hasil * 2)
	fmt.Println(hitung(2, 1))
}