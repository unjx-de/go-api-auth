package auth

import (
	"testing"
	"time"
)

func TestAuth_CreateAndValidateJWT(t *testing.T) {
	token := a.createJWT(shortJwtExpiry)
	err := a.validateJWT(token)
	if err != nil {
		t.Errorf("%v\n", err)
	}
}

func TestAuth_CreateAndValidateJWTAfterExpiry(t *testing.T) {
	token := a.createJWT(shortJwtExpiry)
	time.Sleep(shortJwtExpiry)
	err := a.validateJWT(token)
	if err.Error() != "Token is expired" {
		t.Errorf("expected token expiration\n")
	}
}

func TestAuth_InvalidJWT(t *testing.T) {
	err := a.validateJWT("test")
	if err.Error() != "token contains an invalid number of segments" {
		t.Errorf("expected invalid token\n")
	}
	err = a.validateJWT("")
	if err.Error() != "token contains an invalid number of segments" {
		t.Errorf("expected invalid token\n")
	}
}
