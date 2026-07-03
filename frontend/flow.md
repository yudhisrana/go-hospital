# Flow Penggunaan Frontend Go Hospital

Panduan singkat untuk mendemokan alur pelayanan pasien menggunakan frontend sederhana.
Semua instruksi ini diasumsikan backend (API) berjalan di `http://localhost:8080`.

## Prasyarat
1. Jalankan backend (project Go) terlebih dahulu:

```bash
cd <project-root>
go mod tidy
go run cmd/api/main.go
```

2. Siapkan database (jika belum):

```bash
sqlite3 app.db < setup_visits.sql
```

3. Jalankan static server untuk frontend (opsional, dapat juga buka file `frontend/index.html` langsung):

```bash
cd frontend
python3 -m http.server 3000
# lalu buka http://localhost:3000
```

---

## Langkah Demo (End-to-end)

1. Buka halaman frontend: `http://localhost:3000`.

2. Dasbor: klik `Refresh Dashboard` untuk melihat statistik kunjungan.

3. Jika belum ada pasien, buat pasien baru lewat API (atau gunakan cURL):

```bash
curl -X POST http://localhost:8080/patients \
  -H "Content-Type: application/json" \
  -d '{
    "nik": "3217050002960003",
    "name": "Demo Patient",
    "age": 30,
    "gender": "male",
    "birth_date": "1996-01-01T00:00:00Z",
    "address": "Jl Demo No.1",
    "phone_number": "081234567890"
  }'
```

Catat `id` pasien dari response (akan digunakan sebagai `ID Pasien`).

4. Tab "Daftar Kunjungan": masukkan `ID Pasien` dan `Gejala / Keluhan` lalu klik **Daftarkan Kunjungan**. Anda akan menerima `visit_id` di response.

5. Tab "Daftar Kunjungan": klik tombol *Semua Kunjungan* untuk melihat daftar, atau filter berdasarkan status (Terdaftar / Menunggu Apotek / Selesai).

6. Tab "Periksa Pasien": masukkan `ID Kunjungan` yang didapat saat registrasi, isi `Diagnosis` dan tambahkan satu atau lebih `Resep Obat` melalui tombol **+ Tambah Obat**, lalu klik **Simpan Pemeriksaan**.

7. Setelah pemeriksaan, status kunjungan akan berubah menjadi `waiting_pharmacy` (Menunggu Apotek). Verifikasi di tab Daftar Kunjungan.

8. Tab "Ambil Obat": masukkan `ID Kunjungan` dan klik **Selesai & Berikan Obat** — ini akan menandai kunjungan sebagai `completed`.

9. Kembali ke Dasbor → klik `Refresh Dashboard` untuk melihat statistik terkini.

---

## Catatan untuk Presentasi
- Pastikan backend dan database berjalan sebelum demo.
- Gunakan Postman atau terminal untuk membuat pasien jika belum ada (step 3).
- Frontend ini hanya untuk demo alur; tidak ada autentikasi.
- Jika ingin otomatis membuat pasien demo, saya bisa tambahkan tombol "Buat Demo Patient" di UI.

---

Jika Anda ingin, saya bisa tambahkan fitur tombol "Buat Demo Patient" pada UI untuk mempermudah demo (tanpa perlu cURL).
