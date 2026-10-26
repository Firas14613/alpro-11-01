package main

import "fmt"

func main() {
	var mobil string
	var A, B, posA, posB int
	var menang = ""

	for i := 1; i <= 10; i++ {
		fmt.Scan(&mobil)

		if mobil == "A" {
			A++
			posA++
			if posA == 4 && menang == "" {
				menang = "A"
			}
		} else if mobil == "B" {
			B++
			posB++
			if posB == 4 && menang == "" {
				menang = "B"
			}
		}
	}
	if A != 5 || B !=5 {
		fmt.Println("tidak valid")
	} else {
		fmt.Println(menang)
	}
}
