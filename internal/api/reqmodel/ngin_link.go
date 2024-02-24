package reqmodel

type SocialLink struct {
	PlatformName string `json:"platform_name"`
	URL          string `json:"url"`
}

type SocialLinksStruct struct {
	SocialLinks []*SocialLink `json:"social_links"`
}
