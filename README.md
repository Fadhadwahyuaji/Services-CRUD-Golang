# Employee Management System - CRUD API

![Go](https://img.shields.io/badge/Go-1.22.5-blue.svg)
![Gin](https://img.shields.io/badge/Gin-Web_Framework-green.svg)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-Database-blue.svg)
![License](https://img.shields.io/badge/License-FADHAD_WAHYU_AJI-red.svg)

## 📋 Deskripsi

Employee Management System adalah aplikasi REST API yang dibangun menggunakan Go (Golang) dengan framework Gin untuk mengelola data karyawan, pengguna, peran (roles), dan posisi dalam sebuah organisasi. Aplikasi ini menyediakan sistem autentikasi JWT dan operasi CRUD lengkap untuk semua entitas.

## 🚀 Teknologi yang Digunakan

### Backend Framework & Language

- **Go (Golang)** `v1.22.5` - Bahasa pemrograman utama
- **Gin Web Framework** `v1.10.0` - HTTP web framework untuk Go
- **SQLx** `v1.4.0` - Extensions to database/sql untuk Go
- **JWT-Go** `v3.2.0` - JSON Web Token implementation untuk Go

### Database

- **PostgreSQL** - Database relasional untuk penyimpanan data
- **PostgreSQL Driver** `lib/pq v1.10.9` - Driver PostgreSQL untuk Go

### Security & Authentication

- **bcrypt** - Password hashing
- **JWT (JSON Web Tokens)** - Autentikasi dan autorisasi
- **crypto/rand** - Secure random number generation

### Development Tools

- **Go Modules** - Dependency management
- **Gin Static Files** - Static file serving

## 🎯 Fitur Utama

### 🔐 Autentikasi & Autorisasi

- **Registrasi Pengguna** - Pendaftaran akun baru dengan enkripsi password
- **Login System** - Autentikasi dengan email dan password
- **JWT Tokens** - Token berbasis JWT untuk otorisasi
- **Middleware Protection** - Perlindungan endpoint dengan middleware autentikasi
- **Role-based Access** - Sistem peran untuk kontrol akses

### 👥 Manajemen Pengguna (Users)

- **Create User** - Tambah pengguna baru
- **Get All Users** - Lihat daftar semua pengguna aktif
- **Get User by ID** - Detail pengguna berdasarkan ID
- **Update User** - Perbarui informasi pengguna
- **Delete User** - Hapus pengguna (soft delete)

### 👔 Manajemen Karyawan (Employees)

- **Create Employee** - Tambah data karyawan baru
- **Get All Employees** - Lihat daftar semua karyawan aktif
- **Get Employee by ID** - Detail karyawan berdasarkan ID
- **Update Employee** - Perbarui informasi karyawan
- **Delete Employee** - Hapus karyawan (soft delete)

### 🎭 Manajemen Peran (Roles)

- **Create Role** - Tambah peran baru
- **Get All Roles** - Lihat daftar semua peran aktif
- **Get Role by ID** - Detail peran berdasarkan ID
- **Update Role** - Perbarui informasi peran
- **Delete Role** - Hapus peran (soft delete)

### 💼 Manajemen Posisi (Positions)

- **Create Position** - Tambah posisi baru
- **Get All Positions** - Lihat daftar semua posisi aktif
- **Get Position by ID** - Detail posisi berdasarkan ID
- **Update Position** - Perbarui informasi posisi
- **Delete Position** - Hapus posisi (soft delete)

### 🛡️ Keamanan Data

- **Soft Delete** - Data tidak dihapus permanen, hanya ditandai sebagai terhapus
- **Password Hashing** - Enkripsi password menggunakan bcrypt
- **Input Validation** - Validasi input untuk semua endpoint
- **SQL Injection Protection** - Menggunakan prepared statements

## 📊 Struktur Database

### Tabel Users

```sql
id (int, primary key)
name (string)
email (string, unique)
username (string, unique)
password (string, hashed)
role_id (int, foreign key)
is_deleted (boolean)
created_at (timestamp)
updated_at (timestamp)
```

### Tabel Employees

```sql
id (int, primary key)
name (string)
phone_number (string)
email (string)
position_id (int, foreign key)
is_deleted (boolean)
created_at (timestamp)
updated_at (timestamp)
```

### Tabel Roles

```sql
id (int, primary key)
role_name (string)
is_deleted (boolean)
created_at (timestamp)
updated_at (timestamp)
```

### Tabel Positions

```sql
id (int, primary key)
position_name (string)
is_deleted (boolean)
created_at (timestamp)
updated_at (timestamp)
```

## 🛠️ Instalasi dan Setup

### Prerequisites

- Go 1.22.5 atau lebih tinggi
- PostgreSQL database
- Git

### Langkah Instalasi

1. **Clone Repository**

   ```bash
   git clone <repository-url>
   cd crud-golang
   ```

2. **Install Dependencies**

   ```bash
   go mod download
   ```

3. **Setup Database**

   - Buat database PostgreSQL dengan nama `db_karyawan_sangkuriang`
   - Update connection string di `db/db.go` sesuai konfigurasi database Anda

   ```go
   connStr := "user=your_user dbname=db_karyawan_sangkuriang sslmode=disable password=your_password"
   ```

4. **Buat Tabel Database**
   Jalankan script SQL untuk membuat tabel yang diperlukan (users, employees, roles, positions)

5. **Generate Secret Key (Opsional)**

   ```bash
   go run generateSecretKey.go
   ```

6. **Build dan Run Application**

   ```bash
   go build -o main.exe
   ./main.exe
   ```

   Atau jalankan langsung:

   ```bash
   go run main.go
   ```

Application akan berjalan di `http://localhost:8080`

## 📡 API Endpoints

### Authentication Endpoints

```http
POST /register          # Registrasi pengguna baru
POST /login            # Login pengguna
```

### Protected Endpoints (Memerlukan JWT Token)

#### Users Management

```http
POST   /users          # Create user baru
GET    /users          # Get semua users
GET    /users/:id      # Get user by ID
PUT    /users-update/:id # Update user
DELETE /users/:id      # Delete user
```

#### Employee Management

```http
POST   /employee       # Create employee baru
GET    /employee       # Get semua employees
GET    /employee/:id   # Get employee by ID
PUT    /employee/:id   # Update employee
DELETE /employee/:id   # Delete employee
```

#### Role Management

```http
POST   /role          # Create role baru
GET    /role          # Get semua roles
GET    /role/:id      # Get role by ID
PUT    /role/:id      # Update role
DELETE /role/:id      # Delete role
```

#### Position Management

```http
POST   /position      # Create position baru
GET    /position      # Get semua positions
GET    /position/:id  # Get position by ID
PUT    /position/:id  # Update position
DELETE /position/:id  # Delete position
```

## 📝 Contoh Penggunaan

### 1. Registrasi Pengguna

```json
POST /register
{
    "name": "John Doe",
    "email": "john@example.com",
    "username": "johndoe",
    "password": "password123",
    "role_id": 1
}
```

### 2. Login

```json
POST /login
{
    "email": "john@example.com",
    "password": "password123"
}
```

### 3. Create Employee (dengan JWT Token)

```json
POST /employee
Headers: Authorization: Bearer <jwt_token>
{
    "name": "Jane Smith",
    "phone_number": "081234567890",
    "email": "jane@company.com",
    "position_id": 1
}
```

## 🏗️ Arsitektur Proyek

```
crud-golang/
├── main.go                 # Entry point aplikasi
├── go.mod                  # Dependencies management
├── go.sum                  # Dependencies checksums
├── generateSecretKey.go    # Utility untuk generate secret key
├── db/
│   └── db.go              # Database connection dan konfigurasi
├── models/
│   ├── models.go          # Struct definitions untuk semua models
│   └── auth.go            # Models untuk authentication
├── handlers/
│   ├── auth.go            # Authentication handlers
│   ├── user_handlers.go   # User CRUD handlers
│   ├── employee_handlers.go # Employee CRUD handlers
│   ├── role_handlers.go   # Role CRUD handlers
│   └── position_handlers.go # Position CRUD handlers
├── middleware/
│   └── auth.go            # Authentication middleware
├── routers/
│   └── router.go          # Route definitions
└── view/
    ├── index.html         # Frontend view (opsional)
    ├── css/
    │   └── style.css      # Styling
    └── js/
        └── script.js      # JavaScript frontend
```

## 🔧 Konfigurasi

### Environment Variables

Anda dapat mengatur konfigurasi melalui environment variables:

- `DB_HOST` - Database host
- `DB_PORT` - Database port
- `DB_USER` - Database username
- `DB_PASSWORD` - Database password
- `DB_NAME` - Database name
- `JWT_SECRET` - JWT secret key
- `PORT` - Application port (default: 8080)

### Database Connection

Update file `db/db.go` untuk mengubah string koneksi database:

```go
connStr := "user=username dbname=database_name sslmode=disable password=password host=localhost port=5432"
```

## 🧪 Testing

### Manual Testing Tools

Untuk testing API endpoints, Anda dapat menggunakan:

- **Postman** - Import collection dan test semua endpoints
- **cURL** - Command line testing
- **Thunder Client** (VS Code extension)
- **Insomnia** - REST client

Contoh test dengan cURL:

```bash
# Test registrasi
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{"name":"Test User","email":"test@example.com","username":"testuser","password":"password123","role_id":1}'

# Test login
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"password123"}'
```

### 🔗 Integration Testing dengan Laravel

API ini telah diintegrasikan dan ditest menggunakan aplikasi frontend Laravel yang berfungsi sebagai consumer untuk semua endpoint yang tersedia.

**🌐 Repository Laravel Frontend:**  
[![Laravel Integration](https://img.shields.io/badge/Laravel-Frontend_Integration-red.svg)](https://github.com/Fadhadwahyuaji/Laravel-Go)

**Link:** https://github.com/Fadhadwahyuaji/Laravel-Go

#### Fitur Testing Laravel Integration:

- ✅ **Authentication Flow** - Testing complete login/register dengan JWT
- ✅ **CRUD Operations** - Testing semua operasi Create, Read, Update, Delete
- ✅ **User Management** - Interface untuk mengelola data users
- ✅ **Employee Management** - Interface untuk mengelola data karyawan
- ✅ **Role & Position Management** - Interface untuk master data roles dan positions
- ✅ **API Integration** - Konsumsi semua endpoint Go API menggunakan HTTP Client Laravel
- ✅ **Error Handling** - Testing response handling untuk success dan error cases
- ✅ **Token Management** - Testing JWT token lifecycle dan refresh mechanism

#### Teknologi Laravel Integration:

- **Laravel Framework** - PHP web framework untuk frontend
- **Guzzle HTTP Client** - Untuk consuming Go API endpoints
- **Laravel Authentication** - Session management terintegrasi dengan JWT
- **Blade Templates** - UI untuk testing semua fitur API
- **Bootstrap/CSS** - Responsive UI design

Proyek Laravel ini memberikan bukti bahwa Go API berfungsi dengan sempurna dan dapat dikonsumsi oleh aplikasi frontend modern.

## 🐛 Bug Reports & Feature Requests

Jika Anda menemukan bug atau ingin request fitur baru, silakan buat issue di repository ini dengan detail yang lengkap.

## 📞 Kontak

**Developer:** FADHAD WAHYU AJI  
**Email:** [fadhadwahyuaji@gmail.com]  
**LinkedIn:** [[fadhadwahyuaji](https://www.linkedin.com/in/fadhadwahyuaji/)]

## 📄 Lisensi

Proyek ini dimiliki oleh **FADHAD WAHYU AJI**. Semua hak cipta dilindungi undang-undang.

```
Copyright (c) 2024 FADHAD WAHYU AJI

Hak cipta proyek ini sepenuhnya dimiliki oleh FADHAD WAHYU AJI.
Penggunaan, modifikasi, dan distribusi tanpa izin tertulis dari pemilik dilarang.
```

---

## 🌟 Acknowledgments

- Terima kasih kepada komunitas Go Indonesia
- Gin Web Framework contributors
- PostgreSQL team
- Open source community

---

**⭐ Jika proyek ini membantu Anda, jangan lupa untuk memberikan star!**
