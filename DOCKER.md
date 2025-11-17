# 🐳 Docker Setup Guide

Panduan menjalankan aplikasi absensi menggunakan Docker.

## Prerequisites

- Docker installed
- Docker Compose installed

## 🚀 Cara Menjalankan dengan Docker

### Opsi 1: Menggunakan Docker Compose (Recommended)

Jalankan semua services (Database + Backend + Frontend) dengan satu perintah:

```bash
# Jalankan semua services
docker-compose up -d

# Lihat logs
docker-compose logs -f

# Stop semua services
docker-compose down

# Stop dan hapus semua data
docker-compose down -v
```

Aplikasi akan tersedia di:
- Frontend: http://localhost:3000
- Backend API: http://localhost:8080
- PostgreSQL: localhost:5432

### Opsi 2: Menjalankan Satu per Satu

#### 1. PostgreSQL
```bash
docker run -d \
  --name absensi-db \
  -e POSTGRES_PASSWORD=postgres \
  -e POSTGRES_DB=absensi_db \
  -p 5432:5432 \
  postgres:15-alpine
```

#### 2. Backend
```bash
# Build image
docker build -t absensi-backend ./backend

# Run container
docker run -d \
  --name absensi-backend \
  -p 8080:8080 \
  -e DB_HOST=host.docker.internal \
  absensi-backend
```

#### 3. Frontend
```bash
# Build image
docker build -t absensi-frontend ./frontend

# Run container
docker run -d \
  --name absensi-frontend \
  -p 3000:80 \
  absensi-frontend
```

## 📋 Docker Commands Cheatsheet

```bash
# Lihat semua container yang berjalan
docker ps

# Lihat semua container (termasuk yang berhenti)
docker ps -a

# Stop container
docker stop absensi-backend

# Start container
docker start absensi-backend

# Restart container
docker restart absensi-backend

# Lihat logs
docker logs absensi-backend

# Lihat logs secara real-time
docker logs -f absensi-backend

# Masuk ke dalam container
docker exec -it absensi-backend sh

# Hapus container
docker rm absensi-backend

# Hapus image
docker rmi absensi-backend
```

## 🔧 Docker Compose Commands

```bash
# Start semua services
docker-compose up

# Start di background (detached mode)
docker-compose up -d

# Rebuild dan start
docker-compose up --build

# Stop semua services
docker-compose down

# Stop dan hapus volumes
docker-compose down -v

# Lihat status services
docker-compose ps

# Lihat logs semua services
docker-compose logs

# Lihat logs service tertentu
docker-compose logs backend

# Follow logs
docker-compose logs -f

# Restart service tertentu
docker-compose restart backend

# Scale service (jika perlu)
docker-compose up -d --scale backend=3
```

## 🐛 Troubleshooting Docker

### Container tidak bisa connect ke database
```bash
# Pastikan network sudah terbuat
docker network ls

# Cek IP database
docker inspect absensi-db | grep IPAddress

# Pastikan health check database OK
docker-compose ps
```

### Port sudah digunakan
```bash
# Lihat port yang digunakan
docker ps

# Ganti port di docker-compose.yml
# Contoh: "3001:3000" untuk frontend
```

### Container crash/restart terus
```bash
# Lihat logs untuk error
docker-compose logs backend

# Cek health check
docker inspect absensi-backend
```

### Clear semua Docker resources
```bash
# HATI-HATI: Ini akan menghapus SEMUA container dan image!
docker system prune -a

# Hapus hanya yang tidak digunakan
docker system prune
```

## 💾 Backup Database di Docker

```bash
# Backup
docker exec absensi-db pg_dump -U postgres absensi_db > backup.sql

# Restore
docker exec -i absensi-db psql -U postgres absensi_db < backup.sql
```

## 🔐 Environment Variables

Edit file `.env` atau di `docker-compose.yml`:

```yaml
environment:
  DB_HOST: postgres
  DB_PORT: 5432
  DB_USER: postgres
  DB_PASSWORD: your_secure_password
  DB_NAME: absensi_db
  JWT_SECRET: your_jwt_secret_key
```

## 📊 Monitoring

```bash
# Resource usage
docker stats

# Disk usage
docker system df

# Inspect container
docker inspect absensi-backend

# View container processes
docker top absensi-backend
```

## 🚀 Production Tips

1. **Gunakan .env file** untuk credentials
2. **Set resource limits** di docker-compose.yml
3. **Enable health checks** untuk auto-restart
4. **Use multi-stage builds** untuk image yang lebih kecil
5. **Implement logging** dengan volume mount
6. **Regular backups** database dengan cron job

## 📝 Notes

- Perubahan code memerlukan rebuild: `docker-compose up --build`
- Data database disimpan di Docker volume: `postgres_data`
- Volume akan persist meskipun container dihapus
- Untuk development, gunakan volume mount untuk hot-reload

## ✅ Verifikasi Setup

```bash
# Cek semua services running
docker-compose ps

# Test backend API
curl http://localhost:8080/api/

# Test frontend
curl http://localhost:3000

# Test database connection
docker exec -it absensi-db psql -U postgres -d absensi_db -c "SELECT NOW();"
```

---

Happy Dockerizing! 🐳
