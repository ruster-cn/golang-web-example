package paramtypes

import "errors"

type ClusterMetaRequestParam struct {
	ID          int64             `json:"id,omitempty"`
	Name        *string           `json:"name,omitempty"`        // 集群名称
	Region      *string           `json:"region,omitempty"`      // region
	Az          *string           `json:"az,omitempty"`          // az
	Department  *string           `json:"department,omitempty"`  // 所属部门
	Environment *string           `json:"environment,omitempty"` // 所属环境
	KubeConf    string            `json:"kube_conf,omitempty"`   // kubeconfig文件内容
	Attribute   map[string]string `json:"attribute,omitempty"`   // 属性
}

func (param *ClusterMetaRequestParam) Validate() error {
	if param.Name == nil && param.Region == nil && param.Az == nil && param.Department == nil && param.Environment == nil {
		return errors.New("must specify a query condition")
	}
	return nil
}

func (param *ClusterMetaRequestParam) ValidateFindClusterParam() error {
	if param.Name == nil && param.Region == nil && param.Az == nil && param.Department == nil && param.Environment == nil && param.Attribute == nil {
		return errors.New("must specify a query condition")
	}
	return nil
}

type ClusterMetasList []ClusterMetaRequestParam
