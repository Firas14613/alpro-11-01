package main

import "fmt"

func main () {
	var poin int

	for {
		if poin >= 50 {
			break
		}
		fmt.Print("Masukan Poin: ")
		fmt.Scan(&poin)
	}

	if poin > 200{
		fmt.Println("Gold user")
	} else if poin >= 100 {
		fmt.Println("Platinum user")
	} else {
		fmt.Println("Silver user")
	}
}