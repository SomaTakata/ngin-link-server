package model

import "gorm.io/gorm"

type UserProgrammingLanguage struct {
	gorm.Model
	UserID              uint `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ProgrammingLanguage string
}
