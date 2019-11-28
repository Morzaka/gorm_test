package models

import (
	"database/sql/driver"
	"github.com/google/uuid"
)

type role string

const (
	AllPermission role = "admin"
	CanBLockUsers role = "superUser"
)

func (r *role) Scan(value interface{}) error {
	*r = role(value.([]byte))
	return nil
}
func (r role) Value() (driver.Value, error) {
	return string(r), nil
}

type Admin struct {
	//ID   uint      `gorm:"type:uint;primary_key;"`
	IdentityID uuid.UUID `gorm:"type:uuid"`
	Role role      `sql:"type:Role"`
}
