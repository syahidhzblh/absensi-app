# 📡 API Documentation

Dokumentasi lengkap untuk Absensi API.

## Base URL

```
http://localhost:8080/api
```

## Authentication

API menggunakan JWT (JSON Web Token) untuk autentikasi.

**Header Format:**
```
Authorization: Bearer <your_jwt_token>
```

## Response Format

### Success Response
```json
{
  "message": "Success message",
  "data": { ... }
}
```

### Error Response
```json
{
  "error": "Error message"
}
```

---

## 🔓 Public Endpoints

### 1. Register User

Membuat user baru.

**Endpoint:** `POST /api/register`

**Request Body:**
```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "password123"
}
```

**Success Response (201):**
```json
{
  "message": "User registered successfully",
  "token": "eyJhbGciOiJIUzI1NiIs...",
  "user": {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com"
  }
}
```

**Error Response (400):**
```json
{
  "error": "Email already exists"
}
```

**cURL Example:**
```bash
curl -X POST http://localhost:8080/api/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "password123"
  }'
```

---

### 2. Login

Login user yang sudah terdaftar.

**Endpoint:** `POST /api/login`

**Request Body:**
```json
{
  "email": "john@example.com",
  "password": "password123"
}
```

**Success Response (200):**
```json
{
  "message": "Login successful",
  "token": "eyJhbGciOiJIUzI1NiIs...",
  "user": {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com"
  }
}
```

**Error Response (401):**
```json
{
  "error": "Invalid credentials"
}
```

**cURL Example:**
```bash
curl -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "password123"
  }'
```

---

## 🔐 Protected Endpoints

Semua endpoint berikut memerlukan JWT token di header.

### 3. Check In

Mencatat waktu masuk karyawan.

**Endpoint:** `POST /api/check-in`

**Headers:**
```
Authorization: Bearer <your_token>
```

**Request Body:** _(empty)_

**Success Response (200):**
```json
{
  "message": "Check-in successful",
  "id": 1
}
```

**Error Response (400):**
```json
{
  "error": "Already checked in today"
}
```

**Rules:**
- Hanya bisa check-in sekali per hari
- Waktu check-in otomatis tercatat saat request

**cURL Example:**
```bash
curl -X POST http://localhost:8080/api/check-in \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIs..."
```

---

### 4. Check Out

Mencatat waktu pulang karyawan.

**Endpoint:** `POST /api/check-out`

**Headers:**
```
Authorization: Bearer <your_token>
```

**Request Body:** _(empty)_

**Success Response (200):**
```json
{
  "message": "Check-out successful"
}
```

**Error Response (400):**
```json
{
  "error": "No active check-in found or already checked out"
}
```

**Rules:**
- Hanya bisa check-out jika sudah check-in hari ini
- Hanya bisa check-out sekali per hari
- Waktu check-out otomatis tercatat saat request

**cURL Example:**
```bash
curl -X POST http://localhost:8080/api/check-out \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIs..."
```

---

### 5. Get Today Status

Mendapatkan status absensi hari ini.

**Endpoint:** `GET /api/today-status`

**Headers:**
```
Authorization: Bearer <your_token>
```

**Success Response (200) - Belum Check In:**
```json
{
  "checked_in": false
}
```

**Success Response (200) - Sudah Check In:**
```json
{
  "checked_in": true,
  "checked_out": false,
  "attendance": {
    "id": 1,
    "user_id": 1,
    "check_in": "2024-11-17T08:30:00Z",
    "check_out": null,
    "date": "2024-11-17"
  }
}
```

**Success Response (200) - Sudah Check Out:**
```json
{
  "checked_in": true,
  "checked_out": true,
  "attendance": {
    "id": 1,
    "user_id": 1,
    "check_in": "2024-11-17T08:30:00Z",
    "check_out": "2024-11-17T17:00:00Z",
    "date": "2024-11-17"
  }
}
```

**cURL Example:**
```bash
curl -X GET http://localhost:8080/api/today-status \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIs..."
```

---

### 6. Get Attendances

Mendapatkan riwayat absensi user (30 hari terakhir).

**Endpoint:** `GET /api/attendances`

**Headers:**
```
Authorization: Bearer <your_token>
```

**Success Response (200):**
```json
[
  {
    "id": 3,
    "user_id": 1,
    "user_name": "John Doe",
    "check_in": "2024-11-17T08:30:00Z",
    "check_out": "2024-11-17T17:00:00Z",
    "date": "2024-11-17"
  },
  {
    "id": 2,
    "user_id": 1,
    "user_name": "John Doe",
    "check_in": "2024-11-16T08:25:00Z",
    "check_out": "2024-11-16T17:05:00Z",
    "date": "2024-11-16"
  },
  {
    "id": 1,
    "user_id": 1,
    "user_name": "John Doe",
    "check_in": "2024-11-15T08:35:00Z",
    "check_out": null,
    "date": "2024-11-15"
  }
]
```

**Notes:**
- Data diurutkan dari tanggal terbaru
- Maksimal 30 record terakhir
- `check_out` bisa `null` jika belum check-out

**cURL Example:**
```bash
curl -X GET http://localhost:8080/api/attendances \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIs..."
```

---

## 🔒 Error Codes

| Code | Description |
|------|-------------|
| 200 | Success |
| 201 | Created |
| 400 | Bad Request |
| 401 | Unauthorized |
| 500 | Internal Server Error |

## 🔑 JWT Token

**Token Lifetime:** 24 hours

**Token Payload:**
```json
{
  "user_id": 1,
  "exp": 1700234567
}
```

**How to Use:**
1. Login atau Register untuk mendapatkan token
2. Simpan token di localStorage atau secure storage
3. Sertakan token di header Authorization untuk setiap request ke protected endpoint
4. Token akan expire setelah 24 jam
5. Login ulang untuk mendapatkan token baru

---

## 📝 Example Workflow

### 1. Registrasi User Baru
```bash
curl -X POST http://localhost:8080/api/register \
  -H "Content-Type: application/json" \
  -d '{"name":"Jane Doe","email":"jane@example.com","password":"secure123"}'

# Response: Dapatkan token
```

### 2. Check In di Pagi Hari
```bash
curl -X POST http://localhost:8080/api/check-in \
  -H "Authorization: Bearer <token_from_login>"
  
# Response: {"message":"Check-in successful","id":1}
```

### 3. Cek Status Hari Ini
```bash
curl -X GET http://localhost:8080/api/today-status \
  -H "Authorization: Bearer <token_from_login>"
  
# Response: Status check-in dan check-out
```

### 4. Check Out di Sore Hari
```bash
curl -X POST http://localhost:8080/api/check-out \
  -H "Authorization: Bearer <token_from_login>"
  
# Response: {"message":"Check-out successful"}
```

### 5. Lihat Riwayat Absensi
```bash
curl -X GET http://localhost:8080/api/attendances \
  -H "Authorization: Bearer <token_from_login>"
  
# Response: Array of attendance records
```

---

## 🧪 Testing dengan Postman

### Setup:
1. Import collection atau buat request baru
2. Set base URL: `http://localhost:8080`
3. Create environment variable: `token`

### Test Flow:
1. **Register** → Save token ke variable
2. **Login** → Update token variable
3. **Check In** → Use `{{token}}` in Authorization
4. **Get Today Status** → Verify check-in
5. **Check Out** → Complete attendance
6. **Get Attendances** → View history

---

## 🛡️ Security Notes

1. **Password**: Disimpan dengan bcrypt hashing
2. **JWT Secret**: Ganti di production dengan secret yang kuat
3. **CORS**: Dikonfigurasi untuk allow http://localhost:3000
4. **Token**: Expires dalam 24 jam
5. **Database**: Password tidak pernah dikembalikan dalam response

---

## 📞 Support

Jika ada pertanyaan tentang API, silakan buat issue di repository.

Happy Coding! 🚀
