# Expreso API 🚗

Expreso is a ride-sharing API built in Go using Gin and PostgreSQL.

## 🛠 Tech Stack

- **Go** 1.24+
- **Gin** (HTTP framework)
- **GORM** (ORM for PostgreSQL)
- **PostgreSQL** (via Docker)
- **Docker** & **Docker Compose**
- **.env** for environment variables

---

## 🚀 Getting Started

### 📦 Requirements

- [Go](https://go.dev/dl/)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

### Folder Structure

```bash
api/
├── main.go                # Entry point
├── controllers/           # HTTP handlers
├── models/                # GORM models
├── config/                # DB connection and other config
├── utils/                 # Helper functions (e.g. env)
├── .env.example           # Environment template
└── Dockerfile             # Builds the Go app
