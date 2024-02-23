package db

import (
	"github.com/SomaTakata/ngin-link-server/internal/api/dbmodel"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func NewDB() *gorm.DB {
	dsn := os.Getenv("DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&dbmodel.User{})
	db.AutoMigrate(&dbmodel.UserProgrammingLanguage{})
	db.AutoMigrate(&dbmodel.UserSocialLink{})
	db.AutoMigrate(&dbmodel.UserLinkTreeCollection{})

	return db
}
