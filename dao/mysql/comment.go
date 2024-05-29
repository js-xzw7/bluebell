package mysql

import (
	"bluebell/models"
	"reflect"
	"strings"
)

func AddCommont(comment models.CommentInfo) (err error) {
	tx := db.Create(&models.Comment{
		ParentId: comment.ParentId,
		PostId:   comment.PostId,
		AuthorId: comment.AuthorId,
		Content:  comment.Content,
	})

	if err = tx.Error; err != nil {
		return
	}

	return
}

func GetCommentListByIds(ids []uint64) (comments []*models.CommentInfo, err error) {
	err = db.Model(&models.Comment{}).Where("id in (?)", ids).Find(&comments).Error
	return
}

func GetCommentList(comment models.CommentInfo) (comments []*models.CommentInfo, err error) {
	tx := db.Model(&models.Comment{})

	if !reflect.DeepEqual(comment, &models.CommentInfo{}) {
		if comment.ParentId != 0 {
			tx.Where("parent_id = ?", comment.ParentId)
		}

		if comment.AuthorId != 0 {
			tx.Where("author_id = ?", comment.AuthorId)
		}

		if comment.PostId != 0 {
			tx.Where("post_id = ?", comment.PostId)
		}

		if comment.Content != "" {
			var content strings.Builder
			content.WriteByte('%')
			content.WriteString(comment.Content)
			content.WriteByte('%')
			where := content.String()

			tx.Where("content like ?", where)
		}
	}

	err = tx.Find(&comments).Order("id desc").Error
	return
}
