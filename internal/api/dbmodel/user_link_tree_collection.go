package dbmodel

import "gorm.io/gorm"

type UserLinkTreeCollection struct {
	gorm.Model
	UserID              uint   `gorm:"not null"`
	CollectedNginLinkID string `gorm:"not null"` //FIXME: NginLinkIDに対する外部キーにする
	User                User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	//FIXME: UserIDとCollectedNginLinkIDでUniqueにする
}
