package models

import (
	"database/sql/driver"
	"github.com/google/uuid"
)

type hobby string

const (
	Sport      hobby = "sport"
	Embroidery hobby = "embroidery"
)

func (h *hobby) Scan(value interface{}) error {
	*h = hobby(value.([]byte))
	return nil
}
func (h hobby) Value() (driver.Value, error) {
	return string(h), nil
}

type Hobby struct {
	ID        uuid.UUID `gorm:"type:uuid"`
	HobbyType hobby     `sql:"type:hobby"`
	Comment   string    `gorm:"size:255"`
}
