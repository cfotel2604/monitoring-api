FROM golang:1.24-alpine AS build

WORKDIR /src

COPY go.mod ./
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /out/monitoring-api ./cmd/api

FROM alpine:3.21

WORKDIR /app
COPY --from=build /out/monitoring-api /app/monitoring-api

EXPOSE 8080
ENTRYPOINT ["/app/monitoring-api"]
