package modelconverter

import (
	"github.com/SomaTakata/ngin-link-server/internal/api/dbmodel"
	"github.com/SomaTakata/ngin-link-server/internal/api/model"
)

func NginLinkExchangeHistoriesFromDBModels(
	user *dbmodel.User,
	userLinkTreeCollections []*dbmodel.UserLinkTreeCollection,
) []*model.NginLinkExchangeHistory {
	nginLinkExchangeHistories := make([]*model.NginLinkExchangeHistory, len(userLinkTreeCollections))
	for i, userLinkTreeCollection := range userLinkTreeCollections {
		nginLinkExchangeHistories[i] = NginLinkExchangeHistoryFromDBModel(user, userLinkTreeCollection)
	}
	return nginLinkExchangeHistories
}

func NginLinkExchangeHistoryFromDBModel(
	user *dbmodel.User,
	userLinkTreeCollection *dbmodel.UserLinkTreeCollection,
) *model.NginLinkExchangeHistory {
	return &model.NginLinkExchangeHistory{
		ClerkID:             user.ClerkID,
		ExchangedNginLinkID: userLinkTreeCollection.CollectedNginLinkID,
	}
}
