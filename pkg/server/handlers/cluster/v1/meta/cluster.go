package meta

import (
	"github.com/op-server/pkg/server/dao"
	"github.com/op-server/pkg/server/services/cluster"
)

type ClusterMetaHandler struct {
	service *cluster.ClusterService
}

func NewClusterMetaHandler(dao *dao.BaseDao) *ClusterMetaHandler {
	return &ClusterMetaHandler{
		service: cluster.NewClusterService(dao),
	}
}
