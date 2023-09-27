package entity

import "github.com/google/uuid"

type Friends struct {
	BaseModel
	UserID uuid.UUID `gorm:"type:uuid;not null"`
	// User   User      `gorm:"foreignKey:UserID;references:ID"`
	// FriendID uuid.UUID `gorm:"type:uuid;not null"`
	// Friend   User      `gorm:"foreignKey:FriendID;references:ID"`
}
