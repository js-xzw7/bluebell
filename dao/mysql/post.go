package mysql

import (
	"bluebell/models"
	"errors"
	"reflect"

	"gorm.io/gorm"
)

func AddPost(post models.PostReq) error {
	result := db.Create(&models.Post{
		Title:       post.Title,
		Content:     post.Content,
		AuthorId:    post.AuthorId,
		CommunityId: post.CommunityId,
		Status:      post.Status,
	})

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetPostById(id uint64) (post *models.Post, err error) {
	row := db.First(&post, id)

	if row.Error != nil && !errors.Is(row.Error, gorm.ErrRecordNotFound) {
		return nil, row.Error
	}

	return post, nil
}

func GetPostListByIds(ids []uint64) (posts []*models.Post, err error) {
	rows := db.Find(&posts, ids)

	if rows.Error != nil && errors.Is(rows.Error, gorm.ErrRecordNotFound) {
		return nil, rows.Error
	}

	return posts, nil
}

func GetPostList(post models.PostReq) (posts []*models.Post, err error) {
	tx := db.Model(&models.Post{})

	if !reflect.DeepEqual(post, models.PostReq{}) {
		if post.Title != "" {
			tx.Where("title = ?", post.Title)
		}

		if post.AuthorId != 0 {
			tx.Where("author_id = ?", post.AuthorId)
		}

		if post.CommunityId != 0 {
			tx.Where("community_id = ?", post.CommunityId)
		}

	}

	tx.Find(&posts).Order("id desc")

	if tx.Error != nil && errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil, tx.Error
	}

	return posts, nil
}
