package dbmodel

import "gorm.io/gorm"

type UserSocialLink struct {
	gorm.Model
	UserID        uint
	PlatformName  string `gorm:"not null"`
	SocialLinkURL string `gorm:"not null"`
	User          User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
