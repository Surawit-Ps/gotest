# Task Management API

โปรเจคนี้พัฒนาด้วยโครงสร้างแบบ **Hexagonal Architecture** เพื่อสำหรับรองรับการขยายต่อและสามารถต่อยอดได้

---

## 📁 Project Structure

```
golangTest/
├── adapter/              # External adapters
│   ├── handler/         # HTTP handler (input adapter)
│   │   ├── response.go
│   │   ├── taskHandler.go
│   │   └── taskHandler_test.go
│   └── repository/      # Database adapter (output adapter)
│       ├── taskRepositorySql.go
│       └── taskRepository_test.go
├── core/                # Business logic (framework independent)
│   ├── entity/         # Domain models
│   ├── port/           # Interfaces (contracts)
│   └── service/        # Business logic
├── pkg/                 # Shared packages
│   └── errs/           # Error handling
├── routes/              # Route configuration
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
├── main.go
└── tasks.db
```

---

## 🚀 Running Project

### วิธีที่ 1: Using Go Command
```bash
cd golangTest
go run main.go
# หรือ
go run .
```

### วิธีที่ 2: Build then Run
```bash
go build -o main.exe main.go
./main.exe
```

### วิธีที่ 3: Using Docker
```bash
cd golangTest
docker compose up --build
```

---

## 📚 API Documentation 

---

## 📚 API Documentation

### 1️⃣ CREATE TASK

**Request:**
```bash
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Complete Project",
    "description": "Finish the Go project with tests and Docker",
    "assign_name": "John"
  }'
```

**Response:**
```json
{
  "success": true,
  "message": "task created successfully",
  "data": {
    "title": "Complete Project",
    "description": "Finish the Go project with tests and Docker",
    "assign_name": "John"
  }
}
```

---

### 2️⃣ GET ALL TASKS (with pagination)

**Request:**
```bash
curl -X GET "http://localhost:8080/tasks?page=1&limit=5" \
  -H "Content-Type: application/json"
```

**Response:**
```json
{
  "success": true,
  "message": "Success",
  "data": [
    {
      "id": "db76290e-f52a-475b-b0a4-27e8d75cb1a8",
      "title": "Task 1",
      "description": "Description for Task 1",
      "status": "todo",
      "assign_name": "Alice",
      "created_at": "2026-04-22T01:34:43.0760504+07:00",
      "update_at": "2026-04-22T01:34:43.0760504+07:00"
    },
    {
      "id": "e7f86858-2be6-437b-9347-f3e749c3e40f",
      "title": "Task 2",
      "description": "Description for Task 2",
      "status": "todo",
      "assign_name": "Bob",
      "created_at": "2026-04-22T01:34:43.0786007+07:00",
      "update_at": "2026-04-22T01:34:43.0786007+07:00"
    },
    {
      "id": "8a8b7a5c-1d02-439e-80dd-88e2615dafb2",
      "title": "Task 3",
      "description": "Description for Task 3",
      "status": "todo",
      "assign_name": "Charlie",
      "created_at": "2026-04-22T01:34:43.0807149+07:00",
      "update_at": "2026-04-22T01:34:43.0807149+07:00"
    },
    {
      "id": "0f40ee49-a29f-4434-b5c8-5ffa8a0fe968",
      "title": "Task 4",
      "description": "Description for Task 4",
      "status": "todo",
      "assign_name": "David",
      "created_at": "2026-04-22T01:34:43.0822491+07:00",
      "update_at": "2026-04-22T01:34:43.0822491+07:00"
    },
    {
      "id": "57355b74-0a1c-4ff4-92c7-ed43cfa3217a",
      "title": "Task 5",
      "description": "Description for Task 5",
      "status": "todo",
      "assign_name": "Eve",
      "created_at": "2026-04-22T01:34:43.0842885+07:00",
      "update_at": "2026-04-22T01:34:43.0842885+07:00"
    }
  ]
}
```

---

### 3️⃣ GET TASKS BY ASSIGN_NAME

**Request:**
```bash
curl -X GET "http://localhost:8080/tasks?assign_name=Alice&page=1&limit=5" \
  -H "Content-Type: application/json"
```

**Response:**
```json
{
  "success": true,
  "message": "Success",
  "data": [
    {
      "id": "db76290e-f52a-475b-b0a4-27e8d75cb1a8",
      "title": "Task 1",
      "description": "Description for Task 1",
      "status": "todo",
      "assign_name": "Alice",
      "created_at": "2026-04-22T01:34:43.0760504+07:00",
      "update_at": "2026-04-22T01:34:43.0760504+07:00"
    },
    {
      "id": "cc14a8ee-a365-4ccd-b22b-b67f6a9509ef",
      "title": "Task 27",
      "description": "Description for Task 27",
      "status": "todo",
      "assign_name": "Alice",
      "created_at": "2026-04-22T01:34:43.1260996+07:00",
      "update_at": "2026-04-22T01:34:43.1260996+07:00"
    }
  ]
}
```

---

### 4️⃣ GET TASKS BY STATUS

**Request:**
```bash
curl -X GET "http://localhost:8080/tasks?status=todo&page=1&limit=5" \
  -H "Content-Type: application/json"
```

**Response:**
```json
{
  "success": true,
  "message": "Success",
  "data": [
    {
      "id": "db76290e-f52a-475b-b0a4-27e8d75cb1a8",
      "title": "Task 1",
      "description": "Description for Task 1",
      "status": "todo",
      "assign_name": "Alice",
      "created_at": "2026-04-22T01:34:43.0760504+07:00",
      "update_at": "2026-04-22T01:34:43.0760504+07:00"
    },
    {
      "id": "e7f86858-2be6-437b-9347-f3e749c3e40f",
      "title": "Task 2",
      "description": "Description for Task 2",
      "status": "todo",
      "assign_name": "Bob",
      "created_at": "2026-04-22T01:34:43.0786007+07:00",
      "update_at": "2026-04-22T01:34:43.0786007+07:00"
    }
  ]
}
```

---

### 5️⃣ GET SINGLE TASK

**Request:**
```bash
curl -X GET "http://localhost:8080/tasks/e7f86858-2be6-437b-9347-f3e749c3e40f" \
  -H "Content-Type: application/json"
```

**Response:**
```json
{
  "success": true,
  "message": "Success",
  "data": {
    "id": "e7f86858-2be6-437b-9347-f3e749c3e40f",
    "title": "Task 2",
    "description": "Description for Task 2",
    "status": "todo",
    "assign_name": "Bob",
    "created_at": "2026-04-22T01:34:43.0786007+07:00",
    "update_at": "2026-04-22T01:34:43.0786007+07:00"
  }
}
```

---

### 6️⃣ UPDATE TASK

**Request:**
```bash
curl -X PUT "http://localhost:8080/tasks/e7f86858-2be6-437b-9347-f3e749c3e40f" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Updated Task Title",
    "description": "Updated description",
    "assign_name": "UpdatedName"
  }'
```

**Response:**
```json
{
  "success": true,
  "message": "Success",
  "data": "Task updated successfully"
}
```

---

### 7️⃣ UPDATE TASK STATUS

**Request:**
```bash
curl -X PATCH "http://localhost:8080/tasks/e7f86858-2be6-437b-9347-f3e749c3e40f/in_progress" \
  -H "Content-Type: application/json"
```

**Status Options:** `todo`, `in_progress`, `done`

**Response:**
```json
{
  "success": true,
  "message": "Success",
  "data": "Task updated successfully"
}
```

---

### 8️⃣ TEST PAGINATION

**Request:**
```bash
curl -X GET "http://localhost:8080/tasks?page=1&limit=3" \
  -H "Content-Type: application/json"
```

**Response:**
```json
{
  "success": true,
  "message": "Success",
  "data": [
    {
      "id": "db76290e-f52a-475b-b0a4-27e8d75cb1a8",
      "title": "Task 1",
      "description": "Description for Task 1",
      "status": "todo",
      "assign_name": "Alice",
      "created_at": "2026-04-22T01:34:43.0760504+07:00",
      "update_at": "2026-04-22T01:34:43.0760504+07:00"
    },
    {
      "id": "e7f86858-2be6-437b-9347-f3e749c3e40f",
      "title": "Updated Task Title",
      "description": "Updated description",
      "status": "todo",
      "assign_name": "UpdatedName",
      "created_at": "2026-04-22T01:34:43.0786007+07:00",
      "update_at": "2026-04-22T01:47:52.7328198+07:00"
    },
    {
      "id": "8a8b7a5c-1d02-439e-80dd-88e2615dafb2",
      "title": "Task 3",
      "description": "Description for Task 3",
      "status": "todo",
      "assign_name": "Charlie",
      "created_at": "2026-04-22T01:34:43.0807149+07:00",
      "update_at": "2026-04-22T01:34:43.0807149+07:00"
    }
  ]
}
```

---

### 9️⃣ FILTER COMBINATIONS

**Request:**
```bash
curl -X GET "http://localhost:8080/tasks?assign_name=Bob&status=done&page=1&limit=5" \
  -H "Content-Type: application/json"
```

**Response:**
```json
{
  "success": true,
  "message": "Success",
  "data": [
    {
      "id": "54f36897-62e5-4379-a702-f7f57c2283b7",
      "title": "Task 28",
      "description": "Description for Task 28",
      "status": "todo",
      "assign_name": "Bob",
      "created_at": "2026-04-22T01:34:43.1289464+07:00",
      "update_at": "2026-04-22T01:34:43.1289464+07:00"
    }
  ]
}
```

---

## 💡 Assumptions

1. เมื่อระบบมีการขนายเพิ่มเช่น user_id ที่จะใช้เป็น id อ้างอิงกับ task ในสถานการณ์อย่างการมอบหมายงานให้แต่ละบุคคลสามารถติดตามได้ง่าย

2. เมื่อการใช้งานโดยบุคคลผู้เป็นเจ้าของ task อาจจะต้องเพิ่มการ login และ role สำหรับสิทธิ่ในการเข้าถึง


================== At least 1 thing you'd improve with more time ==================

 - ทำ validate เนื่องในชั้นของ handler ได้มีการใช้คำสั่ง ShouldBindJSON จึงเป็นการ validate input ไปในตัว โดยvalidate ที่อยากทำเช่น assing name ที่ยอมแค่ตัวอักษร
 - ทำ response เพิ่่มเติม เช่นจำยนวนของ status = todo หรือ จำนวนของtaskตาม assign_name เพราะอาจจะนำไปใช้ต่อให้UIได้ 