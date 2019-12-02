package models

import (
	"database/sql/driver"
)

type Hobby string

const (
	Sport      Hobby = "sport"
	Embroidery Hobby = "embroidery"
)

func (h *Hobby) Scan(value interface{}) error {
	*h = Hobby(value.([]byte))
	return nil
}
func (h Hobby) Value() (driver.Value, error) {
	return string(h), nil
}

type HobbyType struct {
	ID      int    `gorm:"type:int"`
	UserID  int    `gorm:"foreignkey:ID"`
	Hobby   Hobby  `sql:"type:hobby"`
	Comment string `gorm:"size:255"`
}
