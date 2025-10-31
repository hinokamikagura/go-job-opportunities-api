# 🚀 GoJob Opportunities API

**GoJob Opportunities API** is a modern, production-ready backend built with **Golang**, designed to manage and serve job listings efficiently.
The project demonstrates professional backend practices — leveraging **Gin** for routing, **GORM** for ORM and database interactions, **SQLite** for lightweight persistence, and **Swagger** for self-documented API endpoints.

---

## ✨ Key Features

* 🔧 Built with **Golang**, one of the highest-paying and most in-demand backend languages
* ⚡ **Gin** framework for fast, efficient, and clean HTTP routing
* 🧱 **GORM + SQLite** for reliable data persistence with minimal setup
* 📘 **Swagger** integration for real-time API documentation and testing
* 🗂️ Modular, maintainable project structure following Go best practices
* 🔍 Full CRUD operations for managing job opportunities
* ✅ Automated testing to ensure stability and performance
* 🐳 **Docker & Docker Compose** support for rapid deployment

---

## ⚙️ Installation & Setup

### 1️⃣ Clone the repository

```bash
git clone https://github.com/hinokamikagura/go-job-opportunities-api.git
cd go-job-opportunities-api
```

### 2️⃣ Install dependencies

```bash
go mod download
```

### 3️⃣ Build the application

```bash
go build
```

### 4️⃣ Run the API

```bash
./main
```

The API runs by default on **[http://localhost:8080](http://localhost:8080)**

---

## 🧰 Makefile Commands

This project includes a `Makefile` to simplify common development tasks:

| Command              | Description                                        |
| -------------------- | -------------------------------------------------- |
| `make run`           | Run the API without regenerating Swagger docs      |
| `make run-with-docs` | Generate Swagger docs and run the API              |
| `make build`         | Compile and build an executable (`gopportunities`) |
| `make test`          | Run all Go tests                                   |
| `make docs`          | Generate Swagger documentation                     |
| `make clean`         | Remove build artifacts and documentation           |

Example:

```bash
make run
```

---

## 🐳 Docker & Deployment

Containerize and deploy effortlessly using Docker or Docker Compose.

### Docker

```bash
docker build -t gojob-api .
docker run -p 8080:8080 -e PORT=8080 gojob-api
```

### Docker Compose

```bash
docker compose build
docker compose up
```

To stop and clean up:

```bash
docker compose down
```

---

## 🧪 API Documentation

Once running, Swagger UI is available at:

```
http://localhost:8080/swagger/index.html
```

You can use it to test, visualize, and explore all API endpoints interactively.

---

## 🧩 Technologies Used

* **Go (Golang)** — backend language
* **Gin** — high-performance HTTP router
* **GORM** — ORM for Go
* **SQLite** — embedded database
* **Swagger (swaggo)** — auto-generated API docs
* **Docker** — containerization

---

## 🤝 Contributing

Contributions are welcome!
Please follow the standard Git workflow:

1. Fork the repository
2. Create a new branch:

   ```bash
   git checkout -b feature/your-feature-name
   ```
3. Commit changes using **Conventional Commits**
4. Push your branch and open a **Pull Request**

---

## ⚖️ License

Licensed under the **MIT License** — see [`LICENSE.md`](LICENSE.md) for full details.

---

## 👨‍💻 Author

Created by **[CryptoSlayer](https://github.com/hinokamikagura)**

> “Clean, modular Go code built for scalability and learning.”
