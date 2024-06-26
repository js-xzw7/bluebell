package models

import (
	"bluebell/pkg/sonyflake"
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
	CommunityId  uint64         `json:"communityId" gorm:"column:community_id; type:bigint(20); not null; comment:社区id(外键关联community表)"`
	Introduction string         `json:"introduction" gorm:"column:introduction; type:varchar(200); comment:社区简介"`
	CreatedAt    time.Time      `json:"-" gorm:"column:created_at; autoCreateTime"`
	UpdatedAt    time.Time      `json:"-" gorm:"column:updated_at; autoUpdateTime"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`
}

func (c *Community) BeforeCreate(tx *gorm.DB) (err error) {
	id, err := sonyflake.GetId()
	if err != nil {
		return
	}
	c.Id = id
	return
}

type CreateCommunityRequest struct {
	Name         string `form:"name" binding:"required"`
	Introduction string `form:"introduction" binding:"required"`
}

type CommunityInfo struct {
	Id           uint64 `json:"id" form:"id" binding:"omitempty"`
	Name         string `json:"name" form:"name" binding:"omitempty"`
	Introduction string `json:"introduction" form:"introduction" binding:"omitempty"`
}
