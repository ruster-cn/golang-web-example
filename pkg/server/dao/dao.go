package dao

import "gorm.io/gorm"

type BaseDao struct {
	orm *gorm.DB
}

func NewBaseDao(orm *gorm.DB) *BaseDao {
	return &BaseDao{orm: orm}
}
