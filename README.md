# Expreso 🚗

Expreso is a ride-sharing platform built for Nicaragua.

This is the main project repository. It will eventually contain:

- `api/` – The backend API (Go + Gin)
- `mobile/` – Mobile app [coming soon]
- `web/` – Admin dashboard [coming soon]

---

## 🐳 Running the Project

To start the API and PostgreSQL database using Docker:

```bash
docker compose up --build -d
```

This will:

- Start the PostgreSQL database
- Build and run the Go API (in api/)
- Create tables automatically via GORM

✅ [Visit: http://localhost:8080](http://localhost:8080)

---

## ⚙️ Environment Configuration

Create a .env file inside the api/ folder:

```bash
PORT=8080
DB_HOST=db
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=expreso
JWT_SECRET=supersecretkey
```

You can copy the default template:
```bash
cp api/.env.example api/.env
```

---

## 📘 Subproject Docs

- [api/README.md](./api/README.md) – API setup, structure, and endpoints

---

## 🧑‍💻 Authors

Made with 💙 by the Expreso Team
