FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o goods-service ./cmd/main.go

FROM gcr.io/distroless/base-debian11

WORKDIR /app

COPY --from=builder /app/goods-service /app/goods-service
COPY --from=builder /app/config/local.yaml /app/config/local.yaml
COPY --from=builder /app/internal/db/schema.sql /app/internal/db/schema.sql

EXPOSE 50052

CMD ["/app/goods-service", "--config=config/local.yaml"]