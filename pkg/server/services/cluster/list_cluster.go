package cluster

import (
	"github.com/op-server/pkg/server/handlers/cluster/v1/meta/paramtypes"
	"github.com/op-server/pkg/server/model"
)

func (service *ClusterService) List() (paramtypes.ClusterMetasList, error) {
	meta := &model.ClusterMeta{}
	//list clusters by other conditions for excluding attributes
	metas, err := service.dao.FindCluster(meta)
	if err != nil {
		return nil, err
	}

	//filter clusters by attribute
	var resps paramtypes.ClusterMetasList
	for _, meta := range metas {
		resps = append(resps, modelClusterMeta2ClusterMetaRequestParam(&meta))
	}
	return resps, nil
}
