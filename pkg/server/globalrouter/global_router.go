package globalrouter

import (
	"github.com/op-server/pkg/server/config"
	"github.com/op-server/pkg/server/globalrouter/types"

	"github.com/op-server/pkg/server/dao"

	"github.com/gin-gonic/gin"
	"github.com/op-server/pkg/server/middleware"

	//NOTICE:Register routing
	_ "github.com/op-server/pkg/server/handlers/cluster/v1/meta"
	_ "github.com/op-server/pkg/server/handlers/ping"
)

func NewGinEngine(dao *dao.BaseDao, config *config.HTTPServerConfiguration) *gin.Engine {
	router := gin.Default()
	router.Use(middleware.MiddleWares()...)
	rootGroup := router.Group("/api/op")
	for _, fn := range types.RouterGroupList {
		fn(rootGroup, dao, config)
	}
	return router
}
