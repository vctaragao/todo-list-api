# To-do list in Go

### Description

This app has the porpuse to be used for the study of the Go programming language and the go way of doing things.

## Features To be Implemented

### DOMAIN - CRUD Task

- [x] Create Task
  - [x] Functional Test (Integrating route and database)
  - [x] Unit test for the Use Case
- [ ] Remove Task
  - [ ] Functional Test (Integrating route and database)
  - [ ] Unit test for the Use Case
- [ ] Update Task
  - [ ] Functional Test (Integrating route and database)
  - [ ] Unit test for the Use Case
- [ ] List Tasks
  - [ ] Functional Test (Integrating route and database)
  - [ ] Unit test for the Use Case
- [ ] Show Task
  - [ ] Functional Test (Integrating route and database)
  - [ ] Unit test for the Use Case

### Storage

- [ ] Persist the informations into a MySql or Postgress database
- [x] Use Repositories to create the abstraction layer between the Core Model and the db

### Http

- [x] Use [Echo](https://echo.labstack.com/) as the framework for the web/http layer
- [ ] Add Validation to the routes

### Testing

- [ ] Create a Functional testsuite that roolbacks the database after every test

## Architeture/Desing

The ideia of the archtecture/desing of this project is focus um using three main principels:

- Hexagonal Architecture
- Domain Driven Desing
- Screamming Architecture

### Hexagonal Architecture

Its being used as a guidence for the folder struct of the project:

![Onion representation of the archtecture of the project](https://user-images.githubusercontent.com/26884793/219226513-28d48bfe-2cc1-4112-a720-f9a438890ae8.png)

### Folder Struct

```bash
.
├── api
│   └── http
│       ├── create_task.go
│       └── router.go
├── cmd
│   └── server
│       └── main.go
├── go.mod
├── go.sum
├── internal
│   ├── application
│   │   ├── create_task
│   │   │   ├── create.go
│   │   │   ├── create_task_test.go
│   │   │   └── task_dto.go
│   │   ├── entity
│   │   │   └── task.go
│   │   └── repository.go
│   └── todo_list.go
├── Makefile
├── README.md
├── storage
│   ├── dummy_adapter.go
│   ├── schemas
│   │   └── task.go
│   └── sqlite_adapter.go
├── tests
│   └── functional
│       ├── create_task_test.go
│       └── test_app.go
└── todo_list.db
```

- **Framewors layer**
  - `/api`: Folder that will contain the different entrypoints of the project (http, cli, grpc, etc)
  - `/storage`: Folder that will container the differentes ways of persistance that the project has (Mysql, Memory, Redis, etc.)
  - `/cmd`: Folder for holding the main files of the project.
  - `/tests`: Folder for the tests that verify the funcionalities of the application integrating all of its layers
  - `/tests/functional`: Functional tests
- **Application Layer**
  - `internal/todo_list.go`: File to serve as a Facade for the application as the name follows the name of the project
  - `/internal/application/*.go`: Files to be for general use inside the application layer (repository.go, etc.)
  - `/internal/application/create_trask`: Use-case for the creation of a task
- **Domain Layer**
  - `/internal/application/entity`: Folder holding the Entityes of the project

### Domain Driven Desing

Its being used as a guidance for the files organization and desing of the project. Patterns being used by proposed by DDD

- Entity
- Repositories
- Domain Services

### Screaming Archtecture

The application of the Scremming Architeture is being mainly used in the naming of the usecases folder inside the Application Layer, and in its files.
