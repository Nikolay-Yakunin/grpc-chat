package auth

import (
	context "context"
	"errors"
	"testing"
)

/* ------------- Mock ------------- */
type mockRepository struct {
	users map[string]string
}

func newMockRepository() *mockRepository {
	return &mockRepository{
		users: make(map[string]string),
	}
}

func (r *mockRepository) Create(name, passwordHash string) error {
	if _, exists := r.users[name]; exists {
		return errors.New("user already exists")
	}
	r.users[name] = passwordHash
	return nil
}

func (r *mockRepository) ReadHash(name string) (string, error) {
	hash, exists := r.users[name]
	if !exists {
		return "", errors.New("user not found")
	}
	return hash, nil
}
/* ------------- Mock ------------- */

func TestAuthService_Register(t *testing.T) {
	mockRepo := newMockRepository()
	service := NewAuthService(mockRepo)

	// Успешная регистрация
	token, err := service.Register("user1", "password123")
	if err != nil {
		t.Errorf("unexpected error during registration: %v", err)
	}
	if token == "" {
		t.Errorf("token is empty string")
	}
}

func TestAuthService_RegisterExistUser(t *testing.T) {
	mockRepo := newMockRepository()
	service := NewAuthService(mockRepo)

	// Успешная регистрация
	token, err := service.Register("user1", "password123")
	if err != nil {
		t.Errorf("unexpected error during registration: %v", err)
	}
	if token == "" {
		t.Errorf("token is empty string")
	}

	// Попытка повторной регистрации
	token, err = service.Register("user1", "password123")
	if err == nil {
		t.Errorf("expected error when registering existing user")
	}
	if token != "" {
		t.Errorf("token is not empty string")
	}
}

func TestAuthHandler_Register(t *testing.T) {
	mockRepo := newMockRepository()
	service := NewAuthService(mockRepo)
	handler := NewAuthHandler(service)

	req := &RegisterRequest{
		Name:     "user1",
		Password: "secret123",
	}
	// Успешная регистрация
	resp, err := handler.Register(context.Background(), req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Message != "registered" {
		t.Errorf("expected 'registered', got %s", resp.Message)
	}
	if resp.Token == "" {
		t.Errorf("expected token, got empty string")
	}
}

func TestAuthHandler_RegisterExistUser(t *testing.T) {
	mockRepo := newMockRepository()
	service := NewAuthService(mockRepo)
	handler := NewAuthHandler(service)

	req := &RegisterRequest{
		Name:     "user1",
		Password: "secret123",
	}

	// Успешная регистрация
	resp, err := handler.Register(context.Background(), req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Message != "registered" {
		t.Errorf("expected 'registered', got %s", resp.Message)
	}
	if resp.Token == "" {
		t.Errorf("expected token, got empty string")
	}
	// Попытка повторной регистрации
	resp, err = handler.Register(context.Background(), req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Message == "registered" {
		t.Errorf("expected error, got 'registered'")
	}
	if resp.Token != "" {
		t.Errorf("expected empty string, got token")
	}
}