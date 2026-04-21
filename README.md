โปรเจคนี้พัฒนาด้วยโครงสร้างแบบ hexagonal architecture เพื่อสำหรับรองรับการขยายต่อและสามารถต่อยอดได้


└── 📁 golangTest
    ├── 📁 adapter
    │   ├── 📁 handler
    │   │   ├── 🐹 response.go /// helper response + errpr handler function
    │   │   ├── 🐹 taskHandler.go /// http adapter
    │   │   └── 🐹 taskHandler_test.go
    │   └── 📁 repository
    │       ├── 🐹 taskRepositorySql.go ///database adapter
    │       └── 🐹 taskRepository_test.go
    ├── 📁 core
    │   ├── 📁 entity
    │   │   └── 🐹 task.go /// base entity 
    │   ├── 📁 port
    │   │   └── 🐹 taskPort.go
    │   └── 📁 service
    │       ├── 🐹 taskService.go /// bussiness logic
    │       └── 🐹 taskService_test.go
    ├── 📁 pkg
    │   └── 📁 errs
    │       └── 🐹 errs.go /// errors
    ├── 📁 routes
    │   └── 🐹 routes.go
    ├── 🐳 Dockerfile
    ├── ⚙️ docker-compose.yml
    ├── 📄 go.mod
    ├── 📄 go.sum
    ├── 🐹 main.go
    └── 📄 tasks.db

Running Project
    1.use go command
        cd เข้า folder ของ gotest
        ใช้คำสั่ง go run main.go หรือ go run .

        หรืออีกวิธีคือ go build -o main.exe main.go เพื่อสร้าง file สำหรับการใช้งานโดยที่ไม่ต้อง compile ซ้ำ
        จากนั้นใข้คำสั่ง ./main.exe เพื่อรันโปรเจค

    2.docker 
        cd เข้า folder ของ gotest
        ใช้คำสั่ง docker compose up --build 

api 

# ============ CREATE TASK ============
+ curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Complete Project",
    "description": "Finish the Go project with tests and Docker",
    "assign_name": "John"
  }'

 response {
    "success": true,
    "message": "task created successfully",
    "data": {
        "title": "Complete Project",
        "description": "Finish the Go project with tests and Docker",
        "assign_name": "John"
    }
}

# ============ GET ALL TASKS ============
+ curl -X GET "http://localhost:8080/tasks?page=1&limit=5" \
  -H "Content-Type: application/json"

  response {
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

# ============ GET TASKS BY ASSIGN_NAME ============
+ curl -X GET "http://localhost:8080/tasks?assign_name=Alice&page=1&limit=5" \
  -H "Content-Type: application/json"

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

# ============ GET TASKS BY STATUS ============
+ curl -X GET "http://localhost:8080/tasks?status=todo&page=1&limit=5" \
  -H "Content-Type: application/json"

response {
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

# ============ GET SINGLE TASK ============
+ curl -X GET "http://localhost:8080/tasks/e7f86858-2be6-437b-9347-f3e749c3e40f" 
  -H "Content-Type: application/json"

response {
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

# ============ UPDATE TASK ============
+ curl -X PUT "http://localhost:8080/tasks/e7f86858-2be6-437b-9347-f3e749c3e40f" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Updated Task Title",
    "description": "Updated description",
    "assign_name": "UpdatedName"
  }'

response {"success":true,"message":"Success","data":"Task updated successfully"}

# ============ UPDATE TASK STATUS ============
# Status can be: pending, in_progress, done
+ curl -X PATCH "http://localhost:8080/tasks/e7f86858-2be6-437b-9347-f3e749c3e40f/in_progress" \    
  -H "Content-Type: application/json"

response {"success":true,"message":"Success","data":"Task updated successfully"}

# ============ TEST PAGINATION ============
+ curl -X GET "http://localhost:8080/tasks?page=1&limit=3" \
  -H "Content-Type: application/json"

response {
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


# ============ FILTER COMBINATIONS ============
+ curl -X GET "http://localhost:8080/tasks?assign_name=Bob&status=done&page=1&limit=5" \
  -H "Content-Type: application/json" 

response {
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

=================== At least 2 assumptions you made ===============================

1. เมื่อระบบมีการขนายเพิ่มเช่น user_id ที่จะใช้เป็น id อ้างอิงกับ task ในสถานการณ์อย่างการมอบหมายงานให้แต่ละบุคคลสามารถติดตามได้ง่าย

2. เมื่อการใช้งานโดยบุคคลผู้เป็นเจ้าของ task อาจจะต้องเพิ่มการ login และ role สำหรับสิทธิ่ในการเข้าถึง


================== At least 1 thing you'd improve with more time ==================

 - ทำ validate เนื่องในชั้นของ handler ได้มีการใช้คำสั่ง ShouldBindJSON จึงเป็นการ validate input ไปในตัว โดยvalidate ที่อยากทำเช่น assing name ที่ยอมแค่ตัวอักษร
 - ทำ response เพิ่่มเติม เช่นจำยนวนของ status = todo หรือ จำนวนของtaskตาม assign_name เพราะอาจจะนำไปใช้ต่อให้UIได้ 