package model

type NginLink struct {
	NginLinkID string  `json:"ngin_link_id"`
	Links      []*Link `json:"links"`
}

type Link struct {
	PlatformName string `json:"platform_name"`
	LinkURL      string `json:"link_url"`
}
