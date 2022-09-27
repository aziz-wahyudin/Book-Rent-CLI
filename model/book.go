package model

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Book struct {
	IdBook     int `gorm:"primaryKey; autoIncrement:true; "`
	Name       string
	Status     string
	Owner      int
	Rent_By    int
	Created_at time.Time `gorm:"autoCreateTime"`
	Updated_at time.Time `gorm:"autoCreateTime"`
}

type BookModel struct {
	DB *gorm.DB
}

func (gm BookModel) GetAll() ([]Book, error) {
	var res []Book
	err := gm.DB.Table("Book").Select("IdBook", "Name", "Email", "Status", "Owner", "Rent_By", "Created_at", "Updated_at").Model(&Book{}).Find(&res).Error
	if err != nil {
		fmt.Println("error on querry", err.Error())
		return nil, err
	}
	return res, nil
}
func (mm BookModel) Insert(newData Book) (Book, error) {
	err := mm.DB.Save(&newData).Error
	if err != nil {
		fmt.Println("error on insert", err.Error())
		return Book{}, err
	}
	return newData, nil
}
