package model

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"time"
)

// Base contains common columns for all tables.
type Base struct {
	ID        uuid.UUID  `gorm:"primary_key;"`  // TODO uuid能避免数据被轻易遍历，但性能上会有更多消耗
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
// BeforeCreate will set a UUID rather than numeric ID.
func (base *Base) BeforeCreate(scope *gorm.Scope) error {
	id:= uuid.NewV4()
	return scope.SetColumn("ID", id)
}
