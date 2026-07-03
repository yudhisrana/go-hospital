# Go Hospital

`Go Hospital` adalah API web service untuk manajemen klinik / rumah sakit yang membantu mengatur alur pelayanan pasien dari pendaftaran hingga pengambilan obat.

Proyek ini dirancang agar proses pelayanan menjadi lebih rapi, terstruktur, dan mudah dipantau oleh resepsionis, dokter, apoteker, dan admin.

---

## Tujuan Sistem

Sistem ini dibuat untuk menangani kebutuhan utama berikut:

- Pendaftaran pasien
- Pengelolaan antrean pemeriksaan dokter
- Pencatatan pemeriksaan medis
- Pembuatan resep obat digital
- Pengelolaan antrean pengambilan obat
- Pengelolaan data master rumah sakit / klinik
- Monitoring status pelayanan pasien

---

## Alur Pelayanan Pasien

Berikut alur kerja utama pada `Go Hospital`:

1. Pasien datang ke klinik / rumah sakit.
2. Pasien daftar ke resepsionis.
3. Resepsionis mengisi data pasien dan keluhan / gejala sakit.
4. Pasien mendapatkan nomor antrean pemeriksaan dokter.
5. Pasien menunggu sampai dipanggil dokter.
6. Dokter memeriksa pasien di ruangan pemeriksaan.
7. Setelah pemeriksaan selesai, dokter menulis diagnosis dan resep obat.
8. Resep dikirim ke bagian farmasi / apotek.
9. Pasien keluar dari ruang dokter dan mendapatkan nomor antrean pengambilan obat.
10. Apoteker menyiapkan obat sesuai resep.
11. Pasien dipanggil untuk mengambil obat.
12. Setelah obat diterima, pasien pulang dan proses selesai.

### Flow Singkat

```text
Pasien datang
   -> daftar ke resepsionis
   -> isi data & gejala sakit
   -> dapat nomor antrean dokter
   -> menunggu dipanggil dokter
   -> diperiksa dokter
   -> dokter menulis resep obat
   -> masuk antrean apotek
   -> ambil obat
   -> pulang / selesai
```

---

## Fitur yang Dibutuhkan

### 1. Modul Pendaftaran Pasien

- Tambah data pasien baru
- Cari pasien lama berdasarkan NIK / nama / nomor telepon
- Simpan data keluhan / gejala sakit
- Buat kunjungan pasien
- Generate nomor antrean otomatis

### 2. Modul Antrean Dokter

- Menampilkan daftar pasien yang menunggu
- Memanggil pasien berdasarkan nomor antrean
- Ubah status antrean: menunggu, dipanggil, diperiksa, selesai
- Filter antrean berdasarkan dokter / poli / tanggal

### 3. Modul Pemeriksaan Dokter

- Input hasil pemeriksaan
- Input diagnosis pasien
- Input catatan medis
- Simpan rekam medis pasien
- Buat resep obat digital
- Hubungkan resep dengan kunjungan pasien

### 4. Modul Apotek / Farmasi

- Menampilkan antrean resep obat
- Menyiapkan obat sesuai resep
- Mengurangi stok obat saat obat keluar
- Menandai resep sudah diserahkan
- Generate nomor antrean pengambilan obat

### 5. Modul Manajemen Obat

- CRUD data obat
- Data nama obat, kode, satuan, stok, harga
- Monitoring stok minimum
- Riwayat penggunaan obat

### 6. Modul Admin / Master Data

- CRUD data user sistem
- CRUD data dokter
- CRUD data pasien
- CRUD data poli / ruangan
- Pengaturan role dan hak akses
- Monitoring aktivitas sistem

### 7. Modul Laporan

- Laporan jumlah pasien harian / mingguan / bulanan
- Laporan antrean pemeriksaan
- Laporan penggunaan obat
- Laporan resep yang telah diselesaikan
- Statistik pelayanan per dokter / poli

---

## Role Pengguna

### Admin

- Mengelola master data
- Mengatur user dan role
- Melihat laporan keseluruhan sistem

### Resepsionis

- Mendaftarkan pasien
- Mengisi data keluhan awal
- Membuat antrean dokter
- Melihat status kunjungan pasien

### Dokter

- Melihat antrean pasien
- Memeriksa pasien
- Menulis diagnosis
- Membuat resep obat

### Apoteker

- Melihat antrean resep
- Menyiapkan obat
- Mengelola stok obat
- Menyerahkan obat ke pasien

---

## Status Proses Pasien

Contoh status yang dapat digunakan dalam sistem:

- `registered` = pasien sudah daftar
- `waiting_doctor` = menunggu dokter
- `being_examined` = sedang diperiksa dokter
- `prescription_created` = resep sudah dibuat
- `waiting_pharmacy` = menunggu obat di apotek
- `medicine_dispensed` = obat sudah diberikan
- `completed` = proses selesai

---

## Entitas / Data Utama

### Pasien

- id
- nik
- nama
- umur
- jenis kelamin
- tanggal lahir
- alamat
- nomor telepon

### User

- id
- nama
- username
- password
- role

### Dokter

- id
- user_id
- nama dokter
- spesialisasi
- poli / ruangan

### Kunjungan

- id
- pasien_id
- dokter_id
- keluhan
- nomor antrean
- status
- tanggal kunjungan

### Rekam Medis

- id
- kunjungan_id
- diagnosis
- catatan dokter
- tanggal pemeriksaan

### Resep

- id
- rekam_medis_id
- status
- tanggal dibuat

### Resep Detail

- id
- resep_id
- obat_id
- jumlah
- aturan pakai

### Obat

- id
- nama obat
- kode obat
- stok
- satuan
- harga

---

## Rancangan Endpoint API

Berikut contoh endpoint yang bisa digunakan pada web service ini.

### Auth

- `POST /api/v1/auth/login`
- `POST /api/v1/auth/logout`
- `GET /api/v1/auth/me`

### Pasien

- `GET /api/v1/patients`
- `GET /api/v1/patients/:id`
- `POST /api/v1/patients`
- `PUT /api/v1/patients/:id`
- `DELETE /api/v1/patients/:id`

### Kunjungan / Antrean

- `GET /api/v1/visits`
- `GET /api/v1/visits/:id`
- `POST /api/v1/visits`
- `PATCH /api/v1/visits/:id/status`
- `GET /api/v1/queues/doctor`
- `GET /api/v1/queues/pharmacy`

### Dokter

- `POST /api/v1/doctor/examinations`
- `GET /api/v1/doctor/examinations/:visit_id`
- `POST /api/v1/doctor/prescriptions`
- `GET /api/v1/doctor/prescriptions/:id`

### Apotek

- `GET /api/v1/pharmacy/prescriptions`
- `GET /api/v1/pharmacy/prescriptions/:id`
- `PATCH /api/v1/pharmacy/prescriptions/:id/dispense`
- `GET /api/v1/pharmacy/medicines`
- `POST /api/v1/pharmacy/medicines`
- `PUT /api/v1/pharmacy/medicines/:id`
- `PATCH /api/v1/pharmacy/medicines/:id/stock`

### Admin / Master Data

- `GET /api/v1/users`
- `POST /api/v1/users`
- `GET /api/v1/doctors`
- `POST /api/v1/doctors`
- `GET /api/v1/polies`
- `POST /api/v1/polies`

---

## Struktur Proyek

Struktur folder pada repository ini mengikuti pendekatan clean architecture.

```text
go-hospital/
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── application/
│   ├── domain/
│   ├── infra/
│   │   ├── config/
│   │   └── persistence/
│   └── interface/
│       └── http/
├── pkg/
├── go.mod
└── readme.md
```

---

## Teknologi yang Disarankan

- Bahasa: Go
- HTTP framework: Gin / Fiber / Chi
- Database: PostgreSQL / MySQL
- ORM / Query Builder: GORM / SQLX / Ent
- Auth: JWT
- Dokumentasi API: Swagger / OpenAPI
- Migration: golang-migrate / goose

---

## Contoh Skenario Penggunaan

### Skenario 1 - Pasien Baru

1. Pasien datang ke resepsionis.
2. Resepsionis membuat data pasien baru.
3. Resepsionis memasukkan keluhan pasien.
4. Sistem membuat antrean dokter.
5. Dokter memeriksa pasien.
6. Dokter membuat resep.
7. Apoteker menyiapkan obat.
8. Pasien menerima obat dan selesai.

### Skenario 2 - Pasien Lama

1. Pasien datang dan datanya sudah tersimpan.
2. Resepsionis mencari data pasien.
3. Resepsionis membuat kunjungan baru.
4. Proses berjalan seperti antrean dokter, pemeriksaan, resep, dan apotek.

---

## Fitur Tambahan yang Bisa Dikembangkan

- Notifikasi antrean ke pasien
- Integrasi SMS / WhatsApp
- Cetak nomor antrean
- Cetak resep obat
- Dashboard real-time
- Audit log aktivitas user
- Multi cabang klinik / rumah sakit
- Integrasi pembayaran / kasir

---

## Cara Menjalankan Project

Jika project ini sudah memiliki source code API, biasanya langkah umum menjalankannya adalah:

```bash
go mod tidy
go run cmd/api/main.go
```

Jika menggunakan file environment:

```env
PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=secret
DB_NAME=go_hospital
JWT_SECRET=your_secret_key
```

