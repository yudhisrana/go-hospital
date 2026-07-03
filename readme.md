# Go Hospital - Hospital Management System
Go Hospital adalah API web service untuk manajemen klinik/rumah sakit yang membantu mengatur alur pelayanan pasien dari pendaftaran hingga pengambilan obat.
Proyek ini dirancang menggunakan **Clean Architecture (Domain-Driven Design)** dengan Fiber web framework dan SQLite database.
---
## 🏥 Alur Pelayanan Pasien
```
Pasien datang
   → Daftar ke Resepsionis (Register Visit)
   → Isi data & gejala sakit
   → Dapat nomor antrean dokter (Auto-generated)
   → Menunggu dipanggil dokter
   → Diperiksa dokter (Doctor Examine)
   → Dokter menulis diagnosis & resep obat
   → Masuk antrean apotek
   → Ambil obat (Pharmacy Dispense)
   → Pulang / Selesai
```
### Status Workflow
- `registered` - Pasien sudah daftar, menunggu dokter
- `waiting_pharmacy` - Pemeriksaan selesai, menunggu obat
- `completed` - Visit selesai, pasien pulang
---
## 📋 API Endpoints
### Patient Management
```
GET    /patients              # Lihat semua pasien
GET    /patients/:id          # Detail pasien
POST   /patients              # Tambah pasien baru
PUT    /patients/:id          # Update pasien
DELETE /patients/:id          # Hapus pasien
```
### Visit Management (Workflow Rumah Sakit)
```
POST   /visits/register       # Resepsionis: Daftar pasien baru
PATCH  /visits/:id/examine    # Dokter: Periksa & buat resep
PATCH  /visits/:id/dispense   # Apoteker: Berikan obat
GET    /visits                # Lihat semua visit
GET    /visits/:id            # Detail visit
GET    /visits/status/:status # Filter berdasarkan status
```
---
## 🗂️ Struktur Proyek (Clean Architecture)
```
go-hospital/
├── cmd/
│   └── api/
│       └── main.go                          # Entry point aplikasi
├── internal/
│   ├── application/
│   │   ├── patient/
│   │   │   ├── dto/                         # Data Transfer Object
│   │   │   └── usecase/                     # Business Logic
│   │   └── visit/
│   │       ├── dto/
│   │       └── usecase/
│   ├── domain/
│   │   ├── patient/
│   │   │   ├── entity/                      # Patient entity
│   │   │   ├── repository/                  # Interface repository
│   │   │   └── valueobject/                 # Value objects (NIK, Name, dll)
│   │   └── visit/
│   │       ├── entity/                      # Visit entity
│   │       └── repository/
│   ├── infra/
│   │   ├── config/                          # Configuration
│   │   └── persistence/
│   │       └── relational/
│   │           ├── sqlite/
│   │           │   ├── patient_repo/        # Patient repository impl
│   │           │   └── visit_repo/          # Visit repository impl
│   │           └── migrations/              # Database migrations
│   └── interface/
│       └── http/
│           ├── fiber.go                     # HTTP server setup
│           ├── handler/
│           │   ├── patient_handler/         # Patient HTTP handlers
│           │   └── visit_handler/           # Visit HTTP handlers
│           └── routes/                      # Route definitions
├── pkg/
│   └── response/                            # Response formatting
├── frontend/                                # HTML/CSS/JS frontend
├── go.mod
├── go.sum
├── app.db                                   # SQLite database
└── readme.md
```
### Penjelasan Layer
- **Domain**: Entities, Value Objects, Repository interfaces (business rules)
- **Application**: Use Cases, DTOs (orchestration logic)
- **Infrastructure**: Repository implementations, Database, Config (technical details)
- **Interface**: HTTP handlers, Routes (client communication)
---
## 🚀 Cara Menjalankan
### 1. Install Dependencies
```bash
go mod tidy
```
### 2. Setup Database
```bash
sqlite3 app.db < setup_visits.sql
```
### 3. Jalankan Server
```bash
go run cmd/api/main.go
```
Server akan berjalan di `http://localhost:8080`
---
## 📦 Teknologi yang Digunakan
- **Bahasa**: Go 1.26+
- **HTTP Framework**: Fiber v3
- **Database**: SQLite
- **Architecture**: Clean Architecture / DDD
- **UUID**: google/uuid
- **Environment**: godotenv
---
## 🧪 Test API dengan cURL
### 1. Register Visit (Resepsionis)
```bash
curl -X POST http://localhost:8080/visits/register \
  -H "Content-Type: application/json" \
  -d '{
    "patient_id": "a79f9a44-36a8-4592-ac4b-9fd963e5fc59",
    "symptoms": "Demam tinggi, batuk"
  }'
```
### 2. Examine Patient (Dokter)
```bash
curl -X PATCH http://localhost:8080/visits/\{visit_id\}/examine \
  -H "Content-Type: application/json" \
  -d '{
    "diagnosis": "Flu biasa",
    "prescriptions": [
      {"medicine": "Paracetamol", "dosage": "500mg x 3x", "quantity": 12},
      {"medicine": "Amoxicillin", "dosage": "500mg x 2x", "quantity": 10}
    ]
  }'
```
### 3. Dispense Medicine (Apoteker)
```bash
curl -X PATCH http://localhost:8080/visits/\{visit_id\}/dispense
```
### 4. View Visit Details
```bash
curl http://localhost:8080/visits/\{visit_id\}
```
---
## 📊 Database Schema
### Tabel `patients`
```sql
id, nik, name, age, gender, birth_date, address, phone, created_at, updated_at
```
### Tabel `visits`
```sql
id, patient_id, status, queue_number, symptoms, diagnosis, prescription (JSON),
registered_at, examined_at, dispensed_at, created_at, updated_at
```
### Tabel `medicines`
```sql
id, name, description, price, stock, created_at, updated_at
```
---
## ✨ Fitur Implementasi
### ✅ Phase 1 - Patient Management
- CRUD Pasien
- Validasi input (NIK, Age, Birth Date, Phone, Gender)
- Error handling komprehensif
- Response wrapper standard
### ✅ Phase 2 - Hospital Workflow
- Registrasi visit dengan auto-increment queue number
- Pemeriksaan dokter dengan diagnosis & resep
- Dispensing obat oleh apoteker
- State machine: registered → waiting_pharmacy → completed
- Filter visit berdasarkan status
### 🎯 Nilai Plus
- **Clean Architecture**: Separation of concerns yang jelas
- **DDD Pattern**: Domain-driven entities dan repositories
- **State-Driven**: Clear status transitions
- **Proper HTTP**: Menggunakan method yang sesuai (POST, PATCH)
- **Professional API**: Standard response format, error handling
---
## 🔗 Integrasi Frontend
Frontend sudah tersedia di folder `/frontend` dengan interface untuk:
- ✅ Daftar pasien
- ✅ Register visit
- ✅ Examine patient
- ✅ Dispense medicine
- ✅ View all visits & status
Akses di `http://localhost:3000` (jika menggunakan live server)
---
## 📝 Catatan Penting
- Database otomatis dibuat saat pertama kali menjalankan `setup_visits.sql`
- Queue number di-generate otomatis per hari (reset setiap hari baru)
- Prescription disimpan sebagai JSON text untuk fleksibilitas
- Semua timestamp menggunakan RFC3339 format
- API menggunakan UUID untuk semua ID
---
## 🎓 Untuk Presentasi
Project ini mendemonstrasikan:
1. **Complete Workflow** - Flow rumah sakit end-to-end
2. **Clean Architecture** - Struktur kode yang maintainable & scalable
3. **State Machine** - State transitions yang jelas
4. **RESTful API** - Proper HTTP methods & status codes
5. **Data Validation** - Comprehensive input validation
6. **Error Handling** - Structured error responses
---
**Status**: ✅ Ready for Presentation
