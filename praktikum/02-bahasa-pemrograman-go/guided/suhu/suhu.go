package main

import "fmt"

func main() {
	var celcius float64

	fmt.Scan(&celcius)
	
	reamur := celcius * 4.0 / 5.0
	fmt.Print(reamur, " ")

	fahrenheit := (celcius * 9.0 / 5.0) + 32
	fmt.Print(fahrenheit, " ")

	kelvin := celcius + 273.15
	fmt.Print(kelvin)

}