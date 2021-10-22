package meta

import (
	"github.com/gin-gonic/gin"
	"github.com/op-server/pkg/logger"
	"github.com/op-server/pkg/server/config"
	"github.com/op-server/pkg/server/core/resp"
	"github.com/op-server/pkg/server/dao"
	"github.com/op-server/pkg/server/globalrouter/types"
)

func init() {
	types.InsertAddRouterGroupFunc(AddClusterHandlerRouterGroup)
}

func AddClusterHandlerRouterGroup(router *gin.RouterGroup, dao *dao.BaseDao, _ *config.HTTPServerConfiguration) {
	logger.Info("add cluster handler router group")
	handler := NewClusterMetaHandler(dao)
	v1Group := router.Group("/cluster/v1/meta")

	v1Group.POST("register", resp.WithJSONResp(handler.register))
	v1Group.PUT("update", resp.WithJSONResp(handler.updateCluster))
	v1Group.POST("find", resp.WithJSONResp(handler.find))
	v1Group.GET("list", resp.WithJSONResp(handler.list))
	v1Group.DELETE("delete/:id", resp.WithJSONResp(handler.deleteClusterByID))

}
