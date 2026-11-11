package main

import "fmt"

func main() {
	var input string
	var M, K, T int

	for  {
		fmt.Scan(&input)

		if input == "M" {
			M++
		} else if input == "K" {
			K++
		} else if input == "T" {
			T++
		} else if input != "M" && input != "K" && input != "T" {
			break
		}
	}
	fmt.Printf("M: %d\nK: %d\nT: %d\n", M, K, T)
}