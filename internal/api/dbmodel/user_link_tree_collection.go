package dbmodel

import "gorm.io/gorm"

type UserLinkTreeCollection struct {
	gorm.Model
	UserID              uint `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CollectedNginLinkID uint `gorm:"not null"` //FIXME: NginLinkIDに対する外部キーにする
	//FIXME: UserIDとCollectedNginLinkIDでUniqueにする
}
