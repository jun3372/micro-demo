package model

import (
	"github.com/jinzhu/gorm"
)

type Member struct {
	gorm.Model
	Avatar   string `json:"avatar" gorm:"type:varchar(100);"`
	Nickname string `json:"nickname" gorm:"type:varchar(100);"`
	Username string `json:"username" gorm:"type:varchar(100);unique_index"`
	Password string `json:"password" gorm:"type:varchar(256)"`
}
