package main

import (
    "database/sql"
    "log"
    "net"
    "os"

    "google.golang.org/grpc"
    _ "github.com/lib/pq"

    "github.com/Nikolay-Yakunin/grpc-chat/internal/auth"
)


func main() {
    dbURL := os.Getenv("DATABASE_URL")
    if dbURL == "" {
        dbURL = "postgres://user:password@localhost/dbname?sslmode=disable"
    }

    port := os.Getenv("GRPC_PORT")
    if port == "" {
        port = "50051"
    }

    db, err := sql.Open("postgres", dbURL)
    if err != nil {
        log.Fatal("Не удалось подключиться к базе данных:", err)
    }
    defer db.Close()

    if err := db.Ping(); err != nil {
        log.Fatal("Не удалось выполнить ping к базе данных:", err)
		}

    repo := auth.NewPostgresRepository(db)
    service := auth.NewAuthService(repo)
    handler := auth.NewAuthHandler(service)

    listener, err := net.Listen("tcp", ":"+port)
    if err != nil {
        log.Fatal("Не удалось запустить слушатель:", err)
    }

    grpcServer := grpc.NewServer()
    auth.RegisterAuthServiceServer(grpcServer, handler)

    log.Println("gRPC сервер запущен на порту", port)
    if err := grpcServer.Serve(listener); err != nil {
        log.Fatal("Ошибка запуска gRPC сервера:", err)
    }
}