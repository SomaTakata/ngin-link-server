package dbmodel

import "gorm.io/gorm"

type UserProgrammingLanguage struct {
	gorm.Model
	UserID              uint `gorm:"not null"`
	ProgrammingLanguage string
	User                User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
