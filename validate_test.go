package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"testing"
)

func TestAuth_PasswordEquals(t *testing.T) {
	ret := a.PasswordEquals("test")
	if ret != false {
		t.Errorf("expected password not to be equal\n")
	}
	a.Password = HashPassword("test")
	ret = a.PasswordEquals("test")
	if ret != true {
		t.Errorf("expected password to be equal\n")
	}
}

func TestAuth_HeaderAuthIsValid(t *testing.T) {
	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	c.Request = req

	ret := a.HeaderAuthIsValid(c)
	if ret != false {
		t.Errorf("expected header to be invalid\n")
	}

	req.Header.Set(authHeader, "test")
	ret = a.HeaderAuthIsValid(c)
	if ret != false {
		t.Errorf("expected header to be invalid\n")
	}

	token := a.createJWT(shortJwtExpiry)
	req.Header.Set(authHeader, token)
	ret = a.HeaderAuthIsValid(c)
	if ret != true {
		t.Errorf("expected header to be valid\n")
	}
}

func TestAuth_TokenAuthIsValid(t *testing.T) {
	req := &http.Request{
		URL: &url.URL{},
	}
	c.Request = req

	ret := a.TokenAuthIsValid(c)
	if ret != false {
		t.Errorf("expected token to be invalid\n")
	}

	c.Params = []gin.Param{{
		Key:   paramName,
		Value: "test",
	}}
	ret = a.TokenAuthIsValid(c)
	if ret != false {
		t.Errorf("expected token to be invalid\n")
	}

	token := a.createJWT(shortJwtExpiry)
	c.Params = []gin.Param{{
		Key:   paramName,
		Value: token,
	}}
	ret = a.TokenAuthIsValid(c)
	if ret != true {
		t.Errorf("expected token to be valid\n")
	}
}

func TestAuth_CookieAuthIsValid(t *testing.T) {
	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	c.Request = req

	ret := a.CookieAuthIsValid(c)
	if ret != false {
		t.Errorf("expected cookie to be invalid\n")
	}

	req.Header.Set("Cookie", sessionCookieName+"=test")
	ret = a.CookieAuthIsValid(c)
	if ret != false {
		t.Errorf("expected cookie to be invalid\n")
	}

	token := a.createJWT(shortJwtExpiry)
	req.Header.Set("Cookie", sessionCookieName+"="+token)
	ret = a.CookieAuthIsValid(c)
	if ret != true {
		t.Errorf("expected cookie to be valid\n")
	}
}
