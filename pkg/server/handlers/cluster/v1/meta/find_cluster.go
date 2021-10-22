package meta

import (
	"github.com/op-server/pkg/server/core/resp"
	"github.com/op-server/pkg/server/handlers/cluster/v1/meta/paramtypes"

	"github.com/gin-gonic/gin"
)

func (handler *ClusterMetaHandler) find(ctx *gin.Context) (interface{}, error) {
	var param paramtypes.ClusterMetaRequestParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		return nil, resp.NewError(resp.ParamError, err.Error())
	}
	if err := param.ValidateFindClusterParam(); err != nil {
		return nil, resp.NewError(resp.ParamError, err.Error())
	}
	metas, err := handler.service.FindCluster(&param)
	if err != nil {
		return nil, resp.NewError(resp.ServerError, err.Error())
	}
	return metas, nil
}
