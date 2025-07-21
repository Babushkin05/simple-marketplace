FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o api-gateway ./cmd/main.go

FROM gcr.io/distroless/base-debian11

WORKDIR /app

COPY --from=builder /app/api-gateway /app/api-gateway
COPY --from=builder /app/config/local.yaml /app/config/local.yaml

EXPOSE 8080

CMD ["/app/api-gateway", "--config=config/local.yaml"]