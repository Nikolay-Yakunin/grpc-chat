package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"

    "github.com/Nikolay-Yakunin/grpc-chat/pkg/config"
)

func GenToken(name string) (string, error) {
    env := config.NewENV()

    claims := jwt.MapClaims{
        "sub": name,
        "exp": time.Now().Add(24 * time.Hour).Unix(),
        "iat": time.Now().Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    secretKey := []byte(env.JWT_SECRET)

    return token.SignedString(secretKey)
}