package model

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type User struct {
	User_Id    int `gorm:"primaryKey"`
	Name       string
	Email      string
	Password   string
	Created_at time.Time `gorm:"autoCreateTime"`
	Updated_at time.Time `gorm:"autoCreateTime"`
	Books      []Book    `gorm:"foreignKey:User_Id"`
	Rents      []Rent    `gorm:"foreignKey:User_Id"`
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

func (um UserModel) Insert(newData User) (User, error) {
	err := um.DB.Save(&newData).Error
	if err != nil {
		fmt.Println("error on registrasi", err.Error())
		return User{}, err
	}
	return newData, nil
}

func (um UserModel) Find(Email, Password string) ([]User, error) {
	var res []User
	err := um.DB.Table("users").Select("User_Id", "Name", "Email").Where("Email = ? AND Password = ?", Email, Password).Model(&User{}).Find(&res).Error
	if err != nil {
		fmt.Println("error on querry", err.Error())
		return nil, err
	}
	return res, nil
}

func (um UserModel) Update(newData User) (User, error) {
	err := um.DB.Where("User_Id = ?", newData.User_Id).Updates(&newData).Error
	if err != nil {
		fmt.Println("error on update", err.Error())
		return User{}, err
	}
	return newData, nil
}

func (um UserModel) DeleteAccount(User_Id int) (User, error) {
	err := um.DB.Where("user_id = ?", User_Id).Delete(&User{}).Error
	if err != nil {
		fmt.Println("error on delete account", err.Error())
		return User{}, err
	}
	return User{}, nil
}
