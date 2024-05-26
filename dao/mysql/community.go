package mysql

import "bluebell/models"

func AddCommunity(cml *models.CreateCommunityRequest) (err error) {
	exists := db.First(&models.Community{Name: cml.Name})
	if exists.RowsAffected > 0 {
		return ErrorExit
	}

	community := models.Community{Name: cml.Name}
	res := db.Create(&community)
	if err := res.Error; err != nil {
		return ErrorInsertFailed
	}

	res = db.Create(&models.CommunityDetail{
		CommunityId:  community.Id,
		Introduction: cml.Introduction,
	})

	if err := res.Error; err != nil {
		return ErrorInsertFailed
	}

	return
}

func GetCommunityList(cmy models.Community) (communityList []*models.Community, err error) {

	rows := db.Find(communityList).Where(cmy)

	if err = rows.Error; err != nil {
		return
	}

	return
}
