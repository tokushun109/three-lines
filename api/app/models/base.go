package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var Db *gorm.DB

// ID, CreatedAt, UpdatedAt, DeletedAtのfieldを持つDefaultModelを継承
type DefaultModel struct {
	ID        *uint          `gorm:"primary_key" json:"-"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// 暗号(Hash)化
func passwordEncrypt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

// 暗号(Hash)と入力された平パスワードの比較
func compareHashAndPassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
