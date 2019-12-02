package models

import (
	"database/sql/driver"
)

type Role string

const (
	AllPermission Role = "admin"
	CanBLockUsers Role = "superUser"
)

func (r *Role) Scan(value interface{}) error {
	*r = Role(value.([]byte))
	return nil
}
func (r Role) Value() (driver.Value, error) {
	return string(r), nil
}

type Admin struct {
	ID   int  `gorm:"type:int;primary_key;"`
	Role Role `sql:"type:Role"`
}
