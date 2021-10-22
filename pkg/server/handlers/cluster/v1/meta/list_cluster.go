package meta

import (
	"github.com/gin-gonic/gin"
	"github.com/op-server/pkg/server/core/resp"
)

func (handler *ClusterMetaHandler) list(_ *gin.Context) (interface{}, error) {
	metas, err := handler.service.List()
	if err != nil {
		return nil, resp.NewError(resp.ServerError, err.Error())
	}
	return metas, nil
}
