package main

import "fmt"

func main() {
	var p int
	var dewasa, kecil, takBerangkat int

	fmt.Scan(&p)

	if p > 0 && p <= 5 {
		dewasa = 1
	} else if p <= 10 {
		dewasa = 2
	} else {
		dewasa = 3
	}

	sisa := p - dewasa * 5

	if sisa > 0 {
		if sisa <= 2 {
			kecil = 1
		} else if sisa <= 4 {
			kecil = 2
		} else if sisa <= 6 {
			kecil = 3
		} else if sisa <= 8 {
			kecil = 4
		} else if sisa <= 10 {
			kecil = 5
		} else {
			kecil = 5
			takBerangkat = sisa - 10
		}
	}

	if dewasa > 0 && kecil > 0 && takBerangkat > 0 {
        fmt.Printf("Dewasa %d, kecil %d, dan %d tak berangkat\n", dewasa, kecil, takBerangkat)
    } else if dewasa > 0 && kecil > 0 {
        fmt.Printf("Dewasa %d, kecil %d\n", dewasa, kecil)
    } else if dewasa > 0 && takBerangkat > 0 {
        fmt.Printf("Dewasa %d, dan %d tak berangkat\n", dewasa, takBerangkat)
    } else if dewasa > 0 {
        fmt.Printf("Dewasa %d\n", dewasa)
    }
}

// Program Avatar
// kamus
//     p, dewasa, kecil, takBerangkat : integer
// algoritma
//     input(p)
//     if p > 0 && p <= 5 then
//         dewasa <- 1
//     else if p <= 10 then
//         dewasa <- 2
//     else
//         dewasa <- 3
//     endif
//     sisa <- p - dewasa * 5
//     if sisa > 0 then
//         if sisa <= 2 then
//             kecil <- 1
//         else if sisa <= 4 then
//             kecil <- 2
//         else if sisa <= 6 then
//             kecil <- 3
//         else if sisa <= 8 then
//             kecil <- 4
//         else if sisa <= 10 then
//             kecil <- 5
//         else
//             kecil <- 5
//             takBerangkat <- sisa - 10
//         endif
//     endif
//     if dewasa > 0 && kecil > 0 && takBerangkat > 0 then
//         print("Dewasa ", dewasa, ", kecil ", kecil, ", dan ", takBerangkat, " tak berangkat")
//     else if dewasa > 0 && kecil > 0 then
//         print("Dewasa ", dewasa, ", kecil ", kecil)
//     else if dewasa > 0 && takBerangkat > 0 then
//         print("Dewasa ", dewasa, ", dan ", takBerangkat, " tak berangkat")
//     else if dewasa > 0 then
//         print("Dewasa ", dewasa)
//     endif
// endprogram