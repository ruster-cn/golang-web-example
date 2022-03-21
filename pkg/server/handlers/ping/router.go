package ping

import (
	"github.com/gin-gonic/gin"
	"github.com/op-server/pkg/server/config"
	"github.com/op-server/pkg/server/dao"
	"github.com/op-server/pkg/server/globalrouter/types"

	log "github.com/ruster-cn/zap-log-wrapper"
)

func init() {
	types.InsertAddRouterGroupFunc(AddPingHandlerRouterGroup)
}

func AddPingHandlerRouterGroup(router *gin.RouterGroup, _ *dao.BaseDao, _ *config.HTTPServerConfiguration) {
	log.Info("add ping handler router group")
	controller := NewPingHandler()
	v1 := router.Group("/v1")
	{
		v1.GET("ping", controller.ping)
	}
}
