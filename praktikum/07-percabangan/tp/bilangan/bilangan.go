package main

import "fmt"

func main() {
	var bil int

	fmt.Scanln(&bil)

	if bil == 0 {
		fmt.Println("Nol")
		return
	}else if bil % 2 == 0 {
		fmt.Print("Genap, ")
	} else {
		fmt.Print("Ganjil, ")
	}

	if bil > 0 {
		fmt.Println("Positif")
	} else if bil < 0 {
		fmt.Println("Negatif")
	} else {
		fmt.Println("Nol")
	}
}