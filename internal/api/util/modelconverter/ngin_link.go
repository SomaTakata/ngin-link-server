package modelconverter

import (
	"github.com/SomaTakata/ngin-link-server/internal/api/dbmodel"
	"github.com/SomaTakata/ngin-link-server/internal/api/httpmodel"
	"github.com/SomaTakata/ngin-link-server/internal/api/model"
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

func SocialLinksFromHTTPModels(
	reqSocialLinks []*httpmodel.SocialLink,
) []*model.SocialLink {
	socialLinks := make([]*model.SocialLink, len(reqSocialLinks))
	for i, reqSocialLink := range reqSocialLinks {
		socialLinks[i] = SocialLinkFromHTTPModel(reqSocialLink)
	}
	return socialLinks
}

func SocialLinkFromHTTPModel(
	reqSocialLink *httpmodel.SocialLink,
) *model.SocialLink {
	return &model.SocialLink{
		PlatformName: reqSocialLink.PlatformName,
		URL:          reqSocialLink.URL,
	}
}

func SocialLinksToHTTPModels(
	socialLinks []*model.SocialLink,
) []*httpmodel.SocialLink {
	resSocialLinks := make([]*httpmodel.SocialLink, len(socialLinks))
	for i, socialLink := range socialLinks {
		resSocialLinks[i] = SocialLinkToHTTPModel(socialLink)
	}
	return resSocialLinks
}

func SocialLinkToHTTPModel(
	socialLink *model.SocialLink,
) *httpmodel.SocialLink {
	return &httpmodel.SocialLink{
		PlatformName: socialLink.PlatformName,
		URL:          socialLink.URL,
	}
}

func SocialLinksToStructHTTPModel(
	socialLinks []*model.SocialLink,
) httpmodel.SocialLinksStruct {
	resSocialLinks := make([]*httpmodel.SocialLink, len(socialLinks))
	for i, socialLink := range socialLinks {
		resSocialLinks[i] = SocialLinkToHTTPModel(socialLink)
	}
	return httpmodel.SocialLinksStruct{
		SocialLinks: resSocialLinks,
	}
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
