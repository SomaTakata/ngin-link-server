package modelconverter

import (
	"github.com/SomaTakata/ngin-link-server/internal/api/dbmodel"
	"github.com/SomaTakata/ngin-link-server/internal/api/model"
)

func NginLinkFromDBModels(
	user *dbmodel.User,
	userSocialLinks []*dbmodel.UserSocialLink,
) *model.NginLink {
	return &model.NginLink{
		NginLinkID: user.NginLinkID,
		Links:      LinksFromDBModels(userSocialLinks),
	}
}

func LinksFromDBModels(
	userSocialLinks []*dbmodel.UserSocialLink,
) []*model.Link {
	links := make([]*model.Link, len(userSocialLinks))
	for i, userSocialLink := range userSocialLinks {
		links[i] = LinkFromDBModel(userSocialLink)
	}
	return links
}

func LinkFromDBModel(
	userSocialLink *dbmodel.UserSocialLink,
) *model.Link {
	return &model.Link{
		PlatformName: userSocialLink.PlatformName,
		LinkURL:      userSocialLink.LinkURL,
	}
}
