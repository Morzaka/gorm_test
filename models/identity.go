package models

import "github.com/google/uuid"

type Identity struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key;"`
	Login      string    `gorm:"type:varchar(100);unique;"`
	Password   string    `gorm:"type:varchar(50)"`
	Salt       string    `gorm:"type:varchar(50)"`
	Blocked    bool      `gorm:"type:bool"`
	User       User      `gorm:"foreignkey:IdentityID"`
	Admin      Admin     `gorm:"foreignkey:IdentityID"`
}
