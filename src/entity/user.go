package entity

import (
	"time"

	"github.com/varunnayal/migrations/src/enum"
)

type User struct {
	BaseModel
	Name     string          `gorm:"type:varchar(100);not null"`
	Email    string          `gorm:"type:varchar(100);not null"`
	Phone    string          `gorm:"type:varchar(15);not null"`
	Status   enum.UserStatus `gorm:"type:user_status_enum;not null"`
	DOB      time.Time       `gorm:"type:date;not null"`
	SecPhone string          `gorm:"type:varchar(15);not null"`
}
