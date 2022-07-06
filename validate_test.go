package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/url"
	"testing"
)

func TestAuth_PasswordEquals(t *testing.T) {
	ret := a.PasswordEquals("test")
	assert.Equal(t, false, ret)
	a.Password = HashPassword("test")
	ret = a.PasswordEquals("test")
	assert.Equal(t, true, ret)
}

func TestAuth_HeaderAuthIsValid(t *testing.T) {
	req := &http.Request{URL: &url.URL{}, Header: make(http.Header)}
	c.Request = req

	ret := a.HeaderAuthIsValid(c)
	assert.Equal(t, false, ret)

	req.Header.Set(authHeader, "test")
	ret = a.HeaderAuthIsValid(c)
	assert.Equal(t, false, ret)

	token := a.createJWT(shortJwtExpiry)
	req.Header.Set(authHeader, token)
	ret = a.HeaderAuthIsValid(c)
	assert.Equal(t, true, ret)
}

func TestAuth_TokenAuthIsValid(t *testing.T) {
	req := &http.Request{URL: &url.URL{}}
	c.Request = req

	ret := a.TokenAuthIsValid(c)
	assert.Equal(t, false, ret)

	c.Params = []gin.Param{{
		Key:   paramName,
		Value: "test",
	}}
	ret = a.TokenAuthIsValid(c)
	assert.Equal(t, false, ret)

	token := a.createJWT(shortJwtExpiry)
	c.Params = []gin.Param{{
		Key:   paramName,
		Value: token,
	}}
	ret = a.TokenAuthIsValid(c)
	assert.Equal(t, true, ret)
}

func TestAuth_CookieAuthIsValid(t *testing.T) {
	req := &http.Request{URL: &url.URL{}, Header: make(http.Header)}
	c.Request = req

	ret := a.CookieAuthIsValid(c)
	assert.Equal(t, false, ret)

	req.Header.Set("Cookie", sessionCookieName+"=test")
	ret = a.CookieAuthIsValid(c)
	assert.Equal(t, false, ret)

	token := a.createJWT(shortJwtExpiry)
	req.Header.Set("Cookie", sessionCookieName+"="+token)
	ret = a.CookieAuthIsValid(c)
	assert.Equal(t, true, ret)
}
