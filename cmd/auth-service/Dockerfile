FROM golang:1.24.4-alpine AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY internal/auth ./internal/auth
COPY cmd/auth-service ./cmd/auth-service
COPY pkg/token ./pkg/token

RUN CGO_ENABLED=0 GOOS=linux go build -o auth-service ./cmd/auth-service/main.go


FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache ca-certificates && \
    addgroup -g 1001 -S appgroup && \
    adduser -u 1001 -S appuser -G appgroup

COPY --from=builder /app/auth-service .

RUN chown -R appuser:appgroup /app
USER appuser

CMD ["./auth-service"]