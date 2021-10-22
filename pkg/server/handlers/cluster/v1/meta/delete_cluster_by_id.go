package meta

import (
	"github.com/gin-gonic/gin"
	"github.com/op-server/pkg/server/core/resp"
)

func (handler *ClusterMetaHandler) deleteClusterByID(ctx *gin.Context) (interface{}, error) {
	clusterID, err := handler.QueryClusterID(ctx)
	if err != nil {
		return nil, resp.NewError(resp.ParamError, err.Error())
	}
	if err := handler.service.DeleteClusterByID(clusterID); err != nil {
		return nil, resp.NewError(resp.ServerError, err.Error())
	}
	return nil, nil
}
