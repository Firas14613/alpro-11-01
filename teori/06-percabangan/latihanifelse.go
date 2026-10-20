package main

import "fmt"

func main() {
	var nilai int

	fmt.Scanln(&nilai)

	if nilai < 0 || nilai > 100 {
        fmt.Println("Nilai tidak valid")
    } else if nilai < 50 {
        fmt.Println("E")
    } else if nilai < 70 {
        fmt.Println("C")
    } else if nilai < 85 {
        fmt.Println("B")
    } else {
        fmt.Println("A")
    }
}