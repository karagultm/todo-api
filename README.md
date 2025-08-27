# 📝 Todo API (Go + Echo + GORM + SQLite)

A simple **RESTful Todo API** application.  
Developed in Go using the [Echo](https://echo.labstack.com/) framework, with [GORM](https://gorm.io/) integrated for SQLite database support.  

---

## 🚀 Features
- ✅ List all todos (`GET /todos`)  
- ➕ Add a new todo (`POST /todos`)  
- 🔄 Update / toggle todo status (`PUT /todos/:id`)  
- ❌ Delete a todo (`DELETE /todos/:id`)  
- 🗄️ Persistent data storage with SQLite  

---

## 🛠️ Installation

### Requirements
- [Go](https://go.dev/dl/) (1.21+ recommended)  
- [Git](https://git-scm.com/)  

### Steps
```bash
# Clone the repository
git clone https://github.com/karagultm/todo-api.git

# Navigate into the directory
cd todo-api

# Download dependencies
go mod tidy

# Run the application
go run .
```

---

## 📡 API Endpoints

### 1️⃣ Get all todos

```http
GET /todos
```

### 2️⃣ Add a new todo

```http
POST /todos
Content-Type: application/json

{
  "content": "Upload your projects to GitHub!",
  "done": false
}
```

### 3️⃣ Toggle todo status

```http
PUT /todos/1
```

### 4️⃣ Delete a todo

```http
DELETE /todos/1
```

---

## 🛠 Technologies Used

- [Go](https://go.dev/)  
- [Echo Framework](https://echo.labstack.com/)  
- [GORM](https://gorm.io/)  
- [SQLite](https://www.sqlite.org/)

---

📄 For the Turkish version of this README, check [README.tr.md](./README.tr.md) 🇹🇷
