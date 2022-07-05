package auth

import (
	"testing"
)

func TestAuth_GetShortSessionToken(t *testing.T) {
	token := a.GetShortSessionToken()
	err := a.validateJWT(token)
	if err != nil {
		t.Errorf("%v\n", err)
	}
}

func TestAuth_NoPasswordSet(t *testing.T) {
	a.Password = HashPassword("test")
	ret := a.NoPasswordSet()
	if ret != false {
		t.Errorf("expected password to be set\n")
	}
	a.Password = [32]byte{}
	ret = a.NoPasswordSet()
	if ret != true {
		t.Errorf("expected no password to be set\n")
	}
}
