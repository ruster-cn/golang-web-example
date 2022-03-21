package cluster

import (
	"encoding/json"

	"github.com/op-server/pkg/server/handlers/cluster/v1/meta/paramtypes"
	"github.com/op-server/pkg/server/model"

	log "github.com/ruster-cn/zap-log-wrapper"
)

func (service *ClusterService) FindCluster(param *paramtypes.ClusterMetaRequestParam) (paramtypes.ClusterMetasList, error) {
	meta := &model.ClusterMeta{
		Name:        *param.Name,
		Region:      *param.Region,
		Az:          *param.Az,
		Department:  *param.Department,
		Environment: *param.Environment,
	}
	log.Infof("find cluster condition is:%v", param)
	//list clusters by other conditions for excluding attributes
	metas, err := service.dao.FindCluster(meta)
	if err != nil {
		return nil, err
	}

	//filter clusters by attribute
	var resps paramtypes.ClusterMetasList

	if param.Attribute != nil && len(param.Attribute) > 0 {
		for _, meta := range metas {
			go func(meta model.ClusterMeta) {
				if meta.Attribute != nil {
					var attr map[string]string
					if err := json.Unmarshal(meta.Attribute, &attr); err != nil {
						log.Warnf("unmarshal cluster %s attribute fail,%v", meta.Name, err)
						return
					}
					for k, v1 := range param.Attribute {
						if v2, ok := attr[k]; !ok || v1 != v2 {
							break
						}
					}
					service.lock.Lock()
					defer service.lock.Unlock()
					resps = append(resps, modelClusterMeta2ClusterMetaRequestParam(&meta))
				}

			}(meta)
		}
	} else {
		for _, meta := range metas {
			resps = append(resps, modelClusterMeta2ClusterMetaRequestParam(&meta))
		}
	}

	return resps, nil

}

func modelClusterMeta2ClusterMetaRequestParam(data *model.ClusterMeta) paramtypes.ClusterMetaRequestParam {
	temp := paramtypes.ClusterMetaRequestParam{
		ID:          int64(data.ID),
		Name:        &[]string{data.Name}[0],
		Region:      &[]string{data.Region}[0],
		Az:          &[]string{data.Az}[0],
		Department:  &[]string{data.Department}[0],
		Environment: &[]string{data.Environment}[0],
		KubeConf:    data.KubeConf,
	}
	if err := json.Unmarshal(data.Attribute, &temp.Attribute); err != nil {
		log.Warnf("unmarshal cluster %s attribute fail,%v", data.Name, err)
	}
	return temp
}
