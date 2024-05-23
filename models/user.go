package models

import (
	"bluebell/pkg/crypto"
	"bluebell/pkg/sonyflake"
	"encoding/json"
	"errors"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserId   uint64 `json:"userId" gorm:"column:user_id;"`
	UserName string `json:"username" gorm:"unique;column:username; type:varchar(64); not null;"`
	Password string `json:"password" gorm:"column:password; type:varchar(64); not null;"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	userId, err := sonyflake.GetId()
	if err != nil {
		return
	}

	password := crypto.Encrypt(u.Password)
	u.UserId = userId
	u.Password = password
	return
}

func (u *User) UnmarshalJSON(data []byte) (err error) {
	required := struct {
		UserName string `json:"username" db:"username"`
		Password string `json:"password" db:"password"`
	}{}

	err = json.Unmarshal(data, &required)

	if err != nil {
		return
	} else if len(required.UserName) == 0 {
		err = errors.New("缺少必填字段username")
	} else if len(required.Password) == 0 {
		err = errors.New("缺少必填字段password")
	} else {
		u.UserName = required.UserName
		u.Password = required.Password
	}

	return
}

type RegisterForm struct {
	UserName        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

func (r *RegisterForm) UnmarshalJSON(data []byte) (err error) {
	required := struct {
		UserName        string `json:"username"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirmPassword"`
	}{}

	err = json.Unmarshal(data, &required)

	if err != nil {
		return
	} else if len(required.UserName) == 0 {
		err = errors.New("缺少必填字段 username")
	} else if len(required.Password) == 0 {
		err = errors.New("缺少必填字段 password")
	} else if required.ConfirmPassword != required.Password {
		err = errors.New("两次密码不一致")
	} else {
		r.UserName = required.UserName
		r.Password = required.Password
		r.ConfirmPassword = required.ConfirmPassword
	}

	return
}
