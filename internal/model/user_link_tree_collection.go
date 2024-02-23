package model

import "gorm.io/gorm"

type UserLinkTreeCollection struct {
	gorm.Model
	UserID              uint `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CollectedNginLinkID uint //FIXME: NginLinkIDに対する外部キーにする
	//FIXME: UserIDとCollectedNginLinkIDでUniqueにする
}
