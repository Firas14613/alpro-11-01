package main

import "fmt"

func main() {
	var lamp = "kuning"

	if lamp == "merah" {
		fmt.Println("Berhenti")
	} else if lamp == "kuning" {
		fmt.Println("Hati-hati/Pelan pelan")
	} else if lamp == "hijau" {
		fmt.Println("jalan")
	} else {
		fmt.Println("Lamput tidak terdaftar")
	}
}