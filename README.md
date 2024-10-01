# TodoList API

This test project is a **Todo List API** built using **Gin** as the web framework and **MongoDB** as the database. It supports basic CRUD operations for managing todos and is structured for scalability and maintainability.


## Table of Contents

 -  [Features](#features)
-   [Prerequisites](#prerequisites)
-   [Installation](#installation)
-   [Environment Variables](#environment-variables)
-   [Run the Application](#run-the-application)
-   [Run the Tests](#run-the-tests)
-   [API Endpoints](#api-endpoints)
-   [Project Structure](#project-structure)
-   [License](#license)

## Features

-   **Create** a todo
-   **Read** a single or all todos
-   **Update** a todo
-   **Delete** a todo
-   RESTful API using **Gin** framework
-   MongoDB for data persistence
-   Environment variables for configuration
-   Test coverage with unit tests

## Prerequisites

Ensure you have the following installed:
-   **Docker** (for running MongoDB and Go project locally)

## Installation

 1. Clone the repository:
	 
        git clone https://github.com/krystian-dajka/Mend-test-todolist.git`
	 
 2. Create .env file
 
    	MONGO_URI=mongodb://<username>:<password>@localhost:27017
    	DB_NAME=TodosDB
	Replace `<username>` and `<password>` with your actual MongoDB credentials.
 4. Build Docker image and launch the container
	 
        docker compose up -d
	 The server should start, and you can interact with it using Postman, Curl, or any API client on: http://localhost:8080

## Run the Tests

Unit tests are written using the Go testing package. You can run the tests using the following command:

    go test ./... -v

## Project Structure

    Mend-test-todolist/
	│
	├── config/
	│   └── db.go                      # Database connection configurations
	│
	├── controllers/
	│   ├── auth/
	│   │   ├── getMe.go               # Get current user details
	│   │   ├── login.go               # Handle user login
	│   │   ├── logout.go              # Handle user logout
	│   │   └── register.go            # Handle user registration
	│   ├── todo/
	│   │   ├── createTodo.go          # Handler to create a new todo
	│   │   ├── deleteTodo.go          # Handler to delete a todo
	│   │   ├── getAllTodos.go         # Handler to fetch all todos
	│   │   └── updateTodo.go          # Handler to update a todo
	│
	├── db_data/                       # Directory for database data
	│
	├── middleware/                    # Directory for middleware (not visible, possibly empty)
	│
	├── models/
	│   ├── todo.go                    # Todo model definition
	│   └── user.go                    # User model definition
	│
	├── routes/
	│   └── routes.go                  # All the application routes
	│
	├── tests/
	│   ├── auth/
	│   │   └── register_test.go       # Unit test for user registration
	│   ├── todo/
	│   │   └── CRUDTodo_test.go       # Unit test for CRUD operations on Todos
	│
	├── util/
	│   └── response.go                # Utility for JSON response formatting
	│
	├── .env                           # Environment variables (database URIs, etc.)
	├── .gitignore                     # Git ignore file
	├── docker-compose.yml             # Docker Compose configuration for the project
	├── Dockerfile                     # Docker configuration for the project
	├── gin-bin/                       # Binary directory for gin (hot-reloading for Go)
	├── go.mod                         # Go module dependencies
	├── go.sum                         # Go module dependency lock file
	└── main.go                        # Entry point for the application


## API Endpoints

### Create a Todo

    POST /api/v1/todos
    

 - **Request Body**:
	 ```
	 {
	  "title": "New Todo",
	  "description": "Todo description",
	  "done": false
	}
	```

  
  ### Get All Todos

    GET /api/v1/todos
    
    
  ### Update a Todo

    PUT /api/v1/todos/:id
    
   - **Request Body**:
	 ```
	 {
	  "title": "New Todo",
	  "description": "Todo description",
	  "done": false
	 }
	 ```
   
   
  ### Delete a Todo

    DELETE /api/v1/todos/:id
