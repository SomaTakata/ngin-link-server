package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ClerkID         string `gorm:"not null;index:index_clerk_id,unique"`
	NginLinkID      string `gorm:"not null;index:index_ngin_link_id,unique"`
	Username        string `gorm:"not null"`
	ProfileImageURI string
	UniversityName  string
}
