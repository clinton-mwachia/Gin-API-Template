# Task Manager API (Golang + Gin + MongoDB)

## 📌 Overview

This is an **API** built using **Golang**, **Gin framework**, and **MongoDB**. It supports user authentication, authorization, and full CRUD operations for tasks and users. Custom validation is implemented without external libraries.

## 📂 Project Structure

```
├── handlers/           # Contains API route handlers
│   ├── taskHandler.go  # CRUD operations for tasks
│   ├── userHandler.go  # CRUD operations for users
│
├── middleware/         # Authentication middleware
│   ├── authMiddleware.go
│
├── models/             # Database models
│   ├── task.go
│   ├── user.go
│
├── helpers/            # Custom helpers
│   ├── validator.go    # Custom validators
│
├── utils/              # Utility functions
│   ├── database.go     # MongoDB connection setup
│   ├── jwt.go          # JWT token generation and validation
│
├── routes/             # API route definitions
│   ├── routes.go
│
├── .env                # Environment variables
├── main.go             # Entry point of the application
└── README.md           # Documentation
```

## 🚀 Features

✅ **User Authentication & Authorization** using JWT  
✅ **CRUD for Tasks & Users** (Create, Read, Update, Delete)  
✅ **Pagination for Tasks**  
✅ **MongoDB Integration**  
✅ **Custom Validators (No external packages)**

## 🛠 Setup Instructions

### 1️⃣ Install Dependencies

Ensure you have Go installed. Then, run:

```sh
go clone https://github.com/clinton-mwachia/Gin-API-Template.git
```

```sh
go mod tidy
```

### 2️⃣ Setup Environment Variables

Create a `.env` file in the root directory and add:

```env
MONGO_URI=mongodb://localhost:27017
dbName=db_name
JWT_SECRET=your_secret_key
```

### 3️⃣ Run the Application

```sh
go run main.go
```

The server will start at **`http://localhost:8080`**

## 🔥 API Endpoints

### Authentication

| Method | Endpoint    | Description       |
| ------ | ----------- | ----------------- |
| POST   | `/login`    | User login        |
| POST   | `/register` | User registration |

### Tasks

| Method | Endpoint           | Description         |
| ------ | ------------------ | ------------------- |
| POST   | `/task`            | Create a new task   |
| GET    | `/users`           | Get all tasks       |
| GET    | `/tasks/paginated` | Get paginated tasks |
| GET    | `/task/:id`        | Get task by ID      |
| PUT    | `/task/:id`        | Update a task       |
| DELETE | `/task/:id`        | Delete a task       |

### Users

| Method | Endpoint    | Description    |
| ------ | ----------- | -------------- |
| GET    | `/user/:id` | Get user by ID |
| DELETE | `/user/:id` | Delete a user  |
| UPDATE | `/user/:id` | Update a user  |
| GET    | `/users`    | Get all user   |

## 📌 Contribution

Feel free to fork and contribute! Submit a pull request with your updates.

## 📜 License

This project is open-source under the MIT License.
