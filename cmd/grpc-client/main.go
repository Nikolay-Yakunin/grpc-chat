package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/Nikolay-Yakunin/grpc-chat/internal/auth"
)

func main() {
	mode := flag.String("mode", "", "register или login")
	name := flag.String("name", "", "Имя пользователя")
	password := flag.String("password", "", "Пароль")
	addr := flag.String("addr", "auth-service:50051", "Адрес gRPC сервера")
	flag.Parse()

	if *mode != "register" && *mode != "login" {
		fmt.Println("Использование: -mode register|login -name <имя> -password <пароль>")
		os.Exit(1)
	}


	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Не удалось подключиться: %v", err)
	}
	defer conn.Close()

	client := auth.NewAuthServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	switch *mode {
	case "register":
		req := &auth.RegisterRequest{
			Name:     *name,
			Password: *password,
		}
		resp, err := client.Register(ctx, req)
		if err != nil {
			log.Fatalf("Ошибка регистрации: %v", err)
		}
		fmt.Printf("Ответ: %s\nТокен: %s\n", resp.Message, resp.Token)
	case "login":
		req := &auth.LoginRequest{
			Name:     *name,
			Password: *password,
		}
		resp, err := client.Login(ctx, req)
		if err != nil {
			log.Fatalf("Ошибка входа: %v", err)
		}
		fmt.Printf("Ответ: %s\nТокен: %s\n", resp.Message, resp.Token)
	}
}
