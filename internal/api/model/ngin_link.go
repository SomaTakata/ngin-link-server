package model

type NginLink struct {
	NginLinkID  string
	SocialLinks []*SocialLink
}

type SocialLink struct {
	PlatformName string
	URL          string
}
