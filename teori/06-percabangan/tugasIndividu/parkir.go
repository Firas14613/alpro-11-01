package main

import "fmt"

func main() {
	var h1, m1, h2, m2 int

	fmt.Scan(&h1, &m1, &h2, &m2)

	var H1 int
	if h1 >= 7 && h1 <= 11 {
		H1 = h1
	} else if h1 == 12 {
		H1 = 12
	} else if h1 >= 1 && h1 <= 5 {
		H1 = h1 + 12
	} else {
		fmt.Println("Jam tidak terdaftar")
		return
	}

	var H2 int
	if h2 >= 7 && h2 <= 11 {
		H2 = h2
	} else if h2 == 12 {
		H2 = 12
	} else if h2 >= 1 && h2 <= 5 {
		H2 = h2 + 12
	} else {
		fmt.Println("Jam tidak terdaftar")
		return
	}

	totalMenit1 := (H1 * 60) + m1
	totalMenit2 := (H2 * 60) + m2

	selisihMenit := totalMenit2 - totalMenit1

	hh := selisihMenit / 60
	mm := selisihMenit % 60

	fmt.Printf("%d jam %d menit", hh, mm)
}


// program parkir
// kamus
//     h1, m1, h2, m2 : integer
//     H1, H2 : integer
//     totalMenit1, totalMenit2, selisihMenit : integer
//     hh, mm : integer
// algoritma
//     input(h1, m1, h2, m2)

//     if h1 >= 7 && h1 <= 11 then
//         H1 <- h1
//     else if h1 == 12 then
//         H1 <- 12
//     else if h1 >= 1 && h1 <= 5 then
//         H1 <- h1 + 12
//     else
//         print("Jam tidak terdaftar")
//         stop
//     endif

//     if h2 >= 7 && h2 <= 11 then
//         H2 <- h2
//     else if h2 == 12 then
//         H2 <- 12
//     else if h2 >= 1 && h2 <= 5 then
//         H2 <- h2 + 12
//     else
//         print("Jam tidak terdaftar")
//         stop
//     endif

//     totalMenit1 <- (H1 * 60) + m1
//     totalMenit2 <- (H2 * 60) + m2

//     selisihMenit <- totalMenit2 - totalMenit1

//     hh <- selisihMenit / 60
//     mm <- selisihMenit % 60

//     print(hh, " jam ", mm, " menit")
// endprogram