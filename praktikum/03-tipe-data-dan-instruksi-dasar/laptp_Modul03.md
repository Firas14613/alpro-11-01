# <h1 align="center">Tugas Pendahuluan Modul 3 - Variable dan Operator</h1>
<p align="center">Firas Abdurrahman Sandro - 109092600009</p>

### 1. Sisa Kue

```go
package main

import "fmt"

func main() {
	var a, b int

	fmt.Scan(&a, &b)

	fmt.Println(a % b)
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](https://github.com/Firas14613/alpro-11-01/blob/main/praktikum/03-tipe-data-dan-instruksi-dasar/tp/sisa/output.png)


#### Deskripsi
Program untuk menghitung berapa banyak sisa kue setelah dibagi kepada seluruh anggota keluarga. Dengan input `x` untuk jumlah keluarga dan `y` untuk jumlah kue.

### 2. bool.go

```go
package main

import "fmt"

func main() {
	var status bool
	fmt.Scan(&status)

	fmt.Println(status)
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](https://github.com/Firas14613/alpro-11-01/blob/main/praktikum/03-tipe-data-dan-instruksi-dasar/tp/bool/output.png)


#### Deskripsi
Program untuk membaca nilai boolean dari inputan user. Jika inputan user `true` maka output menjadi `true`, tapi jika selain `true` output akan menjadi `false`.

### 3. konversi.go

```go
package main

import "fmt"

func main() {
	var mil float64
	const kilo float64 = 1.6

	fmt.Scan(&mil)
	fmt.Printf("%.1f", mil * kilo)
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](https://github.com/Firas14613/alpro-11-01/blob/main/praktikum/03-tipe-data-dan-instruksi-dasar/tp/konversi/output.png)


#### Deskripsi
Program untuk mengkonversi nilai jarak `mil` ke `kilometer` dengan operasi `mil * kilometer` dan menggunakan `%.1f` agar hasil desimal di belakang kome hanya 1 angka.

## Kesimpulan
Kami belajar menggunakan variabel dan operator dalam pemrograman Go untuk menyimpan data dan melakukan proses perhitungan. Kita juga belajar cara membaca input, menggunakan operator modulo, boolean, dan konversi satuan, serta menampilkan hasil dengan format yang tepat. Jadi, praktikum ini membantu kami memahami dasar-dasar penggunaan variabel, tipe data, dan operator dalam program Go.