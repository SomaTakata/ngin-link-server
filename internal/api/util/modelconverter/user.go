package modelconverter

import (
	"github.com/SomaTakata/ngin-link-server/internal/api/dbmodel"
	"github.com/SomaTakata/ngin-link-server/internal/api/model"
)

func UserFromDBModels(
	user *dbmodel.User,
	userProgrammingLanguages []*dbmodel.UserProgrammingLanguage,
) *model.User {
	return &model.User{
		ClerkID:              user.ClerkID,
		NginLinkID:           user.NginLinkID,
		Username:             user.Username,
		ProfileImageURI:      user.ProfileImageURI,
		UniversityName:       user.UniversityName,
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
