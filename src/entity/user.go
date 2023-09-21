package entity

import "time"

type User struct {
	BaseModel
	Name  string    `gorm:"type:varchar(100);not null"`
	Email string    `gorm:"type:varchar(100);not null"`
	Phone string    `gorm:"type:varchar(15);not null"`
	DOB   time.Time `gorm:"type:date;not null"`
}
