# Sypnasis Golang Test E-commerce

This project is a simple e-commerce backend built with Go, utilizing the Gin framework for handling HTTP requests, Gorm for object-relational mapping, Redis for caching, and PostgreSQL as the database.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go (1.15 or later)
- PostgreSQL
- Redis

### Installing

1. Clone the repository to your local machine:

```sh
git clone https://github.com/yourusername/sypnasis-golang-test-ecommerce.git
cd sypnasis-golang-test-ecommerce
```

2. Install the Go dependencies:
```
go mod tidy
```

3. Set up your .env file with the necessary environment variables:
```
DATABASE_URL = "host= user= password== dbname= port="
REDIS_ADDR = ""
JWT_KEY = ""
REDIS_PASSWORD = ""
```

4. Run the migrations to set up your database schema:
```
go run migrate/migrate.go
```

5. Running the server
```
go run main.go
```
The server will start on http://localhost:3000.

## Endpoints

- **User Authentication**
  - POST `/signup` - User signup
  - POST `/login` - User login

- **Product Management**
  - GET `/products` - List all products
  - GET `/product/:id` - Get a single product by ID
  - POST `/product` - Create or update a product

- **Category Management**
  - POST `/category` - Create or update a category

- **Order Management**
  - POST `/order` - Create an order
  - GET `/orders` - Get list of user orders
  - GET `/order/:id` - Get a single order by ID

- **Cart Management**
  - POST `/cart` - Add or update cart
  - GET `/carts` - Get user's cart items

## Built With

- [Go](https://golang.org/) - The Go Programming Language
- [Gin](https://github.com/gin-gonic/gin) - HTTP web framework
- [Gorm](https://gorm.io/) - The fantastic ORM library for Golang
- [PostgreSQL](https://www.postgresql.org/) - The World's Most Advanced Open Source Relational Database

## Authors

- **Irfan Muhammad Guvian** - *Initial work* - [irfanguvian](https://github.com/irfanguvian)

## License