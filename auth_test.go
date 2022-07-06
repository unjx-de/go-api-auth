package auth

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuth_GetShortSessionToken(t *testing.T) {
	token := a.GetShortSessionToken()
	err := a.validateJWT(token)
	assert.Equal(t, nil, err, "expected token to be valid")
}

func TestAuth_NoPasswordSet(t *testing.T) {
	a.Password = HashPassword("test")
	ret := a.NoPasswordSet()
	assert.Equal(t, false, ret, "expected password not to be set")
	a.Password = [32]byte{}
	ret = a.NoPasswordSet()
	assert.Equal(t, true, ret, "expected password to be set")
}
