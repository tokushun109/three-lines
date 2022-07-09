package connector

import (
	"api/app/config"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var Db *gorm.DB

// ID, CreatedAt, UpdatedAt, DeletedAtのfieldを持つDefaultModelを継承
type DefaultModel struct {
	ID        *uint          `gorm:"primary_key" json:"-"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func GenerateUuid() (uuidString string, err error) {
	uuidObj, err := uuid.NewRandom()
	uuidString = uuidObj.String()
	return uuidString, err
}

func gormConnect() *gorm.DB {
	// DB接続設定読み込み
	DBUser := config.Config.DBUser
	DBPass := config.Config.DBPass
	Protocol := config.Config.Protocol
	DBName := config.Config.DBName
	SQL_CONNECT := DBUser + ":" + DBPass + "@" + Protocol + "/" + "?charset=utf8mb4&parseTime=True&loc=Asia%2FTokyo"
	DB_CONNECT := DBUser + ":" + DBPass + "@" + Protocol + "/" + DBName + "?charset=utf8&parseTime=True&loc=Asia%2FTokyo"

	// SQLに接続
	dbConnection, err := gorm.Open(mysql.Open(SQL_CONNECT), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	}})
	if err != nil {
		log.Fatalln(err)
	}
	// DBの作成
	cmdCreateDB := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", DBName)
	dbConnection.Exec(cmdCreateDB)

	// DBに接続
	Db, err := gorm.Open(mysql.Open(DB_CONNECT), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	}})
	if err != nil {
		log.Fatalln(err)
	}
	// ローカルの場合はクエリを出力する
	if config.Config.Env == "local" {
		Db = Db.Debug()
	}
	return Db
}

func init() {
	Db = gormConnect()
}
