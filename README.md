# Go system auth

## EndPoints

| Methods | Urls           | Actions                   |
| :------ | :------------- | :------------------------ |
| POST    | /auth/signUp   | register user             |
| GET     | /auth/enable   | anable user               |
| POST    | /auth/signIn   | login user and return jwt |
| GET     | /products/save | testing auth              |

## Installation

To use this project, you need to follow these steps:

1. Clone the repository: `git clone https://github.com/RenanAlmeida225/go-system-auth.git`
2. Install the dependencies: `go mod download`
3. Build the application: `go build`
4. Run the application: `./main`

## Used Tools

This project uses the following tools:

- [Golang](https://golang.org/) for backend development
- [Go-Gin](https://github.com/gin-gonic/gin) for route management
- [GoORM](https://gorm.io/) for database communication
- [Postgres](https://github.com/go-gorm/postgres) as the database
- [JWT](https://github.com/golang-jwt/jwt/v5) for generate and verify JWT
- [Uuid](https://github.com/google/uuid) for generate uuid
- [Crypto](https://golang.org/x/crypto) for generate hash
- [Godotenv](https://github.com/joho/godotenv) for environment variables
