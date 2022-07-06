package auth

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAuth_CreateAndValidateJWT(t *testing.T) {
	token := a.createJWT(shortJwtExpiry)
	err := a.validateJWT(token)
	assert.Equal(t, nil, err, "expected token to be valid")
}

func TestAuth_CreateAndValidateJWTAfterExpiry(t *testing.T) {
	token := a.createJWT(shortJwtExpiry)
	time.Sleep(shortJwtExpiry)
	err := a.validateJWT(token)
	assert.Equal(t, "Token is expired", err.Error(), "expected token to be expired")
}

func TestAuth_InvalidJWT(t *testing.T) {
	err := a.validateJWT("")
	assert.NotEqual(t, nil, err, "expected token to be invalid")
	err = a.validateJWT("test")
	assert.NotEqual(t, nil, err, "expected token to be invalid")
}
