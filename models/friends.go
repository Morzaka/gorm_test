package models

import (
	"time"
)

type Friend struct {
	UserFrom int `gorm:"type:int"`
	UserTo   int `gorm:"type:int"`
	Send     time.Time
	Approved bool
}
