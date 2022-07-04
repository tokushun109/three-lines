package config

import (
	"os"
)

type ConfigList struct {
	Port        string
	Sql         string
	LogFile     string
	Env         string
	DBUser      string
	DBName      string
	DBPass      string
	Protocol    string
	ApiBaseUrl  string
	CreatorName string
}

func init() {
	LoadConfig()
}

var Config ConfigList

func LoadConfig() {
	Config = ConfigList{
		Env:         os.Getenv("ENV"),
		Sql:         os.Getenv("SQL"),
		Port:        os.Getenv("PORT"),
		DBUser:      os.Getenv("DB_USER"),
		DBName:      os.Getenv("DB_NAME"),
		DBPass:      os.Getenv("DB_PASS"),
		Protocol:    os.Getenv("PROTOCOL"),
		ApiBaseUrl:  os.Getenv("API_BASE_URL"),
		CreatorName: os.Getenv("CREATOR_NAME"),
	}
}
