package auth

import (
	"net/http"
	"net/url"
	"testing"
)

func TestAuth_SetSessionCookie(t *testing.T) {
	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	c.Request = req
	a.SetSessionCookie(c)
}

func TestAuth_DeleteSessionCookie(t *testing.T) {
	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	c.Request = req
	a.DeleteSessionCookie(c)
}
