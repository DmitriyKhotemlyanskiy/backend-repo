# backend-repo
# Booking System — Backend Service

Backend REST API service built with **Go (Golang)**, **Gin framework**, and **MongoDB**.

## 🚀 Tech Stack
- **Language:** Go 1.21+
- **Framework:** Gin (`github.com/gin-gonic/gin`)
- **Database Driver:** MongoDB Go Driver (`go.mongodb.org/mongo-driver`)

## 📌 API Endpoints
- `GET /api/hotels` — Retrieve list of available hotels (including details & descriptions).
- `POST /api/reservations` — Create a new hotel reservation.
- `GET /api/reservations/lookup?search=<email_or_name>` — Search reservations.
- `DELETE /api/reservations/:id` — Cancel an existing reservation.

## 🐳 Docker Build
To build the backend image manually:
```bash
docker build -t dimakhot/backend-app:latest .

- **The server run on port:** 8085
