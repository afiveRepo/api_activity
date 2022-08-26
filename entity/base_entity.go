package entity

import (
	"time"

	"gorm.io/gorm"
)

const FORMAT_TIME = "2006-01-02 15:04:05"

type BaseEntity struct {
	ID        int64     `gorm:"primary_key:auto_increment" json:"id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func (base *BaseEntity) BeforeCreate(tx *gorm.DB) error {
	tx.Statement.SetColumn("CreatedAt", time.Now().Format(FORMAT_TIME))
	return nil
}
func (base *BaseEntity) BeforeUpdate(tx *gorm.DB) error {
	tx.Statement.SetColumn("UpdatedAt", time.Now().Format(FORMAT_TIME))
	return nil
}
func (base *BaseEntity) BeforeDelete(tx *gorm.DB) error {
	tx.Statement.SetColumn("DeletedAt", time.Now().Format(FORMAT_TIME))
	return nil
}
