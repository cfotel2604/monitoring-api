# monitoring-api

HTTP API trung gian giữa Grafana (Infinity datasource) và ClickHouse.

## Mục tiêu ban đầu

- Grafana không truy cập trực tiếp ClickHouse.
- Grafana Infinity gọi HTTP endpoint của Monitoring API.
- Monitoring API kiểm soát query, validate input và trả JSON phù hợp cho Grafana.

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
