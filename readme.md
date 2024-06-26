# Meetup style events REST API Demo

## Overview
This is a demo project showcasing a RESTful API for managing events and users for a meetup style app. The project aims to demonstrate essential skills in developing secure and efficient web APIs using Go, along with JWT-based authentication, database interactions, and middleware usage.

## Features
- User management
    - Create user accounts with email and password
    - Securely hash and store passwords
    - Authenticate users using JWT tokens
- Event Management
    - Create, update delete and retrieve events
    - Access control is enforced (only users can update and delete their own events)
- Event registration
    - Users can sign up for events and cancel their registrations
    - Maintain a many-to-many relationship between users and events
- Authentication
    - Use JWT tokens for secure user authentication
    - Middleware to verify and authorize user requests

## Technologies used

- **Go**: Programming language
- **Gin**: HTTP web framework
- **SQLite**: Database
- **JWT**: JSON Web Tokens for authentication
- **Bcrypt**: Password hashing

## Project structure

```
├── db
│   └── db.go          // Database initialization and table creation
├── middleware
│   └── auth.go        // JWT authentication middleware
├── models
│   ├── event.go       // Event model and database interactions
│   └── user.go        // User model and database interactions
├── routes
│   ├── events.go      // Handlers for event-related API endpoints 
|   ├── register.go    // Handlers for managing registration logic
|   ├── routes.go      // API route handlers
|   └── users.go       // Handlers for user-related API endpoints
├── utils
|   ├── hash.go        // Hashing and comparing hashed passwords
|   └── jwt.go         // Generating and verifying jwt tokens
├── api.db             // SQLite database file
└── main.go            // Entry point of the application
```

## Getting started

### Prerequisite
- Go 1.2.2 or higher installed on your machine
- SQLite installed on your machine

### Installation
1. Clone the repository:
```
git clone https://github.com/JohnPalmgren/events-app-api.git
cd events-app-api
```

2. Install dependencies
```
go mod tidy
```

3. Set up environment variables
- Create a `.env` file in the root directory of your project
- Add the following line to the `.env` file:
```
JWT_KEY=your_secret_key
```

4. Run the application

```
go run main.go
```

## Conclusion
This demo project demonstrates the implementation of a secure RESTful API for the core functionally of an events / groups application using Go. It covers key concepts such as JWT authentication, password hashing, and database interactions. 