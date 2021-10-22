package cluster

import (
	"sync"

	"github.com/op-server/pkg/server/dao"
)

type ClusterService struct {
	lock sync.Mutex
	dao  *dao.BaseDao
}

func NewClusterService(dao *dao.BaseDao) *ClusterService {
	return &ClusterService{
		dao: dao,
	}
}
