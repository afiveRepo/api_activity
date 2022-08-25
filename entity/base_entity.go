package entity

import (
	"time"

	"gorm.io/gorm"
)

const FORMAT_TIME = "2006-01-02 15:04:05"

type BaseEntity struct {
	ID        int64     `gorm:"primary_key:auto_increment" json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt time.Time `json:"-"`
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
