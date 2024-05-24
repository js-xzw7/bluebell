package mysql

import (
	"bluebell/models"
	"bluebell/settings"
	"fmt"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

// 初始化mysql连接
func Init(cfg *settings.MySQLConfig) (err error) {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
	db, err = gorm.Open(mysql.Open(dns), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_",                              // table name prefix, table for `User` would be `t_users`
			SingularTable: true,                              // use singular table name, table for `User` would be `user` with this option enabled
			NoLowerCase:   false,                             // skip the snake_casing of names
			NameReplacer:  strings.NewReplacer("CID", "Cid"), // use name replacer to change struct/field name before convert it to db name
		},
	})
	if err != nil {
		return
	}

	sqlDB, err := db.DB()
	//sets the maximum number of connections in the idle connection pool
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	//sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	//sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	db.AutoMigrate(&models.User{}, &models.Community{}, &models.CommunityDetail{})
	return
}
