package types

import (
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/op-server/pkg/server/config"
	"github.com/op-server/pkg/server/dao"
)

var RouterGroupList []AddRouterGroupFunc

//NOTICE: every module implements its own AddRouterGroupFunc to register url
type AddRouterGroupFunc func(routerGroup *gin.RouterGroup, dao *dao.BaseDao, config *config.HTTPServerConfiguration)

//NOTICE: every module add init function,add AddRouterGroupFunc to routerGroupList,
func InsertAddRouterGroupFunc(fn AddRouterGroupFunc) {
	lock := sync.Mutex{}
	lock.Lock()
	defer lock.Unlock()
	RouterGroupList = append(RouterGroupList, fn)
}
