package auth

import (
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/url"
	"testing"
	"time"
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

	token := a.createJWT(1 * time.Second)
	req.Header.Set(authHeader, token)
	ret = a.HeaderAuthIsValid(c)
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

	token := a.createJWT(1 * time.Second)
	req.Header.Set("Cookie", sessionCookieName+"="+token)
	ret = a.CookieAuthIsValid(c)
	assert.Equal(t, true, ret)
}
