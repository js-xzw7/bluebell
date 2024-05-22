package mysql

import (
	"bluebell/models"
	"bluebell/pkg/sonyflake"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
)

const secret = "jszxw7.com"

func encryptPassword(data []byte) string {
	h := md5.New()
	h.Write([]byte(secret))
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}

func Register(user *models.User) (err error) {
	sqlStr := "select count(user_id) from user where username = ?; "

	var count int64
	err = db.Get(&count, sqlStr, user.UserName)
	if err != nil && err != sql.ErrNoRows {
		return
	}

	if count > 0 {
		//用户已存在
		return ErrorUserExit
	}

	userId, err := sonyflake.GetId()

	if err != nil {
		return ErrorGenIDFailed
	}

	//生成加密密码
	password := encryptPassword([]byte(user.Password))
	//插入用户
	sqlStr = "insert into user(user_id,username,password) values (?,?,?)"
	_, err = db.Exec(sqlStr, userId, user.UserName, password)

	return
}

func Login(user *models.User) (err error) {
	originPassword := user.Password

	sqlStr := "select user_id, username, password form user where username = ?"
	err = db.Get(user, sqlStr, user.UserName)
	if err != nil && err != sql.ErrNoRows {
		return
	}

	if err == sql.ErrNoRows {
		return ErrorUserNotExit
	}

	//生成机密密码比较
	password := encryptPassword([]byte(originPassword))
	if password != user.Password {
		return ErrorPasswordWrong
	}

	return
}

func GetuserById(idStr string) (user *models.User, err error) {
	sqlStr := "select * from user where id = ?"
	err = db.Get(user, sqlStr, idStr)
	if err != nil && err != sql.ErrNoRows {
		return
	}

	if err == sql.ErrNoRows {
		return nil, ErrorUserNotExit
	}

	return
}
