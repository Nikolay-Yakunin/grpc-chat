package auth

import (
	"context"
)

type AuthHandler struct {
    svc Service
    UnimplementedAuthServiceServer
}

func NewAuthHandler(scv Service) *AuthHandler {
	return &AuthHandler{svc: scv}
}

func (h *AuthHandler) Register(ctx context.Context, req *RegisterRequest) (*RegisterResponse, error) {
	token, err := h.svc.Register(req.GetName(), req.GetPassword())

	if err != nil {
		return &RegisterResponse{
			Token: "",
			Message: err.Error(),
		}, nil
	}

	if token == "" {
		return &RegisterResponse{
			Token: "",
			Message: "service error: empty token",
		}, nil
	}

	return &RegisterResponse{
		Token: token,
		Message: "registered",
	}, nil
}

func (h *AuthHandler)Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	token, err := h.svc.Login(req.GetName(), req.GetPassword())

	if err != nil {
		return &LoginResponse{
			Token: "",
			Message: err.Error(),
		}, nil
	}

	if token == "" {
		return &LoginResponse{
			Token: "",
			Message: "service error: empty token",
		}, nil
	}

	return &LoginResponse{
		Token: token,
		Message: "log in",
	}, nil 
} 