package auth

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func (a *Auth) PasswordEquals(input string) bool {
	return a.Password == HashPassword(input)
}

func (a *Auth) CookieAuthIsValid(c *gin.Context) bool {
	token, err := c.Cookie(sessionCookieName)
	if err != nil {
		return false
	}
	return a.validateJWT(token) == nil
}

func (a *Auth) HeaderAuthIsValid(c *gin.Context) bool {
	token := strings.TrimPrefix(c.GetHeader(authHeader), bearerPrefix)
	if token == "" {
		return false
	}
	return a.validateJWT(token) == nil
}
