package connector

import (
	"api/app/config"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// ID, CreatedAt, UpdatedAt, DeletedAtのfieldを持つDefaultModelを継承
type DefaultModel struct {
	ID        *uint          `gorm:"primary_key" json:"-"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func GormConnect() *gorm.DB {
	// DB接続設定読み込み
	DBUser := config.Config.DBUser
	DBPass := config.Config.DBPass
	Protocol := config.Config.Protocol
	DBName := config.Config.DBName
	DB_CONNECT := DBUser + ":" + DBPass + "@" + Protocol + "/" + DBName + "?charset=utf8&parseTime=True&loc=Asia%2FTokyo"

	// DBに接続
	db, err := gorm.Open(mysql.Open(DB_CONNECT), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	}})
	if err != nil {
		log.Fatalln(err)
	}
	// ローカルの場合はクエリを出力する
	if config.Config.Env == "local" {
		db = db.Debug()
	}
	return db
}
