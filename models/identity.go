package models

type Identity struct {
	ID       int    `gorm:"type:int;primary_key;"`
	Login    string `gorm:"type:varchar(50);unique;"`
	Password string `gorm:"type:varchar(50)"`
	Salt     string `gorm:"type:varchar(50)"`
	Blocked  bool   `gorm:"type:bool"`
}
