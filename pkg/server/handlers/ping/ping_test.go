package ping

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/op-server/pkg/logger"

	"github.com/stretchr/testify/suite"
)

type pingControllerTestSuit struct {
	suite.Suite
	controller *PingHandler
	router     *gin.Engine
}

func (suite *pingControllerTestSuit) SetupTest() {
	suite.controller = NewPingHandler()
	suite.router = gin.Default()
	AddPingHandlerRouterGroup(suite.router.Group("/api/op"), nil, nil)
}

func (suite *pingControllerTestSuit) TestPingHandler() {
	req, err := http.NewRequest("GET", "/api/op/v1/ping", nil)
	if err != nil {
		logger.Fatal(err.Error())
	}
	rr := httptest.NewRecorder()
	suite.router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		logger.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := "ok"
	if rr.Body.String() != expected {
		logger.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestPingController(t *testing.T) {
	suite.Run(t, &pingControllerTestSuit{})
}
