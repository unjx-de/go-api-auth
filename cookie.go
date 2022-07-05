package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (a *Auth) SetSessionCookie(c *gin.Context) {
	c.SetSameSite(http.SameSiteStrictMode)
	c.SetCookie(sessionCookieName, a.createJWT(longJwtExpiry), int(longJwtExpiry.Seconds()), "/", strings.Split(c.Request.Host, ":")[0], true, true)
}

func (a *Auth) DeleteSessionCookie(c *gin.Context) {
	c.SetCookie(sessionCookieName, "none", -1, "/", strings.Split(c.Request.Host, ":")[0], true, true)
}
