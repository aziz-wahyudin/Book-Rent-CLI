package model

import (
	"fmt"
	"log"
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

func (bm BookModel) ShowBook() []Book {
	var BookList = []Book{}
	if err := bm.DB.Find(&BookList).Error; err != nil {
		log.Print(err)
		return nil
	}
	return BookList
}

func (bm BookModel) Show(User_Id int) ([]Book, error) {
	var res []Book
	err := bm.DB.Where("User_Id = ?", User_Id).Find(&res).Error
	if err != nil {
		fmt.Println("error on query", err.Error())
		return nil, err
	}
	return res, nil

}

func (bm BookModel) Input(newData Book) (Book, error) {
	err := bm.DB.Save(&newData).Error
	if err != nil {
		fmt.Println("error on adding new book", err.Error())
		return Book{}, err
	}
	return newData, nil
}
