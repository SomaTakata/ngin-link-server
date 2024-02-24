package modelconverter

import (
	"github.com/SomaTakata/ngin-link-server/internal/api/dbmodel"
	"github.com/SomaTakata/ngin-link-server/internal/api/model"
)

func NginLinkExchangeHistoryFromDBModels(
	user *dbmodel.User,
	userNginLinkCollection []*dbmodel.UserNginLinkCollection,
) *model.NginLinkExchangeHistory {
	return &model.NginLinkExchangeHistory{
		ClerkID:              user.ClerkID,
		ExchangedNginLinkIDs: ExchangedNginLinkIDsFromDBModels(userNginLinkCollection),
	}
}

func ExchangedNginLinkIDsFromDBModels(
	userNginLinkCollections []*dbmodel.UserNginLinkCollection,
) []string {
	exchangedNginLinkIDs := make([]string, len(userNginLinkCollections))
	for i, userNginLinkCollection := range userNginLinkCollections {
		exchangedNginLinkIDs[i] = userNginLinkCollection.CollectedNginLinkID
	}
	return exchangedNginLinkIDs
}
