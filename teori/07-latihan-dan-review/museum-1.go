package main

import "fmt"

func main() {
	var pengunjung, totalPengunjung, jumlahHari, hariSebelumnya, kenaikan int
	var rataRata float64
	
	for {
		fmt.Scan(&pengunjung)

		if pengunjung < 0 || pengunjung > 200{
			break
		}
		totalPengunjung += pengunjung
		jumlahHari++

		if hariSebelumnya != 0 {
			if pengunjung > hariSebelumnya{
				kenaikan++
			}
		}

		hariSebelumnya = pengunjung
	}
	if jumlahHari > 0 {
		rataRata = float64(totalPengunjung) / float64(jumlahHari)
	}

	fmt.Printf("%d %.2f", kenaikan, rataRata)
}