# E-commerce Microservice Application using Golang

## 📌 Overview
This project is an **e-commerce microservice architecture** built with **Golang**.  
It demonstrates how multiple services can interact using **gRPC**, with a **GraphQL gateway** aggregating data from all services.

### Services
- **Account Service** – manages user accounts.
- **Catalog Service** – manages products.
- **Order Service** – manages orders and coordinates with Account & Catalog services.
- **GraphQL Gateway** – single entry point for clients to interact with all services.
- **PostgreSQL** – used by Account and Order services for persistence.
- **Elasticsearch** – used by Catalog service for product search.

---

## 🏗 Architecture
[ Client ]
↓
[ GraphQL Gateway ]
↓ gRPC calls
+-------------+ +-------------+ +-------------+
| Account Svc | | Catalog Svc | | Order Svc |
| PostgreSQL | | Elasticsearch| | PostgreSQL |
+-------------+ +-------------+ +-------------+

---

## 🚀 Getting Started

### Prerequisites
- [Docker Desktop](https://www.docker.com/products/docker-desktop) installed & running
- [Go 1.23+](https://go.dev/dl/) (only if you plan to run locally)

---

### Run with Docker Compose
```bash
# Clone the repository
git clone https://github.com/yourusername/E-commerce-microservice-application-using-Golang.git
cd E-commerce-microservice-application-using-Golang

# Build and start services
docker compose up -d --build

# View logs for GraphQL gateway
docker compose logs -f graphql


GraphQL Playground will be available at:
➡ http://localhost:8000/playground (or / depending on your router config)