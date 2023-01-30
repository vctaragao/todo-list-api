# To-do list in Go

### Description

This app has the porpuse to be used for the study of the Go programming language and the go way of doing things.

## Features To be Implemented

### DOMAIN - CRUD Task

- [x] Create Task
- [ ] Remove Task
- [ ] Update Task
- [ ] List Tasks
- [ ] Show Task

### Storage

- [ ] Persist the informations into a MySql or Postgress database
- [ ] Use Repositories to create the abstraction layer between the Core Model and the db

### Http

- [x] Use [Echo](https://echo.labstack.com/) as the framework for the web/http layer

## Architeture

The ideia of the archtecture of this project is focus um using three main principels:

- Hexagonal Architecture
- Domain Driven Desing
- Screamming Architecture

### Hexagonal Architecture

Its being used as a guidence for the folder struct of the project:

![Captura de tela de 2023-01-29 20-22-13](https://user-images.githubusercontent.com/26884793/215372234-8433523b-5082-4335-854f-370a59c95586.png)

### Folder Struct
```bash
.
├── api
│   └── http
├── cmd
│   └── server
├── internal
│   ├── create
│   └── entity
└── storage
```
- **Framewors layer**
  - `/api`: Folder that will contain the different entrypoints of the project (http, cli, grpc, etc)
  - `/storage`: Folder that will container the differentes ways of persistance that the project has (Mysql, Memory, Redis, etc.)
  - `/cmd`: Folder for holding the main files of the project.
- **Application Layer**
  - `/internal/create`: Use-case for the creation of a task
  - `/internal/*.go`: Files to be for general use inside the application layer (repository.go, etc.)
- **Domain Layer**
  - `/internal/entity`: Folder holding the Entityes of the project 

