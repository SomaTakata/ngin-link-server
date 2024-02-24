package repository

import (
	"github.com/SomaTakata/ngin-link-server/internal/api/dbmodel"
	"github.com/SomaTakata/ngin-link-server/internal/api/model"
	"github.com/SomaTakata/ngin-link-server/internal/api/util/modelconverter"
	"gorm.io/gorm"
)

type LinkRepository interface {
	GetExchangeHistory(clerkID string) (*model.NginLinkExchangeHistory, error)
	Update(clerkID string, socialLinks []*model.SocialLink) ([]*model.SocialLink, error)
	CreateExchangeHistory(clerkID string, nginLinkID string) (*model.NginLinkExchangeHistory, error)
}

func NewLinkRepository(db *gorm.DB) LinkRepository {
	return &linkRepository{db}
}

type linkRepository struct {
	db *gorm.DB
}

func (r linkRepository) GetExchangeHistory(clerkID string) (*model.NginLinkExchangeHistory, error) {
	var collectedNginLinkIDs []string

	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&dbmodel.UserNginLinkCollection{}).
			Select("collected_ngin_link_id").
			Joins("left join users ON user_ngin_link_collections.user_id = users.id AND users.clerk_id = ?", clerkID).
			Scan(&collectedNginLinkIDs).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	nginLinkExchangeHistory := &model.NginLinkExchangeHistory{
		ClerkID:              clerkID,
		ExchangedNginLinkIDs: collectedNginLinkIDs,
	}
	return nginLinkExchangeHistory, nil
}

func (r linkRepository) Update(clerkID string, socialLinks []*model.SocialLink) ([]*model.SocialLink, error) {
	//DELETE&INSERTで対応
	var dbUser *dbmodel.User
	dbUserSocialLinks := modelconverter.SocialLinksToDBModels(socialLinks)

	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&dbUser, "clerk_id = ?", clerkID).Error; err != nil {
			return err
		}

		if err := tx.Where("user_id = ?", dbUser.ID).Delete(&dbmodel.UserSocialLink{}).Error; err != nil {
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

	newSocialLinks := modelconverter.SocialLinksFromDBModels(dbUserSocialLinks)
	return newSocialLinks, nil
}

func (r linkRepository) CreateExchangeHistory(clerkID string, nginLinkID string) (*model.NginLinkExchangeHistory, error) {
	var dbUser *dbmodel.User
	var collectedNginLinkIDs []string

	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&dbUser, "clerk_id = ?", clerkID).Error; err != nil {
			return err
		}

		userNginLinkCollection := &dbmodel.UserNginLinkCollection{
			UserID:              dbUser.ID,
			CollectedNginLinkID: nginLinkID,
		}
		if err := tx.Create(userNginLinkCollection).Error; err != nil {
			return err
		}

		if err := tx.Model(&dbmodel.UserNginLinkCollection{}).
			Select("collected_ngin_link_id").
			Where("user_id = ?", dbUser.ID).
			Scan(&collectedNginLinkIDs).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	nginLinkExchangeHistory := &model.NginLinkExchangeHistory{
		ClerkID:              clerkID,
		ExchangedNginLinkIDs: collectedNginLinkIDs,
	}
	return nginLinkExchangeHistory, nil
}
