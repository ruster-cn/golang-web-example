package dao

import "github.com/op-server/pkg/server/model"

func (dao *BaseDao) SaveCluster(meta *model.ClusterMeta) (*model.ClusterMeta, error) {
	result := dao.orm.Create(meta)
	if result.Error != nil {
		return nil, result.Error
	}

	return meta, nil
}

func (dao *BaseDao) UpdateClusterByID(id int, meta *model.ClusterMeta) (*model.ClusterMeta, error) {
	result := dao.orm.Model(&meta).Where("id=?", id).Updates(meta)
	if result.Error != nil {
		return nil, result.Error
	}
	dao.orm.Where("id=?", id, meta)
	if result.Error != nil {
		return nil, result.Error
	}
	return meta, nil
}

func (dao *BaseDao) DeleteClusterByID(id int) error {
	meta := model.ClusterMeta{}
	result := dao.orm.Where("id=?", id).Delete(meta)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (dao *BaseDao) FindCluster(condition *model.ClusterMeta) ([]model.ClusterMeta, error) {
	var meta []model.ClusterMeta
	result := dao.orm.Where(condition).Find(&meta)
	if result.Error != nil {
		return nil, result.Error
	}
	return meta, nil
}
