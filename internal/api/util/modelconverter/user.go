package modelconverter

import (
	"github.com/SomaTakata/ngin-link-server/internal/api/dbmodel"
	"github.com/SomaTakata/ngin-link-server/internal/api/model"
	"github.com/SomaTakata/ngin-link-server/internal/api/reqmodel"
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

func UserFromCreateUserReqModel(createUser *reqmodel.CreateUser) *model.User {
	return &model.User{
		NginLink: &model.NginLink{
			NginLinkID:  createUser.NginLinkID,
			SocialLinks: SocialLinksFromUserReqModel(createUser.SocialLinks),
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
