package modelconverter

import (
	"github.com/SomaTakata/ngin-link-server/internal/api/dbmodel"
	"github.com/SomaTakata/ngin-link-server/internal/api/model"
)

func NginLinkExchangeHistoryFromDBModels(
	user *dbmodel.User,
	userLinkTreeCollection []*dbmodel.UserLinkTreeCollection,
) *model.NginLinkExchangeHistory {
	return &model.NginLinkExchangeHistory{
		ClerkID:              user.ClerkID,
		ExchangedNginLinkIDs: ExchangedNginLinkIDsFromDBModels(userLinkTreeCollection),
	}
}

func ExchangedNginLinkIDsFromDBModels(
	userLinkTreeCollections []*dbmodel.UserLinkTreeCollection,
) []string {
	exchangedNginLinkIDs := make([]string, len(userLinkTreeCollections))
	for i, userLinkTreeCollection := range userLinkTreeCollections {
		exchangedNginLinkIDs[i] = userLinkTreeCollection.CollectedNginLinkID
	}
	return exchangedNginLinkIDs
}
