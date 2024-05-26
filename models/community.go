package models

import (
	"time"

	"gorm.io/gorm"
)

// 试用gorm has one
type Community struct {
	Id              uint64 `json:"id" gorm:"primary_key; cloumn:id; type:bigint(20); not null; comment:社区id"`
	Name            string `json:"name" gorm:"column:name; type:varchar(50); not null; comment:社区名称"`
	CommunityDetail CommunityDetail
	CreatedAt       time.Time      `json:"-" gorm:"column:created_at; autoCreateTime"`
	UpdatedAt       time.Time      `json:"-" gorm:"column:updated_at; autoUpdateTime"`
	DeletedAt       gorm.DeletedAt `json:"-" gorm:"index"`
}

type CommunityDetail struct {
	CommunityId  uint64         `json:"communityId" gorm:"cloumn:community_id; type:bingint(20); not null; comment:社区id(外键关联community表)"`
	Introduction string         `json:"introduction" gorm:"cloumn:introduction; type:varchar(200); comment:社区简介"`
	CreatedAt    time.Time      `json:"-" gorm:"column:created_at; autoCreateTime"`
	UpdatedAt    time.Time      `json:"-" gorm:"column:updated_at; autoUpdateTime"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`
}

type CreateCommunityRequest struct {
	Name         string `json:"name" binding:"required"`
	Introduction string `json:"introduction" binding:"required"`
}

type CommunityInfo struct {
	Id           uint64 `json:"id"`
	Name         string `json:"name"`
	Introduction string `json:"introduction"`
}
