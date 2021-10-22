package ping

import (
	"net/http"

	"github.com/op-server/pkg/server/services/ping"

	"github.com/gin-gonic/gin"
)

type PingHandler struct {
	service *ping.PingService
}

func NewPingHandler() *PingHandler {
	return &PingHandler{
		service: ping.NewPingService(),
	}
}

func (ping *PingHandler) ping(ctx *gin.Context) {
	result, err := ping.service.Ping()
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.String(http.StatusOK, result)
}
