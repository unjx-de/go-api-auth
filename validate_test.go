package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestAuth_PasswordEquals(t *testing.T) {
	ret := a.PasswordEquals("test")
	assert.Equal(t, false, ret, "expected password not to be equal")
	a.Password = HashPassword("test")
	ret = a.PasswordEquals("test")
	assert.Equal(t, true, ret, "expected password to be equal")
}

func TestAuth_HeaderAuthIsValid(t *testing.T) {
	req := &http.Request{URL: &url.URL{}, Header: make(http.Header)}
	c.Request = req

	ret := a.HeaderAuthIsValid(c)
	assert.Equal(t, false, ret, "expected header to be invalid")

	req.Header.Set(authHeader, "test")
	ret = a.HeaderAuthIsValid(c)
	assert.Equal(t, false, ret, "expected header to be invalid")

	token := a.createJWT(shortJwtExpiry)
	req.Header.Set(authHeader, token)
	ret = a.HeaderAuthIsValid(c)
	assert.Equal(t, true, ret, "expected header to be valid")
}

func TestAuth_TokenAuthIsValid(t *testing.T) {
	req := &http.Request{URL: &url.URL{}}
	c.Request = req

	ret := a.TokenAuthIsValid(c)
	assert.Equal(t, false, ret, "expected token to be invalid")

	c.Params = []gin.Param{{
		Key:   paramName,
		Value: "test",
	}}
	ret = a.TokenAuthIsValid(c)
	assert.Equal(t, false, ret, "expected token to be invalid")

	token := a.createJWT(shortJwtExpiry)
	c.Params = []gin.Param{{
		Key:   paramName,
		Value: token,
	}}
	ret = a.TokenAuthIsValid(c)
	assert.Equal(t, true, ret, "expected token to be valid")
}

func TestAuth_CookieAuthIsValid(t *testing.T) {
	req := &http.Request{URL: &url.URL{}, Header: make(http.Header)}
	c.Request = req

	ret := a.CookieAuthIsValid(c)
	assert.Equal(t, false, ret, "expected cookie to be invalid")

	req.Header.Set("Cookie", sessionCookieName+"=test")
	ret = a.CookieAuthIsValid(c)
	assert.Equal(t, false, ret, "expected cookie to be invalid")

	token := a.createJWT(shortJwtExpiry)
	req.Header.Set("Cookie", sessionCookieName+"="+token)
	ret = a.CookieAuthIsValid(c)
	assert.Equal(t, true, ret, "expected cookie to be valid")
}
