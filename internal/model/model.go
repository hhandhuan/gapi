package model

import "gorm.io/gorm"

type BaseModel struct {
	ID        uint `gorm:"primarykey" json:"id"`
	CreatedAt uint `gorm:"gorm:autoCreateTime|column:created_at" json:"created_at"`
	UpdatedAt uint `gorm:"gorm:autoUpdateTime|column:updated_at" json:"updated_at"`
	DeletedAt uint `gorm:"column:deleted_at" json:"deleted_at"`
}

func Trash(db *gorm.DB) *gorm.DB {
	return db.Where("deleted_at != ?", 0)
}

func UnTrash(db *gorm.DB) *gorm.DB {
	return db.Where("deleted_at = ?", 0)
}

func (m *BaseModel) IsEmpty() bool {
	return m.ID <= 0
}
