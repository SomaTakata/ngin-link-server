package modelconverter

import (
	"github.com/SomaTakata/ngin-link-server/internal/api/dbmodel"
	"github.com/SomaTakata/ngin-link-server/internal/api/httpmodel"
	"github.com/SomaTakata/ngin-link-server/internal/api/model"
)

func UserFromDBModels(
	user *dbmodel.User,
	userProgrammingLanguages []*dbmodel.UserProgrammingLanguage,
	userSocialLinks []*dbmodel.UserSocialLink,
) *model.User {
	return &model.User{
		ClerkID: user.ClerkID,
		NginLink: &model.NginLink{
			NginLinkID:  user.NginLinkID,
			SocialLinks: SocialLinksFromDBModels(userSocialLinks),
		},
		Username:             user.Username,
		ProfileImageURL:      user.ProfileImageURL,
		Description:          user.Description,
		ProgrammingLanguages: ProgrammingLanguagesFromDBModels(userProgrammingLanguages),
		JobRole:              user.JobRole,
	}
}

func ProgrammingLanguagesFromDBModels(
	userProgrammingLanguages []*dbmodel.UserProgrammingLanguage,
) []string {
	programmingLanguages := make([]string, len(userProgrammingLanguages))
	for i, userProgrammingLanguage := range userProgrammingLanguages {
		programmingLanguages[i] = userProgrammingLanguage.ProgrammingLanguage
	}
	return programmingLanguages
}

func UserFromCreateUserReqModel(createUser *httpmodel.CreateUser) *model.User {
	return &model.User{
		NginLink: &model.NginLink{
			NginLinkID:  createUser.NginLinkID,
			SocialLinks: SocialLinksFromHTTPModels(createUser.SocialLinks),
		},
		Username:             createUser.Username,
		ProfileImageURL:      createUser.ProfileImageURL,
		Description:          createUser.Description,
		ProgrammingLanguages: createUser.ProgrammingLanguages,
		JobRole:              createUser.JobRole,
	}
}

func UserToDBModel(user *model.User) *dbmodel.User {
	return &dbmodel.User{
		ClerkID:         user.ClerkID,
		NginLinkID:      user.NginLink.NginLinkID,
		Username:        user.Username,
		ProfileImageURL: user.ProfileImageURL,
		Description:     user.Description,
		JobRole:         user.JobRole,
	}
}

func UserProgrammingLanguagesToDBModels(programmingLanguages []string) []*dbmodel.UserProgrammingLanguage {
	//UserIDはRepository内で設定する
	userProgrammingLanguages := make([]*dbmodel.UserProgrammingLanguage, len(programmingLanguages))
	for i, programmingLanguage := range programmingLanguages {
		userProgrammingLanguages[i] = &dbmodel.UserProgrammingLanguage{
			ProgrammingLanguage: programmingLanguage,
		}
	}
	return userProgrammingLanguages
}

func UserSocialLinksToDBModels(socialLinks []*model.SocialLink) []*dbmodel.UserSocialLink {
	//UserIDはRepository内で設定する
	userSocialLinks := make([]*dbmodel.UserSocialLink, len(socialLinks))
	for i, socialLink := range socialLinks {
		userSocialLinks[i] = &dbmodel.UserSocialLink{
			PlatformName:  socialLink.PlatformName,
			SocialLinkURL: socialLink.URL,
		}
	}
	return userSocialLinks
}

func UserToHTTPModel(user *model.User) *httpmodel.User {
	return &httpmodel.User{
		ClerkID: user.ClerkID,
		NginLink: &httpmodel.NginLink{
			NginLinkID:  user.NginLink.NginLinkID,
			SocialLinks: SocialLinksToHTTPModels(user.NginLink.SocialLinks),
		},
		Username:             user.Username,
		ProfileImageURL:      user.ProfileImageURL,
		Description:          user.Description,
		ProgrammingLanguages: user.ProgrammingLanguages,
		JobRole:              user.JobRole,
	}
}
