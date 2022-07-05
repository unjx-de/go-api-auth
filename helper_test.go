package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
)

var a = Auth{Secret: SecretGenerator()}
var w = httptest.NewRecorder()
var c, _ = gin.CreateTestContext(w)

type keyValueInterface struct {
	Key   string
	Value string
}

func PerformRequest(r http.Handler, method, path string, headers ...keyValueInterface) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, nil)
	for _, h := range headers {
		req.Header.Add(h.Key, h.Value)
	}
	writer := httptest.NewRecorder()
	r.ServeHTTP(writer, req)
	return writer
}
