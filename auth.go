package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"time"
)

type Auth struct {
	Password [32]byte
	Secret   []byte
}

const (
	JwtExpiry         = 30 * 24 * time.Hour
	sessionCookieName = "session"
	secretLength      = 40
	authHeader        = "Authorization"
	bearerPrefix      = "Bearer "
)

func SecretGenerator() []byte {
	b := make([]byte, secretLength)
	_, _ = rand.Read(b)
	return b
}

func HashPassword(password string) [32]byte {
	return sha256.Sum256([]byte(password))
}

func (a *Auth) NoPasswordSet() bool {
	return a.Password == [32]byte{}
}
