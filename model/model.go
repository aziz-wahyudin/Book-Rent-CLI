package model

import (
	"fmt"
	"gorm.io/gorm"
)
type Book struct {
	IdBook	int
	Name	string
	Status	string
	Owner	int
	Rent_By	int
	Created_at	
}

type MuridModel struct {
	db *gorm.DB
}

func  (mc )