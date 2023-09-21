package entity

import "github.com/varunnayal/migrations/src/enum"

type Category struct {
	BaseModel
	IsActive bool          `gorm:"default:true"`
	Name     string        `gorm:"type:varchar(100);not null"`
	Type     enum.Category `gorm:"type:category_name_enum;"`
}
