package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestAuth_CookieAuthRequired(t *testing.T) {
	router := gin.New()
	router.Use(a.CookieAuthRequired())
	router.GET("/test", func(c *gin.Context) {
		c.Status(200)
	})

	w = PerformRequest(router, "GET", "/test")
	assert.Equal(t, http.StatusOK, w.Code, "expected request to return StatusOK")

	a.Password = HashPassword("test")
	w = PerformRequest(router, "GET", "/test")
	assert.Equal(t, http.StatusUnauthorized, w.Code, "expected request to return StatusUnauthorized")

	token := a.createJWT(shortJwtExpiry)
	w = PerformRequest(router, "GET", "/test",
		keyValueInterface{"Cookie", sessionCookieName + "=" + token},
	)
	assert.Equal(t, http.StatusOK, w.Code, "expected request to return StatusOK")
	a.Password = [32]byte{}
}

func TestAuth_TokenParamAuth(t *testing.T) {
	router := gin.New()
	router.Use(a.TokenParamAuth())
	router.GET("/test", func(c *gin.Context) {
		c.Status(200)
	})

	w = PerformRequest(router, "GET", "/test")
	assert.Equal(t, http.StatusOK, w.Code, "expected request to return StatusOK")

	a.Password = HashPassword("test")
	w = PerformRequest(router, "GET", "/test")
	assert.Equal(t, http.StatusUnauthorized, w.Code, "expected request to return StatusUnauthorized")
	a.Password = [32]byte{}
}

func TestAuth_TokenHeaderAuth(t *testing.T) {
	router := gin.New()
	router.Use(a.TokenHeaderAuth())
	router.GET("/test", func(c *gin.Context) {
		c.Status(200)
	})

	w = PerformRequest(router, "GET", "/test")
	assert.Equal(t, http.StatusOK, w.Code, "expected request to return StatusOK")

	a.Password = HashPassword("test")
	w = PerformRequest(router, "GET", "/test")
	assert.Equal(t, http.StatusUnauthorized, w.Code, "expected request to return StatusUnauthorized")

	token := a.createJWT(shortJwtExpiry)
	w = PerformRequest(router, "GET", "/test",
		keyValueInterface{authHeader, token},
	)
	assert.Equal(t, http.StatusOK, w.Code, "expected request to return StatusOK")
	a.Password = [32]byte{}
}
