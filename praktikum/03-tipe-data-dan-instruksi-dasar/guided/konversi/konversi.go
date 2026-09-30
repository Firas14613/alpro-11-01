package main

import "fmt"

func main() {
	var celcius float64

	fmt.Print("Masukan suhu: ")
	fmt.Scan(&celcius)

	k := celcius + 273

	fmt.Println("Kelvin: ", k)
}