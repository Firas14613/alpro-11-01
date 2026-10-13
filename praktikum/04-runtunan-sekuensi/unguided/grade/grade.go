package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var nilai float64
	var nama string
	var grade string

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	nama = scanner.Text()
	fmt.Scanln(&nilai)

	switch {
	case nilai >= 90 && nilai <= 100:
		grade = "A"
	case nilai >= 80 && nilai < 90:
		grade = "B"
	case nilai >= 70 && nilai < 80:
		grade = "C"
	case nilai >= 60 && nilai < 70:
		grade = "D"
	default:
		grade = "E"
	}
	fmt.Printf("%s mendapatkan nilai %s\n", nama, grade)
}
