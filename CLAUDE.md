# inventory-service

HTTP service returning product stock information. Part of a 3-service test setup.

## Build & Run
```bash
go run main.go
```

Listens on `:8081` (override with `PORT` env var).

## Endpoints
- `GET /stock?product_id=PROD-001` — returns stock info (product name, quantity, warehouse)
- `GET /health` — health check

## Test Data
- `PROD-001`: Widget A, qty 150, WH-EAST
- `PROD-002`: Widget B, qty 0, WH-WEST (out of stock)
- `PROD-003`: Gadget C, qty 42, WH-EAST
