package main

import (
    "database/sql"
    "log"
    "net"

    "google.golang.org/grpc"
    _ "github.com/lib/pq"

    "github.com/Nikolay-Yakunin/grpc-chat/internal/auth"
    "github.com/Nikolay-Yakunin/grpc-chat/pkg/config"
)


func main() {
    env := config.NewENV()
    if env.DATABASE_URL == "" {
        env.DATABASE_URL = "postgres://user:password@localhost/dbname?sslmode=disable"
    }

    if env.GRPC_PORT == "" {
        env.GRPC_PORT = "50051"
    }

    db, err := sql.Open("postgres", env.DATABASE_URL)
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

    listener, err := net.Listen("tcp", ":"+env.GRPC_PORT)
    if err != nil {
        log.Fatal("Не удалось запустить слушатель:", err)
    }

    grpcServer := grpc.NewServer()
    auth.RegisterAuthServiceServer(grpcServer, handler)

    log.Println("gRPC сервер запущен на порту", env.GRPC_PORT)
    if err := grpcServer.Serve(listener); err != nil {
        log.Fatal("Ошибка запуска gRPC сервера:", err)
    }
}