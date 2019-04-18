package models

import (
	"github.com/jinzhu/gorm"
	"time"

	"github.com/guregu/null"
)

type CoreUser struct {
	ID                 int         `gorm:"column:id;primary_key" json:"id"`
	Password           string      `gorm:"column:password" json:"password"`
	LastLogin          null.Time   `gorm:"column:last_login" json:"last_login"`
	IsSuperuser        int         `gorm:"column:is_superuser" json:"is_superuser"`
	Username           string      `gorm:"column:username" json:"username"`
	FirstName          string      `gorm:"column:first_name" json:"first_name"`
	LastName           string      `gorm:"column:last_name" json:"last_name"`
	Email              string      `gorm:"column:email" json:"email"`
	IsStaff            int         `gorm:"column:is_staff" json:"is_staff"`
	IsActive           int         `gorm:"column:is_active" json:"is_active"`
	DateJoined         time.Time   `gorm:"column:date_joined" json:"date_joined"`
	FullName           null.String `gorm:"column:full_name" json:"full_name"`
	ReferralTrackingID string      `gorm:"column:referral_tracking_id" json:"referral_tracking_id"`
}

func (c *CoreUser) TableName() string {
	return "core_user"
}

func (c *CoreUser) FirstOrCreate(db *gorm.DB) *CoreUser {
	db.FirstOrCreate(c, "id = ?", c.ID)

	return c
}