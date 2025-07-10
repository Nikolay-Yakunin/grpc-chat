package token

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenToken(name string) (string, error) {
    claims := jwt.MapClaims{
        "sub": name,
        "exp": time.Now().Add(24 * time.Hour).Unix(),
        "iat": time.Now().Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    secretKey := []byte(os.Getenv("JWT_SECRET"))

    return token.SignedString(secretKey)
}