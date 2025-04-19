# This is the method to follow to run and check the assignment

# Sales Data API (Golang + Gin + PostgreSQL)

A RESTful API to load and analyze sales data from a CSV file.

## nECESSARY THINGS BEFORE PROCEEDING
- Go 1.20+
- PostgreSQL 13+
- VS Code (optional)
- Place your `sales_data.csv` file in the `data/` folder

## hOW TO Start??

```bash
go mod tidy
psql -U postgres -f schema.sql  # to Create DB and tables
go run cmd/main.go              # to Start the server at http://localhost:8080
```

##  API Endpoints

| Route                 | Method | Query Params                   | Description                         |
|----------------------|--------|--------------------------------|-------------------------------------|
| `/refresh`           | POST   | None                           | Refresh data from CSV               |
| `/revenue/total`     | GET    | `start`, `end` (YYYY-MM-DD)    | Total revenue for date range        |
| `/revenue/product`   | GET    | `start`, `end`                 | Revenue by product                  |
| `/revenue/category`  | GET    | `start`, `end`                 | Revenue by category                 |
| `/revenue/region`    | GET    | `start`, `end`                 | Revenue by region                   |

## Sample CURL Requests

```bash

curl -X POST http://localhost:8080/refresh

curl "http://localhost:8080/revenue/total?start=2024-01-01&end=2024-06-30"

curl "http://localhost:8080/revenue/product?start=2024-01-01&end=2024-06-30"

curl "http://localhost:8080/revenue/category?start=2024-01-01&end=2024-06-30"

curl "http://localhost:8080/revenue/region?start=2024-01-01&end=2024-06-30"
```

# please make these PS------

- Make sure your PostgreSQL credentials are correctly set in `config/database.go`
- Logs (if enabled) can be written to `logs/refresh.log`
- The CSV must follow the format used in `sales_data.csv`


We can discuss if you find any issues in running above project and make request.