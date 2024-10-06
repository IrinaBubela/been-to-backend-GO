# go-been-to

**go-been-to** is a backend web application built with Golang that provides user authentication using JWT (JSON Web Tokens) and MongoDB as the database. It includes routes for user signup, login, and middleware to protect certain routes using JWT authentication.

## Features

- User signup with hashed passwords
- User login with JWT token generation
- User can select and keep list of countries visited by him saved in database
- JWT-based authentication middleware for protecting routes
- MongoDB for storing user information

## Technologies

- **Go**: Backend programming language
- **Gin**: HTTP web framework for Go
- **MongoDB**: NoSQL database for storing user data
- **JWT**: JSON Web Token for user authentication
- **Bcrypt**: Password hashing and comparison

## Requirements

Before you begin, ensure you have the following installed on your system:

- [Go](https://golang.org/doc/install) (version 1.19 or later)
- [MongoDB](https://www.mongodb.com/try/download/community) (local or cloud-based MongoDB instance)

## Installation

### Step 1: Clone the repository

```bash
git clone https://github.com/IrinaBubela/been-to-backend-GO.git
cd go-been-to
