package auth

import (
	"errors"
	"strings"
	"golang.org/x/crypto/bcrypt"

	"github.com/Nikolay-Yakunin/grpc-chat/pkg/token"
)

type service struct {
	repo Repository
}

func NewAuthService(repo Repository) Service {
	return  &service{repo: repo}
}

func (s *service) Register(name, password string) (string,  error) {
	name = strings.TrimSpace(name)
	password = strings.TrimSpace(password)

	if name == "" {
		return "", errors.New("argument 'name' is required")
	}

	if password == "" {
		return "", errors.New("argument 'password' is required")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	token, err := token.GenToken(name)

	if err != nil {
    return "", err
	}

	err = s.repo.Create(name, string(hash))

	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *service)Login(name, password string) (string, error) {
	name = strings.TrimSpace(name)
	password = strings.TrimSpace(password)

	if name == "" {
		return "", errors.New("argument 'name' is required")
	}

	if password == "" {
		return "", errors.New("argument 'password' is required")
	}

	hash, err := s.repo.ReadHash(name)

	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	if err != nil {
		return "", err
	}

	token, err := token.GenToken(name)

	if err != nil {
		return "", err
	}

	return token, nil
}
