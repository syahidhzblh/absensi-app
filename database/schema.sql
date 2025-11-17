-- Buat Database
CREATE DATABASE absensi_db;

-- Connect ke database
\c absensi_db;

-- Buat Tabel Users
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Buat Tabel Attendances
CREATE TABLE IF NOT EXISTS attendances (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    check_in TIMESTAMP NOT NULL,
    check_out TIMESTAMP,
    date DATE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Index untuk performa
CREATE INDEX idx_attendances_user_id ON attendances(user_id);
CREATE INDEX idx_attendances_date ON attendances(date);
CREATE INDEX idx_users_email ON users(email);

-- Contoh query untuk melihat data
-- SELECT * FROM users;
-- SELECT * FROM attendances;

-- Query untuk melihat absensi hari ini
-- SELECT 
--     u.name,
--     a.check_in,
--     a.check_out,
--     a.date
-- FROM attendances a
-- JOIN users u ON a.user_id = u.id
-- WHERE a.date = CURRENT_DATE;

-- Query untuk melihat riwayat absensi user tertentu
-- SELECT 
--     a.date,
--     a.check_in,
--     a.check_out,
--     CASE 
--         WHEN a.check_out IS NOT NULL THEN 'Selesai'
--         ELSE 'Belum Check Out'
--     END as status
-- FROM attendances a
-- WHERE a.user_id = 1
-- ORDER BY a.date DESC
-- LIMIT 30;
