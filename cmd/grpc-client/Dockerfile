FROM golang:1.24.4-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY internal/ ./internal/
COPY cmd/grpc-client ./cmd/grpc-client
COPY pkg ./pkg

RUN CGO_ENABLED=0 GOOS=linux go build -o grpc-client ./cmd/grpc-client/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/grpc-client .

ENTRYPOINT ["./grpc-client"]