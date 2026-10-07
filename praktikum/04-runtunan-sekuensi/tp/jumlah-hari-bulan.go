package main

import "fmt"

func main() {
	var tahun int
	var bulan string
	var tahunKabisat bool

	fmt.Scanln(&tahun)
	fmt.Scanln(&bulan)
	tahunKabisat = tahun%4 == 0 && (tahun%100 != 0 || tahun%400 == 0)

	if bulan == "Jan" || bulan == "Mar" || bulan == "Mei" || bulan == "Jul" || bulan == "Agu" || bulan == "Okt" || bulan == "Des" {
		fmt.Println(31)
	} else if bulan == "Apr" || bulan == "Jun" || bulan == "Sep" || bulan == "Nov" {
		fmt.Println(30)
	} else if bulan == "Feb" {
		if tahunKabisat {
			fmt.Println(29)
		} else {
			fmt.Println(28)
		}
	} else {
		fmt.Println("-")
	}
}