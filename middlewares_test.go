package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"net/http"
	"testing"
	"time"
)

func TestAuth_CookieAuthRequired(t *testing.T) {
	router := gin.New()
	router.Use(a.CookieAuthRequired())
	router.GET("/test", func(c *gin.Context) {
		c.Status(200)
	})

	w = PerformRequest(router, "GET", "/test")
	assert.Equal(t, http.StatusOK, w.Code)

	a.Password = HashPassword("test")
	w = PerformRequest(router, "GET", "/test")
	assert.Equal(t, http.StatusUnauthorized, w.Code)

	token := a.createJWT(1 * time.Second)
	w = PerformRequest(router, "GET", "/test",
		keyValueInterface{"Cookie", sessionCookieName + "=" + token},
	)
	assert.Equal(t, http.StatusOK, w.Code)
	a.Password = [32]byte{}
}

func TestAuth_TokenHeaderAuth(t *testing.T) {
	router := gin.New()
	router.Use(a.HeaderAuthRequired())
	router.GET("/test", func(c *gin.Context) {
		c.Status(200)
	})

	w = PerformRequest(router, "GET", "/test")
	assert.Equal(t, http.StatusOK, w.Code)

	a.Password = HashPassword("test")
	w = PerformRequest(router, "GET", "/test")
	assert.Equal(t, http.StatusUnauthorized, w.Code)

	token := a.createJWT(1 * time.Second)
	w = PerformRequest(router, "GET", "/test",
		keyValueInterface{authHeader, token},
	)
	assert.Equal(t, http.StatusOK, w.Code)
	a.Password = [32]byte{}
}
