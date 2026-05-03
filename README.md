# 📝 Todo API (Golang + Docker)

Simple RESTful API untuk mengelola todo dengan fitur authentication & authorization berbasis JWT.

---

## 🚀 Features

* 🔐 JWT Authentication (Login & Register)
* 👤 Role-based Authorization (Admin & User)
* ✅ CRUD Todo
* 🐳 Dockerized (multi-container: app + database)
* 📦 Database Migration
* 🔑 Password hashing dengan bcrypt

---

## 🛠 Tech Stack

* Golang (net/http, chi router)
* PostgreSQL
* Docker & Docker Compose
* JWT (github.com/golang-jwt/jwt)
* bcrypt (golang.org/x/crypto)

---

## 🌐 Live Demo

👉 http://YOUR_SERVER_IP:8080

---

## ⚙️ Environment Variables

Buat file `.env`:

```
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=todo_db
DB_HOST=db
DB_PORT=5432
JWT_SECRET=supersecretkey
```

---

## 🐳 Run with Docker

```bash
docker-compose up --build
```

---

## 📦 API Endpoints

### Auth

* POST /auth/register
* POST /auth/login

### Todos

* GET /todos
* GET /todos/{id}
* POST /todos
* PUT /todos/{id}
* DELETE /todos/{id} (Admin only)

---

## 🔐 Authorization

Gunakan header:

```
Authorization: Bearer <your_token>
```

---

## 🧪 Example Request (Login)

```
POST /auth/login
{
  "email": "admin@mail.com",
  "password": "123456"
}
```

---

## 📌 Notes

* Role default: `user`
* Admin bisa delete todo
* Database auto migrate saat container start

---

## 👨‍💻 Author

Dicky Kurniawan
