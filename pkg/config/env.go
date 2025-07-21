package config

import (
	"os"
)

type ENV struct {
	DATABASE_URL string
	GRPC_PORT string
	JWT_SECRET string
}

func NewENV() ENV {
	var env ENV
	
	env.DATABASE_URL = os.Getenv("DATABASE_URL")
	env.GRPC_PORT = os.Getenv("GRPC_PORT")
	env.JWT_SECRET = os.Getenv("JWT_SECRET")

	return env
}