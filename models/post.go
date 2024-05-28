package models

import (
	"bluebell/pkg/sonyflake"
	"time"

	"gorm.io/gorm"
)

type Post struct {
	Id          uint64         `json:"id" gorm:"column:id;type:bigint(20);primary_key;not null;auto_increment;"`
	Title       string         `json:"title" gorm:"column:title;type:varchar(50);not null; bingding"`
	Content     string         `json:"content" gorm:"column:content;type:longtext;not null;"`
	AuthorId    uint64         `json:"authorId" gorm:"column:author_id; type:bigint(20); not null;"`
	CommunityId uint64         `json:"communityId" gorm:"column:community_id; type:bigint(20); not null;"`
	Community   Community      `json:"-" gorm:"foreignkey:CommunityId;"`
	Status      int32          `json:"status" gorm:"column:status; type:int not null; default:0;"`
	CreatedAt   time.Time      `json:"-" gorm:"column:crated_at; autoCreateTime"`
	UpdatedAt   time.Time      `json:"-" gorm:"column:updated_at; autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"column:deleted_at;index"`
}

type PostReq struct {
	Id          uint64 `json:"id" binding:"omitempty"`
	Title       string `json:"title" binding:"required"`
	Content     string `json:"content"  binding:"required"`
	AuthorId    uint64 `json:"authorId"  binding:"required"`
	CommunityId uint64 `json:"communityId"  binding:"required"`
	Status      int32  `json:"status"  binding:"omitempty"`
}

type PostIds struct {
	Ids []uint64 `json:"ids" binding:"required"`
}

func (p *Post) BeforeCreate(tx *gorm.DB) (err error) {
	id, err := sonyflake.GetId()
	if err != nil {
		return
	}

	p.Id = id
	return
}
