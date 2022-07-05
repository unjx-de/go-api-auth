package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"time"
)

type Auth struct {
	Password [32]byte `mapstructure:"PASSWORD"`
	Secret   []byte   `mapstructure:"SECRET"`
}

const (
	shortJwtExpiry    = 2 * time.Second
	longJwtExpiry     = 30 * 24 * time.Hour
	sessionCookieName = "session"
	secretLength      = 40
	authHeader        = "Authorization"
	paramName         = "token"
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

func (a *Auth) GetShortSessionToken() string {
	return a.createJWT(shortJwtExpiry)
}

func (a *Auth) NoPasswordSet() bool {
	return a.Password == [32]byte{}
}
