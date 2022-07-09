package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (a *Auth) CookieAuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if a.NoPasswordSet() || a.CookieAuthIsValid(c) {
			c.Next()
		} else {
			c.Status(http.StatusUnauthorized)
			c.Abort()
		}
	}
}

func (a *Auth) HeaderAuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if a.NoPasswordSet() || a.HeaderAuthIsValid(c) {
			c.Next()
		} else {
			c.Status(http.StatusUnauthorized)
			c.Abort()
		}
	}
}
