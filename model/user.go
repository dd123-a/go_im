package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	UserName string
	PasswordDigest string
	Email string `gorm:"unique"`
	Avatar string 	`gorm:"size:1000"`
	Phone string
	Status string
}

const (
	PassWordCost =12
	Active string="active"
)

func (user *User) SetPassword(password string) error {
	bytes,err:=bcrypt.GenerateFromPassword([]byte(password),PassWordCost)
	if err!=nil{
		return err
	}
	user.PasswordDigest=string(bytes)
	return nil
}

func (user *User) CheckPassword(password string) bool {
	err:=bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest),[]byte(password))
	return err==nil
}

func (user *User) AvatarURL() string {
	signedGetURL :=user.Avatar
	return signedGetURL
}
























