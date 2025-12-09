# Go Zakat - Zakat Management System API

RESTful API untuk sistem manajemen Zakat, Infaq, dan Sadaqah (ZIS) yang dibangun dengan Go, Gin, dan PostgreSQL.

## 🚀 Tech Stack

- **Go** 1.25
- **Gin** - HTTP web framework
- **PostgreSQL** - Database with pgx driver
- **JWT** - Authentication (Access + Refresh Token)
- **Google OAuth2** - Social login (Web & Mobile)
- **Docker** - Containerization Support

## 📋 Features

### ✅ Main Features
- **Authentication**: Register, Login, Google OAuth2, JWT Refresh Token.
- **Master Data**: Manajemen Muzakki, Asnaf (8 Golongan), Mustahiq, dan Program Penyaluran.
- **Transactions**: Pencatatan Penerimaan (Zakat/Infaq/Sadaqah) dan Penyaluran Dana.
- **Reports**: Laporan Penghimpunan, Penyaluran, dan Saldo Dana (Income, Distribution, Fund Balance).
- **API Documentation**: Terintegrasi dengan Swagger UI.

## 🏗️ Project Structure

```
go-zakat/
├── cmd/
│   └── api/
│       └── main.go                 # Application entry point
├── internal/
│   ├── delivery/
│   │   └── http/
│   │       ├── dto/                # Data Transfer Objects
│   │       ├── handler/            # HTTP handlers
│   │       └── middleware/         # Middleware (auth, cors, etc)
│   ├── domain/
│   │   ├── entity/                 # Domain entities
│   │   └── repository/             # Repository interfaces
│   ├── infrastructure/
│   │   ├── database/               # Database connection
│   │   ├── oauth/                  # OAuth state management
│   │   └── service/                # External services (Google, JWT)
│   ├── repository/
│   │   └── postgres/               # PostgreSQL implementations
│   └── usecase/                    # Business logic
├── migrations/                     # Database migrations (9 migrations)
├── pkg/
│   ├── config/                     # Config implementations
│   ├── database/                   # Database implementations
│   ├── logger/                     # Logger implementations
│   └── response/                   # Standardized API responses
├── docs/                           # Swagger documentation
├── .env                            # Environment variables
└── go.mod                          # Go dependencies
```

## 🛠️ Setup

Direkomendasikan menggunakan Docker untuk kemudahan instalasi dan deployment.

### Prerequisites
- Docker & Docker Compose
- Git

### Installation via Docker

1. **Clone repository**
   ```bash
   git clone https://github.com/Go-Zakat/go-zakat-be.git
   cd go-zakat-be
   ```

2. **Setup environment variables**
   ```bash
   cp .env_example .env
   ```
   Sesuaikan konfigurasi di `.env` (Database credentials di env akan diabaikan jika menggunakan default docker-compose, namun tetap diperlukan untuk aplikasi).


3. **Run Application**
   ```bash
   docker-compose up -d --build
   ```

   Aplikasi akan berjalan di: `http://localhost:8080`
   Database PostgreSQL akan berjalan di port `5432`

## 📚 API Documentation

Swagger documentation tersedia di: `http://localhost:8080/swagger/index.html`

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License
This project is licensed under the MIT License.

## 👥 Authors
- Muhammad Dila