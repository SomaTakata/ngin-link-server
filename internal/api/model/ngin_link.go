package model

type NginLink struct {
	NginLinkID string
	Links      []Link
}

type Link struct {
	PlatformName string
	LinkURL      string
}
