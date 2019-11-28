package models

import "github.com/google/uuid"

type User struct {
	ID         uint      `gorm:"type:uint; serial;"`
	IdentityID uuid.UUID `gorm:"type:uuid"`
	FName      string    `gorm:"not null"`
	LName      string    `gorm:"not null"`
	Bio        string    `gorm:"size:255"`
	Hobbies    []Hobby   `gorm:"foreignkey:ID"`
	Friends    []Friend  `gorm:"many2many:friends;"`
}
