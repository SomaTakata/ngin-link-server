package resmodel

type NginLinkExchangeHistory struct {
	ClerkID              string   `json:"clerk_id"`
	ExchangedNginLinkIDs []string `json:"exchanged_ngin_link_ids"`
}
