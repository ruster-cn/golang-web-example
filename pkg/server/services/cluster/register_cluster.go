package cluster

import (
	"encoding/json"
	"fmt"

	"github.com/op-server/pkg/server/handlers/cluster/v1/meta/paramtypes"

	"github.com/op-server/pkg/server/model"
)

func (service *ClusterService) Register(param *paramtypes.ClusterMetaRequestParam) (*paramtypes.ClusterMetaRequestParam, error) {
	meta := &model.ClusterMeta{
		Name:        *param.Name,
		Region:      *param.Region,
		Az:          *param.Az,
		Department:  *param.Department,
		Environment: *param.Environment,
		KubeConf:    param.KubeConf,
	}
	attributeByte, err := json.Marshal(param.Attribute)
	if err != nil {
		return nil, fmt.Errorf("marshal cluster attribute fail,%v", err)
	}
	meta.Attribute = attributeByte

	newMeta, err := service.dao.SaveCluster(meta)
	if err != nil {
		return nil, fmt.Errorf("save cluster meta info fail,%v", err)
	}
	result := modelClusterMeta2ClusterMetaRequestParam(newMeta)
	return &result, nil
}
