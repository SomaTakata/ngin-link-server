package modelconverter

import (
	"github.com/SomaTakata/ngin-link-server/internal/api/dbmodel"
	"github.com/SomaTakata/ngin-link-server/internal/api/model"
	"github.com/SomaTakata/ngin-link-server/internal/api/reqmodel"
	"github.com/SomaTakata/ngin-link-server/internal/api/resmodel"
)

func NginLinkFromDBModels(
	user *dbmodel.User,
	userSocialLinks []*dbmodel.UserSocialLink,
) *model.NginLink {
	return &model.NginLink{
		NginLinkID:  user.NginLinkID,
		SocialLinks: SocialLinksFromDBModels(userSocialLinks),
	}
}

func SocialLinksFromDBModels(
	userSocialLinks []*dbmodel.UserSocialLink,
) []*model.SocialLink {
	socialLinks := make([]*model.SocialLink, len(userSocialLinks))
	for i, userSocialLink := range userSocialLinks {
		socialLinks[i] = SocialLinkFromDBModel(userSocialLink)
	}
	return socialLinks
}

func SocialLinkFromDBModel(
	userSocialLink *dbmodel.UserSocialLink,
) *model.SocialLink {
	return &model.SocialLink{
		PlatformName: userSocialLink.PlatformName,
		URL:          userSocialLink.SocialLinkURL,
	}
}

func SocialLinksFromReqModels(
	reqSocialLinks []*reqmodel.SocialLink,
) []*model.SocialLink {
	socialLinks := make([]*model.SocialLink, len(reqSocialLinks))
	for i, reqSocialLink := range reqSocialLinks {
		socialLinks[i] = &model.SocialLink{
			PlatformName: reqSocialLink.PlatformName,
			URL:          reqSocialLink.URL,
		}
	}
	return socialLinks
}

func SocialLinksToResModels(
	socialLinks []*model.SocialLink,
) []*resmodel.SocialLink {
	resSocialLinks := make([]*resmodel.SocialLink, len(socialLinks))
	for i, socialLink := range socialLinks {
		resSocialLinks[i] = &resmodel.SocialLink{
			PlatformName: socialLink.PlatformName,
			URL:          socialLink.URL,
		}
	}
	return resSocialLinks
}

func SocialLinksToDBModels(
	socialLinks []*model.SocialLink,
) []*dbmodel.UserSocialLink {
	userSocialLinks := make([]*dbmodel.UserSocialLink, len(socialLinks))
	for i, socialLink := range socialLinks {
		userSocialLinks[i] = SocialLinkToDBModel(socialLink)
	}
	return userSocialLinks
}

func SocialLinkToDBModel(
	socialLink *model.SocialLink,
) *dbmodel.UserSocialLink {
	//UserIDはRepository内で設定する
	return &dbmodel.UserSocialLink{
		PlatformName:  socialLink.PlatformName,
		SocialLinkURL: socialLink.URL,
	}
}
