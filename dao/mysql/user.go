package mysql

import (
	"bluebell/models"
	"bluebell/pkg/crypto"
	"database/sql"
)

func Register(user *models.User) (err error) {
	//判断用户是否存在
	var count int64
	db.Model(&models.User{}).Where("name = ?", user.Name).Count(&count)
	if count > 0 {
		//用户已存在
		return ErrorUserExit
	}

	result := db.Create(user)
	if result.Error != nil {
		return ErrorInsertFailed
	}
	return
}

func Login(user *models.User) (err error) {
	originPassword := user.Password

	row := db.Where("name = ?", user.Name).First(&user)
	if row.Error != nil && row.Error != sql.ErrNoRows {
		return
	}

	if err == sql.ErrNoRows {
		return ErrorUserNotExit
	}

	//生成机密密码比较
	password := crypto.Encrypt(originPassword)
	if password != user.Password {
		return ErrorPasswordWrong
	}

	return
}

// func GetuserById(idStr string) (user *models.User, err error) {
// 	sqlStr := "select * from user where id = ?"
// 	row := db.QueryRow(sqlStr, idStr)
// 	err = row.Scan(user)
// 	if err != nil && err != sql.ErrNoRows {
// 		return
// 	}

// 	if err == sql.ErrNoRows {
// 		return nil, ErrorUserNotExit
// 	}

// 	return
// }
