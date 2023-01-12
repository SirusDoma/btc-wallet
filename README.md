# BTC Wallet
A simple application to simulate btc transaction bookkeeping.  
Written in Go and uses PostgreSQL as persistence storage.

## Setup with Docker
- Copy `.env.sample` to `.env` and adjust the configuration to your liking.
- Run `docker compose up`, it will automatically provision PostgreSQL based on `.env` configuration.
- Application is served with `8080` port.

Note that if you change the `DATABASE_HOST` variable, make sure to update the database name as well in `docker-compose.yml`

## Setup locally
- Copy `.env.sample` to `.env` and adjust the configuration to your liking.
- Create application database within your PostgreSQL local instance.
- Run `go run cmd/app/main.go`.
- Application is served with `8080` port.

## Endpoints
### `GET /top-ups?{startDatetime}&{endDatetime}&{offset}&{limit}`
Query Params:
- `startDatetime`: Specify the earliest datetime of Top-up to fetch.
- `endDatetime`: Specify the latest datetime of Top-up to fetch.
- `offset`: (OPTIONAL) Specify offset pagination.
- `limit`: (OPTIONAL) Specify limit pagination.

Return: Sum of recorded Top-up, grouped by hour.

### `POST /top-ups`
Request Body (JSON):
- `amount`: Specify the amount of topup

Example:
```json
{
  "amount": 1001.1
}
```

Return: Created Top-up object.
