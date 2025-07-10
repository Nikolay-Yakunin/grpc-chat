package auth

type AuthUser struct {
	UserName string
	PasswordHash string
}

type Service interface {
	Register(name, password string) (string, error)
	Login(name, password string) (string, error)
}

type Repository interface {
	Create(name, passwordHash string) error
	ReadHash(name string) (string, error)
}