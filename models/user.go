package models

import (
	"bluebell/pkg/crypto"
	"bluebell/pkg/sonyflake"
	"encoding/json"
	"errors"
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        uint64         `json:"id" gorm:"column:id;primarykey"`
	Name      string         `json:"name" gorm:"unique;column:name; type:varchar(64); not null;"`
	Password  string         `json:"password" gorm:"column:password; type:varchar(64); not null;"`
	CreatedAt time.Time      `json:"-" gorm:"column:created_at; autoCreateTime"`
	UpdatedAt time.Time      `json:"-" gorm:"column:updated_at; autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	userId, err := sonyflake.GetId()
	if err != nil {
		return
	}

	password := crypto.Encrypt(u.Password)
	u.Id = userId
	u.Password = password
	return
}

func (u *User) UnmarshalJSON(data []byte) (err error) {
	required := struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}{}

	err = json.Unmarshal(data, &required)

	if err != nil {
		return
	} else if len(required.Name) == 0 {
		err = errors.New("缺少必填字段username")
	} else if len(required.Password) == 0 {
		err = errors.New("缺少必填字段password")
	} else {
		u.Name = required.Name
		u.Password = required.Password
	}

	return
}

type RegisterForm struct {
	Name            string `json:"name"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

func (r *RegisterForm) UnmarshalJSON(data []byte) (err error) {
	required := struct {
		Name            string `json:"name"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirmPassword"`
	}{}

	err = json.Unmarshal(data, &required)

	if err != nil {
		return
	} else if len(required.Name) == 0 {
		err = errors.New("缺少必填字段 name")
	} else if len(required.Password) == 0 {
		err = errors.New("缺少必填字段 password")
	} else if required.ConfirmPassword != required.Password {
		err = errors.New("两次密码不一致")
	} else {
		r.Name = required.Name
		r.Password = required.Password
		r.ConfirmPassword = required.ConfirmPassword
	}

	return
}
