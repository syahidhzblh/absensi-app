# Aplikasi Absensi - React, Golang, PostgreSQL

Aplikasi absensi sederhana dengan fitur:
- ✅ Registrasi dan Login
- ✅ Check-in dan Check-out harian
- ✅ Riwayat absensi
- ✅ Dashboard modern

## Tech Stack
- **Frontend**: React 18
- **Backend**: Golang (Gin Framework)
- **Database**: PostgreSQL

## Prerequisites
- Go 1.21 atau lebih tinggi
- Node.js 16 atau lebih tinggi
- PostgreSQL 12 atau lebih tinggi

## Setup Database

1. Install PostgreSQL (jika belum ada)

2. Buat database baru:
```bash
psql -U postgres
CREATE DATABASE absensi_db;
\q
```

3. Database akan otomatis membuat tabel saat aplikasi backend dijalankan

## Setup Backend (Golang)

1. Masuk ke folder backend:
```bash
cd backend
```

2. Install dependencies:
```bash
go mod download
```

3. Jalankan server:
```bash
go run main.go
```

Server akan berjalan di `http://localhost:8080`

## Setup Frontend (React)

1. Masuk ke folder frontend:
```bash
cd frontend
```

2. Install dependencies:
```bash
npm install
```

3. Jalankan aplikasi:
```bash
npm start
```

Aplikasi akan terbuka di `http://localhost:3000`

## Konfigurasi

### Backend Database Connection
Edit di file `backend/main.go` pada bagian `initDB()`:
```go
connStr := "host=localhost port=5432 user=postgres password=postgres dbname=absensi_db sslmode=disable"
```

Sesuaikan dengan konfigurasi PostgreSQL Anda:
- `user`: username PostgreSQL
- `password`: password PostgreSQL
- `dbname`: nama database

### JWT Secret
Edit di file `backend/main.go`:
```go
var jwtSecret = []byte("your-secret-key-change-in-production")
```

**PENTING**: Ganti dengan secret key yang aman untuk production!

## API Endpoints

### Public Endpoints
- `POST /api/register` - Registrasi user baru
- `POST /api/login` - Login user

### Protected Endpoints (Memerlukan Token)
- `POST /api/check-in` - Check-in absensi
- `POST /api/check-out` - Check-out absensi
- `GET /api/today-status` - Status absensi hari ini
- `GET /api/attendances` - Riwayat absensi (30 hari terakhir)

## Struktur Database

### Tabel `users`
- `id` (SERIAL PRIMARY KEY)
- `name` (VARCHAR)
- `email` (VARCHAR UNIQUE)
- `password` (VARCHAR - hashed)
- `created_at` (TIMESTAMP)

### Tabel `attendances`
- `id` (SERIAL PRIMARY KEY)
- `user_id` (INTEGER - FK ke users)
- `check_in` (TIMESTAMP)
- `check_out` (TIMESTAMP - nullable)
- `date` (DATE)
- `created_at` (TIMESTAMP)

## Cara Menggunakan

1. **Registrasi**: Buat akun baru dengan nama, email, dan password
2. **Login**: Login dengan email dan password
3. **Check-in**: Klik tombol "Check In" untuk mencatat waktu masuk
4. **Check-out**: Klik tombol "Check Out" untuk mencatat waktu pulang
5. **Lihat Riwayat**: Scroll ke bawah untuk melihat riwayat absensi

## Fitur

### Frontend
- ✨ UI modern dengan gradient background
- 📱 Responsive design
- 🔐 Protected routes dengan React Router
- 💾 LocalStorage untuk menyimpan token
- ⚡ Real-time status update

### Backend
- 🔒 JWT Authentication
- 🔐 Password hashing dengan bcrypt
- 🛡️ CORS enabled
- ✅ Input validation
- 📊 RESTful API design

## Troubleshooting

### Backend tidak bisa connect ke database
- Pastikan PostgreSQL sudah running
- Cek kredensial database di `initDB()`
- Pastikan database `absensi_db` sudah dibuat

### Frontend tidak bisa connect ke backend
- Pastikan backend sudah running di port 8080
- Cek CORS configuration di backend
- Pastikan API_URL di frontend sudah benar

### Error "Already checked in today"
- Anda hanya bisa check-in sekali per hari
- Gunakan check-out untuk menyelesaikan absensi hari ini

## Development

### Run Backend dengan Auto-reload
Install air untuk hot-reload:
```bash
go install github.com/cosmtrek/air@latest
cd backend
air
```

### Build untuk Production

Backend:
```bash
cd backend
go build -o absensi-server main.go
./absensi-server
```

Frontend:
```bash
cd frontend
npm run build
# Deploy folder build ke web server
```

## Security Notes

⚠️ **IMPORTANT untuk Production**:
1. Ganti JWT secret key dengan nilai yang aman
2. Gunakan HTTPS untuk production
3. Aktifkan SSL untuk koneksi PostgreSQL
4. Implementasi rate limiting
5. Tambahkan validasi input yang lebih ketat
6. Gunakan environment variables untuk credentials

## License

MIT License - Bebas digunakan untuk pembelajaran dan produksi

## Support

Jika ada pertanyaan atau masalah, silakan buat issue di repository ini.

---
Dibuat dengan ❤️ menggunakan React, Golang, dan PostgreSQL
