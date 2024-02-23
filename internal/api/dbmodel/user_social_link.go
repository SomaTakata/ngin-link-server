package dbmodel

import "gorm.io/gorm"

type UserSocialLink struct {
	gorm.Model
	UserID       uint   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	PlatformName string `gorm:"not null"`
	LinkURL      string `gorm:"not null"`
}
