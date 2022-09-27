package model

import (
	"time"
)

type Rent struct {
	Rent_id     int `gorm:"primaryKey"`
	User_Id     int
	IdBook      int
	Return_date time.Time
	Created_at  time.Time `gorm:"autoCreateTime"`
	Updated_at  time.Time `gorm:"autoCreateTime"`
}
