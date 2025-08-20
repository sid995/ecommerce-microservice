# E-Commerce Microservices Architecture

A modern e-commerce platform built using microservices architecture with Go, gRPC, and GraphQL. This project demonstrates the implementation of a distributed system with separate services for account management, product catalog, and order processing.

## Project Overview

The project consists of the following microservices:

1. **Account Service**: Handles user account management
2. **Catalog Service**: Manages product inventory and details
3. **Order Service**: Processes customer orders
4. **GraphQL API Gateway**: Provides a unified API interface for clients

## Tech Stack

- **Backend**: Go (Golang)
- **API Gateway**: GraphQL with gqlgen
- **Inter-Service Communication**: gRPC
- **Database**: PostgreSQL (for accounts and orders)
- **Search**: Elasticsearch (for product catalog)
- **Configuration**: Environment variables with envconfig
- **Dependencies**: Managed with Go modules

## Architecture
```
                ┌─────────────┐
                │    Client   │
                └──────┬──────┘
                       │
              ┌────────┴────────┐
              │     GraphQL     │
              │     Gateway     │
              └───┬────┬────┬───┘
                  │    │    │
          ┌───────┘    │    └───────┐
          │            │            │
   ┌──────┴───┐  ┌─────┴────┐  ┌────┴────┐
   │ Acc Svc  │  │ Cat Svc  │  │ Ord Svc │
   └──────────┘  └──────────┘  └─────────┘
```

## Services Description

### Account Service
- User account management
- CRUD operations for user accounts
- gRPC endpoints for account operations

### Catalog Service
- Product management
- Product search and filtering
- Elasticsearch integration for efficient product search

### Order Service
- Order processing
- Order history tracking
- Handles order status updates

### GraphQL Gateway
- Unified API layer
- Query resolution
- Data aggregation from multiple services

## API Features

### Queries
- `accounts`: Fetch user accounts with pagination
- `products`: Search and retrieve products with filtering

### Mutations
- `createAccount`: Create new user accounts
- `createProduct`: Add new products to the catalog
- `createOrder`: Process new orders

## Setup Instructions

1. Clone the repository:
   ```bash
   git clone https://github.com/sid995/ecommerce-microservice.git
   cd ecommerce-microservice
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Set up environment variables:
   ```bash
   export ACCOUNT_SERVICE_URL=localhost:50051
   export CATALOG_SERVICE_URL=localhost:50052
   export ORDER_SERVICE_URL=localhost:50053
   ```

4. Start the services:
   ```bash
   # Start Account Service
   cd account/cmd/account
   go run main.go

   # Start Catalog Service
   cd ../../catalog/cmd/catalog
   go run main.go

   # Start Order Service
   cd ../../order/cmd/order
   go run main.go

   # Start GraphQL Gateway
   cd ../../graphql
   go run .
   ```

5. Access the GraphQL Playground:
   - Open `http://localhost:8080/playground` in your browser
   - The GraphQL endpoint is available at `http://localhost:8080/graphql`

## Example Queries

### Fetch Accounts
```graphql
query {
  accounts(pagination: { skip: 0, take: 10 }) {
    id
    name
    orders {
      id
      createdAt
      totalPrice
    }
  }
}
```

### Create Product
```graphql
mutation {
  createProduct(product: {
    name: "Example Product"
    description: "Product description"
    price: 99.99
  }) {
    id
    name
    price
  }
}
```

### Create Order
```graphql
mutation {
  createOrder(order: {
    accountId: "account123"
    products: [
      { id: "product123", quantity: 2 }
    ]
  }) {
    id
    totalPrice
    createdAt
  }
}
```

## Project Structure
```
.
├── account/           # Account service
│   ├── cmd/          # Service entry point
│   ├── pb/           # Generated protobuf files
│   └── ...          # Service implementation
├── catalog/          # Catalog service
│   ├── cmd/          # Service entry point
│   ├── pb/           # Generated protobuf files
│   └── ...          # Service implementation
├── order/            # Order service
│   ├── cmd/          # Service entry point
│   ├── pb/           # Generated protobuf files
│   └── ...          # Service implementation
├── graphql/          # GraphQL gateway
│   ├── schema.graphql # GraphQL schema
│   └── ...          # Resolver implementations
├── go.mod           # Go module file
└── go.sum           # Go module checksum
```

## Contributing

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a new Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.
