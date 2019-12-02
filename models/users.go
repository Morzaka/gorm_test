package models

type User struct {
	ID      int         `gorm:"type:int; primary_key"`
	FName   string      `gorm:"varchar(50); not null"`
	LName   string      `gorm:"varchar(50); not null"`
	Bio     string      `gorm:"size:255"`
	Hobbies []HobbyType `gorm:"foreignkey:ID"`
	Friends []Friend    `gorm:"many2many:friends;"`
}
