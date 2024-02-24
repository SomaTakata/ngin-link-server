package repository

import (
	"github.com/SomaTakata/ngin-link-server/internal/api/dbmodel"
	"github.com/SomaTakata/ngin-link-server/internal/api/model"
	"github.com/SomaTakata/ngin-link-server/internal/api/util/modelconverter"
	"gorm.io/gorm"
)

type UserRepository interface {
	Get(clerkID string) (*model.User, error)
	Create(user *model.User) (*model.User, error)
	Update(user *model.User) (*model.User, error)
	GetByNginLinkID(nginLinkID string) (*model.User, error)
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

type userRepository struct {
	db *gorm.DB
}

func (r userRepository) Get(clerkID string) (*model.User, error) {
	var dbUser *dbmodel.User
	var dbUserProgrammingLanguages []*dbmodel.UserProgrammingLanguage
	var dbUserSocialLinks []*dbmodel.UserSocialLink

	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&dbUser, "clerk_id = ?", clerkID).Error; err != nil {
			return err
		}

		if err := tx.Find(&dbUserProgrammingLanguages, "user_id = ?", dbUser.ID).Error; err != nil {
			return err
		}

		if err := tx.Find(&dbUserSocialLinks, "user_id = ?", dbUser.ID).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	user := modelconverter.UserFromDBModels(dbUser, dbUserProgrammingLanguages, dbUserSocialLinks)
	return user, nil
}

func (r userRepository) Create(user *model.User) (*model.User, error) {
	dbUser := modelconverter.UserToDBModel(user)
	dbUserProgrammingLanguages := modelconverter.UserProgrammingLanguagesToDBModels(user.ProgrammingLanguages)
	dbUserSocialLinks := modelconverter.UserSocialLinksToDBModels(user.NginLink.SocialLinks)

	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(dbUser).Error; err != nil {
			return err
		}

		for _, dbUserProgrammingLanguage := range dbUserProgrammingLanguages {
			dbUserProgrammingLanguage.UserID = dbUser.ID
		}
		if err := tx.Create(dbUserProgrammingLanguages).Error; err != nil {
			return err
		}

		for _, dbUserSocialLink := range dbUserSocialLinks {
			dbUserSocialLink.UserID = dbUser.ID
		}
		if err := tx.Create(dbUserSocialLinks).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	newUser := modelconverter.UserFromDBModels(dbUser, dbUserProgrammingLanguages, dbUserSocialLinks)
	return newUser, nil
}

func (r userRepository) Update(user *model.User) (*model.User, error) {
	//TODO
	return &model.User{}, nil
}

func (r userRepository) GetByNginLinkID(nginLinkID string) (*model.User, error) {
	var dbUser *dbmodel.User
	var dbUserProgrammingLanguages []*dbmodel.UserProgrammingLanguage
	var dbUserSocialLinks []*dbmodel.UserSocialLink

	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&dbUser, "ngin_link_id = ?", nginLinkID).Error; err != nil {
			return err
		}

		if err := tx.Find(&dbUserProgrammingLanguages, "user_id = ?", dbUser.ID).Error; err != nil {
			return err
		}

		if err := tx.Find(&dbUserSocialLinks, "user_id = ?", dbUser.ID).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	user := modelconverter.UserFromDBModels(dbUser, dbUserProgrammingLanguages, dbUserSocialLinks)
	return user, nil
}
