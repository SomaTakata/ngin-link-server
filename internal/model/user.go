package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	NginLinkID     string `gorm:"index:index_ngin_link_id,unique"`
	Username       string
	UniversityName string
}
