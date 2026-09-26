# <h1 align="center">Laporan Praktikum Modul [Nomor Modul] - [Judul Modul/Topik]</h1>
<p align="center">Firas Abdurrahman Sandro - 109092600009</p>

## Dasar Teori

### A. Struktur Pemrograman Go
Berdasarkan Modul 2 Bahasa Pemrograman Go, kerangka program yang di tulis dalam bahasa Go selalu mempunyai dua komponen, yaitu `package main` yang merupakan penanda bahwa file ini adalah program utama dan `func main()` yang merupakan kode utama dari sebuah program go.

Ada pula komentar, walau bukan bagian dari kode program, komentar dapat tulis dengan di awali `//` untuk komentar 1 baris, sedangkan jika ingin komentar beberapa baris sekaligus bisa menggunakan `/*` dan diakhiri dengan `*/`.

Contoh program go (contoh `helloworld.go`)
```go
package main
import "fmt"

func main(){
    var salam = "Selamat datang"
    fmt.Println(salam)
}
```

###### Outputnya
```go
Selamat datang
```

### B. Koding, Kompolasi, dan Eksekusi Go

#### 1. Koding
Dalam penulisan program Go terdapat beberapa ketentuan dasar, antara lain:
#### a. Format Peyimpanan Text
Sama seperti bahasa pemrograman yang lain, Go juga harus di tulis menggunakan _text editor_ seperti _VS Code_, _Goland_, atau _Notepad++_. Program Go juga harus disimpan dalam _plain text_ bukan menggunakan format seperti `.pdf`, `.docx` atau lainnya

#### b. Ekstensi File Program
Setiap _file_ Go wajib disimpan dengan ekstensi `.go`, Nama _file_ bebas di isi dengan apa saja tapi disarankan untuk menggunakan nama sesuai isi fungsi dari _file_ tersebut.

#### c. Struktur Folder
Setiap program lengkap Go harus disimpan dalam satu folder Go tersendiri, dan nama folder merupakan nama dari program tersebut. Secara prinsip program Go dapat di pecah menjadi beberapa _file_ `.go`, selama masih disimpan dalam satu folder yang sama.

#### 2. Kompilasi
Beberapa bahasa pemrograman dirancang sebagai _Interprenter_ yang mana akan mengecek kode saat sedang dijalankan, jadi jika ada error di tengah jalan, baru ketahuan saat baris itu sedang dijalankan. Sedangan beberapa bahasa seperti Go dirancang sebeagai _Kompilator_, yang akan mengecek seluruh kode terlebih dahulu baru bisa di jalankan, jika ada bagian yang salah, maka program tidak mau berjalan, kode tersebut harus diubah menjadi _file_ yang bisa dieksekusi.

#### a. Cara Mengkompilasi dan Menjalankan Program Go
1. Buka Command Prompt, PowerShell atau terminal komputer masing masing.
2. Masuk ke directori project Go dengan perintah `cd`, misal `cd D:\TEL-U\AlgoPemro\Golang\alpro-11-01`.
3. Ketik perintah `go build` untuk memprores kodenya atau `go build namafile.go` jika hanya ingin memprores salah satu file program saja.
4. Jika ada error, terminal akan memunculkan pesan error, perbaiki kode error tersebut lalu lakukan `go build` lagi.
5. Jika tidak ada error sama sekali, maka akan tercipta file `.exe`
6. Jalankan file `.exe` lewat terminal.

### Catatan
1. `go build` mengkompilasi program yang ada dalam folder menjadi sebuah program.
2. `go build file.go` mengkompilasi program sumber `file.go` saja
3. `go fmt`membaca semua program dalam folder dan menformat penulisannya agar sesuai dengan standar penulisan program Go.
4. `go clean` membersihkan file file dalam folder hingga hanya menisakan program nya saja.

<!-- Tambahkan poin A, B, C, ... atau sub-topik 1, 2, 3, ... sesuai kebutuhan modul -->

## Guided

### 1. skor.go

```go
package main

import "fmt"

func main() {
	var (
		nama string
		skorMtk int
		skorBhsInggris int
		totalSkor int
		rataRata int
	)

	// nama = "Hong Gil_Dong"
	// skorMtk = 96
	// skorBhsInggris = 82
	// totalSkor = skorMtk + skorBhsInggris
	// rataRata = (totalSkor) / 2
	// fmt.Println(nama)
	// fmt.Println(totalSkor)
	// fmt.Println(rataRata)

	// nama = "Rina Amalia"
	// skorMtk = 88
	// skorBhsInggris = 75
	// totalSkor = skorMtk + skorBhsInggris
	// rataRata = (totalSkor) / 2
	// fmt.Println(nama)
	// fmt.Println(totalSkor)
	// fmt.Println(rataRata)

	fmt.Scan(&nama)
	fmt.Scan(&skorMtk)
	fmt.Scan(&skorBhsInggris)

	totalSkor = skorMtk + skorBhsInggris
	rataRata = totalSkor / 2

	fmt.Println(nama)
	fmt.Println(totalSkor)
	fmt.Println(rataRata)
}
```
#### Deskripsi
Program Go yang menginput `nama`, `skorBhsInggris`, `skorMTK` dan menampilkan `nama`, `totalSkor`, `rataRata`. `nama` disimpan dengan data `String`, skor disimpan dengan int dengan skala skor 1-100. Lalu program menghitung `totalSkore` dari `skorBhsInggris` dan `skorMTK`, dan `rataRata` menghitung `totalSkor` dibagi 2, Dan menampilakan `nama`, `totalSkor`, dan `rataRata`.

##### Input
    Hong_Gil-dong
    96
    82

##### Output
    Hong_Gil-dong
    178
    81

### 2. tukar.go

```go
package main

import "fmt"

func main() {
	var (
		a int
		b int
	)

	fmt.Println(a)
	fmt.Println(b)

	fmt.Scan(&a)
	fmt.Scan(&b)

	a, b = b, a

	fmt.Println(a)
	fmt.Println(b)
}
```
#### Deskripsi
Program yang menukarkan nilai 2 bilangan bulat a dan b menjadi b dan a.

##### Input
	1
	2

##### Output
	2
	1

### 3. lingkaran.go

```go
package main

import "fmt"

func main() {
	const phi float64 = 3.14
	var jari float64
	

	fmt.Scan(&jari)

	luas := phi * jari * jari
	fmt.Println(luas)
}
```
#### Deskripsi
Program yang menghitung luas lingkaran dalam bilangan real. Menggunakan variabel `pi` dengan isi `3.14`, rumus luas `pi * r * r ` dan data disimpan dengan tipe data `float64`.

##### Input
	5

##### Output
	78.5

### 4. suhu.go

```go
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
```
#### Deskripsi
Program yang mengkonversi suhu dejarat Celcius ke Reamur, Fahrenheit, dan Kelvin. Data di simpan dengan bilangan real atau `float64`.

##### Input
	37.5

##### Output
	30 99.5 310.65

<!-- Tambahkan blok file/kode lain sesuai jumlah file pada soal guided -->

## Unguided

### 1. kalkulator.go

```go
package main

import "fmt"

func main() {
	var (
		a int
		b int
	)

	for {
		fmt.Scanln(&a)
		fmt.Scanln(&b)

		if b == 0 {
			fmt.Println("Input Ulang")
			continue
		}

		fmt.Println(a+b, " ", a-b, " ", a*b, " ", a/b, " ", a%b)
		break
	}
	// tambah := a+b
	// kurang := a-b
	// kali := a*b
	// bagi := a/b
	// sisa := a%b

	// if b==0 {
	// 	fmt.Println("Input Ulang")
	// }

	// fmt.Println(a+b)
	// fmt.Println(a-b)
	// fmt.Println(a*b)
	// fmt.Println(a/b)
	// fmt.Println(a%b)
	// fmt.Println(a+b, " ", a-b, " ", a*b, " ", a/b, " ", a%b)

}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](https://github.com/renwxyz/alpro-11-01/blob/main/praktikum/02-bahasa-pemrograman-go/unguided/kalkulator/output.png)


#### Deskripsi
Kalkulator sederhana input bilangan bulat angka A dan B, dan program menghitung tambah, kurang, kali, bagi, dan sisa dari kedua bilangan tersebut. Bilangan B tidak boleh 0.

### 2. cacahuang.go

```go
package main

import "fmt"

func main() {
	var a int

	fmt.Scan(&a)

	lembar1 := a / 10000
	sisa1 := a % 10000

	lembar2 := sisa1 / 5000
	sisa2 := sisa1 % 5000

	lembar3 := sisa2 / 1000

	fmt.Println(lembar1, " ", lembar2, " ", lembar3)

	// pertama menghitung jumlah lembar uang 10000, lalu sisa dari pembagian tersebut disimpan di sisa1
	// kedua menghitung jumlah lembar uang 5000, lalu sisa dari pembagian tersebut disimpan di sisa2
	// ketiga menghitung jumlah lembar uang 1000, lalu sisa dari pembagian tersebut disimpan di sisa3
}
```

##### Output
![Screenshot Output Unguided](unguided/[nama_soal]/output.png)

#### Deskripsi
Program menghitung pecahan uang yang di butuhkan dari nominal rupiah. Menentukan beberapa jumlah lembar pecahan uang sesedikit mungkin dari Rp.10000, Rp.5000, Rp.1000. Menggunakan operasi `/` untuk menentukan jumlah lembar setiap pecahan dan `%` untuk mendapatkan sisa dari pembagian.

<!-- Duplikasi blok "### [nama_soal]" sesuai jumlah folder soal di dalam unguided -->


## Kesimpulan
Berdasarkan praktikum yang telah dilakukan pada Modul 2 ini, dapat disimpulkan bahwa:
1. Program dalam bahasa Go memiliki struktur utama yang wajib ada, yaitu `package main` sebagai penanda program utama dan fungsi `func main()` sebagai blok kode utama yang dieksekusi pertama kali.
2. Go merupakan bahasa yang menggunakan _kompilator_, di mana kode program yang disimpan dalam berkas berekstensi `.go` harus dikompilasi terlebih dahulu menggunakan perintah seperti `go build` menjadi file yang dapat dieksekusi `.exe` tanpa error.
3. Fungsi input dan output dapat dijalankan dengan memanfaatkan package `fmt`, seperti fungsi `fmt.Scan` atau `fmt.Scanln` untuk menerima masukan data dan `fmt.Print` atau `fmt.Println` untuk menampilkan data atau hasil keluaran ke terminal.
4. Penggunaan tipe data dasar `int`, `float64`, `string`, serta operator aritmatika `+`, `-`, `*`, `/`, `%` sangat penting dalam menyelesaikan operasi matematika dan logika, seperti perhitungan nilai rata rata, penukaran nilai variabel, perhitungan luas lingkaran, konversi satuan suhu, kalkulator sederhana, hingga pemecahan nominal uang ke beberapa pecahan lembar terkecil.

## Referensi
1. Materi Modul 2 - Pemrograman Bahasa Go
<!-- Tambahkan nomor referensi berikutnya sesuai kebutuhan -->
