# monitoring-api

HTTP API trung gian giữa Grafana (Infinity datasource) và ClickHouse.

## Mục tiêu ban đầu

- Grafana không truy cập trực tiếp ClickHouse.
- Grafana Infinity gọi HTTP endpoint của Monitoring API.
- Monitoring API kiểm soát query, validate input và trả JSON phù hợp cho Grafana.
- Ưu tiên contract API rõ ràng trước, chưa tối ưu tải sớm.

## Kiến trúc

```text
Grafana
  |
  | Infinity datasource / HTTP
  v
Monitoring API
  |
  | controlled SQL
  v
ClickHouse
```

## Endpoints hiện tại

- `GET /health`
- `GET /api/v1/metrics/summary`
- `GET /api/v1/metrics/timeseries`

Hai metrics endpoint hiện trả dữ liệu mock để kiểm chứng luồng Grafana Infinity -> Monitoring API trước khi nối ClickHouse thật.

## Chạy local

```bash
go run ./cmd/api
```

Mặc định API listen tại `:8080`.

## Chạy bằng Docker Compose

```bash
docker compose up --build
```

Test nhanh:

```bash
curl http://localhost:8080/health
curl http://localhost:8080/api/v1/metrics/summary
curl http://localhost:8080/api/v1/metrics/timeseries
```

## Bước tiếp theo

1. Chạy build/test local để xác nhận môi trường Go/Docker.
2. Cho Grafana Infinity gọi thử các endpoint mock.
3. Chốt JSON contract phù hợp với panel Grafana.
4. Thêm ClickHouse config + repository và thay dữ liệu mock bằng controlled SQL.
