package meta

import (
	"github.com/op-server/pkg/server/core/resp"
	"github.com/op-server/pkg/server/handlers/cluster/v1/meta/paramtypes"

	"github.com/gin-gonic/gin"
)

func (handler *ClusterMetaHandler) register(ctx *gin.Context) (interface{}, error) {
	var param paramtypes.ClusterMetaRequestParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		return nil, resp.NewError(resp.ParamError, err.Error())
	}

	if err := param.Validate(); err != nil {
		return nil, resp.NewError(resp.ParamError, err.Error())
	}

	meta, err := handler.service.Register(&param)
	if err != nil {
		return nil, resp.NewError(resp.ServerError, err.Error())
	}

	return meta, nil
}
