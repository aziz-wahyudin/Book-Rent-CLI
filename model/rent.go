package model

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Rent struct {
	Rent_id     int `gorm:"primaryKey"`
	User_Id     int
	IdBook      int
	Return_date time.Time
	Created_at  time.Time `gorm:"autoCreateTime"`
	Updated_at  time.Time `gorm:"autoCreateTime"`
}

type RentModel struct {
	DB *gorm.DB
}

func (rm RentModel) Input(newData Rent) (Rent, error) {
	err := rm.DB.Save(&newData).Error
	if err != nil {
		fmt.Println("error on borrowing a book", err.Error())
		return Rent{}, err
	}
	return newData, nil
}

func (rm RentModel) DeleteAccount(User_Id int) (Rent, error) {
	err := rm.DB.Where("user_id = ?", User_Id).Delete(&Rent{}).Error
	if err != nil {
		fmt.Println("error on delete account rent", err.Error())
		return Rent{}, err
	}
	return Rent{}, nil
}
