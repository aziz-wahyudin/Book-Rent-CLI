package model

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Book struct {
	IdBook     int `gorm:"primaryKey; autoIncrement:true; "`
	User_Id    int
	Name       string
	Status     string
	Owner      int
	Rent_By    int
	Created_at time.Time `gorm:"autoCreateTime"`
	Updated_at time.Time `gorm:"autoCreateTime"`
	Rents      []Rent    `gorm:"foreignKey:IdBook"`
}

type BookModel struct {
	DB *gorm.DB
}

func (bm BookModel) GetAll() ([]Book, error) {
	var res []Book
	err := bm.DB.Table("Book").Select("IdBook", "Name", "Email", "Status", "Owner", "Rent_By", "Created_at", "Updated_at").Model(&Book{}).Find(&res).Error
	if err != nil {
		fmt.Println("error on querry", err.Error())
		return nil, err
	}
	return res, nil
}
func (bm BookModel) Insert(newData Book) (Book, error) {
	err := bm.DB.Save(&newData).Error
	if err != nil {
		fmt.Println("error on insert", err.Error())
		return Book{}, err
	}
	return newData, nil
}

func (bm BookModel) Show(User_Id int) ([]Book, error) {
	var res []Book
	err := bm.DB.Where("User_Id = ?", User_Id).Find(&res).Error
	if err != nil {
		fmt.Println("error on query", err.Error())
		return nil, err
	}
	return nil, err

}

/*
// Update with conditions
db.Model(&User{}).Where("active = ?", true).Update("name", "hello")
// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE active=true;

err := um.DB.Save(&newData).Error
	if err != nil {
		fmt.Println("error on registrasi", err.Error())
		return User{}, err
	}
	return newData, nil
}
*/

func (bm BookModel) Input(newData Book) (Book, error) {
	err := bm.DB.Save(&newData).Error
	if err != nil {
		fmt.Println("error on adding new book", err.Error())
		return Book{}, err
	}
	return newData, nil
}
