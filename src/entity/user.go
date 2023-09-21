package entity

type User struct {
	BaseModel
	Name  string `gorm:"type:varchar(100);not null"`
	Email string `gorm:"type:varchar(100);not null"`
}
