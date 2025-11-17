# Quick Start Guide - Aplikasi Absensi

Panduan cepat untuk menjalankan aplikasi absensi.

## 🚀 Langkah Cepat (Quick Setup)

### 1️⃣ Setup Database (5 menit)

```bash
# Login ke PostgreSQL
psql -U postgres

# Buat database
CREATE DATABASE absensi_db;

# Keluar
\q
```

### 2️⃣ Jalankan Backend (2 menit)

```bash
# Masuk ke folder backend
cd backend

# Download dependencies
go mod download

# Jalankan server
go run main.go
```

✅ Backend berjalan di `http://localhost:8080`

### 3️⃣ Jalankan Frontend (3 menit)

Buka terminal baru:

```bash
# Masuk ke folder frontend
cd frontend

# Install dependencies
npm install

# Jalankan aplikasi
npm start
```

✅ Aplikasi terbuka otomatis di browser `http://localhost:3000`

## 📱 Cara Menggunakan

1. **Registrasi** - Buat akun baru
2. **Login** - Masuk dengan akun yang sudah dibuat
3. **Check In** - Klik tombol hijau untuk absen masuk
4. **Check Out** - Klik tombol kuning untuk absen pulang
5. **Lihat Riwayat** - Scroll ke bawah untuk melihat riwayat

## ⚙️ Konfigurasi Database

Jika PostgreSQL Anda menggunakan username/password yang berbeda:

Edit file `backend/main.go` pada baris 40:
```go
connStr := "host=localhost port=5432 user=GANTI_USER password=GANTI_PASSWORD dbname=absensi_db sslmode=disable"
```

## 🆘 Troubleshooting

### ❌ Backend error: "connection refused"
**Solusi**: Pastikan PostgreSQL sudah running
```bash
# Linux/Mac
sudo service postgresql start

# Windows
# Jalankan PostgreSQL dari Services
```

### ❌ Backend error: "database does not exist"
**Solusi**: Buat database terlebih dahulu
```bash
psql -U postgres -c "CREATE DATABASE absensi_db;"
```

### ❌ Frontend error saat npm install
**Solusi**: Pastikan Node.js sudah terinstall
```bash
node --version  # Harus versi 16 atau lebih tinggi
npm --version
```

### ❌ Frontend tidak bisa connect ke backend
**Solusi**: 
1. Pastikan backend sudah running di port 8080
2. Cek di browser: http://localhost:8080/api/
3. Pastikan tidak ada firewall yang memblokir

## 📋 Checklist Setup

- [ ] PostgreSQL terinstall dan running
- [ ] Database `absensi_db` sudah dibuat
- [ ] Go 1.21+ terinstall
- [ ] Node.js 16+ terinstall
- [ ] Backend running di port 8080
- [ ] Frontend running di port 3000
- [ ] Browser terbuka ke http://localhost:3000

## 🎯 Default Credentials

Tidak ada default user. Anda harus registrasi terlebih dahulu.

## 📊 Port yang Digunakan

- Backend: `8080`
- Frontend: `3000`
- PostgreSQL: `5432`

## 💡 Tips

1. **Auto-reload Backend**: Install `air` untuk hot-reload
   ```bash
   go install github.com/cosmtrek/air@latest
   cd backend && air
   ```

2. **Lihat Log Database**: 
   ```bash
   psql -U postgres -d absensi_db
   SELECT * FROM users;
   SELECT * FROM attendances;
   ```

3. **Reset Database**:
   ```sql
   TRUNCATE TABLE attendances CASCADE;
   TRUNCATE TABLE users CASCADE;
   ```

## 🔐 Keamanan

⚠️ **Untuk Production**:
- Ganti JWT secret di `backend/main.go` baris 16
- Gunakan HTTPS
- Tambahkan rate limiting
- Gunakan environment variables

## ✨ Fitur Aplikasi

✅ Registrasi dan Login dengan JWT
✅ Check-in dan Check-out harian
✅ Validasi (1 check-in per hari)
✅ Riwayat absensi 30 hari terakhir
✅ UI modern dan responsive
✅ Password hashing dengan bcrypt
✅ Protected API routes

---

Selamat menggunakan! 🎉

Jika ada masalah, lihat README.md untuk dokumentasi lengkap.
