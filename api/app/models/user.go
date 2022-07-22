package models

import (
	"api/app/db/connector"
	"errors"
	"fmt"
)

type User struct {
	DefaultModel
	LoginID  string `json:"loginId" validate:"min=1,max=20"`
	Password string `json:"password" validate:"min=1,max=20"`
}

func (u *User) CreateUser() error {
	user := User{}

	db := connector.GormConnect()

	db.Where("login_id = ?", u.LoginID).
		Limit(1).
		Find(&user)

	if user.ID != nil {
		err := errors.New("login_idが既に登録されています。")
		fmt.Println(err)
		return err
	}

	encryptPw, err := passwordEncrypt(u.Password)
	if err != nil {
		fmt.Println(err)
		return err
	}

	user = User{LoginID: u.LoginID, Password: encryptPw}
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Create(&user).Error; err != nil {
		fmt.Println(err)
		return err
	}
	tx.Commit()

	return nil
}
