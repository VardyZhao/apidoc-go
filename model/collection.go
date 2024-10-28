package model

import (
	"gorm.io/gorm"
)

type Collection struct {
	gorm.Model
	Name string `gorm:"type:varchar(255);not null;comment:名称"`
}
