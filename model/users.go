package model

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type User struct {
	User_Id    int
	Name       string
	Email      string
	Password   string
	Created_at time.Time
	Updated_at time.Time
}

type UserModel struct {
	DB *gorm.DB
}

func (um UserModel) GetAll() ([]User, error) {
	var res []User
	err := um.DB.Table("users").Select("User_Id", "Name", "Email", "Password", "Created_at", "Updated_at").Model(&User{}).Find(&res).Error
	if err != nil {
		fmt.Println("error on querry", err.Error())
		return nil, err
	}
	return res, nil
}
