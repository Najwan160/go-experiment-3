package entity

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	CreatedBy *string        `gorm:"column:created_by" json:"created_by"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedBy *string        `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt *time.Time     `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

var ModelColumns = struct {
	DeletedAt string
}{
	DeletedAt: "deleted_at",
}
