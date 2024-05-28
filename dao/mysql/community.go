package mysql

import (
	"bluebell/models"
	"errors"
	"reflect"

	"gorm.io/gorm"
)

func AddCommunity(cml *models.CreateCommunityRequest) (err error) {
	exists := db.First(&models.Community{Name: cml.Name})
	if exists.RowsAffected > 0 {
		return ErrorExit
	}

	// 插入community表
	res := db.Create(&models.Community{
		Name: cml.Name,
		CommunityDetail: models.CommunityDetail{
			Introduction: cml.Introduction,
		},
	})

	// community := models.Community{Name: cml.Name}
	// res := db.Create(&community)
	// if err := res.Error; err != nil {
	// 	return ErrorInsertFailed
	// }

	// res = db.Create(&models.CommunityDetail{
	// 	CommunityId:  community.Id,
	// 	Introduction: cml.Introduction,
	// })

	if err := res.Error; err != nil {
		return ErrorInsertFailed
	}

	return
}

func GetCommunityList(cmy models.CommunityInfo) (communityList []*models.CommunityInfo, err error) {

	rows := db.Model(&models.Community{}).Select("id", "name", "CommunityDetail.Introduction as introduction")

	if !reflect.DeepEqual(cmy, models.CommunityInfo{}) {
		if cmy.Name != "" {
			rows.Where("name = ?", cmy.Name)
		}

		if cmy.Introduction != "" {
			rows.Where("CommunityDetail.Introduction = ?", cmy.Introduction)
		}
	}

	rows.Joins("CommunityDetail").Find(&communityList).Order("id desc")

	if err = rows.Error; err != nil {
		return
	}

	return
}

func GetCommunityById(id uint64) (community *models.CommunityInfo, err error) {
	row := db.Model(&models.Community{}).Select("id", "name", "CommunityDetail.Introduction as introduction").Joins("CommunityDetail").First(&community, id)

	if row.Error != nil && !errors.Is(row.Error, gorm.ErrRecordNotFound) {
		return nil, row.Error
	}

	return
}
