package model

import (
	"gorm.io/gorm"
)

type CollectionItem struct {
	gorm.Model
	Name string `gorm:"type:varchar(255);not null;comment:名称"`
	Path string `gorm:"type:varchar(255);not null;comment:路径"`
	Pid  uint64 `gorm:"not null;default:0;comment:父id;index:idx-cid-pid-type;priority:2"`
	Cid  uint64 `gorm:"not null;default:0;comment:所属集合id;index:idx-cid-pid-type;priority:1"`
	Type uint8  `gorm:"not null;default:0;comment:类型，1=目录，2=请求;index:idx-cid-pid-type;priority:3"`
}
