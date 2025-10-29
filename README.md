# CRUD Golang - Sistem Manajemen Karyawan

Aplikasi REST API untuk sistem manajemen karyawan yang dibangun menggunakan Go (Golang) dengan framework Gin dan PostgreSQL sebagai database.

## 📋 Daftar Isi

- [Fitur](#-fitur)
- [Teknologi yang Digunakan](#-teknologi-yang-digunakan)
- [Struktur Proyek](#-struktur-proyek)
- [Prasyarat](#-prasyarat)
- [Instalasi](#-instalasi)
- [Konfigurasi Database](#-konfigurasi-database)
- [Menjalankan Aplikasi](#-menjalankan-aplikasi)
- [API Endpoints](#-api-endpoints)
- [Autentikasi](#-autentikasi)
- [Model Data](#-model-data)
- [Kontribusi](#-kontribusi)
- [Lisensi](#-lisensi)

## ✨ Fitur

- **Autentikasi & Autorisasi**
  - Registrasi pengguna baru
  - Login dengan JWT (JSON Web Token)
  - Password hashing menggunakan bcrypt
  - Middleware autentikasi untuk endpoint yang dilindungi

- **Manajemen CRUD Lengkap**
  - **Users**: Kelola data pengguna sistem
  - **Employees**: Kelola data karyawan
  - **Roles**: Kelola peran/role pengguna
  - **Positions**: Kelola jabatan karyawan

- **Soft Delete**: Implementasi soft delete untuk semua entitas
- **RESTful API**: Endpoint yang mengikuti standar REST
- **Validasi Input**: Validasi data input menggunakan Gin binding

## 🛠 Teknologi yang Digunakan

### Backend
- **Go 1.22.5** - Bahasa pemrograman
- **Gin v1.10.0** - Web framework
- **PostgreSQL** - Database relasional
- **sqlx v1.4.0** - Extension untuk database/sql
- **JWT (dgrijalva/jwt-go)** - Autentikasi token
- **bcrypt** - Password hashing

### Libraries Utama
```
github.com/gin-gonic/gin v1.10.0
github.com/jmoiron/sqlx v1.4.0
github.com/lib/pq v1.10.9
golang.org/x/crypto v0.23.0
github.com/dgrijalva/jwt-go v3.2.0+incompatible
```

## 📁 Struktur Proyek

```
crud-golang/
├── db/                     # Konfigurasi database
│   └── db.go              # Inisialisasi koneksi database
├── handlers/              # Handler untuk endpoint API
│   ├── auth.go           # Handler autentikasi (login, register)
│   ├── user_handlers.go  # Handler CRUD user
│   ├── employee_handlers.go  # Handler CRUD employee
│   ├── role_handlers.go  # Handler CRUD role
│   └── position_handlers.go  # Handler CRUD position
├── middleware/           # Middleware aplikasi
│   └── auth.go          # Middleware autentikasi JWT
├── models/              # Model data
│   ├── models.go       # Model untuk User, Employee, Role, Position
│   └── auth.go         # Model untuk Login dan Register
├── routers/            # Konfigurasi routing
│   └── router.go      # Setup semua routes
├── view/              # Frontend (optional)
│   ├── index.html    # Halaman utama
│   ├── css/         # Stylesheet
│   └── js/          # JavaScript files
├── main.go           # Entry point aplikasi
├── go.mod           # Go modules
├── go.sum          # Go dependencies checksum
└── generateSecretKey.go  # Utility untuk generate JWT secret
```

## 📦 Prasyarat

Sebelum menjalankan aplikasi, pastikan Anda telah menginstal:

- **Go** (versi 1.22.5 atau lebih baru) - [Download](https://golang.org/dl/)
- **PostgreSQL** (versi 12 atau lebih baru) - [Download](https://www.postgresql.org/download/)
- **Git** - [Download](https://git-scm.com/)

## 🚀 Instalasi

1. **Clone repository**
   ```bash
   git clone <repository-url>
   cd crud-golang
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Verifikasi instalasi**
   ```bash
   go mod verify
   ```

## 💾 Konfigurasi Database

1. **Buat database PostgreSQL**
   ```sql
   CREATE DATABASE db_karyawan_sangkuriang;
   ```

2. **Buat tabel yang diperlukan**

   ```sql
   -- Tabel Roles
   CREATE TABLE roles (
       id SERIAL PRIMARY KEY,
       role_name VARCHAR(100) NOT NULL,
       is_deleted BOOLEAN DEFAULT FALSE,
       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
   );

   -- Tabel Positions
   CREATE TABLE positions (
       id SERIAL PRIMARY KEY,
       position_name VARCHAR(100) NOT NULL,
       is_deleted BOOLEAN DEFAULT FALSE,
       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
   );

   -- Tabel Users
   CREATE TABLE users (
       id SERIAL PRIMARY KEY,
       name VARCHAR(255) NOT NULL,
       email VARCHAR(255) UNIQUE NOT NULL,
       username VARCHAR(100) UNIQUE NOT NULL,
       password VARCHAR(255) NOT NULL,
       role_id INTEGER REFERENCES roles(id),
       is_deleted BOOLEAN DEFAULT FALSE,
       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
   );

   -- Tabel Employees
   CREATE TABLE employees (
       id SERIAL PRIMARY KEY,
       name VARCHAR(255) NOT NULL,
       phone_number VARCHAR(20),
       email VARCHAR(255) UNIQUE NOT NULL,
       position_id INTEGER REFERENCES positions(id),
       is_deleted BOOLEAN DEFAULT FALSE,
       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
   );
   ```

3. **Update konfigurasi database** di `db/db.go`
   ```go
   connStr := "user=YOUR_USERNAME dbname=db_karyawan_sangkuriang sslmode=disable password=YOUR_PASSWORD"
   ```
   
   Sesuaikan dengan kredensial PostgreSQL Anda:
   - `YOUR_USERNAME`: Username PostgreSQL
   - `YOUR_PASSWORD`: Password PostgreSQL

4. **Insert data awal (optional)**
   ```sql
   -- Insert default roles
   INSERT INTO roles (role_name) VALUES 
       ('Admin'),
       ('Manager'),
       ('Staff');

   -- Insert default positions
   INSERT INTO positions (position_name) VALUES 
       ('CEO'),
       ('Manager'),
       ('Developer'),
       ('Designer'),
       ('HR Staff');
   ```

## ▶️ Menjalankan Aplikasi

1. **Jalankan aplikasi**
   ```bash
   go run main.go
   ```

2. **Build aplikasi (optional)**
   ```bash
   go build -o main.exe
   ./main.exe
   ```

3. **Aplikasi akan berjalan di**: `http://localhost:8080`

## 🔌 API Endpoints

### Autentikasi

| Method | Endpoint | Deskripsi | Auth Required |
|--------|----------|-----------|---------------|
| POST | `/register` | Registrasi pengguna baru | ❌ |
| POST | `/login` | Login pengguna | ❌ |

**Request Body Register:**
```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "username": "johndoe",
  "password": "password123",
  "role_id": 1
}
```

**Request Body Login:**
```json
{
  "username": "johndoe",
  "password": "password123"
}
```

**Response Login:**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

### Users

| Method | Endpoint | Deskripsi | Auth Required |
|--------|----------|-----------|---------------|
| POST | `/users` | Buat user baru | ✅ |
| GET | `/users` | Ambil semua users | ✅ |
| GET | `/users/:id` | Ambil user berdasarkan ID | ❌ |
| PUT | `/users-update/:id` | Update user | ✅ |
| DELETE | `/users/:id` | Hapus user (soft delete) | ✅ |

### Employees

| Method | Endpoint | Deskripsi | Auth Required |
|--------|----------|-----------|---------------|
| POST | `/employee` | Buat employee baru | ✅ |
| GET | `/employee` | Ambil semua employees | ✅ |
| GET | `/employee/:id` | Ambil employee berdasarkan ID | ✅ |
| PUT | `/employee/:id` | Update employee | ✅ |
| DELETE | `/employee/:id` | Hapus employee (soft delete) | ✅ |

**Request Body Employee:**
```json
{
  "name": "Jane Smith",
  "phone_number": "081234567890",
  "email": "jane@example.com",
  "position_id": 3
}
```

### Roles

| Method | Endpoint | Deskripsi | Auth Required |
|--------|----------|-----------|---------------|
| POST | `/role` | Buat role baru | ✅ |
| GET | `/role` | Ambil semua roles | ✅ |
| GET | `/role/:id` | Ambil role berdasarkan ID | ✅ |
| PUT | `/role/:id` | Update role | ✅ |
| DELETE | `/role/:id` | Hapus role (soft delete) | ✅ |

**Request Body Role:**
```json
{
  "role_name": "Supervisor"
}
```

### Positions

| Method | Endpoint | Deskripsi | Auth Required |
|--------|----------|-----------|---------------|
| POST | `/position` | Buat position baru | ✅ |
| GET | `/position` | Ambil semua positions | ✅ |
| GET | `/position/:id` | Ambil position berdasarkan ID | ✅ |
| PUT | `/position/:id` | Update position | ✅ |
| DELETE | `/position/:id` | Hapus position (soft delete) | ✅ |

**Request Body Position:**
```json
{
  "position_name": "Senior Developer"
}
```

## 🔐 Autentikasi

Aplikasi ini menggunakan JWT (JSON Web Token) untuk autentikasi.

### Cara Menggunakan Token

1. **Login** untuk mendapatkan token
2. Tambahkan token ke header setiap request yang memerlukan autentikasi:
   ```
   Authorization: Bearer <your-jwt-token>
   ```

### Contoh dengan cURL
```bash
# Login
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"username":"johndoe","password":"password123"}'

# Gunakan token yang didapat
curl -X GET http://localhost:8080/users \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
```

### Contoh dengan Postman
1. Pilih tab **Authorization**
2. Pilih Type: **Bearer Token**
3. Paste token di kolom Token

## 📊 Model Data

### User
```go
type User struct {
    ID        int    `json:"id"`
    Name      string `json:"name"`
    Email     string `json:"email"`
    Username  string `json:"username"`
    Password  string `json:"password"`
    RoleID    int    `json:"role_id"`
    IsDeleted bool   `json:"is_deleted"`
    UpdatedAt string `json:"updated_at"`
    CreatedAt string `json:"created_at"`
}
```

### Employee
```go
type Employee struct {
    ID          int    `json:"id"`
    Name        string `json:"name"`
    PhoneNumber string `json:"phone_number"`
    Email       string `json:"email"`
    PositionID  int    `json:"position_id"`
    IsDeleted   bool   `json:"is_deleted"`
    UpdatedAt   string `json:"updated_at"`
    CreatedAt   string `json:"created_at"`
}
```

### Role
```go
type Role struct {
    ID        int    `json:"id"`
    RoleName  string `json:"role_name"`
    IsDeleted bool   `json:"is_deleted"`
    UpdatedAt string `json:"updated_at"`
    CreatedAt string `json:"created_at"`
}
```

### Position
```go
type Position struct {
    ID           int    `json:"id"`
    PositionName string `json:"position_name"`
    IsDeleted    bool   `json:"is_deleted"`
    UpdatedAt    string `json:"updated_at"`
    CreatedAt    string `json:"created_at"`
}
```

## 🧪 Testing

Untuk testing API, Anda dapat menggunakan:
- **Postman** - [Download](https://www.postman.com/downloads/)
- **Thunder Client** (VS Code Extension)
- **cURL** (Command Line)
- **Insomnia** - [Download](https://insomnia.rest/download)

## 🔧 Development

### Generate Secret Key Baru

Jika Anda ingin menggenerate secret key baru untuk JWT:

```bash
go run generateSecretKey.go
```

Kemudian update `jwtKey` di `handlers/auth.go` dengan key yang baru.

### Hot Reload (Development)

Install Air untuk hot reload:
```bash
go install github.com/cosmtrek/air@latest
air
```

## 📝 Catatan Penting

1. **Security**: 
   - Jangan commit file dengan kredensial database asli
   - Gunakan environment variables untuk konfigurasi sensitif
   - Update JWT secret key untuk production

2. **Soft Delete**: 
   - Semua delete operation adalah soft delete
   - Data tidak benar-benar dihapus, hanya di-mark dengan `is_deleted = true`

3. **Password**: 
   - Password di-hash menggunakan bcrypt
   - Tidak pernah disimpan dalam plain text

## 🤝 Kontribusi

Kontribusi selalu diterima! Silakan:
1. Fork repository ini
2. Buat branch baru (`git checkout -b feature/AmazingFeature`)
3. Commit perubahan Anda (`git commit -m 'Add some AmazingFeature'`)
4. Push ke branch (`git push origin feature/AmazingFeature`)
5. Buat Pull Request

## 📄 Lisensi

Proyek ini dibuat untuk keperluan magang dan pembelajaran.

## 👥 Author

Dikembangkan sebagai bagian dari proyek magang.

## 📞 Kontak & Support

Jika Anda memiliki pertanyaan atau menemukan bug, silakan buat issue di repository ini.

---

**Happy Coding! 🚀**
