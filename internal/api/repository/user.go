package repository

import (
	"github.com/SomaTakata/ngin-link-server/internal/api/model"
	"github.com/SomaTakata/ngin-link-server/internal/api/util/modelconverter"
	"gorm.io/gorm"
)

type UserRepository interface {
	Get(clerkID string) (*model.User, error)
	Create(user *model.User) (*model.User, error)
	Update(user *model.User) (*model.User, error)
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

type userRepository struct {
	db *gorm.DB
}

func (r userRepository) Get(clerkID string) (*model.User, error) {
	//TODO
	return &model.User{}, nil
}

func (r userRepository) Create(user *model.User) (*model.User, error) {
	dbUser := modelconverter.UserToDBModel(user)
	dbUserProgrammingLanguages := modelconverter.UserProgrammingLanguagesToDBModels(user.ProgrammingLanguages)
	dbUserSocialLinks := modelconverter.UserSocialLinksToDBModels(user.NginLink.SocialLinks)

	r.db.Transaction(func(tx *gorm.DB) error {
		tx.Create(dbUser)

		for _, dbUserProgrammingLanguage := range dbUserProgrammingLanguages {
			dbUserProgrammingLanguage.UserID = dbUser.ID
		}
		tx.Create(dbUserProgrammingLanguages)

		for _, dbUserSocialLink := range dbUserSocialLinks {
			dbUserSocialLink.UserID = dbUser.ID
		}
		tx.Create(dbUserSocialLinks)

		return nil
	})

	newUser := modelconverter.UserFromDBModels(dbUser, dbUserProgrammingLanguages, dbUserSocialLinks)
	return newUser, nil
}

func (r userRepository) Update(user *model.User) (*model.User, error) {
	//TODO
	return &model.User{}, nil
}
