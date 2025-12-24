# Student Management REST API

A robust RESTful API built with **Go (Golang)** designed to manage student records. This project demonstrates the implementation of a production-ready **Standard Go Project Layout**, focusing on modularity, clean architecture, and scalability.

## 🚀 Features

* **Create Student:** Register new students with details (Name, Email, Age, etc.).
* **Read Student:** Retrieve specific student details by ID or list all students.
* **Update Student:** Modify existing student records.
* **Delete Student:** Remove student records from the database.
* **Modular Architecture:** Clean separation of concerns (Handlers, Models, Storage).
* **Configuration:** Externalized configuration management.

## 🛠️ Tech Stack

* **Language:** Go (Golang)
* **Router:** `net/http` (Standard Lib) *[Update if you used Chi/Gin/Mux]*
* **Database:** *[e.g., SQLite / PostgreSQL / In-Memory]*
* **Config:** YAML/JSON based configuration

## 📂 Project Structure

This project follows the [Standard Go Project Layout](https://github.com/golang-standards/project-layout):

```text
studentsRestAPI/
├── cmd/
│   └── studentsRestAPI/
│       └── main.go       # Entry point of the application
├── config/               # Configuration files and logic
├── internal/             # Private application code
│   ├── models/           # Data structures
│   ├── handlers/         # HTTP request handlers (Controllers)
│   └── storage/          # Database operations (Repository)
├── go.mod                # Module definition
└── README.md             # Project documentation
