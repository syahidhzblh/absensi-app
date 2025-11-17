# 📚 Dokumentasi Aplikasi Absensi

Selamat datang! Ini adalah panduan lengkap untuk Aplikasi Absensi.

## 🚀 Mulai Cepat

**Baru pertama kali?** Mulai dari sini:

1. 📖 **[QUICKSTART.md](QUICKSTART.md)** - Panduan setup tercepat (10 menit)
2. 📘 **[README.md](README.md)** - Dokumentasi utama lengkap

## 📖 Dokumentasi Tersedia

### Untuk Developer

| Dokumen | Deskripsi | Kapan Menggunakan |
|---------|-----------|-------------------|
| **[README.md](README.md)** | Dokumentasi utama lengkap | Setup pertama kali, referensi umum |
| **[QUICKSTART.md](QUICKSTART.md)** | Setup cepat step-by-step | Ingin langsung running |
| **[API.md](API.md)** | Dokumentasi API endpoints | Integrasi, testing, development |
| **[FEATURES.md](FEATURES.md)** | Daftar lengkap fitur | Lihat kemampuan aplikasi |
| **[TROUBLESHOOTING.md](TROUBLESHOOTING.md)** | Solusi masalah umum | Ada error/masalah |

### Untuk Deployment

| Dokumen | Deskripsi | Kapan Menggunakan |
|---------|-----------|-------------------|
| **[DOCKER.md](DOCKER.md)** | Setup dengan Docker | Production deployment |
| **[.env.example](backend/.env.example)** | Template environment vars | Konfigurasi environment |

### Database

| File | Deskripsi | Kapan Menggunakan |
|------|-----------|-------------------|
| **[schema.sql](database/schema.sql)** | SQL schema database | Manual database setup |

## 🎯 Panduan Berdasarkan Kebutuhan

### "Saya ingin setup aplikasi SECEPAT mungkin"
→ Baca: **[QUICKSTART.md](QUICKSTART.md)**

### "Saya ingin memahami cara kerja aplikasi"
→ Baca: **[README.md](README.md)** + **[FEATURES.md](FEATURES.md)**

### "Saya ingin integrasi dengan sistem lain"
→ Baca: **[API.md](API.md)**

### "Saya ingin deploy ke production"
→ Baca: **[DOCKER.md](DOCKER.md)** + **[README.md](README.md)** (Security Notes)

### "Ada error, tidak bisa jalan"
→ Baca: **[TROUBLESHOOTING.md](TROUBLESHOOTING.md)**

### "Saya ingin modifikasi/develop lebih lanjut"
→ Baca: Semua dokumen, mulai dari **[README.md](README.md)**

## 📂 Struktur Project

```
absensi-app/
├── 📄 Dokumentasi
│   ├── README.md              # Dokumentasi utama
│   ├── QUICKSTART.md          # Setup cepat
│   ├── API.md                 # API documentation
│   ├── FEATURES.md            # Daftar fitur
│   ├── TROUBLESHOOTING.md     # Solusi masalah
│   ├── DOCKER.md              # Docker guide
│   └── INDEX.md               # File ini
│
├── 🔧 Backend (Golang)
│   ├── main.go                # Main application
│   ├── go.mod                 # Dependencies
│   ├── Dockerfile             # Docker config
│   └── .env.example           # Environment template
│
├── 🎨 Frontend (React)
│   ├── src/
│   │   ├── App.js             # Main component
│   │   ├── index.js           # Entry point
│   │   ├── index.css          # Styles
│   │   └── components/
│   │       ├── Login.js       # Login page
│   │       ├── Register.js    # Register page
│   │       └── Dashboard.js   # Main dashboard
│   ├── public/
│   │   └── index.html         # HTML template
│   ├── package.json           # Dependencies
│   └── Dockerfile             # Docker config
│
├── 💾 Database
│   └── schema.sql             # Database schema
│
└── 🐳 Docker
    └── docker-compose.yml     # Multi-container setup
```

## 🎓 Learning Path

**Pemula → Expert**

1. ✅ **Setup** - Ikuti QUICKSTART.md
2. ✅ **Explore** - Baca README.md dan FEATURES.md
3. ✅ **Understand** - Pelajari struktur code
4. ✅ **Test** - Gunakan aplikasi, coba semua fitur
5. ✅ **Integrate** - Baca API.md untuk integrasi
6. ✅ **Deploy** - Setup production dengan DOCKER.md
7. ✅ **Customize** - Modifikasi sesuai kebutuhan
8. ✅ **Troubleshoot** - Pelajari TROUBLESHOOTING.md

## 🔑 Informasi Penting

### Tech Stack
- **Frontend**: React 18
- **Backend**: Golang (Gin Framework)
- **Database**: PostgreSQL
- **Auth**: JWT

### Default Ports
- Frontend: `3000`
- Backend: `8080`
- PostgreSQL: `5432`

### Endpoints
- Frontend: http://localhost:3000
- Backend API: http://localhost:8080/api

## 📊 Quick Reference

### Backend Commands
```bash
cd backend
go run main.go              # Development
go build -o server          # Build binary
./server                    # Run binary
```

### Frontend Commands
```bash
cd frontend
npm install                 # Install dependencies
npm start                   # Development
npm run build              # Production build
```

### Database Commands
```bash
psql -U postgres                           # Connect
CREATE DATABASE absensi_db;                # Create DB
\c absensi_db                              # Connect to DB
\dt                                        # List tables
SELECT * FROM users;                       # Query users
SELECT * FROM attendances;                 # Query attendances
```

### Docker Commands
```bash
docker-compose up           # Start all services
docker-compose down         # Stop all services
docker-compose logs -f      # View logs
docker-compose ps           # View status
```

## 🎯 Fitur Utama

- ✅ User Registration & Login
- ✅ Daily Check-in & Check-out
- ✅ Attendance History (30 days)
- ✅ Real-time Status
- ✅ JWT Authentication
- ✅ Password Hashing
- ✅ Responsive UI
- ✅ Error Handling
- ✅ Input Validation
- ✅ Docker Support

## 🔒 Security Highlights

- Bcrypt password hashing
- JWT token authentication
- SQL injection protection
- CORS configuration
- Input validation
- No password in responses

## 📱 Supported Platforms

- ✅ Windows
- ✅ macOS
- ✅ Linux
- ✅ Docker (Any OS)

## 🆘 Need Help?

1. **Error/Bug**: Cek [TROUBLESHOOTING.md](TROUBLESHOOTING.md)
2. **Setup Issue**: Cek [QUICKSTART.md](QUICKSTART.md)
3. **API Question**: Cek [API.md](API.md)
4. **Feature Request**: Cek [FEATURES.md](FEATURES.md)
5. **Still stuck?**: Create an issue dengan detail lengkap

## 📝 Checklist Sebelum Mulai

- [ ] PostgreSQL installed
- [ ] Go 1.21+ installed  
- [ ] Node.js 16+ installed
- [ ] Git installed (optional)
- [ ] Code editor ready
- [ ] Terminal/Command Prompt ready
- [ ] Browser ready

## 🎉 Ready to Start?

1. Baca **[QUICKSTART.md](QUICKSTART.md)**
2. Setup database dan aplikasi
3. Login dan explore!

---

## 📚 Additional Resources

### External Links
- [Go Documentation](https://go.dev/doc/)
- [React Documentation](https://react.dev/)
- [PostgreSQL Documentation](https://www.postgresql.org/docs/)
- [Gin Framework](https://gin-gonic.com/docs/)
- [JWT.io](https://jwt.io/)

### Related Topics
- REST API Design
- React Hooks
- JWT Authentication
- PostgreSQL Best Practices
- Docker Containerization
- Full-stack Development

---

## 📄 License

MIT License - Bebas digunakan untuk pembelajaran dan produksi

---

## 🙏 Acknowledgments

Dibuat dengan ❤️ menggunakan:
- React
- Golang
- PostgreSQL
- Dan bantuan dari berbagai open-source libraries

---

**Happy Coding!** 🚀

Jika ada pertanyaan, jangan ragu untuk membaca dokumentasi atau membuat issue.

Semoga sukses dengan project Anda! 🎯
