package models

import (
	"github.com/google/uuid"

	"time"
)

type Friend struct {
	UserFrom uuid.UUID `gorm:"type:uuid"`
	UserTo   uuid.UUID `gorm:"type:uuid"`
	Send     time.Time
	Approved bool
}