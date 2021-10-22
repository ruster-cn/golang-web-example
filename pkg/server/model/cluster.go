package model

import (
	"gorm.io/gorm"
)

type ClusterMeta struct {
	gorm.Model
	Name        string `grom:"name"`        // 集群名称
	Introduce   string `gorm:"introduce"`   //集群说明
	Region      string `grom:"region"`      // region
	Az          string `grom:"az"`          // az
	Department  string `grom:"department"`  // 所属部门
	Environment string `grom:"environment"` // 所属环境
	KubeConf    string `grom:"kube_conf"`   // kubeconfig文件内容
	Attribute   []byte `grom:"attribute"`   // 属性
}

func (ClusterMeta) TableName() string {
	return "t_cluster_meta"
}
