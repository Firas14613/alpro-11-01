# PRD: Aplikasi Kuis Edukasi Interaktif

*Product Requirements Document, versi 1.0, Tugas Besar Kuliah (Golang)*

## 1. Ringkasan

Aplikasi kuis edukasi berbasis **CLI** dengan arsitektur **client-server**. Pemain memasukkan username, memilih menu *Mulai Main* atau *Leaderboard*, lalu menjawab **15 soal acak** (pilihan A-D) dengan **timer per soal**. Skor dihitung dari jawaban benar ditambah **bonus kecepatan**. Skor tertinggi tiap username disimpan di database dan ditampilkan di leaderboard. Server menyimpan seluruh data dan logika, sehingga banyak pemain di komputer berbeda bisa bermain bersama lewat jaringan (LAN atau Tailscale).

## 2. Latar Belakang dan Tujuan

**Latar belakang.** Kuis sederhana (tampil soal, jawab, lihat hasil) sudah cukup sebagai dasar, tetapi sebagai tugas besar perlu nilai tambah: database, jaringan, keamanan jawaban, pengujian, dan pengalaman pengguna yang baik.

**Tujuan:**

1. Membangun aplikasi kuis yang berjalan **online** (client-server) dan bisa dimainkan banyak pengguna.
2. Menerapkan praktik Go yang baik: struktur paket rapi, interface, *table-driven test*, penanganan error.
3. Menjamin keadilan permainan: kunci jawaban dan penilaian hanya di server.
4. Merancang sistem agar mudah dikembangkan (misalnya ditambah web di masa depan).

**Bukan tujuan (versi 1.0):** akun dengan password, multiplayer real-time, aplikasi mobile.

## 3. Ruang Lingkup

| Masuk lingkup (v1.0) | Di luar lingkup (v1.0) |
| --- | --- |
| Client CLI dan server REST API | Antarmuka web / mobile |
| Login berbasis username (tanpa password) | Autentikasi password, OAuth |
| 15 soal acak per game dari bank 50-100 soal | Mode multiplayer real-time |
| Timer per soal (waktu dinilai server) | Kategori yang dipilih pemain |
| Skor = poin dasar + bonus kecepatan | Sistem achievement / level |
| Acak urutan opsi jawaban | Chat antar pemain |
| Review jawaban setelah game selesai |  |
| Leaderboard berdasarkan skor tertinggi |  |
| Mode admin (kelola soal) |  |
| Unit test logika inti |  |
| UI CLI yang menarik (Bubble Tea) |  |

## 4. Pengguna

| Peran | Deskripsi | Kebutuhan utama |
| --- | --- | --- |
| **Pemain** | Mahasiswa/teman yang menjalankan client | Bermain mudah, adil, melihat peringkat, belajar dari jawaban salah |
| **Admin** | Pemilik server (pengembang) | Menambah/mengubah/menonaktifkan soal tanpa menyentuh database langsung |
| **Penilai** | Dosen/asprak | Menjalankan program dengan mudah, melihat kode terstruktur dan teruji |

## 5. Kebutuhan Fungsional

### 5.1 Login

| ID | Kebutuhan |
| --- | --- |
| FR-01 | Pemain memasukkan username saat client dibuka. |
| FR-02 | Username divalidasi: 3-20 karakter, hanya huruf, angka, dan underscore. Input tidak valid meminta input ulang. |
| FR-03 | Jika username belum ada, server membuatnya. Jika sudah ada, server memakai data lama (skor tertinggi tetap). |
| FR-04 | Username bersifat unik dan tidak membedakan huruf besar/kecil. |

### 5.2 Menu Utama

| ID | Kebutuhan |
| --- | --- |
| FR-05 | Menu berisi: **Mulai Main**, **Leaderboard**, **Keluar**. |
| FR-06 | Menu menampilkan skor tertinggi pemain saat ini. |

### 5.3 Permainan

| ID | Kebutuhan |
| --- | --- |
| FR-07 | Setiap game terdiri dari **tepat 15 soal** yang dipilih acak dari soal berstatus aktif. |
| FR-08 | Soal ditampilkan **satu per satu** dengan format pilihan A, B, C, D. |
| FR-09 | Jika pemain pernah bermain, server **memprioritaskan soal yang belum pernah dilihat** pemain tersebut. Jika soal baru kurang dari 15, sisanya diisi dari soal lama. |
| FR-10 | Urutan opsi A-D **diacak oleh server** per soal per game. Server menyimpan urutan yang ditampilkan. |
| FR-11 | Input selain A-D (termasuk kosong atau lebih dari satu huruf) ditolak dan pemain diminta mengulang. Huruf kecil diterima dan dinormalisasi. |
| FR-12 | Setiap soal punya batas waktu (default 15 detik). Hitung mundur tampil di client. |
| FR-13 | Jika waktu habis tanpa jawaban, soal dihitung salah (0 poin) dan permainan lanjut ke soal berikutnya. |
| FR-14 | Kunci jawaban **tidak pernah dikirim** ke client selama game berlangsung. |
| FR-15 | Satu soal hanya bisa dijawab satu kali. Jawaban kedua untuk soal yang sama ditolak. |

### 5.4 Hasil dan Review

| ID | Kebutuhan |
| --- | --- |
| FR-16 | Setelah soal ke-15, tampil ringkasan: jumlah benar, jumlah salah, **total poin**, akurasi (%), dan **skor tertinggi** pemain. |
| FR-17 | Skor tertinggi hanya berubah jika skor game baru **lebih besar** (90 lalu 80 tetap 90). |
| FR-18 | Pemain dapat membuka **review**: tiap soal berisi jawaban pemain, jawaban benar, dan penjelasan. |
| FR-19 | Review hanya tersedia untuk game berstatus *selesai*. |

### 5.5 Leaderboard

| ID | Kebutuhan |
| --- | --- |
| FR-20 | Menampilkan semua pemain yang pernah bermain beserta skor tertingginya, urut dari yang terbesar. |
| FR-21 | Jika skor sama, urutan ditentukan oleh waktu pencapaian skor tersebut (yang lebih dulu menang). |
| FR-22 | Baris milik pemain saat ini ditandai. |
| FR-23 | Pemain yang belum pernah menyelesaikan game tidak muncul. |

### 5.6 Mode Admin

| ID | Kebutuhan |
| --- | --- |
| FR-24 | Endpoint admin hanya dapat diakses dengan header `X-Admin-Token` yang cocok dengan `ADMIN_TOKEN` di environment server. |
| FR-25 | Admin dapat melihat, menambah, mengubah, dan menonaktifkan soal. |
| FR-26 | Penonaktifan memakai *soft delete* (`is_active = false`), sehingga riwayat game lama dan review tetap utuh. |
| FR-27 | Validasi soal: teks tidak kosong, 4 opsi terisi, jawaban benar A-D, batas waktu positif. |
| FR-28 | (Opsional) Impor soal massal dari file JSON. |

## 6. Aturan Bisnis: Skor

```
poin_dasar = 10          (hanya jika benar)
bonus      = round(5 x sisa_waktu / batas_waktu)
poin_soal  = poin_dasar + bonus      (benar: 10-15, salah/timeout: 0)
```

- Skor maksimal per game: **15 soal x 15 poin = 225**.
- Waktu dihitung **server**: selisih waktu soal dikirim dan jawaban diterima.
- Toleransi jaringan: **1 detik** di atas batas waktu masih dianggap valid (dapat diatur lewat konfigurasi).
- Semua perhitungan memakai integer.
- Fungsi inti: `CalculatePoints(correct bool, elapsed, limit time.Duration) int`, murni dan mudah dites.

## 7. Kebutuhan Non-Fungsional

| Kategori | Kebutuhan |
| --- | --- |
| **Keamanan** | Kunci jawaban tidak dikirim saat game berjalan. Admin dilindungi token dari environment variable. Query database memakai *prepared statement*. Validasi semua input di server. |
| **Keadilan** | Waktu, nilai, dan pemilihan soal hanya di server. Client tidak dipercaya. |
| **Performa** | Respons API normal di bawah 200 ms pada jaringan lokal. Mampu melayani minimal 20 pemain bersamaan. |
| **Keandalan** | Game yang terputus tidak merusak data. Game *ongoing* yang terlalu lama (misalnya lebih dari 30 menit) ditandai kedaluwarsa dan tidak mengubah skor tertinggi. |
| **Portabilitas** | Client dan server dapat di-compile menjadi satu file executable untuk Windows, Linux, dan macOS. SQLite memakai driver pure Go (tanpa CGO). |
| **Kemudahan pakai** | Pesan error berbahasa Indonesia yang jelas. Alamat server diatur lewat flag `--server` atau variabel environment. |
| **Pemeliharaan** | Kode berlapis (handler, service, repository). Akses database lewat interface agar mudah diganti. |
| **Kualitas** | Cakupan test pada logika inti minimal 70%. |

## 8. Arsitektur dan Teknologi

```
[CLI Client] --HTTP/JSON--> [Server Go] ---> [SQLite]
[CLI Client] --HTTP/JSON-->      ^
(banyak pemain)                  |
                     logika, penilaian, timer, data
```

| Komponen | Pilihan | Catatan |
| --- | --- | --- |
| Bahasa | Go (versi stabil terbaru) |  |
| HTTP server | `net/http` standar (routing Go 1.22+) | Bisa diganti Gin/Echo jika diizinkan |
| Database | SQLite via `modernc.org/sqlite` | Diakses lewat interface repository, bisa diganti PostgreSQL |
| TUI | Bubble Tea, Lip Gloss, Bubbles | Dikerjakan belakangan, pisahkan dari logika |
| Testing | paket `testing`, `httptest`, SQLite `:memory:` | Gaya table-driven |
| Deployment | Server lokal, lalu server Tailscale | Opsional: Dockerfile |

**Konfigurasi server (environment):** `PORT`, `DB_PATH`, `ADMIN_TOKEN`, `QUESTION_TIME_LIMIT`, `TIME_TOLERANCE_MS`.

**Konfigurasi client:** flag `--server`, misalnya `./kuis --server http://100.x.y.z:8080`.

## 9. Model Data

**users**

| Kolom | Tipe | Keterangan |
| --- | --- | --- |
| id | INTEGER PK |  |
| username | TEXT UNIQUE | disimpan huruf kecil |
| high\_score | INTEGER | poin tertinggi, default 0 |
| high\_score\_at | DATETIME | waktu rekor tercapai (untuk tie-break) |
| total\_games | INTEGER |  |
| created\_at | DATETIME |  |

**questions**

| Kolom | Tipe | Keterangan |
| --- | --- | --- |
| id | INTEGER PK |  |
| question\_text | TEXT |  |
| option\_a ... option\_d | TEXT |  |
| correct\_answer | TEXT | A/B/C/D (posisi asli) |
| category | TEXT |  |
| difficulty | TEXT | mudah/sedang/sulit |
| time\_limit\_sec | INTEGER | default 15 |
| explanation | TEXT | untuk review |
| is\_active | BOOLEAN | soft delete |

**games**

| Kolom | Tipe | Keterangan |
| --- | --- | --- |
| id | INTEGER PK |  |
| user\_id | INTEGER FK |  |
| status | TEXT | ongoing / finished / expired |
| score | INTEGER | total poin |
| correct\_count | INTEGER |  |
| total\_questions | INTEGER | selalu 15 |
| started\_at, finished\_at | DATETIME |  |

**game\_questions**

| Kolom | Tipe | Keterangan |
| --- | --- | --- |
| id | INTEGER PK |  |
| game\_id, question\_id | INTEGER FK |  |
| position | INTEGER | urutan 1-15 |
| shown\_order | TEXT | mis. "CADB": isi tiap slot A-D menunjuk opsi asli |
| sent\_at | DATETIME | waktu soal dikirim server |
| chosen\_answer | TEXT | huruf yang dipilih pemain (nullable) |
| is\_correct | BOOLEAN |  |
| time\_taken\_ms | INTEGER |  |
| points | INTEGER |  |

## 10. Spesifikasi API

Semua request/response berformat JSON.

| Method | Endpoint | Fungsi | Auth |
| --- | --- | --- | --- |
| POST | `/login` | Buat/ambil user dari username | - |
| POST | `/games` | Mulai game, server memilih 15 soal, mengirim soal #1 | username/user\_id |
| POST | `/games/{id}/answer` | Kirim jawaban soal aktif, terima soal berikutnya atau hasil akhir | user\_id |
| GET | `/games/{id}/review` | Review lengkap (hanya game selesai) | user\_id |
| GET | `/leaderboard` | Daftar skor tertinggi | - |
| GET | `/admin/questions` | Daftar soal | `X-Admin-Token` |
| POST | `/admin/questions` | Tambah soal | `X-Admin-Token` |
| PUT | `/admin/questions/{id}` | Ubah soal | `X-Admin-Token` |
| DELETE | `/admin/questions/{id}` | Nonaktifkan soal | `X-Admin-Token` |

**Contoh respons soal (tanpa kunci jawaban):**

```json
{
  "game_id": 12,
  "position": 3,
  "total": 15,
  "question": "Apa ibu kota Australia?",
  "options": {"A": "Sydney", "B": "Canberra", "C": "Melbourne", "D": "Perth"},
  "time_limit_sec": 15
}
```

**Kode error utama:** 400 (input tidak valid), 401 (token admin salah), 404 (game/soal tidak ada), 409 (soal sudah dijawab / game sudah selesai), 500 (kesalahan server).

## 11. Alur Pengguna (CLI)

1. Buka client, masukkan username.
2. Menu utama: Mulai Main / Leaderboard / Keluar.
3. **Mulai Main:** soal 1 sampai 15 tampil bergantian dengan hitung mundur. Pemain mengetik A-D, input salah diulang, waktu habis lanjut otomatis.
4. Layar hasil: benar, salah, poin, akurasi, skor tertinggi.
5. Pilihan: lihat review, main lagi, atau kembali ke menu.
6. **Leaderboard:** tabel peringkat, tombol kembali.

**Contoh tampilan soal:**

```
╭──────────────────────────────────────╮
│  Soal 3/15              ⏱  09 detik  │
│                                      │
│  Apa ibu kota Australia?             │
│                                      │
│   A. Sydney                          │
│   B. Canberra                        │
│   C. Melbourne                       │
│   D. Perth                           │
╰──────────────────────────────────────╯
 Jawabanmu (A-D): _
```

**Kondisi khusus:** server tidak terjangkau (tampilkan pesan dan opsi coba lagi), koneksi putus saat game (game tidak mengubah skor tertinggi dan dapat kedaluwarsa), pemain menutup client di tengah game.

## 12. Strategi Pengujian

| Area | Kasus uji |
| --- | --- |
| `CalculatePoints` | Benar cepat (15), benar di batas (10), salah (0), timeout, dalam toleransi 1 detik |
| Pemilihan soal | Tepat 15, tanpa duplikat, soal belum pernah dilihat didahulukan, aman jika bank kurang dari 15 |
| Acak opsi | Pemetaan huruf ke jawaban asli selalu benar setelah diacak |
| Skor tertinggi | 90 lalu 80 tetap 90; 70 lalu 95 jadi 95; skor sama dengan rekor tidak mengubah waktu rekor |
| Timer (fake clock) | Tepat waktu, telat, di batas toleransi |
| Validasi input | `a`, `B`, `E`, kosong, `AB`, spasi |
| Handler HTTP | Alur login-game-jawab-review; admin tanpa token ditolak (401); jawab dua kali ditolak (409) |
| Repository | SQLite `:memory:`: leaderboard terurut, tie-break, soft delete |

Pengujian manual: dua laptop di satu jaringan, memainkan game bersamaan.

## 13. Struktur Proyek

```
kuis-edukasi/
├── cmd/
│   ├── server/main.go
│   └── client/main.go
├── internal/
│   ├── model/        # User, Question, Game
│   ├── repository/   # interface + implementasi SQLite
│   ├── service/      # pemilihan soal, skor, aturan game
│   ├── handler/      # HTTP handler + middleware admin
│   ├── clock/        # interface Clock (real & fake)
│   └── ui/           # tampilan CLI
├── data/questions.json   # seed 50-100 soal
├── go.mod
├── .gitignore            # *.db, .env
└── README.md
```

## 14. Rencana Pengerjaan

| Tahap | Pekerjaan | Hasil |
| --- | --- | --- |
| 1 | Logika inti: `CalculatePoints`, pemilihan soal, skor tertinggi + unit test | Logika terbukti benar |
| 2 | SQLite, repository, seed soal dari JSON | Data tersimpan |
| 3 | Server REST, alur per soal, timer server-side (tes di localhost) | API berjalan |
| 4 | Client CLI polos (input A-D, timer) | Game bisa dimainkan |
| 5 | Acak opsi + review jawaban | Fitur edukasi lengkap |
| 6 | Mode admin | Soal bisa dikelola |
| 7 | UI Bubble Tea | Tampilan menarik |
| 8 | Deploy ke server Tailscale, tes multi-laptop, README, laporan | Siap dikumpulkan |

Setiap tahap menghasilkan sesuatu yang berjalan, jadi jika waktu mepet, versi terakhir yang stabil tetap bisa dikumpulkan.

## 15. Risiko dan Mitigasi

| Risiko | Dampak | Mitigasi |
| --- | --- | --- |
| Menulis 50-100 soal beserta penjelasan memakan waktu | Tahap 5 terlambat | Mulai dari 30 soal, tambah bertahap; boleh minta bantuan teman/AI untuk draf lalu diverifikasi |
| Bubble Tea sulit dipelajari | UI terlambat | Pisahkan layer UI, selesaikan CLI polos dulu |
| Latency Tailscale merugikan pemain | Skor tidak adil | Toleransi waktu dapat dikonfigurasi |
| Dosen mensyaratkan library tertentu | Perlu refactor | Pakai `net/http` standar dan interface repository |
| Penilai tidak bisa join Tailscale | Demo gagal | Siapkan cadangan: server lokal + hotspot, atau Tailscale Funnel |
| Pemain curang lewat manipulasi request | Leaderboard tidak valid | Semua penilaian dan waktu di server, kunci jawaban tidak dikirim |

## 16. Kriteria Penerimaan

- [ ] Pemain dapat login dengan username dan username lama mempertahankan skor.
- [ ] Setiap game berisi tepat 15 soal acak; dua game berturut-turut menghasilkan set soal berbeda.
- [ ] Input selain A-D ditolak dan diminta ulang.
- [ ] Timer berjalan; jawaban lewat batas waktu dinilai salah.
- [ ] Poin dihitung dengan bonus kecepatan sesuai rumus.
- [ ] Skor tertinggi tidak turun saat skor game baru lebih kecil.
- [ ] Leaderboard menampilkan semua pemain terurut dengan benar.
- [ ] Review menampilkan jawaban benar dan penjelasan setelah game selesai.
- [ ] Kunci jawaban tidak terlihat di respons API selama game berjalan.
- [ ] Endpoint admin menolak request tanpa token yang valid.
- [ ] Minimal dua klien dari laptop berbeda dapat bermain ke satu server.
- [ ] Unit test logika inti lulus semua.
- [ ] README menjelaskan cara build, menjalankan server, dan menjalankan client.

## 17. Pengembangan ke Depan

- Klien **web** yang memakai API yang sama (CLI tidak perlu diubah).
- Pilihan kategori dan tingkat kesulitan.
- Leaderboard per kategori atau mingguan.
- Autentikasi password atau token sesi.
- Mode duel/multiplayer real-time (WebSocket).
- Migrasi ke PostgreSQL dan deployment dengan Docker.
- Statistik pribadi: akurasi per kategori, grafik riwayat.

## 18. Pertanyaan Terbuka

1. Apakah dosen/asprak membatasi library atau framework (misalnya harus `net/http` standar)?
2. Apakah antarmuka web diperbolehkan untuk tugas ini?
3. Apakah penilaian mensyaratkan database tertentu (SQLite, MySQL, atau PostgreSQL)?
4. Topik/kategori apa yang dipakai untuk bank soal (umum, informatika, atau mata kuliah tertentu)?
5. Bagaimana format demo ke dosen: LAN kelas, Tailscale, atau hanya satu laptop?
