package auth

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestAuth_NoPasswordSet(t *testing.T) {
	a.Password = HashPassword("test")
	ret := a.NoPasswordSet()
	assert.Equal(t, false, ret)
	a.Password = [32]byte{}
	ret = a.NoPasswordSet()
	assert.Equal(t, true, ret)
}
