package requestid

import (
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	middleware_lib "github.com/op-server/pkg/server/middleware/lib"

	"github.com/gin-gonic/gin"

	"github.com/stretchr/testify/assert"
)

func TestRequestID(t *testing.T) {
	assert := assert.New(t)

	router := gin.New()
	router.Use(Middleware(middleware_lib.MethodAndPathSkipper(http.MethodGet, regexp.MustCompile("^/req1"))))

	req1 := httptest.NewRequest(http.MethodGet, "/req1", nil)
	rec1 := httptest.NewRecorder()
	router.ServeHTTP(rec1, req1)
	assert.Equal("", rec1.Header().Get(HeaderXRequestID))

	req2 := httptest.NewRequest(http.MethodGet, "/req2", nil)
	rec2 := httptest.NewRecorder()
	router.ServeHTTP(rec2, req2)
	assert.NotEqual("", rec2.Header().Get(HeaderXRequestID))

	req3 := httptest.NewRequest(http.MethodGet, "/req3", nil)
	req3.Header.Add(HeaderXRequestID, "852803be-e5fe-499b-bbea-c9e5b5f43916")
	rec3 := httptest.NewRecorder()
	router.ServeHTTP(rec3, req3)
	assert.Equal("852803be-e5fe-499b-bbea-c9e5b5f43916", rec3.Header().Get(HeaderXRequestID))
}
