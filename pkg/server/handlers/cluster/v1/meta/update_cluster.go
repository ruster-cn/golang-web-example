package meta

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/op-server/pkg/server/core/resp"

	"github.com/op-server/pkg/server/handlers/cluster/v1/meta/paramtypes"

	"github.com/gin-gonic/gin"
)

func (handler *ClusterMetaHandler) updateCluster(ctx *gin.Context) (interface{}, error) {
	clusterID, err := handler.QueryClusterID(ctx)
	if err != nil {
		return nil, resp.NewError(resp.ParamError, err.Error())
	}
	var param paramtypes.ClusterMetaRequestParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		return nil, resp.NewError(resp.ParamError, err.Error())
	}

	if err := param.Validate(); err != nil {
		return nil, resp.NewError(resp.ParamError, err.Error())
	}
	meta, err := handler.service.UpdateClusterMeta(clusterID, &param)
	if err != nil {
		return nil, resp.NewError(resp.ServerError, err.Error())
	}

	return meta, nil
}

func (handler *ClusterMetaHandler) QueryClusterID(ctx *gin.Context) (int, error) {
	idStr := ctx.Query("id")
	if idStr == "" {
		return 0, errors.New("cluster id is null")
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, fmt.Errorf("get cluster id fail,%v", err)
	}
	return id, nil
}
