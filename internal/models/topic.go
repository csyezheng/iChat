package models

import "time"


type Topic struct {
	Id               time.Time   `gorm:"column:id" json:"id"`
	Uid              time.Time   `gorm:"column:uid" json:"uid"`
	CreatedAt        time.Time   `gorm:"column:create_at" json:"create_at"`
	UpdatedAt        time.Time   `gorm:"column:update_at" json:"update_at"`
	DeletedAt        time.Time   `gorm:"column:delete_at" json:"delete_at"`
	Name             string      `gorm:"column:name" json:"name"`
}