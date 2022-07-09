package auth

import (
	"github.com/go-playground/assert/v2"
	"testing"
	"time"
)

func TestAuth_CreateAndValidateJWT(t *testing.T) {
	token := a.createJWT(1 * time.Second)
	err := a.validateJWT(token)
	assert.Equal(t, nil, err)
}

func TestAuth_CreateAndValidateJWTAfterExpiry(t *testing.T) {
	token := a.createJWT(1 * time.Second)
	time.Sleep(1 * time.Second)
	err := a.validateJWT(token)
	assert.Equal(t, "Token is expired", err.Error())
}

func TestAuth_InvalidJWT(t *testing.T) {
	err := a.validateJWT("")
	assert.NotEqual(t, nil, err)
	err = a.validateJWT("test")
	assert.NotEqual(t, nil, err)
}
