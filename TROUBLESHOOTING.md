# 🔧 Troubleshooting Guide

Solusi untuk masalah umum yang mungkin terjadi.

## 🔴 Backend Issues

### Error: "pq: password authentication failed"

**Penyebab**: Kredensial PostgreSQL salah

**Solusi**:
```bash
# 1. Cek kredensial PostgreSQL Anda
psql -U postgres

# 2. Edit file backend/main.go baris 40
connStr := "host=localhost port=5432 user=YOUR_USER password=YOUR_PASSWORD dbname=absensi_db sslmode=disable"

# 3. Atau reset password postgres
sudo -u postgres psql
ALTER USER postgres PASSWORD 'newpassword';
```

---

### Error: "dial tcp [::1]:5432: connect: connection refused"

**Penyebab**: PostgreSQL tidak running

**Solusi**:
```bash
# Ubuntu/Debian
sudo service postgresql start
sudo service postgresql status

# macOS (dengan Homebrew)
brew services start postgresql

# Windows
# Start PostgreSQL dari Services.msc

# Verify
psql -U postgres -c "SELECT version();"
```

---

### Error: "database does not exist"

**Penyebab**: Database belum dibuat

**Solusi**:
```bash
# Buat database
psql -U postgres -c "CREATE DATABASE absensi_db;"

# Verify
psql -U postgres -l | grep absensi
```

---

### Error: "cannot find package"

**Penyebab**: Dependencies belum terinstall

**Solusi**:
```bash
cd backend
go mod download
go mod tidy

# Jika masih error
rm go.sum
go mod download
```

---

### Error: "bind: address already in use"

**Penyebab**: Port 8080 sudah digunakan

**Solusi**:
```bash
# Cari process yang menggunakan port 8080
# Linux/Mac
lsof -i :8080
kill -9 <PID>

# Windows
netstat -ano | findstr :8080
taskkill /PID <PID> /F

# Atau ganti port di main.go
port := "8081"  # Ganti dari 8080
```

---

## 🔵 Frontend Issues

### Error: "npm: command not found"

**Penyebab**: Node.js belum terinstall

**Solusi**:
```bash
# Download dari https://nodejs.org/
# Atau dengan package manager

# Ubuntu/Debian
curl -fsSL https://deb.nodesource.com/setup_18.x | sudo -E bash -
sudo apt-get install -y nodejs

# macOS
brew install node

# Verify
node --version
npm --version
```

---

### Error: "npm ERR! EACCES: permission denied"

**Penyebab**: Permission issue

**Solusi**:
```bash
# Jangan gunakan sudo!
# Fix npm permissions
mkdir ~/.npm-global
npm config set prefix '~/.npm-global'
echo 'export PATH=~/.npm-global/bin:$PATH' >> ~/.bashrc
source ~/.bashrc

# Atau gunakan nvm
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.0/install.sh | bash
nvm install 18
```

---

### Error: "Module not found"

**Penyebab**: Dependencies belum terinstall

**Solusi**:
```bash
cd frontend
rm -rf node_modules package-lock.json
npm install

# Atau
npm ci  # Untuk clean install
```

---

### Error: "Port 3000 is already in use"

**Penyebab**: Port 3000 sudah digunakan

**Solusi**:
```bash
# Opsi 1: Kill process
# Linux/Mac
lsof -i :3000
kill -9 <PID>

# Windows
netstat -ano | findstr :3000
taskkill /PID <PID> /F

# Opsi 2: Gunakan port lain
PORT=3001 npm start
```

---

### Error: "Failed to fetch" atau "Network Error"

**Penyebab**: Backend tidak running atau CORS issue

**Solusi**:
```bash
# 1. Pastikan backend running
curl http://localhost:8080/api/

# 2. Cek CORS di backend/main.go
AllowOrigins: []string{"http://localhost:3000"}

# 3. Cek API URL di frontend
# Pastikan API_URL benar di components

# 4. Test dengan curl
curl -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{"email":"test@test.com","password":"test123"}'
```

---

## 🟡 Database Issues

### Error: "relation does not exist"

**Penyebab**: Tabel belum dibuat

**Solusi**:
```bash
# Tabel dibuat otomatis saat backend start
# Jika tidak, jalankan manual:
psql -U postgres -d absensi_db -f database/schema.sql

# Atau
psql -U postgres -d absensi_db
\i database/schema.sql
```

---

### Error: "duplicate key value violates unique constraint"

**Penyebab**: Email sudah terdaftar

**Solusi**:
```bash
# Ini expected behavior untuk prevent duplicate email
# Gunakan email lain untuk registrasi

# Atau hapus user lama
psql -U postgres -d absensi_db
DELETE FROM users WHERE email = 'duplicate@email.com';
```

---

### Database Slow/Hanging

**Solusi**:
```bash
# 1. Cek koneksi aktif
psql -U postgres -d absensi_db
SELECT * FROM pg_stat_activity;

# 2. Kill koneksi yang hanging
SELECT pg_terminate_backend(pid) 
FROM pg_stat_activity 
WHERE datname = 'absensi_db' AND pid <> pg_backend_pid();

# 3. Vacuum database
VACUUM ANALYZE;

# 4. Restart PostgreSQL
sudo service postgresql restart
```

---

## 🟢 Authentication Issues

### Error: "Invalid token" atau "Unauthorized"

**Penyebab**: Token expired atau invalid

**Solusi**:
```javascript
// 1. Clear localStorage dan login ulang
localStorage.clear();
// Refresh page dan login

// 2. Cek JWT secret consistency
// Backend dan frontend harus sama

// 3. Cek token expiry (24 jam)
// Login ulang jika sudah expired
```

---

### Auto Logout

**Penyebab**: Token expired

**Solusi**:
```javascript
// Normal behavior - token expires after 24 hours
// Login ulang untuk mendapatkan token baru

// Untuk extend token lifetime, edit backend/main.go:
"exp": time.Now().Add(time.Hour * 48).Unix(), // 48 jam
```

---

### "Already checked in today"

**Penyebab**: Sudah check-in hari ini

**Solusi**:
```bash
# Ini expected behavior
# Gunakan check-out untuk menyelesaikan absensi

# Untuk testing, hapus data hari ini:
psql -U postgres -d absensi_db
DELETE FROM attendances WHERE date = CURRENT_DATE;
```

---

## 🐳 Docker Issues

### Error: "Cannot connect to Docker daemon"

**Penyebab**: Docker tidak running

**Solusi**:
```bash
# Start Docker
# Linux
sudo systemctl start docker
sudo systemctl enable docker

# macOS/Windows
# Start Docker Desktop

# Verify
docker ps
```

---

### Error: "port is already allocated"

**Penyebab**: Port sudah digunakan

**Solusi**:
```bash
# Opsi 1: Stop container yang menggunakan port
docker ps
docker stop <container_name>

# Opsi 2: Ganti port di docker-compose.yml
ports:
  - "8081:8080"  # Ganti 8081 dengan port lain
```

---

### Container Keeps Restarting

**Solusi**:
```bash
# 1. Lihat logs
docker logs absensi-backend

# 2. Cek health status
docker inspect absensi-backend

# 3. Enter container untuk debug
docker exec -it absensi-backend sh

# 4. Rebuild dengan no-cache
docker-compose up --build --force-recreate
```

---

## 💾 Data Issues

### Lost Data After Restart

**Penyebab**: Docker volumes tidak terkonfigurasi

**Solusi**:
```bash
# Data ada di Docker volume
docker volume ls

# Jangan gunakan -v flag saat down
docker-compose down  # Data tetap ada
docker-compose down -v  # Data AKAN HILANG

# Backup data
docker exec absensi-db pg_dump -U postgres absensi_db > backup.sql

# Restore data
docker exec -i absensi-db psql -U postgres absensi_db < backup.sql
```

---

### Reset Database

**Solusi**:
```bash
# Opsi 1: Drop dan recreate
psql -U postgres
DROP DATABASE absensi_db;
CREATE DATABASE absensi_db;
\q

# Opsi 2: Truncate tables
psql -U postgres -d absensi_db
TRUNCATE TABLE attendances CASCADE;
TRUNCATE TABLE users CASCADE;

# Opsi 3: Docker (hapus volume)
docker-compose down -v
docker-compose up -d
```

---

## 🌐 Browser Issues

### CORS Error in Console

**Solusi**:
```go
// Pastikan di backend/main.go:
router.Use(cors.New(cors.Config{
    AllowOrigins:     []string{"http://localhost:3000"},
    AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
    AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
    AllowCredentials: true,
}))
```

---

### Token Not Persisting

**Solusi**:
```javascript
// Cek localStorage
console.log(localStorage.getItem('token'));

// Clear cache dan cookies
// Chrome: Ctrl+Shift+Delete

// Try incognito mode
```

---

## 🔍 General Debugging

### Enable Debug Mode

**Backend:**
```go
// Add di main.go
gin.SetMode(gin.DebugMode)
```

**Frontend:**
```javascript
// Add console logs
console.log('API Response:', response);
console.log('Error:', error);
```

### Check Service Status

```bash
# Backend
curl http://localhost:8080/api/

# Frontend
curl http://localhost:3000

# Database
psql -U postgres -c "SELECT version();"
```

### View Logs

```bash
# Backend logs
# Akan terlihat di terminal saat go run main.go

# Frontend logs
# Buka Browser Console (F12)

# PostgreSQL logs
# Ubuntu
tail -f /var/log/postgresql/postgresql-15-main.log

# Docker
docker logs absensi-backend -f
docker logs absensi-db -f
```

---

## 📞 Getting Help

Jika masalah masih berlanjut:

1. **Cek dokumentasi**:
   - README.md
   - API.md
   - QUICKSTART.md

2. **Search error message**:
   - Stack Overflow
   - GitHub Issues
   - Google

3. **Create issue**:
   - Sertakan error message lengkap
   - Screenshot jika perlu
   - Steps to reproduce
   - Environment details (OS, versions)

4. **Debug systematic**:
   - Isolate the problem
   - Check logs
   - Test each component separately
   - Verify configuration

---

## ✅ Quick Checklist

Sebelum melaporkan bug, pastikan:

- [ ] PostgreSQL running
- [ ] Database created
- [ ] Backend running (port 8080)
- [ ] Frontend running (port 3000)
- [ ] No firewall blocking
- [ ] Correct credentials
- [ ] Dependencies installed
- [ ] Logs checked
- [ ] Environment variables correct
- [ ] Network connectivity OK

---

**Remember**: Most issues can be solved by:
1. Reading error messages carefully
2. Checking logs
3. Verifying configuration
4. Restarting services
5. Clearing cache

Good luck! 🍀
